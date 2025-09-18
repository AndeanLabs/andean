package inti

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"andean/testutil/sample"
	intisimulation "andean/x/inti/simulation"
	"andean/x/inti/types"
)

// avoid unused import issue
var (
	_ = intisimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateLazyBridgeTransfer = "op_weight_msg_lazy_bridge_transfer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateLazyBridgeTransfer int = 100

	opWeightMsgUpdateLazyBridgeTransfer = "op_weight_msg_lazy_bridge_transfer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateLazyBridgeTransfer int = 100

	opWeightMsgDeleteLazyBridgeTransfer = "op_weight_msg_lazy_bridge_transfer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteLazyBridgeTransfer int = 100

	opWeightMsgInitiateBridgeTransfer = "op_weight_msg_initiate_bridge_transfer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitiateBridgeTransfer int = 100

	opWeightMsgConfirmBridgeTransfer = "op_weight_msg_confirm_bridge_transfer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgConfirmBridgeTransfer int = 100

	opWeightMsgCreateLazyTransfer = "op_weight_msg_create_lazy_transfer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateLazyTransfer int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	intiGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		LazyBridgeTransferList: []types.LazyBridgeTransfer{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&intiGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateLazyBridgeTransfer int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateLazyBridgeTransfer, &weightMsgCreateLazyBridgeTransfer, nil,
		func(_ *rand.Rand) {
			weightMsgCreateLazyBridgeTransfer = defaultWeightMsgCreateLazyBridgeTransfer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateLazyBridgeTransfer,
		intisimulation.SimulateMsgCreateLazyBridgeTransfer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateLazyBridgeTransfer int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateLazyBridgeTransfer, &weightMsgUpdateLazyBridgeTransfer, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateLazyBridgeTransfer = defaultWeightMsgUpdateLazyBridgeTransfer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateLazyBridgeTransfer,
		intisimulation.SimulateMsgUpdateLazyBridgeTransfer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteLazyBridgeTransfer int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteLazyBridgeTransfer, &weightMsgDeleteLazyBridgeTransfer, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteLazyBridgeTransfer = defaultWeightMsgDeleteLazyBridgeTransfer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteLazyBridgeTransfer,
		intisimulation.SimulateMsgDeleteLazyBridgeTransfer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgInitiateBridgeTransfer int
	simState.AppParams.GetOrGenerate(opWeightMsgInitiateBridgeTransfer, &weightMsgInitiateBridgeTransfer, nil,
		func(_ *rand.Rand) {
			weightMsgInitiateBridgeTransfer = defaultWeightMsgInitiateBridgeTransfer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInitiateBridgeTransfer,
		intisimulation.SimulateMsgInitiateBridgeTransfer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgConfirmBridgeTransfer int
	simState.AppParams.GetOrGenerate(opWeightMsgConfirmBridgeTransfer, &weightMsgConfirmBridgeTransfer, nil,
		func(_ *rand.Rand) {
			weightMsgConfirmBridgeTransfer = defaultWeightMsgConfirmBridgeTransfer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgConfirmBridgeTransfer,
		intisimulation.SimulateMsgConfirmBridgeTransfer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateLazyTransfer int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateLazyTransfer, &weightMsgCreateLazyTransfer, nil,
		func(_ *rand.Rand) {
			weightMsgCreateLazyTransfer = defaultWeightMsgCreateLazyTransfer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateLazyTransfer,
		intisimulation.SimulateMsgCreateLazyTransfer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateLazyBridgeTransfer,
			defaultWeightMsgCreateLazyBridgeTransfer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				intisimulation.SimulateMsgCreateLazyBridgeTransfer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateLazyBridgeTransfer,
			defaultWeightMsgUpdateLazyBridgeTransfer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				intisimulation.SimulateMsgUpdateLazyBridgeTransfer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteLazyBridgeTransfer,
			defaultWeightMsgDeleteLazyBridgeTransfer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				intisimulation.SimulateMsgDeleteLazyBridgeTransfer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgInitiateBridgeTransfer,
			defaultWeightMsgInitiateBridgeTransfer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				intisimulation.SimulateMsgInitiateBridgeTransfer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgConfirmBridgeTransfer,
			defaultWeightMsgConfirmBridgeTransfer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				intisimulation.SimulateMsgConfirmBridgeTransfer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateLazyTransfer,
			defaultWeightMsgCreateLazyTransfer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				intisimulation.SimulateMsgCreateLazyTransfer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
