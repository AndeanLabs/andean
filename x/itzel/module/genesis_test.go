package itzel_test

import (
	"testing"

	keepertest "andean/testutil/keeper"
	"andean/testutil/nullify"
	itzel "andean/x/itzel/module"
	"andean/x/itzel/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PriceList: []types.Price{
			{
				Source: "0",
			},
			{
				Source: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.ItzelKeeper(t)
	itzel.InitGenesis(ctx, k, genesisState)
	got := itzel.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PriceList, got.PriceList)
	// this line is used by starport scaffolding # genesis/test/assert
}
