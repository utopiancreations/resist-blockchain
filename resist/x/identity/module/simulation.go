package identity

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"resist/testutil/sample"
	identitysimulation "resist/x/identity/simulation"
	"resist/x/identity/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	identityGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		UserProfileMap: []types.UserProfile{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&identityGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgRequestChallenge          = "op_weight_msg_identity"
		defaultWeightMsgRequestChallenge int = 100
	)

	var weightMsgRequestChallenge int
	simState.AppParams.GetOrGenerate(opWeightMsgRequestChallenge, &weightMsgRequestChallenge, nil,
		func(_ *rand.Rand) {
			weightMsgRequestChallenge = defaultWeightMsgRequestChallenge
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestChallenge,
		identitysimulation.SimulateMsgRequestChallenge(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgVerifySignature          = "op_weight_msg_identity"
		defaultWeightMsgVerifySignature int = 100
	)

	var weightMsgVerifySignature int
	simState.AppParams.GetOrGenerate(opWeightMsgVerifySignature, &weightMsgVerifySignature, nil,
		func(_ *rand.Rand) {
			weightMsgVerifySignature = defaultWeightMsgVerifySignature
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVerifySignature,
		identitysimulation.SimulateMsgVerifySignature(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCreateUserProfile          = "op_weight_msg_identity"
		defaultWeightMsgCreateUserProfile int = 100
	)

	var weightMsgCreateUserProfile int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateUserProfile, &weightMsgCreateUserProfile, nil,
		func(_ *rand.Rand) {
			weightMsgCreateUserProfile = defaultWeightMsgCreateUserProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateUserProfile,
		identitysimulation.SimulateMsgCreateUserProfile(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateUserProfile          = "op_weight_msg_identity"
		defaultWeightMsgUpdateUserProfile int = 100
	)

	var weightMsgUpdateUserProfile int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateUserProfile, &weightMsgUpdateUserProfile, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateUserProfile = defaultWeightMsgUpdateUserProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateUserProfile,
		identitysimulation.SimulateMsgUpdateUserProfile(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteUserProfile          = "op_weight_msg_identity"
		defaultWeightMsgDeleteUserProfile int = 100
	)

	var weightMsgDeleteUserProfile int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteUserProfile, &weightMsgDeleteUserProfile, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteUserProfile = defaultWeightMsgDeleteUserProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteUserProfile,
		identitysimulation.SimulateMsgDeleteUserProfile(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
