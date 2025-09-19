package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:         DefaultParams(),
		UserProfileMap: []UserProfile{}}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	userProfileIndexMap := make(map[string]struct{})

	for _, elem := range gs.UserProfileMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := userProfileIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for userProfile")
		}
		userProfileIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
