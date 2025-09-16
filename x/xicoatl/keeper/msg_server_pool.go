package keeper

import (
	"context"

	"andean/x/xicoatl/types"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetPool(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// Get creator address
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Get module address
	moduleAddr := k.AccountKeeper.GetModuleAddress(types.ModuleName)

	// Create coin objects for the initial liquidity
	coinA := sdk.NewCoin(msg.DenomA, math.NewIntFromUint64(msg.AmountA))
	coinB := sdk.NewCoin(msg.DenomB, math.NewIntFromUint64(msg.AmountB))
	liquidity := sdk.NewCoins(coinA, coinB)

	// Transfer liquidity from creator to the module account
	err = k.BankKeeper.SendCoins(ctx, creator, moduleAddr, liquidity)
	if err != nil {
		return nil, err
	}

	var pool = types.Pool{
		Creator: msg.Creator,
		Index:   msg.Index,
		DenomA:  msg.DenomA,
		AmountA: msg.AmountA,
		DenomB:  msg.DenomB,
		AmountB: msg.AmountB,
	}

	k.SetPool(
		ctx,
		pool,
	)
	return &types.MsgCreatePoolResponse{}, nil
}

func (k msgServer) UpdatePool(goCtx context.Context, msg *types.MsgUpdatePool) (*types.MsgUpdatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPool(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var pool = types.Pool{
		Creator: msg.Creator,
		Index:   msg.Index,
		DenomA:  msg.DenomA,
		AmountA: msg.AmountA,
		DenomB:  msg.DenomB,
		AmountB: msg.AmountB,
	}

	k.SetPool(ctx, pool)

	return &types.MsgUpdatePoolResponse{}, nil
}

func (k msgServer) DeletePool(goCtx context.Context, msg *types.MsgDeletePool) (*types.MsgDeletePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPool(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePool(
		ctx,
		msg.Index,
	)

	return &types.MsgDeletePoolResponse{}, nil
}
