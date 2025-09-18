package types

import (
	"testing"

	"andean/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgInitiateBridgeTransfer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgInitiateBridgeTransfer
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgInitiateBridgeTransfer{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgInitiateBridgeTransfer{
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
