package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/VelaChain/vela/testutil/keeper"
	"github.com/VelaChain/vela/testutil/nullify"
	"github.com/VelaChain/vela/x/market/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestLiqProvQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLiqProv(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetLiqProvRequest
		response *types.QueryGetLiqProvResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetLiqProvRequest{
				PoolName: msgs[0].PoolName,
				Address:  msgs[0].Address,
			},
			response: &types.QueryGetLiqProvResponse{LiqProv: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetLiqProvRequest{
				PoolName: msgs[1].PoolName,
				Address:  msgs[1].Address,
			},
			response: &types.QueryGetLiqProvResponse{LiqProv: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetLiqProvRequest{
				PoolName: strconv.Itoa(100000),
				Address:  strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.LiqProv(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestLiqProvQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLiqProv(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllLiqProvRequest {
		return &types.QueryAllLiqProvRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LiqProvAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LiqProv), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LiqProv),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LiqProvAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LiqProv), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LiqProv),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.LiqProvAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.LiqProv),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.LiqProvAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
