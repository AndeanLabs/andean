package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "andean/testutil/keeper"
	"andean/x/inti/keeper"
	"andean/x/inti/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestLazyBridgeTransferMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.IntiKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateLazyBridgeTransfer{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateLazyBridgeTransfer(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetLazyBridgeTransfer(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestLazyBridgeTransferMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateLazyBridgeTransfer
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateLazyBridgeTransfer{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateLazyBridgeTransfer{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateLazyBridgeTransfer{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.IntiKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateLazyBridgeTransfer{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateLazyBridgeTransfer(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateLazyBridgeTransfer(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetLazyBridgeTransfer(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestLazyBridgeTransferMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteLazyBridgeTransfer
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteLazyBridgeTransfer{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteLazyBridgeTransfer{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteLazyBridgeTransfer{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.IntiKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateLazyBridgeTransfer(ctx, &types.MsgCreateLazyBridgeTransfer{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteLazyBridgeTransfer(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetLazyBridgeTransfer(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
