package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:       DefaultParams(),
		UserGroupMap: []UserGroup{}, ContentReportMap: []ContentReport{}, GovernanceProposalMap: []GovernanceProposal{}}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	userGroupIndexMap := make(map[string]struct{})

	for _, elem := range gs.UserGroupMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := userGroupIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for userGroup")
		}
		userGroupIndexMap[index] = struct{}{}
	}
	contentReportIndexMap := make(map[string]struct{})

	for _, elem := range gs.ContentReportMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := contentReportIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for contentReport")
		}
		contentReportIndexMap[index] = struct{}{}
	}
	governanceProposalIndexMap := make(map[string]struct{})

	for _, elem := range gs.GovernanceProposalMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := governanceProposalIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for governanceProposal")
		}
		governanceProposalIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
