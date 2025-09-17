package keeper

import (
	"context"

	"andean/x/xicoatl/types"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Get the pool from the store
	pool, found := k.GetPool(ctx, msg.PoolId)
	if !found {
		return nil, errors.Wrapf(types.ErrPoolNotFound, "pool %s does not exist", msg.PoolId)
	}

	// 2. Determine which token is which in the pool
	var tokenIn, tokenOut sdk.Coin
	var poolTokenInBalance, poolTokenOutBalance math.Int

	if msg.TokenInDenom == pool.DenomA && msg.TokenOutDenom == pool.DenomB {
		poolTokenInBalance = math.NewIntFromUint64(pool.AmountA)
		poolTokenOutBalance = math.NewIntFromUint64(pool.AmountB)
	} else if msg.TokenInDenom == pool.DenomB && msg.TokenOutDenom == pool.DenomA {
		poolTokenInBalance = math.NewIntFromUint64(pool.AmountB)
		poolTokenOutBalance = math.NewIntFromUint64(pool.AmountA)
	} else {
		return nil, errors.Wrapf(types.ErrInvalidTokens, "invalid token pair for pool %s", msg.PoolId)
	}

	// 2.5. Get the oracle price to be used for dynamic fees
	source := msg.TokenInDenom + "/" + msg.TokenOutDenom
	oraclePrice, found := k.ItzelKeeper.GetAggregatedPrice(ctx, source)
	if found {
		k.Logger().Info("Oracle price found for source", "source", source, "price", oraclePrice.Price)
	}

	// 3. Create the input coin object from the message
	tokenIn = sdk.NewCoin(msg.TokenInDenom, math.NewIntFromUint64(msg.TokenInAmount))

	// 4. Calculate the output amount based on the constant product formula (x * y = k)
	// We apply a dynamic fee based on the oracle price
	feePercentage := math.LegacyNewDecWithPrec(3, 3) // Default 0.3% fee
	if found {
		k.Logger().Info("Oracle price found for source", "source", source, "price", oraclePrice.Price)
		feePercentage = math.LegacyNewDecWithPrec(5, 3) // Use 0.5% fee if oracle price is available
	}

	fee := math.LegacyNewDecFromInt(tokenIn.Amount).Mul(feePercentage)
	tokenInAmountAfterFee := math.LegacyNewDecFromInt(tokenIn.Amount).Sub(fee).TruncateInt()

	numerator := poolTokenOutBalance.Mul(tokenInAmountAfterFee)
	denominator := poolTokenInBalance.Add(tokenInAmountAfterFee)
	tokenOutAmount := numerator.Quo(denominator)

	if tokenOutAmount.IsZero() {
		return nil, types.ErrZeroOutput
	}

	// 5. Check for slippage
	if tokenOutAmount.LT(math.NewIntFromUint64(msg.MinTokenOutAmount)) {
		return nil, errors.Wrapf(types.ErrSlippage, "output amount %s is less than minimum %d", tokenOutAmount, msg.MinTokenOutAmount)
	}

	// 6. Perform the token transfers
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	moduleAddr := k.AccountKeeper.GetModuleAddress(types.ModuleName)
	tokenOut = sdk.NewCoin(msg.TokenOutDenom, tokenOutAmount)

	// Send input tokens from user to the module
	err = k.BankKeeper.SendCoins(ctx, creator, moduleAddr, sdk.NewCoins(tokenIn))
	if err != nil {
		return nil, err
	}

	// Send output tokens from the module to the user
	err = k.BankKeeper.SendCoins(ctx, moduleAddr, creator, sdk.NewCoins(tokenOut))
	if err != nil {
		return nil, err
	}

	// 7. Update the pool balances
	if msg.TokenInDenom == pool.DenomA {
		pool.AmountA += msg.TokenInAmount
		pool.AmountB -= tokenOutAmount.Uint64()
	} else {
		pool.AmountB += msg.TokenInAmount
		pool.AmountA -= tokenOutAmount.Uint64()
	}

	k.SetPool(ctx, pool)

	return &types.MsgSwapResponse{Amount: tokenOutAmount.Uint64()}, nil
}
