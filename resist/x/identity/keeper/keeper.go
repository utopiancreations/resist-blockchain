package keeper

import (
	"fmt"
	"strings"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"resist/x/identity/types"
)

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec
	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority []byte

	Schema      collections.Schema
	Params      collections.Item[types.Params]
	UserProfile collections.Map[string, types.UserProfile]
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

		Params:      collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		UserProfile: collections.NewMap(sb, types.UserProfileKey, "userProfile", collections.StringKey, codec.CollValue[types.UserProfile](cdc))}

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

// GetChallenge returns the challenge for a given address.
func (k Keeper) GetChallenge(ctx sdk.Context, address string) (string, error) {
	store := k.storeService.OpenKVStore(ctx)
	challengeKey := fmt.Sprintf("challenge:%s", address)
	challengeData, err := store.Get([]byte(challengeKey))
	if err != nil {
		return "", err
	}
	if challengeData == nil {
		return "", nil
	}

	parts := strings.Split(string(challengeData), ":")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid challenge data")
	}
	return parts[0], nil
}
