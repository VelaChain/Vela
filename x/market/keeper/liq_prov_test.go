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

func createNLiqProv(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LiqProv {
	items := make([]types.LiqProv, n)
	for i := range items {
		items[i].PoolName = strconv.Itoa(i)
		items[i].Address = strconv.Itoa(i)

		keeper.SetLiqProv(ctx, items[i])
	}
	return items
}

func TestLiqProvGet(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNLiqProv(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLiqProv(ctx,
			item.PoolName,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLiqProvRemove(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNLiqProv(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLiqProv(ctx,
			item.PoolName,
			item.Address,
		)
		_, found := keeper.GetLiqProv(ctx,
			item.PoolName,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestLiqProvGetAll(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNLiqProv(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLiqProv(ctx)),
	)
}
