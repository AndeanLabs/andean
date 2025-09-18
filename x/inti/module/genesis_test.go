package inti_test

import (
	"testing"

	keepertest "andean/testutil/keeper"
	"andean/testutil/nullify"
	inti "andean/x/inti/module"
	"andean/x/inti/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LazyBridgeTransferList: []types.LazyBridgeTransfer{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IntiKeeper(t)
	inti.InitGenesis(ctx, k, genesisState)
	got := inti.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LazyBridgeTransferList, got.LazyBridgeTransferList)
	// this line is used by starport scaffolding # genesis/test/assert
}
