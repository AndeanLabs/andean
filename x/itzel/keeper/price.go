package keeper

import (
	"context"

	"andean/x/itzel/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetPrice set a specific price in the store from its index
func (k Keeper) SetPrice(ctx context.Context, price types.Price) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PriceKeyPrefix))
	b := k.cdc.MustMarshal(&price)
	store.Set(types.PriceKey(
		price.Source,
	), b)
}

// GetPrice returns a price from its index
func (k Keeper) GetPrice(
	ctx context.Context,
	source string,

) (val types.Price, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PriceKeyPrefix))

	b := store.Get(types.PriceKey(
		source,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePrice removes a price from the store
func (k Keeper) RemovePrice(
	ctx context.Context,
	source string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PriceKeyPrefix))
	store.Delete(types.PriceKey(
		source,
	))
}

// GetAllPrice returns all price
func (k Keeper) GetAllPrice(ctx context.Context) (list []types.Price) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PriceKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Price
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
