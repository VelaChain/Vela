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

func (k Keeper) LiqProvAll(c context.Context, req *types.QueryAllLiqProvRequest) (*types.QueryAllLiqProvResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var liqProvs []types.LiqProv
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	liqProvStore := prefix.NewStore(store, types.KeyPrefix(types.LiqProvKeyPrefix))

	pageRes, err := query.Paginate(liqProvStore, req.Pagination, func(key []byte, value []byte) error {
		var liqProv types.LiqProv
		if err := k.cdc.Unmarshal(value, &liqProv); err != nil {
			return err
		}

		liqProvs = append(liqProvs, liqProv)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLiqProvResponse{LiqProv: liqProvs, Pagination: pageRes}, nil
}

func (k Keeper) LiqProv(c context.Context, req *types.QueryGetLiqProvRequest) (*types.QueryGetLiqProvResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLiqProv(
		ctx,
		req.PoolName,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLiqProvResponse{LiqProv: val}, nil
}
