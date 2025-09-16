package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "andean/testutil/keeper"
	"andean/x/itzel/types"
)

func TestGetParams(t *testing.T) {
	k, ctx, _ := keepertest.ItzelKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
