package types_test

import (
	"testing"

	"andean/x/itzel/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				PriceList: []types.Price{
					{
						Source: "0",
					},
					{
						Source: "1",
					},
				},
				PriceReportList: []types.PriceReport{
					{
						Source: "0",
						Oracle: "0",
					},
					{
						Source: "1",
						Oracle: "1",
					},
				},
				AggregatedPriceList: []types.AggregatedPrice{
					{
						Source: "0",
					},
					{
						Source: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated price",
			genState: &types.GenesisState{
				PriceList: []types.Price{
					{
						Source: "0",
					},
					{
						Source: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated priceReport",
			genState: &types.GenesisState{
				PriceReportList: []types.PriceReport{
					{
						Source: "0",
						Oracle: "0",
					},
					{
						Source: "0",
						Oracle: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated aggregatedPrice",
			genState: &types.GenesisState{
				AggregatedPriceList: []types.AggregatedPrice{
					{
						Source: "0",
					},
					{
						Source: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
