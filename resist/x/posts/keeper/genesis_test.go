package keeper_test

import (
	"testing"

	"resist/x/posts/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:        types.DefaultParams(),
		SocialPostMap: []types.SocialPost{{Index: "0"}, {Index: "1"}}, VoteMap: []types.Vote{{Index: "0"}, {Index: "1"}}, SourceMap: []types.Source{{Index: "0"}, {Index: "1"}}, PostTagMap: []types.PostTag{{Index: "0"}, {Index: "1"}}}

	f := initFixture(t)
	err := f.keeper.InitGenesis(f.ctx, genesisState)
	require.NoError(t, err)
	got, err := f.keeper.ExportGenesis(f.ctx)
	require.NoError(t, err)
	require.NotNil(t, got)

	require.EqualExportedValues(t, genesisState.Params, got.Params)
	require.EqualExportedValues(t, genesisState.SocialPostMap, got.SocialPostMap)
	require.EqualExportedValues(t, genesisState.VoteMap, got.VoteMap)
	require.EqualExportedValues(t, genesisState.SourceMap, got.SourceMap)
	require.EqualExportedValues(t, genesisState.PostTagMap, got.PostTagMap)

}
