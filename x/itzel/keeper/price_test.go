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

func createNPrice(keeper keeper.Keeper, ctx context.Context, n int) []types.Price {
	items := make([]types.Price, n)
	for i := range items {
		items[i].Source = strconv.Itoa(i)

		keeper.SetPrice(ctx, items[i])
	}
	return items
}

func TestPriceGet(t *testing.T) {
	keeper, ctx, _ := keepertest.ItzelKeeper(t)
	items := createNPrice(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPrice(ctx,
			item.Source,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPriceRemove(t *testing.T) {
	keeper, ctx, _ := keepertest.ItzelKeeper(t)
	items := createNPrice(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePrice(ctx,
			item.Source,
		)
		_, found := keeper.GetPrice(ctx,
			item.Source,
		)
		require.False(t, found)
	}
}

func TestPriceGetAll(t *testing.T) {
	keeper, ctx, _ := keepertest.ItzelKeeper(t)
	items := createNPrice(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPrice(ctx)),
	)
}
