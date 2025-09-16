package itzel

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "andean/api/andean/itzel"
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
					RpcMethod: "PriceAll",
					Use:       "list-price",
					Short:     "List all price",
				},
				{
					RpcMethod:      "Price",
					Use:            "show-price [id]",
					Short:          "Shows a price",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "source"}},
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
					RpcMethod:      "SubmitPrice",
					Use:            "submit-price [source] [price]",
					Short:          "Send a submit-price tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "source"}, {ProtoField: "price"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
