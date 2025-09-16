package itzel

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"andean/testutil/sample"
	itzelsimulation "andean/x/itzel/simulation"
	"andean/x/itzel/types"
)

// avoid unused import issue
var (
	_ = itzelsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgSubmitPrice = "op_weight_msg_submit_price"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitPrice int = 100

	opWeightMsgAggregatePrices = "op_weight_msg_aggregate_prices"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAggregatePrices int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	itzelGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&itzelGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSubmitPrice int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitPrice, &weightMsgSubmitPrice, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitPrice = defaultWeightMsgSubmitPrice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitPrice,
		itzelsimulation.SimulateMsgSubmitPrice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAggregatePrices int
	simState.AppParams.GetOrGenerate(opWeightMsgAggregatePrices, &weightMsgAggregatePrices, nil,
		func(_ *rand.Rand) {
			weightMsgAggregatePrices = defaultWeightMsgAggregatePrices
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAggregatePrices,
		itzelsimulation.SimulateMsgAggregatePrices(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitPrice,
			defaultWeightMsgSubmitPrice,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				itzelsimulation.SimulateMsgSubmitPrice(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAggregatePrices,
			defaultWeightMsgAggregatePrices,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				itzelsimulation.SimulateMsgAggregatePrices(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
