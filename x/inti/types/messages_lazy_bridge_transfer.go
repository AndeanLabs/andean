package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateLazyBridgeTransfer{}

func NewMsgCreateLazyBridgeTransfer(
	creator string,
	index string,
	sourceChain string,
	destChain string,
	amount string,
	status string,

) *MsgCreateLazyBridgeTransfer {
	return &MsgCreateLazyBridgeTransfer{
		Creator:     creator,
		Index:       index,
		SourceChain: sourceChain,
		DestChain:   destChain,
		Amount:      amount,
		Status:      status,
	}
}

func (msg *MsgCreateLazyBridgeTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateLazyBridgeTransfer{}

func NewMsgUpdateLazyBridgeTransfer(
	creator string,
	index string,
	sourceChain string,
	destChain string,
	amount string,
	status string,

) *MsgUpdateLazyBridgeTransfer {
	return &MsgUpdateLazyBridgeTransfer{
		Creator:     creator,
		Index:       index,
		SourceChain: sourceChain,
		DestChain:   destChain,
		Amount:      amount,
		Status:      status,
	}
}

func (msg *MsgUpdateLazyBridgeTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteLazyBridgeTransfer{}

func NewMsgDeleteLazyBridgeTransfer(
	creator string,
	index string,

) *MsgDeleteLazyBridgeTransfer {
	return &MsgDeleteLazyBridgeTransfer{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteLazyBridgeTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
