package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/VelaChain/vela/testutil/keeper"
	"github.com/VelaChain/vela/testutil/nullify"
	"github.com/VelaChain/vela/x/market/keeper"
	"github.com/VelaChain/vela/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNFeeMap(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.FeeMap {
	items := make([]types.FeeMap, n)
	for i := range items {
		items[i].PoolName = strconv.Itoa(i)

		keeper.SetFeeMap(ctx, items[i])
	}
	return items
}

func TestFeeMapGet(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNFeeMap(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFeeMap(ctx,
			item.PoolName,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFeeMapRemove(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNFeeMap(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFeeMap(ctx,
			item.PoolName,
		)
		_, found := keeper.GetFeeMap(ctx,
			item.PoolName,
		)
		require.False(t, found)
	}
}

func TestFeeMapGetAll(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNFeeMap(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFeeMap(ctx)),
	)
}
