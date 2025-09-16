package keeper

import (
	"context"

	"andean/x/itzel/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetAggregatedPrice set a specific aggregatedPrice in the store from its index
func (k Keeper) SetAggregatedPrice(ctx context.Context, aggregatedPrice types.AggregatedPrice) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AggregatedPriceKeyPrefix))
	b := k.cdc.MustMarshal(&aggregatedPrice)
	store.Set(types.AggregatedPriceKey(
		aggregatedPrice.Source,
	), b)
}

// GetAggregatedPrice returns a aggregatedPrice from its index
func (k Keeper) GetAggregatedPrice(
	ctx context.Context,
	source string,

) (val types.AggregatedPrice, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AggregatedPriceKeyPrefix))

	b := store.Get(types.AggregatedPriceKey(
		source,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAggregatedPrice removes a aggregatedPrice from the store
func (k Keeper) RemoveAggregatedPrice(
	ctx context.Context,
	source string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AggregatedPriceKeyPrefix))
	store.Delete(types.AggregatedPriceKey(
		source,
	))
}

// GetAllAggregatedPrice returns all aggregatedPrice
func (k Keeper) GetAllAggregatedPrice(ctx context.Context) (list []types.AggregatedPrice) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AggregatedPriceKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AggregatedPrice
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
