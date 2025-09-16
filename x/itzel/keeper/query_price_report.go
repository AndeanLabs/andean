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

func (k Keeper) PriceReportAll(ctx context.Context, req *types.QueryAllPriceReportRequest) (*types.QueryAllPriceReportResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var priceReports []types.PriceReport

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	priceReportStore := prefix.NewStore(store, types.KeyPrefix(types.PriceReportKeyPrefix))

	pageRes, err := query.Paginate(priceReportStore, req.Pagination, func(key []byte, value []byte) error {
		var priceReport types.PriceReport
		if err := k.cdc.Unmarshal(value, &priceReport); err != nil {
			return err
		}

		priceReports = append(priceReports, priceReport)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPriceReportResponse{PriceReport: priceReports, Pagination: pageRes}, nil
}

func (k Keeper) PriceReport(ctx context.Context, req *types.QueryGetPriceReportRequest) (*types.QueryGetPriceReportResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetPriceReport(
		ctx,
		req.Source,
		req.Oracle,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPriceReportResponse{PriceReport: val}, nil
}
