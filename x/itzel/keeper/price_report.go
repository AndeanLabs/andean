package keeper

import (
	"context"

	"andean/x/itzel/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetPriceReport set a specific priceReport in the store from its index
func (k Keeper) SetPriceReport(ctx context.Context, priceReport types.PriceReport) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PriceReportKeyPrefix))
	b := k.cdc.MustMarshal(&priceReport)
	store.Set(types.PriceReportKey(
		priceReport.Source,
		priceReport.Oracle,
	), b)
}

// GetPriceReport returns a priceReport from its index
func (k Keeper) GetPriceReport(
	ctx context.Context,
	source string,
	oracle string,

) (val types.PriceReport, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PriceReportKeyPrefix))

	b := store.Get(types.PriceReportKey(
		source,
		oracle,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePriceReport removes a priceReport from the store
func (k Keeper) RemovePriceReport(
	ctx context.Context,
	source string,
	oracle string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PriceReportKeyPrefix))
	store.Delete(types.PriceReportKey(
		source,
		oracle,
	))
}

// GetAllPriceReport returns all priceReport
func (k Keeper) GetAllPriceReport(ctx context.Context) (list []types.PriceReport) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PriceReportKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PriceReport
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPriceReportsBySource returns all price reports for a given source
func (k Keeper) GetPriceReportsBySource(ctx context.Context, source string) (list []types.PriceReport) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PriceReportKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte(source+"/"))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PriceReport
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
