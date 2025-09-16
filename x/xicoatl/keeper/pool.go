package keeper

import (
	"context"

	"andean/x/xicoatl/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetPool set a specific pool in the store from its index
func (k Keeper) SetPool(ctx context.Context, pool types.Pool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PoolKeyPrefix))
	b := k.cdc.MustMarshal(&pool)
	store.Set(types.PoolKey(
		pool.Index,
	), b)
}

// GetPool returns a pool from its index
func (k Keeper) GetPool(
	ctx context.Context,
	index string,

) (val types.Pool, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PoolKeyPrefix))

	b := store.Get(types.PoolKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePool removes a pool from the store
func (k Keeper) RemovePool(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PoolKeyPrefix))
	store.Delete(types.PoolKey(
		index,
	))
}

// GetAllPool returns all pool
func (k Keeper) GetAllPool(ctx context.Context) (list []types.Pool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PoolKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Pool
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
