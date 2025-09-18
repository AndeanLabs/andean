package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "andean/testutil/keeper"
	"andean/testutil/nullify"
	"andean/x/inti/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestLazyBridgeTransferQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.IntiKeeper(t)
	msgs := createNLazyBridgeTransfer(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetLazyBridgeTransferRequest
		response *types.QueryGetLazyBridgeTransferResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetLazyBridgeTransferRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetLazyBridgeTransferResponse{LazyBridgeTransfer: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetLazyBridgeTransferRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetLazyBridgeTransferResponse{LazyBridgeTransfer: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetLazyBridgeTransferRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.LazyBridgeTransfer(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestLazyBridgeTransferQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.IntiKeeper(t)
	msgs := createNLazyBridgeTransfer(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllLazyBridgeTransferRequest {
		return &types.QueryAllLazyBridgeTransferRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LazyBridgeTransferAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LazyBridgeTransfer), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LazyBridgeTransfer),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LazyBridgeTransferAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LazyBridgeTransfer), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LazyBridgeTransfer),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.LazyBridgeTransferAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.LazyBridgeTransfer),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.LazyBridgeTransferAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
