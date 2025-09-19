package keeper

import (
	"context"

	"resist/x/usergroups/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	for _, elem := range genState.UserGroupMap {
		if err := k.UserGroup.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}
	for _, elem := range genState.ContentReportMap {
		if err := k.ContentReport.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}
	for _, elem := range genState.GovernanceProposalMap {
		if err := k.GovernanceProposal.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}

	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	if err := k.UserGroup.Walk(ctx, nil, func(_ string, val types.UserGroup) (stop bool, err error) {
		genesis.UserGroupMap = append(genesis.UserGroupMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	if err := k.ContentReport.Walk(ctx, nil, func(_ string, val types.ContentReport) (stop bool, err error) {
		genesis.ContentReportMap = append(genesis.ContentReportMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	if err := k.GovernanceProposal.Walk(ctx, nil, func(_ string, val types.GovernanceProposal) (stop bool, err error) {
		genesis.GovernanceProposalMap = append(genesis.GovernanceProposalMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	return genesis, nil
}
