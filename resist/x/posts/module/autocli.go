package posts

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"resist/x/posts/types"
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
					RpcMethod: "ListSocialPost",
					Use:       "list-social-post",
					Short:     "List all social-post",
				},
				{
					RpcMethod:      "GetSocialPost",
					Use:            "get-social-post [id]",
					Short:          "Gets a social-post",
					Alias:          []string{"show-social-post"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ListVote",
					Use:       "list-vote",
					Short:     "List all vote",
				},
				{
					RpcMethod:      "GetVote",
					Use:            "get-vote [id]",
					Short:          "Gets a vote",
					Alias:          []string{"show-vote"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ListSource",
					Use:       "list-source",
					Short:     "List all source",
				},
				{
					RpcMethod:      "GetSource",
					Use:            "get-source [id]",
					Short:          "Gets a source",
					Alias:          []string{"show-source"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ListPostTag",
					Use:       "list-post-tag",
					Short:     "List all post-tag",
				},
				{
					RpcMethod:      "GetPostTag",
					Use:            "get-post-tag [id]",
					Short:          "Gets a post-tag",
					Alias:          []string{"show-post-tag"},
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
					RpcMethod:      "CreatePost",
					Use:            "create-post [title] [content] [media-url] [media-type] [group-id]",
					Short:          "Send a create-post tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "content"}, {ProtoField: "media_url"}, {ProtoField: "media_type"}, {ProtoField: "group_id"}},
				},
				{
					RpcMethod:      "VotePost",
					Use:            "vote-post [post-id] [vote-type]",
					Short:          "Send a vote-post tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "post_id"}, {ProtoField: "vote_type"}},
				},
				{
					RpcMethod:      "CreateSocialPost",
					Use:            "create-social-post [index] [title] [content] [media-url] [media-type] [group-id] [author] [upvotes] [downvotes] [created-at]",
					Short:          "Create a new social-post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "title"}, {ProtoField: "content"}, {ProtoField: "media_url"}, {ProtoField: "media_type"}, {ProtoField: "group_id"}, {ProtoField: "author"}, {ProtoField: "upvotes"}, {ProtoField: "downvotes"}, {ProtoField: "created_at"}},
				},
				{
					RpcMethod:      "UpdateSocialPost",
					Use:            "update-social-post [index] [title] [content] [media-url] [media-type] [group-id] [author] [upvotes] [downvotes] [created-at]",
					Short:          "Update social-post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "title"}, {ProtoField: "content"}, {ProtoField: "media_url"}, {ProtoField: "media_type"}, {ProtoField: "group_id"}, {ProtoField: "author"}, {ProtoField: "upvotes"}, {ProtoField: "downvotes"}, {ProtoField: "created_at"}},
				},
				{
					RpcMethod:      "DeleteSocialPost",
					Use:            "delete-social-post [index]",
					Short:          "Delete social-post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateVote",
					Use:            "create-vote [index] [voter-address] [post-id] [vote-type] [timestamp]",
					Short:          "Create a new vote",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "voter_address"}, {ProtoField: "post_id"}, {ProtoField: "vote_type"}, {ProtoField: "timestamp"}},
				},
				{
					RpcMethod:      "UpdateVote",
					Use:            "update-vote [index] [voter-address] [post-id] [vote-type] [timestamp]",
					Short:          "Update vote",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "voter_address"}, {ProtoField: "post_id"}, {ProtoField: "vote_type"}, {ProtoField: "timestamp"}},
				},
				{
					RpcMethod:      "DeleteVote",
					Use:            "delete-vote [index]",
					Short:          "Delete vote",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateSource",
					Use:            "create-source [index] [url] [title] [description] [credibility-score] [analysis-summary] [verified]",
					Short:          "Create a new source",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "url"}, {ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "credibility_score"}, {ProtoField: "analysis_summary"}, {ProtoField: "verified"}},
				},
				{
					RpcMethod:      "UpdateSource",
					Use:            "update-source [index] [url] [title] [description] [credibility-score] [analysis-summary] [verified]",
					Short:          "Update source",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "url"}, {ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "credibility_score"}, {ProtoField: "analysis_summary"}, {ProtoField: "verified"}},
				},
				{
					RpcMethod:      "DeleteSource",
					Use:            "delete-source [index]",
					Short:          "Delete source",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreatePostTag",
					Use:            "create-post-tag [index] [post-id] [tag] [category] [similarity-score] [related-posts]",
					Short:          "Create a new post-tag",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "post_id"}, {ProtoField: "tag"}, {ProtoField: "category"}, {ProtoField: "similarity_score"}, {ProtoField: "related_posts"}},
				},
				{
					RpcMethod:      "UpdatePostTag",
					Use:            "update-post-tag [index] [post-id] [tag] [category] [similarity-score] [related-posts]",
					Short:          "Update post-tag",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "post_id"}, {ProtoField: "tag"}, {ProtoField: "category"}, {ProtoField: "similarity_score"}, {ProtoField: "related_posts"}},
				},
				{
					RpcMethod:      "DeletePostTag",
					Use:            "delete-post-tag [index]",
					Short:          "Delete post-tag",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
