package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgConfirmBridgeTransfer{}

func NewMsgConfirmBridgeTransfer(creator string, index string, finalTxHash string) *MsgConfirmBridgeTransfer {
	return &MsgConfirmBridgeTransfer{
		Creator:     creator,
		Index:       index,
		FinalTxHash: finalTxHash,
	}
}

func (msg *MsgConfirmBridgeTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
