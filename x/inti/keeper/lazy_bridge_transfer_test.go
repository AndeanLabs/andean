package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "andean/testutil/keeper"
	"andean/testutil/nullify"
	"andean/x/inti/keeper"
	"andean/x/inti/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNLazyBridgeTransfer(keeper keeper.Keeper, ctx context.Context, n int) []types.LazyBridgeTransfer {
	items := make([]types.LazyBridgeTransfer, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetLazyBridgeTransfer(ctx, items[i])
	}
	return items
}

func TestLazyBridgeTransferGet(t *testing.T) {
	keeper, ctx := keepertest.IntiKeeper(t)
	items := createNLazyBridgeTransfer(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLazyBridgeTransfer(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLazyBridgeTransferRemove(t *testing.T) {
	keeper, ctx := keepertest.IntiKeeper(t)
	items := createNLazyBridgeTransfer(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLazyBridgeTransfer(ctx,
			item.Index,
		)
		_, found := keeper.GetLazyBridgeTransfer(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestLazyBridgeTransferGetAll(t *testing.T) {
	keeper, ctx := keepertest.IntiKeeper(t)
	items := createNLazyBridgeTransfer(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLazyBridgeTransfer(ctx)),
	)
}
