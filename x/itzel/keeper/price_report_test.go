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

func createNPriceReport(keeper keeper.Keeper, ctx context.Context, n int) []types.PriceReport {
	items := make([]types.PriceReport, n)
	for i := range items {
		items[i].Source = strconv.Itoa(i)
		items[i].Oracle = strconv.Itoa(i)

		keeper.SetPriceReport(ctx, items[i])
	}
	return items
}

func TestPriceReportGet(t *testing.T) {
	keeper, ctx, _ := keepertest.ItzelKeeper(t, 1)
	items := createNPriceReport(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPriceReport(ctx,
			item.Source,
			item.Oracle,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPriceReportRemove(t *testing.T) {
	keeper, ctx, _ := keepertest.ItzelKeeper(t, 1)
	items := createNPriceReport(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePriceReport(ctx,
			item.Source,
			item.Oracle,
		)
		_, found := keeper.GetPriceReport(ctx,
			item.Source,
			item.Oracle,
		)
		require.False(t, found)
	}
}

func TestPriceReportGetAll(t *testing.T) {
	keeper, ctx, _ := keepertest.ItzelKeeper(t, 1)
	items := createNPriceReport(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPriceReport(ctx)),
	)
}
