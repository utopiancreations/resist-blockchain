package usergroups

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"resist/x/usergroups/types"
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
					RpcMethod: "ListUserGroup",
					Use:       "list-user-group",
					Short:     "List all user-group",
				},
				{
					RpcMethod:      "GetUserGroup",
					Use:            "get-user-group [id]",
					Short:          "Gets a user-group",
					Alias:          []string{"show-user-group"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ListContentReport",
					Use:       "list-content-report",
					Short:     "List all content-report",
				},
				{
					RpcMethod:      "GetContentReport",
					Use:            "get-content-report [id]",
					Short:          "Gets a content-report",
					Alias:          []string{"show-content-report"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ListGovernanceProposal",
					Use:       "list-governance-proposal",
					Short:     "List all governance-proposal",
				},
				{
					RpcMethod:      "GetGovernanceProposal",
					Use:            "get-governance-proposal [id]",
					Short:          "Gets a governance-proposal",
					Alias:          []string{"show-governance-proposal"},
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
					RpcMethod:      "CreateUserGroup",
					Use:            "create-user-group [index] [name] [description] [admin] [members] [vote-threshold] [created-at]",
					Short:          "Create a new user-group",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "name"}, {ProtoField: "description"}, {ProtoField: "admin"}, {ProtoField: "members"}, {ProtoField: "vote_threshold"}, {ProtoField: "created_at"}},
				},
				{
					RpcMethod:      "UpdateUserGroup",
					Use:            "update-user-group [index] [name] [description] [admin] [members] [vote-threshold] [created-at]",
					Short:          "Update user-group",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "name"}, {ProtoField: "description"}, {ProtoField: "admin"}, {ProtoField: "members"}, {ProtoField: "vote_threshold"}, {ProtoField: "created_at"}},
				},
				{
					RpcMethod:      "DeleteUserGroup",
					Use:            "delete-user-group [index]",
					Short:          "Delete user-group",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateContentReport",
					Use:            "create-content-report [index] [post-id] [reporter] [reason] [evidence] [status] [community-response] [resolution]",
					Short:          "Create a new content-report",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "post_id"}, {ProtoField: "reporter"}, {ProtoField: "reason"}, {ProtoField: "evidence"}, {ProtoField: "status"}, {ProtoField: "community_response"}, {ProtoField: "resolution"}},
				},
				{
					RpcMethod:      "UpdateContentReport",
					Use:            "update-content-report [index] [post-id] [reporter] [reason] [evidence] [status] [community-response] [resolution]",
					Short:          "Update content-report",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "post_id"}, {ProtoField: "reporter"}, {ProtoField: "reason"}, {ProtoField: "evidence"}, {ProtoField: "status"}, {ProtoField: "community_response"}, {ProtoField: "resolution"}},
				},
				{
					RpcMethod:      "DeleteContentReport",
					Use:            "delete-content-report [index]",
					Short:          "Delete content-report",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateGovernanceProposal",
					Use:            "create-governance-proposal [index] [title] [description] [proposer] [proposal-type] [voting-period-start] [voting-period-end] [yes-votes] [no-votes] [abstain-votes] [status]",
					Short:          "Create a new governance-proposal",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "proposer"}, {ProtoField: "proposal_type"}, {ProtoField: "voting_period_start"}, {ProtoField: "voting_period_end"}, {ProtoField: "yes_votes"}, {ProtoField: "no_votes"}, {ProtoField: "abstain_votes"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "UpdateGovernanceProposal",
					Use:            "update-governance-proposal [index] [title] [description] [proposer] [proposal-type] [voting-period-start] [voting-period-end] [yes-votes] [no-votes] [abstain-votes] [status]",
					Short:          "Update governance-proposal",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "proposer"}, {ProtoField: "proposal_type"}, {ProtoField: "voting_period_start"}, {ProtoField: "voting_period_end"}, {ProtoField: "yes_votes"}, {ProtoField: "no_votes"}, {ProtoField: "abstain_votes"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "DeleteGovernanceProposal",
					Use:            "delete-governance-proposal [index]",
					Short:          "Delete governance-proposal",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
