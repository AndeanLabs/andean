package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSwap{}

func NewMsgSwap(creator string, poolId string, tokenInDenom string, tokenInAmount uint64, tokenOutDenom string, minTokenOutAmount uint64) *MsgSwap {
	return &MsgSwap{
		Creator:           creator,
		PoolId:            poolId,
		TokenInDenom:      tokenInDenom,
		TokenInAmount:     tokenInAmount,
		TokenOutDenom:     tokenOutDenom,
		MinTokenOutAmount: minTokenOutAmount,
	}
}

func (msg *MsgSwap) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
