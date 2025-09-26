package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	"resist/x/usergroups/types"
)

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec
	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority []byte

	Schema collections.Schema
	Params collections.Item[types.Params]

	bankKeeper         types.BankKeeper
	UserGroup          collections.Map[string, types.UserGroup]
	ContentReport      collections.Map[string, types.ContentReport]
	GovernanceProposal collections.Map[string, types.GovernanceProposal]
}

func NewKeeper(
	storeService corestore.KVStoreService,
	cdc codec.Codec,
	addressCodec address.Codec,
	authority []byte,

	bankKeeper types.BankKeeper,
) *Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		storeService: storeService,
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,

		bankKeeper: bankKeeper,
		Params:     collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		UserGroup:  collections.NewMap(sb, types.UserGroupKey, "userGroup", collections.StringKey, codec.CollValue[types.UserGroup](cdc)), ContentReport: collections.NewMap(sb, types.ContentReportKey, "contentReport", collections.StringKey, codec.CollValue[types.ContentReport](cdc)), GovernanceProposal: collections.NewMap(sb, types.GovernanceProposalKey, "governanceProposal", collections.StringKey, codec.CollValue[types.GovernanceProposal](cdc))}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return &k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() []byte {
	return k.authority
}
