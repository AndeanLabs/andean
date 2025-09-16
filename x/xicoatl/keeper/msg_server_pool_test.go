package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "andean/testutil/keeper"
	"andean/testutil/sample"
	"andean/x/xicoatl/keeper"
	"andean/x/xicoatl/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPoolMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.XicoatlKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := sample.AccAddress()
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreatePool{
			Creator: creator,
			Index:   strconv.Itoa(i),
			DenomA:  "stake",
			AmountA: 100,
			DenomB:  "token",
			AmountB: 100,
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
	creator := sample.AccAddress()

	tests := []struct {
		desc    string
		request *types.MsgUpdatePool
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdatePool{Creator: creator,
				Index:   strconv.Itoa(0),
				DenomA:  "stake",
				AmountA: 100,
				DenomB:  "token",
				AmountB: 100,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdatePool{Creator: sample.AccAddress(),
				Index:   strconv.Itoa(0),
				DenomA:  "stake",
				AmountA: 100,
				DenomB:  "token",
				AmountB: 100,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdatePool{Creator: creator,
				Index:   strconv.Itoa(100000),
				DenomA:  "stake",
				AmountA: 100,
				DenomB:  "token",
				AmountB: 100,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.XicoatlKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreatePool{
				Creator: creator,
				Index:   strconv.Itoa(0),
				DenomA:  "stake",
				AmountA: 100,
				DenomB:  "token",
				AmountB: 100,
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
	creator := sample.AccAddress()

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
			request: &types.MsgDeletePool{Creator: sample.AccAddress(),
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

			_, err := srv.CreatePool(ctx, &types.MsgCreatePool{
				Creator: creator,
				Index:   strconv.Itoa(0),
				DenomA:  "stake",
				AmountA: 100,
				DenomB:  "token",
				AmountB: 100,
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
