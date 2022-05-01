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

func createNProvider(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Provider {
	items := make([]types.Provider, n)
	for i := range items {
		items[i].DenomA = strconv.Itoa(i)
		items[i].DenomB = strconv.Itoa(i)

		keeper.SetProvider(ctx, items[i])
	}
	return items
}

func TestProviderGet(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNProvider(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProvider(ctx,
			item.DenomA,
			item.DenomB,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestProviderRemove(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNProvider(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProvider(ctx,
			item.DenomA,
			item.DenomB,
		)
		_, found := keeper.GetProvider(ctx,
			item.DenomA,
			item.DenomB,
		)
		require.False(t, found)
	}
}

func TestProviderGetAll(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNProvider(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProvider(ctx)),
	)
}
