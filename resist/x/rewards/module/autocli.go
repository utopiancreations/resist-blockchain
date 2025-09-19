package rewards

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"resist/x/rewards/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "RegisterNode",
					Use:            "register-node [node-id] [node-type] [stake-amount]",
					Short:          "Send a register-node tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "node_id"}, {ProtoField: "node_type"}, {ProtoField: "stake_amount"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
