package types

import (
	"testing"

	"andean/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgConfirmBridgeTransfer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgConfirmBridgeTransfer
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgConfirmBridgeTransfer{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgConfirmBridgeTransfer{
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
