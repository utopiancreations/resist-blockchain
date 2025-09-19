package posts

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"resist/testutil/sample"
	postssimulation "resist/x/posts/simulation"
	"resist/x/posts/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	postsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		SocialPostMap: []types.SocialPost{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}, VoteMap: []types.Vote{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}, SourceMap: []types.Source{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}, PostTagMap: []types.PostTag{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&postsGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreatePost          = "op_weight_msg_posts"
		defaultWeightMsgCreatePost int = 100
	)

	var weightMsgCreatePost int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePost, &weightMsgCreatePost, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePost = defaultWeightMsgCreatePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePost,
		postssimulation.SimulateMsgCreatePost(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgVotePost          = "op_weight_msg_posts"
		defaultWeightMsgVotePost int = 100
	)

	var weightMsgVotePost int
	simState.AppParams.GetOrGenerate(opWeightMsgVotePost, &weightMsgVotePost, nil,
		func(_ *rand.Rand) {
			weightMsgVotePost = defaultWeightMsgVotePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVotePost,
		postssimulation.SimulateMsgVotePost(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCreateSocialPost          = "op_weight_msg_posts"
		defaultWeightMsgCreateSocialPost int = 100
	)

	var weightMsgCreateSocialPost int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateSocialPost, &weightMsgCreateSocialPost, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSocialPost = defaultWeightMsgCreateSocialPost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSocialPost,
		postssimulation.SimulateMsgCreateSocialPost(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateSocialPost          = "op_weight_msg_posts"
		defaultWeightMsgUpdateSocialPost int = 100
	)

	var weightMsgUpdateSocialPost int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateSocialPost, &weightMsgUpdateSocialPost, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSocialPost = defaultWeightMsgUpdateSocialPost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSocialPost,
		postssimulation.SimulateMsgUpdateSocialPost(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteSocialPost          = "op_weight_msg_posts"
		defaultWeightMsgDeleteSocialPost int = 100
	)

	var weightMsgDeleteSocialPost int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteSocialPost, &weightMsgDeleteSocialPost, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSocialPost = defaultWeightMsgDeleteSocialPost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSocialPost,
		postssimulation.SimulateMsgDeleteSocialPost(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCreateVote          = "op_weight_msg_posts"
		defaultWeightMsgCreateVote int = 100
	)

	var weightMsgCreateVote int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateVote, &weightMsgCreateVote, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVote = defaultWeightMsgCreateVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVote,
		postssimulation.SimulateMsgCreateVote(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateVote          = "op_weight_msg_posts"
		defaultWeightMsgUpdateVote int = 100
	)

	var weightMsgUpdateVote int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateVote, &weightMsgUpdateVote, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateVote = defaultWeightMsgUpdateVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateVote,
		postssimulation.SimulateMsgUpdateVote(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteVote          = "op_weight_msg_posts"
		defaultWeightMsgDeleteVote int = 100
	)

	var weightMsgDeleteVote int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteVote, &weightMsgDeleteVote, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteVote = defaultWeightMsgDeleteVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteVote,
		postssimulation.SimulateMsgDeleteVote(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCreateSource          = "op_weight_msg_posts"
		defaultWeightMsgCreateSource int = 100
	)

	var weightMsgCreateSource int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateSource, &weightMsgCreateSource, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSource = defaultWeightMsgCreateSource
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSource,
		postssimulation.SimulateMsgCreateSource(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateSource          = "op_weight_msg_posts"
		defaultWeightMsgUpdateSource int = 100
	)

	var weightMsgUpdateSource int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateSource, &weightMsgUpdateSource, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSource = defaultWeightMsgUpdateSource
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSource,
		postssimulation.SimulateMsgUpdateSource(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteSource          = "op_weight_msg_posts"
		defaultWeightMsgDeleteSource int = 100
	)

	var weightMsgDeleteSource int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteSource, &weightMsgDeleteSource, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSource = defaultWeightMsgDeleteSource
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSource,
		postssimulation.SimulateMsgDeleteSource(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCreatePostTag          = "op_weight_msg_posts"
		defaultWeightMsgCreatePostTag int = 100
	)

	var weightMsgCreatePostTag int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePostTag, &weightMsgCreatePostTag, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePostTag = defaultWeightMsgCreatePostTag
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePostTag,
		postssimulation.SimulateMsgCreatePostTag(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdatePostTag          = "op_weight_msg_posts"
		defaultWeightMsgUpdatePostTag int = 100
	)

	var weightMsgUpdatePostTag int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePostTag, &weightMsgUpdatePostTag, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePostTag = defaultWeightMsgUpdatePostTag
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePostTag,
		postssimulation.SimulateMsgUpdatePostTag(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeletePostTag          = "op_weight_msg_posts"
		defaultWeightMsgDeletePostTag int = 100
	)

	var weightMsgDeletePostTag int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePostTag, &weightMsgDeletePostTag, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePostTag = defaultWeightMsgDeletePostTag
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePostTag,
		postssimulation.SimulateMsgDeletePostTag(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
