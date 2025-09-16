package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "andean/testutil/keeper"
	"andean/x/itzel/keeper"
	"andean/x/itzel/types"
)

func TestAggregatePrices(t *testing.T) {
	k, ctx, authorizedOracles := keepertest.ItzelKeeper(t, 3)
	srv := keeper.NewMsgServerImpl(k)
	source := "test-source"

	// Submit 3 price reports
	_, err := srv.SubmitPrice(ctx, &types.MsgSubmitPrice{Creator: authorizedOracles[0], Source: source, Price: 100})
	require.NoError(t, err)
	_, err = srv.SubmitPrice(ctx, &types.MsgSubmitPrice{Creator: authorizedOracles[1], Source: source, Price: 110})
	require.NoError(t, err)
	_, err = srv.SubmitPrice(ctx, &types.MsgSubmitPrice{Creator: authorizedOracles[2], Source: source, Price: 120})
	require.NoError(t, err)

	// Aggregate the prices
	_, err = srv.AggregatePrices(ctx, &types.MsgAggregatePrices{Creator: authorizedOracles[0], Source: source})
	require.NoError(t, err)

	// Check if the aggregated price is correct
	aggPrice, found := k.GetAggregatedPrice(ctx, source)
	require.True(t, found)
	require.Equal(t, int32(110), aggPrice.Price)
}
