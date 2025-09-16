package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PriceList:           []Price{},
		PriceReportList:     []PriceReport{},
		AggregatedPriceList: []AggregatedPrice{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in price
	priceIndexMap := make(map[string]struct{})

	for _, elem := range gs.PriceList {
		index := string(PriceKey(elem.Source))
		if _, ok := priceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for price")
		}
		priceIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in priceReport
	priceReportIndexMap := make(map[string]struct{})

	for _, elem := range gs.PriceReportList {
		index := string(PriceReportKey(elem.Source, elem.Oracle))
		if _, ok := priceReportIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for priceReport")
		}
		priceReportIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in aggregatedPrice
	aggregatedPriceIndexMap := make(map[string]struct{})

	for _, elem := range gs.AggregatedPriceList {
		index := string(AggregatedPriceKey(elem.Source))
		if _, ok := aggregatedPriceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for aggregatedPrice")
		}
		aggregatedPriceIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
