package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "andean/testutil/keeper"
	"andean/testutil/sample"
	"andean/x/itzel/keeper"
	"andean/x/itzel/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestSubmitPrice(t *testing.T) {
	k, ctx, authorizedOracle := keepertest.ItzelKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	unauthorizedOracle := sample.AccAddress()

	tests := []struct {
		name      string
		msg       *types.MsgSubmitPrice
		expectErr error
	}{
		{
			name: "successful submission from authorized oracle",
			msg: &types.MsgSubmitPrice{
				Creator: authorizedOracle,
				Source:  "test-source",
				Price:   100,
			},
			expectErr: nil,
		},
		{
			name: "failed submission from unauthorized oracle",
			msg: &types.MsgSubmitPrice{
				Creator: unauthorizedOracle,
				Source:  "test-source",
				Price:   100,
			},
			expectErr: sdkerrors.ErrUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.SubmitPrice(ctx, tt.msg)
			if tt.expectErr != nil {
				require.ErrorIs(t, err, tt.expectErr)
			} else {
				require.NoError(t, err)
				// Check if the price was stored
				price, found := k.GetPrice(ctx, tt.msg.Source)
				require.True(t, found)
				require.Equal(t, tt.msg.Price, price.Value)
			}
		})
	}
}
