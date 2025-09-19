package identity

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"resist/x/identity/types"
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
				{
					RpcMethod: "ListUserProfile",
					Use:       "list-user-profile",
					Short:     "List all user-profile",
				},
				{
					RpcMethod:      "GetUserProfile",
					Use:            "get-user-profile [id]",
					Short:          "Gets a user-profile",
					Alias:          []string{"show-user-profile"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
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
					RpcMethod:      "RequestChallenge",
					Use:            "request-challenge [address]",
					Short:          "Send a request-challenge tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod:      "VerifySignature",
					Use:            "verify-signature [challenge] [signature] [address]",
					Short:          "Send a verify-signature tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "challenge"}, {ProtoField: "signature"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "CreateUserProfile",
					Use:            "create-user-profile [index] [display-name] [bio] [avatar-url] [verified] [created-at]",
					Short:          "Create a new user-profile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "display_name"}, {ProtoField: "bio"}, {ProtoField: "avatar_url"}, {ProtoField: "verified"}, {ProtoField: "created_at"}},
				},
				{
					RpcMethod:      "UpdateUserProfile",
					Use:            "update-user-profile [index] [display-name] [bio] [avatar-url] [verified] [created-at]",
					Short:          "Update user-profile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "display_name"}, {ProtoField: "bio"}, {ProtoField: "avatar_url"}, {ProtoField: "verified"}, {ProtoField: "created_at"}},
				},
				{
					RpcMethod:      "DeleteUserProfile",
					Use:            "delete-user-profile [index]",
					Short:          "Delete user-profile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
