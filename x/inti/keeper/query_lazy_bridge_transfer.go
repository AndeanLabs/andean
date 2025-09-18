package keeper

import (
	"context"

	"andean/x/inti/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LazyBridgeTransferAll(ctx context.Context, req *types.QueryAllLazyBridgeTransferRequest) (*types.QueryAllLazyBridgeTransferResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var lazyBridgeTransfers []types.LazyBridgeTransfer

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	lazyBridgeTransferStore := prefix.NewStore(store, types.KeyPrefix(types.LazyBridgeTransferKeyPrefix))

	pageRes, err := query.Paginate(lazyBridgeTransferStore, req.Pagination, func(key []byte, value []byte) error {
		var lazyBridgeTransfer types.LazyBridgeTransfer
		if err := k.cdc.Unmarshal(value, &lazyBridgeTransfer); err != nil {
			return err
		}

		lazyBridgeTransfers = append(lazyBridgeTransfers, lazyBridgeTransfer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLazyBridgeTransferResponse{LazyBridgeTransfer: lazyBridgeTransfers, Pagination: pageRes}, nil
}

func (k Keeper) LazyBridgeTransfer(ctx context.Context, req *types.QueryGetLazyBridgeTransferRequest) (*types.QueryGetLazyBridgeTransferResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetLazyBridgeTransfer(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLazyBridgeTransferResponse{LazyBridgeTransfer: val}, nil
}
