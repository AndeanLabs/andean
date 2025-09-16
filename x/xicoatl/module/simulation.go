package xicoatl

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"andean/testutil/sample"
	xicoatlsimulation "andean/x/xicoatl/simulation"
	"andean/x/xicoatl/types"
)

// avoid unused import issue
var (
	_ = xicoatlsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreatePool = "op_weight_msg_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePool int = 100

	opWeightMsgUpdatePool = "op_weight_msg_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePool int = 100

	opWeightMsgDeletePool = "op_weight_msg_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePool int = 100

	opWeightMsgSwap = "op_weight_msg_swap"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSwap int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	xicoatlGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PoolList: []types.Pool{
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
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&xicoatlGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePool int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePool, &weightMsgCreatePool, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePool = defaultWeightMsgCreatePool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePool,
		xicoatlsimulation.SimulateMsgCreatePool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePool int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePool, &weightMsgUpdatePool, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePool = defaultWeightMsgUpdatePool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePool,
		xicoatlsimulation.SimulateMsgUpdatePool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePool int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePool, &weightMsgDeletePool, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePool = defaultWeightMsgDeletePool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePool,
		xicoatlsimulation.SimulateMsgDeletePool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSwap int
	simState.AppParams.GetOrGenerate(opWeightMsgSwap, &weightMsgSwap, nil,
		func(_ *rand.Rand) {
			weightMsgSwap = defaultWeightMsgSwap
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSwap,
		xicoatlsimulation.SimulateMsgSwap(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePool,
			defaultWeightMsgCreatePool,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				xicoatlsimulation.SimulateMsgCreatePool(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdatePool,
			defaultWeightMsgUpdatePool,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				xicoatlsimulation.SimulateMsgUpdatePool(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeletePool,
			defaultWeightMsgDeletePool,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				xicoatlsimulation.SimulateMsgDeletePool(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSwap,
			defaultWeightMsgSwap,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				xicoatlsimulation.SimulateMsgSwap(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
