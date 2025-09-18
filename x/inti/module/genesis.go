package inti

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"andean/x/inti/keeper"
	"andean/x/inti/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the lazyBridgeTransfer
	for _, elem := range genState.LazyBridgeTransferList {
		k.SetLazyBridgeTransfer(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.LazyBridgeTransferList = k.GetAllLazyBridgeTransfer(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
