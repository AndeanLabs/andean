package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePool{}

func NewMsgCreatePool(
	creator string,
	index string,
	denomA string,
	amountA uint64,
	denomB string,
	amountB uint64,

) *MsgCreatePool {
	return &MsgCreatePool{
		Creator: creator,
		Index:   index,
		DenomA:  denomA,
		AmountA: amountA,
		DenomB:  denomB,
		AmountB: amountB,
	}
}

func (msg *MsgCreatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePool{}

func NewMsgUpdatePool(
	creator string,
	index string,
	denomA string,
	amountA uint64,
	denomB string,
	amountB uint64,

) *MsgUpdatePool {
	return &MsgUpdatePool{
		Creator: creator,
		Index:   index,
		DenomA:  denomA,
		AmountA: amountA,
		DenomB:  denomB,
		AmountB: amountB,
	}
}

func (msg *MsgUpdatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePool{}

func NewMsgDeletePool(
	creator string,
	index string,

) *MsgDeletePool {
	return &MsgDeletePool{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeletePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
