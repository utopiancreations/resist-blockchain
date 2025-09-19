package keeper

import (
	"context"

	"resist/x/identity/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	for _, elem := range genState.UserProfileMap {
		if err := k.UserProfile.Set(ctx, elem.Index, elem); err != nil {
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
	if err := k.UserProfile.Walk(ctx, nil, func(_ string, val types.UserProfile) (stop bool, err error) {
		genesis.UserProfileMap = append(genesis.UserProfileMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	return genesis, nil
}
