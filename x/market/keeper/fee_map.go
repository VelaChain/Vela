package keeper

import (
	"github.com/VelaChain/vela/x/market/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetFeeMap set a specific feeMap in the store from its index
func (k Keeper) SetFeeMap(ctx sdk.Context, feeMap types.FeeMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeMapKeyPrefix))
	b := k.cdc.MustMarshal(&feeMap)
	store.Set(types.FeeMapKey(
		feeMap.PoolName,
	), b)
}

// GetFeeMap returns a feeMap from its index
func (k Keeper) GetFeeMap(
	ctx sdk.Context,
	poolName string,

) (val types.FeeMap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeMapKeyPrefix))

	b := store.Get(types.FeeMapKey(
		poolName,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFeeMap removes a feeMap from the store
func (k Keeper) RemoveFeeMap(
	ctx sdk.Context,
	poolName string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeMapKeyPrefix))
	store.Delete(types.FeeMapKey(
		poolName,
	))
}

// GetAllFeeMap returns all feeMap
func (k Keeper) GetAllFeeMap(ctx sdk.Context) (list []types.FeeMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeMapKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FeeMap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
