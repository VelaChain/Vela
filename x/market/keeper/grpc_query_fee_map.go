package keeper

import (
	"context"

	"github.com/VelaChain/vela/x/market/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FeeMapAll(c context.Context, req *types.QueryAllFeeMapRequest) (*types.QueryAllFeeMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var feeMaps []types.FeeMap
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	feeMapStore := prefix.NewStore(store, types.KeyPrefix(types.FeeMapKeyPrefix))

	pageRes, err := query.Paginate(feeMapStore, req.Pagination, func(key []byte, value []byte) error {
		var feeMap types.FeeMap
		if err := k.cdc.Unmarshal(value, &feeMap); err != nil {
			return err
		}

		feeMaps = append(feeMaps, feeMap)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFeeMapResponse{FeeMap: feeMaps, Pagination: pageRes}, nil
}

func (k Keeper) FeeMap(c context.Context, req *types.QueryGetFeeMapRequest) (*types.QueryGetFeeMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetFeeMap(
		ctx,
		req.PoolName,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetFeeMapResponse{FeeMap: val}, nil
}
