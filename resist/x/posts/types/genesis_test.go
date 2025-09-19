package types_test

import (
	"testing"

	"resist/x/posts/types"

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
			genState: &types.GenesisState{SocialPostMap: []types.SocialPost{{Index: "0"}, {Index: "1"}}, VoteMap: []types.Vote{{Index: "0"}, {Index: "1"}}, SourceMap: []types.Source{{Index: "0"}, {Index: "1"}}, PostTagMap: []types.PostTag{{Index: "0"}, {Index: "1"}}},
			valid:    true,
		}, {
			desc: "duplicated socialPost",
			genState: &types.GenesisState{
				SocialPostMap: []types.SocialPost{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
				VoteMap: []types.Vote{{Index: "0"}, {Index: "1"}}, SourceMap: []types.Source{{Index: "0"}, {Index: "1"}}, PostTagMap: []types.PostTag{{Index: "0"}, {Index: "1"}}},
			valid: false,
		}, {
			desc: "duplicated vote",
			genState: &types.GenesisState{
				VoteMap: []types.Vote{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
				SourceMap: []types.Source{{Index: "0"}, {Index: "1"}}, PostTagMap: []types.PostTag{{Index: "0"}, {Index: "1"}}},
			valid: false,
		}, {
			desc: "duplicated source",
			genState: &types.GenesisState{
				SourceMap: []types.Source{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
				PostTagMap: []types.PostTag{{Index: "0"}, {Index: "1"}}},
			valid: false,
		}, {
			desc: "duplicated postTag",
			genState: &types.GenesisState{
				PostTagMap: []types.PostTag{
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
