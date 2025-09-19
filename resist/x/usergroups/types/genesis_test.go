package types_test

import (
	"testing"

	"resist/x/usergroups/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc:     "valid genesis state",
			genState: &types.GenesisState{UserGroupMap: []types.UserGroup{{Index: "0"}, {Index: "1"}}, ContentReportMap: []types.ContentReport{{Index: "0"}, {Index: "1"}}, GovernanceProposalMap: []types.GovernanceProposal{{Index: "0"}, {Index: "1"}}},
			valid:    true,
		}, {
			desc: "duplicated userGroup",
			genState: &types.GenesisState{
				UserGroupMap: []types.UserGroup{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
				ContentReportMap: []types.ContentReport{{Index: "0"}, {Index: "1"}}, GovernanceProposalMap: []types.GovernanceProposal{{Index: "0"}, {Index: "1"}}},
			valid: false,
		}, {
			desc: "duplicated contentReport",
			genState: &types.GenesisState{
				ContentReportMap: []types.ContentReport{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
				GovernanceProposalMap: []types.GovernanceProposal{{Index: "0"}, {Index: "1"}}},
			valid: false,
		}, {
			desc: "duplicated governanceProposal",
			genState: &types.GenesisState{
				GovernanceProposalMap: []types.GovernanceProposal{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
