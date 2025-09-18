package keeper

import (
	"context"

	"andean/x/inti/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ConfirmBridgeTransfer(goCtx context.Context, msg *types.MsgConfirmBridgeTransfer) (*types.MsgConfirmBridgeTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Implement relayer authorization check. For now, we trust the creator.

	// Find the original transfer record.
	transfer, found := k.GetLazyBridgeTransfer(ctx, msg.Index)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrKeyNotFound, "lazy bridge transfer with index %s not found", msg.Index)
	}

	// Check if the transfer is already completed or failed.
	if transfer.Status != "PENDING" {
		return nil, errors.Wrapf(types.ErrInvalidStatus, "transfer is not in PENDING state, current status: %s", transfer.Status)
	}

	// Update the status to COMPLETE.
	transfer.Status = "COMPLETE"

	// Save the updated transfer object.
	k.SetLazyBridgeTransfer(ctx, transfer)

	k.Logger().Info("confirmed lazy bridge transfer", "index", msg.Index, "relayer", msg.Creator, "finalTxHash", msg.FinalTxHash)

	return &types.MsgConfirmBridgeTransferResponse{}, nil
}
