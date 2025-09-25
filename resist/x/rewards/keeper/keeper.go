package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	"resist/x/rewards/types"
)

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec
	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority []byte

	Schema              collections.Schema
	Params              collections.Item[types.Params]
	Nodes               collections.Map[string, types.Node]
	ResourceOffers      collections.Map[string, types.ResourceOffer]
	ResourceAllocations collections.Map[string, types.ResourceAllocation]
	HubMetrics          collections.Map[string, types.HubMetrics]
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

		Params:              collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		Nodes:               collections.NewMap(sb, types.NodesKey, "nodes", collections.StringKey, codec.CollValue[types.Node](cdc)),
		ResourceOffers:      collections.NewMap(sb, types.ResourceOffersKey, "resource_offers", collections.StringKey, codec.CollValue[types.ResourceOffer](cdc)),
		ResourceAllocations: collections.NewMap(sb, types.ResourceAllocationsKey, "resource_allocations", collections.StringKey, codec.CollValue[types.ResourceAllocation](cdc)),
		HubMetrics:          collections.NewMap(sb, types.HubMetricsKey, "hub_metrics", collections.StringKey, codec.CollValue[types.HubMetrics](cdc)),
	}

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
