package keeper

import (
	"context"

	"andean/x/inti/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateLazyBridgeTransfer(goCtx context.Context, msg *types.MsgCreateLazyBridgeTransfer) (*types.MsgCreateLazyBridgeTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetLazyBridgeTransfer(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var lazyBridgeTransfer = types.LazyBridgeTransfer{
		Creator:     msg.Creator,
		Index:       msg.Index,
		SourceChain: msg.SourceChain,
		DestChain:   msg.DestChain,
		Amount:      msg.Amount,
		Status:      msg.Status,
	}

	k.SetLazyBridgeTransfer(
		ctx,
		lazyBridgeTransfer,
	)
	return &types.MsgCreateLazyBridgeTransferResponse{}, nil
}

func (k msgServer) UpdateLazyBridgeTransfer(goCtx context.Context, msg *types.MsgUpdateLazyBridgeTransfer) (*types.MsgUpdateLazyBridgeTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetLazyBridgeTransfer(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var lazyBridgeTransfer = types.LazyBridgeTransfer{
		Creator:     msg.Creator,
		Index:       msg.Index,
		SourceChain: msg.SourceChain,
		DestChain:   msg.DestChain,
		Amount:      msg.Amount,
		Status:      msg.Status,
	}

	k.SetLazyBridgeTransfer(ctx, lazyBridgeTransfer)

	return &types.MsgUpdateLazyBridgeTransferResponse{}, nil
}

func (k msgServer) DeleteLazyBridgeTransfer(goCtx context.Context, msg *types.MsgDeleteLazyBridgeTransfer) (*types.MsgDeleteLazyBridgeTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetLazyBridgeTransfer(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveLazyBridgeTransfer(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteLazyBridgeTransferResponse{}, nil
}
