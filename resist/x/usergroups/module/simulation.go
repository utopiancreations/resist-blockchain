package usergroups

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"resist/testutil/sample"
	usergroupssimulation "resist/x/usergroups/simulation"
	"resist/x/usergroups/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	usergroupsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		UserGroupMap: []types.UserGroup{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}, ContentReportMap: []types.ContentReport{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}, GovernanceProposalMap: []types.GovernanceProposal{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&usergroupsGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreateUserGroup          = "op_weight_msg_usergroups"
		defaultWeightMsgCreateUserGroup int = 100
	)

	var weightMsgCreateUserGroup int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateUserGroup, &weightMsgCreateUserGroup, nil,
		func(_ *rand.Rand) {
			weightMsgCreateUserGroup = defaultWeightMsgCreateUserGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateUserGroup,
		usergroupssimulation.SimulateMsgCreateUserGroup(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateUserGroup          = "op_weight_msg_usergroups"
		defaultWeightMsgUpdateUserGroup int = 100
	)

	var weightMsgUpdateUserGroup int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateUserGroup, &weightMsgUpdateUserGroup, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateUserGroup = defaultWeightMsgUpdateUserGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateUserGroup,
		usergroupssimulation.SimulateMsgUpdateUserGroup(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteUserGroup          = "op_weight_msg_usergroups"
		defaultWeightMsgDeleteUserGroup int = 100
	)

	var weightMsgDeleteUserGroup int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteUserGroup, &weightMsgDeleteUserGroup, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteUserGroup = defaultWeightMsgDeleteUserGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteUserGroup,
		usergroupssimulation.SimulateMsgDeleteUserGroup(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCreateContentReport          = "op_weight_msg_usergroups"
		defaultWeightMsgCreateContentReport int = 100
	)

	var weightMsgCreateContentReport int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateContentReport, &weightMsgCreateContentReport, nil,
		func(_ *rand.Rand) {
			weightMsgCreateContentReport = defaultWeightMsgCreateContentReport
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateContentReport,
		usergroupssimulation.SimulateMsgCreateContentReport(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateContentReport          = "op_weight_msg_usergroups"
		defaultWeightMsgUpdateContentReport int = 100
	)

	var weightMsgUpdateContentReport int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateContentReport, &weightMsgUpdateContentReport, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateContentReport = defaultWeightMsgUpdateContentReport
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateContentReport,
		usergroupssimulation.SimulateMsgUpdateContentReport(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteContentReport          = "op_weight_msg_usergroups"
		defaultWeightMsgDeleteContentReport int = 100
	)

	var weightMsgDeleteContentReport int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteContentReport, &weightMsgDeleteContentReport, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteContentReport = defaultWeightMsgDeleteContentReport
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteContentReport,
		usergroupssimulation.SimulateMsgDeleteContentReport(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCreateGovernanceProposal          = "op_weight_msg_usergroups"
		defaultWeightMsgCreateGovernanceProposal int = 100
	)

	var weightMsgCreateGovernanceProposal int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateGovernanceProposal, &weightMsgCreateGovernanceProposal, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGovernanceProposal = defaultWeightMsgCreateGovernanceProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateGovernanceProposal,
		usergroupssimulation.SimulateMsgCreateGovernanceProposal(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateGovernanceProposal          = "op_weight_msg_usergroups"
		defaultWeightMsgUpdateGovernanceProposal int = 100
	)

	var weightMsgUpdateGovernanceProposal int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateGovernanceProposal, &weightMsgUpdateGovernanceProposal, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGovernanceProposal = defaultWeightMsgUpdateGovernanceProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateGovernanceProposal,
		usergroupssimulation.SimulateMsgUpdateGovernanceProposal(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteGovernanceProposal          = "op_weight_msg_usergroups"
		defaultWeightMsgDeleteGovernanceProposal int = 100
	)

	var weightMsgDeleteGovernanceProposal int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteGovernanceProposal, &weightMsgDeleteGovernanceProposal, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteGovernanceProposal = defaultWeightMsgDeleteGovernanceProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteGovernanceProposal,
		usergroupssimulation.SimulateMsgDeleteGovernanceProposal(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
