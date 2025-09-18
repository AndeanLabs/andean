package inti

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "andean/api/andean/inti"
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
					RpcMethod: "LazyBridgeTransferAll",
					Use:       "list-lazy-bridge-transfer",
					Short:     "List all lazyBridgeTransfer",
				},
				{
					RpcMethod:      "LazyBridgeTransfer",
					Use:            "show-lazy-bridge-transfer [id]",
					Short:          "Shows a lazyBridgeTransfer",
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
					RpcMethod:      "CreateLazyBridgeTransfer",
					Use:            "create-lazy-bridge-transfer [index] [sourceChain] [destChain] [amount] [status]",
					Short:          "Create a new lazyBridgeTransfer",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "sourceChain"}, {ProtoField: "destChain"}, {ProtoField: "amount"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "UpdateLazyBridgeTransfer",
					Use:            "update-lazy-bridge-transfer [index] [sourceChain] [destChain] [amount] [status]",
					Short:          "Update lazyBridgeTransfer",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "sourceChain"}, {ProtoField: "destChain"}, {ProtoField: "amount"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "DeleteLazyBridgeTransfer",
					Use:            "delete-lazy-bridge-transfer [index]",
					Short:          "Delete lazyBridgeTransfer",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "InitiateBridgeTransfer",
					Use:            "initiate-bridge-transfer [source-chain] [dest-chain] [amount]",
					Short:          "Send a initiate-bridge-transfer tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "sourceChain"}, {ProtoField: "destChain"}, {ProtoField: "amount"}},
				},
				{
					RpcMethod:      "ConfirmBridgeTransfer",
					Use:            "confirm-bridge-transfer [index] [final-tx-hash]",
					Short:          "Send a confirm-bridge-transfer tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "finalTxHash"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
