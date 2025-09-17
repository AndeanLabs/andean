package xicoatl_test

import (
	"testing"

	keepertest "andean/testutil/keeper"
	"andean/testutil/nullify"
	xicoatl "andean/x/xicoatl/module"
	"andean/x/xicoatl/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PoolList: []types.Pool{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _, _, _, _ := keepertest.XicoatlKeeper(t)
	xicoatl.InitGenesis(ctx, k, genesisState)
	got := xicoatl.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PoolList, got.PoolList)
	// this line is used by starport scaffolding # genesis/test/assert
}
