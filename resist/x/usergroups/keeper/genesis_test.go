package keeper_test

import (
	"testing"

	"resist/x/usergroups/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:       types.DefaultParams(),
		UserGroupMap: []types.UserGroup{{Index: "0"}, {Index: "1"}}, ContentReportMap: []types.ContentReport{{Index: "0"}, {Index: "1"}}, GovernanceProposalMap: []types.GovernanceProposal{{Index: "0"}, {Index: "1"}}}

	f := initFixture(t)
	err := f.keeper.InitGenesis(f.ctx, genesisState)
	require.NoError(t, err)
	got, err := f.keeper.ExportGenesis(f.ctx)
	require.NoError(t, err)
	require.NotNil(t, got)

	require.EqualExportedValues(t, genesisState.Params, got.Params)
	require.EqualExportedValues(t, genesisState.UserGroupMap, got.UserGroupMap)
	require.EqualExportedValues(t, genesisState.ContentReportMap, got.ContentReportMap)
	require.EqualExportedValues(t, genesisState.GovernanceProposalMap, got.GovernanceProposalMap)

}
