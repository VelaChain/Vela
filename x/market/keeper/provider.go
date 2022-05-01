package keeper

import (
	"github.com/VelaChain/vela/x/market/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetProvider set a specific provider in the store from its index
func (k Keeper) SetProvider(ctx sdk.Context, provider types.Provider) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderKeyPrefix))
	b := k.cdc.MustMarshal(&provider)
	store.Set(types.ProviderKey(
		provider.DenomA,
		provider.DenomB,
	), b)
}

// GetProvider returns a provider from its index
func (k Keeper) GetProvider(
	ctx sdk.Context,
	denomA string,
	denomB string,

) (val types.Provider, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderKeyPrefix))

	b := store.Get(types.ProviderKey(
		denomA,
		denomB,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProvider removes a provider from the store
func (k Keeper) RemoveProvider(
	ctx sdk.Context,
	denomA string,
	denomB string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderKeyPrefix))
	store.Delete(types.ProviderKey(
		denomA,
		denomB,
	))
}

// GetAllProvider returns all provider
func (k Keeper) GetAllProvider(ctx sdk.Context) (list []types.Provider) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Provider
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
