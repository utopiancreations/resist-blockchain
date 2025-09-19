package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:        DefaultParams(),
		SocialPostMap: []SocialPost{}, VoteMap: []Vote{}, SourceMap: []Source{}, PostTagMap: []PostTag{}}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	socialPostIndexMap := make(map[string]struct{})

	for _, elem := range gs.SocialPostMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := socialPostIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for socialPost")
		}
		socialPostIndexMap[index] = struct{}{}
	}
	voteIndexMap := make(map[string]struct{})

	for _, elem := range gs.VoteMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := voteIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for vote")
		}
		voteIndexMap[index] = struct{}{}
	}
	sourceIndexMap := make(map[string]struct{})

	for _, elem := range gs.SourceMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := sourceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for source")
		}
		sourceIndexMap[index] = struct{}{}
	}
	postTagIndexMap := make(map[string]struct{})

	for _, elem := range gs.PostTagMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := postTagIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for postTag")
		}
		postTagIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
