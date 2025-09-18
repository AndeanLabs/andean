package keeper

import (
	"context"

	"andean/x/inti/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateLazyTransfer(goCtx context.Context, msg *types.MsgCreateLazyTransfer) (*types.MsgCreateLazyTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Convert creator string address to sdk.AccAddress
	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Send coins from creator's account to the module's account (locking the funds)
	moduleAddr := k.AccountKeeper.GetModuleAddress(types.ModuleName)
	err = k.BankKeeper.SendCoins(ctx, creatorAddress, moduleAddr, sdk.NewCoins(msg.Amount))
	if err != nil {
		return nil, err
	}

	// Emit an event to notify off-chain listeners (like relayers)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"LazyTransferCreated",
			sdk.NewAttribute("creator", msg.Creator),
			sdk.NewAttribute("recipient", msg.Recipient),
			sdk.NewAttribute("destination_chain", msg.DestinationChain),
			sdk.NewAttribute("amount", msg.Amount.String()),
		),
	)

	return &types.MsgCreateLazyTransferResponse{}, nil
}
