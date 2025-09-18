package keeper

import (
	"context"
	"fmt"

	"andean/x/inti/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) InitiateBridgeTransfer(goCtx context.Context, msg *types.MsgInitiateBridgeTransfer) (*types.MsgInitiateBridgeTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Validate creator address
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// 2. Parse the amount string into sdk.Coin
	// The format is expected to be "{amount}{denom}", e.g., "1000uandean"
	coin, err := sdk.ParseCoinNormalized(msg.Amount)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid amount string: %v", err)
	}

	// 3. Lock the funds by sending them from the user to the module account
	moduleAddr := k.AccountKeeper.GetModuleAddress(types.ModuleName)
	err = k.BankKeeper.SendCoins(ctx, creator, moduleAddr, sdk.NewCoins(coin))
	if err != nil {
		return nil, err
	}

	// 4. Create and store the lazyBridgeTransfer object
	// We'll use a combination of source, destination, and a nonce/timestamp for the index
	// For now, let's use a simple index. A more robust solution would be needed for production.
	index := fmt.Sprintf("%s-%s-%d", msg.SourceChain, msg.DestChain, ctx.BlockHeight())

	transfer := types.LazyBridgeTransfer{
		Index:       index,
		SourceChain: msg.SourceChain,
		DestChain:   msg.DestChain,
		Amount:      msg.Amount,
		Status:      "PENDING",
		Creator:     msg.Creator,
	}

	k.SetLazyBridgeTransfer(ctx, transfer)

	k.Logger().Info("initiated lazy bridge transfer", "index", index, "from", msg.Creator, "amount", msg.Amount)

	return &types.MsgInitiateBridgeTransferResponse{}, nil
}
