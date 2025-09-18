package types

import (
	"testing"

	"andean/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateLazyBridgeTransfer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateLazyBridgeTransfer
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateLazyBridgeTransfer{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateLazyBridgeTransfer{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateLazyBridgeTransfer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateLazyBridgeTransfer
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateLazyBridgeTransfer{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateLazyBridgeTransfer{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteLazyBridgeTransfer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteLazyBridgeTransfer
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteLazyBridgeTransfer{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteLazyBridgeTransfer{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
