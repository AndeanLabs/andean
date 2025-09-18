package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateLazyTransfer{}

func NewMsgCreateLazyTransfer(creator string, amount sdk.Coin, recipient string, destinationChain string) *MsgCreateLazyTransfer {
	return &MsgCreateLazyTransfer{
		Creator:          creator,
		Amount:           amount,
		Recipient:        recipient,
		DestinationChain: destinationChain,
	}
}

func (msg *MsgCreateLazyTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
