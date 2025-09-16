package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "andean/testutil/keeper"
	"andean/testutil/nullify"
	"andean/x/itzel/keeper"
	"andean/x/itzel/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAggregatedPrice(keeper keeper.Keeper, ctx context.Context, n int) []types.AggregatedPrice {
	items := make([]types.AggregatedPrice, n)
	for i := range items {
		items[i].Source = strconv.Itoa(i)

		keeper.SetAggregatedPrice(ctx, items[i])
	}
	return items
}

func TestAggregatedPriceGet(t *testing.T) {
	keeper, ctx, _ := keepertest.ItzelKeeper(t, 1)
	items := createNAggregatedPrice(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAggregatedPrice(ctx,
			item.Source,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAggregatedPriceRemove(t *testing.T) {
	keeper, ctx, _ := keepertest.ItzelKeeper(t, 1)
	items := createNAggregatedPrice(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAggregatedPrice(ctx,
			item.Source,
		)
		_, found := keeper.GetAggregatedPrice(ctx,
			item.Source,
		)
		require.False(t, found)
	}
}

func TestAggregatedPriceGetAll(t *testing.T) {
	keeper, ctx, _ := keepertest.ItzelKeeper(t, 1)
	items := createNAggregatedPrice(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAggregatedPrice(ctx)),
	)
}
