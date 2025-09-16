package itzel

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"andean/x/itzel/keeper"
	"andean/x/itzel/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the price
	for _, elem := range genState.PriceList {
		k.SetPrice(ctx, elem)
	}
	// Set all the priceReport
	for _, elem := range genState.PriceReportList {
		k.SetPriceReport(ctx, elem)
	}
	// Set all the aggregatedPrice
	for _, elem := range genState.AggregatedPriceList {
		k.SetAggregatedPrice(ctx, elem)
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

	genesis.PriceList = k.GetAllPrice(ctx)
	genesis.PriceReportList = k.GetAllPriceReport(ctx)
	genesis.AggregatedPriceList = k.GetAllAggregatedPrice(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
