package xicoatl

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "andean/api/andean/xicoatl"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "PoolAll",
					Use:       "list-pool",
					Short:     "List all pool",
				},
				{
					RpcMethod:      "Pool",
					Use:            "show-pool [id]",
					Short:          "Shows a pool",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreatePool",
					Use:            "create-pool [index] [denomA] [amountA] [denomB] [amountB]",
					Short:          "Create a new pool",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "denomA"}, {ProtoField: "amountA"}, {ProtoField: "denomB"}, {ProtoField: "amountB"}},
				},
				{
					RpcMethod:      "UpdatePool",
					Use:            "update-pool [index] [denomA] [amountA] [denomB] [amountB]",
					Short:          "Update pool",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "denomA"}, {ProtoField: "amountA"}, {ProtoField: "denomB"}, {ProtoField: "amountB"}},
				},
				{
					RpcMethod:      "DeletePool",
					Use:            "delete-pool [index]",
					Short:          "Delete pool",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "Swap",
					Use:            "swap [pool-id] [token-in-denom] [token-in-amount] [token-out-denom] [min-token-out-amount]",
					Short:          "Send a swap tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "poolId"}, {ProtoField: "tokenInDenom"}, {ProtoField: "tokenInAmount"}, {ProtoField: "tokenOutDenom"}, {ProtoField: "minTokenOutAmount"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
