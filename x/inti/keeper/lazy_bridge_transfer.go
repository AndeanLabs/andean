package keeper

import (
	"context"

	"andean/x/inti/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetLazyBridgeTransfer set a specific lazyBridgeTransfer in the store from its index
func (k Keeper) SetLazyBridgeTransfer(ctx context.Context, lazyBridgeTransfer types.LazyBridgeTransfer) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LazyBridgeTransferKeyPrefix))
	b := k.cdc.MustMarshal(&lazyBridgeTransfer)
	store.Set(types.LazyBridgeTransferKey(
		lazyBridgeTransfer.Index,
	), b)
}

// GetLazyBridgeTransfer returns a lazyBridgeTransfer from its index
func (k Keeper) GetLazyBridgeTransfer(
	ctx context.Context,
	index string,

) (val types.LazyBridgeTransfer, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LazyBridgeTransferKeyPrefix))

	b := store.Get(types.LazyBridgeTransferKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLazyBridgeTransfer removes a lazyBridgeTransfer from the store
func (k Keeper) RemoveLazyBridgeTransfer(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LazyBridgeTransferKeyPrefix))
	store.Delete(types.LazyBridgeTransferKey(
		index,
	))
}

// GetAllLazyBridgeTransfer returns all lazyBridgeTransfer
func (k Keeper) GetAllLazyBridgeTransfer(ctx context.Context) (list []types.LazyBridgeTransfer) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LazyBridgeTransferKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LazyBridgeTransfer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
