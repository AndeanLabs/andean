package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "andean/testutil/keeper"
	"andean/x/xicoatl/keeper"
	"andean/x/xicoatl/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPoolMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.XicoatlKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreatePool{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreatePool(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetPool(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestPoolMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdatePool
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdatePool{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdatePool{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdatePool{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.XicoatlKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreatePool{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreatePool(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdatePool(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetPool(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestPoolMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeletePool
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeletePool{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeletePool{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeletePool{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.XicoatlKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreatePool(ctx, &types.MsgCreatePool{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeletePool(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetPool(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
