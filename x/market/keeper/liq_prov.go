package keeper

import (
	"github.com/VelaChain/vela/x/market/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLiqProv set a specific liqProv in the store from its index
func (k Keeper) SetLiqProv(ctx sdk.Context, liqProv types.LiqProv) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LiqProvKeyPrefix))
	b := k.cdc.MustMarshal(&liqProv)
	store.Set(types.LiqProvKey(
		liqProv.PoolName,
		liqProv.Address,
	), b)
}

// GetLiqProv returns a liqProv from its index
func (k Keeper) GetLiqProv(
	ctx sdk.Context,
	poolName string,
	address string,

) (val types.LiqProv, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LiqProvKeyPrefix))

	b := store.Get(types.LiqProvKey(
		poolName,
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLiqProv removes a liqProv from the store
func (k Keeper) RemoveLiqProv(
	ctx sdk.Context,
	poolName string,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LiqProvKeyPrefix))
	store.Delete(types.LiqProvKey(
		poolName,
		address,
	))
}

// GetAllLiqProv returns all liqProv
func (k Keeper) GetAllLiqProv(ctx sdk.Context) (list []types.LiqProv) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LiqProvKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LiqProv
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
