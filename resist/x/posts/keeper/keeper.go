package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	"resist/x/posts/types"
)

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec
	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority []byte

	Schema     collections.Schema
	Params     collections.Item[types.Params]
	SocialPost collections.Map[string, types.SocialPost]
	Vote       collections.Map[string, types.Vote]
	Source     collections.Map[string, types.Source]
	PostTag    collections.Map[string, types.PostTag]
}

func NewKeeper(
	storeService corestore.KVStoreService,
	cdc codec.Codec,
	addressCodec address.Codec,
	authority []byte,

) Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		storeService: storeService,
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,

		Params:     collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		SocialPost: collections.NewMap(sb, types.SocialPostKey, "socialPost", collections.StringKey, codec.CollValue[types.SocialPost](cdc)), Vote: collections.NewMap(sb, types.VoteKey, "vote", collections.StringKey, codec.CollValue[types.Vote](cdc)), Source: collections.NewMap(sb, types.SourceKey, "source", collections.StringKey, codec.CollValue[types.Source](cdc)), PostTag: collections.NewMap(sb, types.PostTagKey, "postTag", collections.StringKey, codec.CollValue[types.PostTag](cdc))}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() []byte {
	return k.authority
}
