package keeper

import (
	"context"

	"resist/x/posts/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	for _, elem := range genState.SocialPostMap {
		if err := k.SocialPost.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}
	for _, elem := range genState.VoteMap {
		if err := k.Vote.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}
	for _, elem := range genState.SourceMap {
		if err := k.Source.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}
	for _, elem := range genState.PostTagMap {
		if err := k.PostTag.Set(ctx, elem.Index, elem); err != nil {
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
	if err := k.SocialPost.Walk(ctx, nil, func(_ string, val types.SocialPost) (stop bool, err error) {
		genesis.SocialPostMap = append(genesis.SocialPostMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	if err := k.Vote.Walk(ctx, nil, func(_ string, val types.Vote) (stop bool, err error) {
		genesis.VoteMap = append(genesis.VoteMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	if err := k.Source.Walk(ctx, nil, func(_ string, val types.Source) (stop bool, err error) {
		genesis.SourceMap = append(genesis.SourceMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	if err := k.PostTag.Walk(ctx, nil, func(_ string, val types.PostTag) (stop bool, err error) {
		genesis.PostTagMap = append(genesis.PostTagMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	return genesis, nil
}
