package keeper

import (
	"context"

	"andean/x/itzel/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AggregatedPriceAll(ctx context.Context, req *types.QueryAllAggregatedPriceRequest) (*types.QueryAllAggregatedPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var aggregatedPrices []types.AggregatedPrice

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	aggregatedPriceStore := prefix.NewStore(store, types.KeyPrefix(types.AggregatedPriceKeyPrefix))

	pageRes, err := query.Paginate(aggregatedPriceStore, req.Pagination, func(key []byte, value []byte) error {
		var aggregatedPrice types.AggregatedPrice
		if err := k.cdc.Unmarshal(value, &aggregatedPrice); err != nil {
			return err
		}

		aggregatedPrices = append(aggregatedPrices, aggregatedPrice)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAggregatedPriceResponse{AggregatedPrice: aggregatedPrices, Pagination: pageRes}, nil
}

func (k Keeper) AggregatedPrice(ctx context.Context, req *types.QueryGetAggregatedPriceRequest) (*types.QueryGetAggregatedPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetAggregatedPrice(
		ctx,
		req.Source,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAggregatedPriceResponse{AggregatedPrice: val}, nil
}
