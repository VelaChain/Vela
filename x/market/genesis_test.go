package market_test

import (
	"testing"

	keepertest "github.com/VelaChain/vela/testutil/keeper"
	"github.com/VelaChain/vela/testutil/nullify"
	"github.com/VelaChain/vela/x/market"
	"github.com/VelaChain/vela/x/market/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		PoolList: []types.Pool{
			{
				DenomA: "0",
				DenomB: "0",
			},
			{
				DenomA: "1",
				DenomB: "1",
			},
		},
		ProviderList: []types.Provider{
			{
				DenomA: "0",
				DenomB: "0",
			},
			{
				DenomA: "1",
				DenomB: "1",
			},
		},
		LiqProvList: []types.LiqProv{
			{
				PoolName: "0",
				Address:  "0",
			},
			{
				PoolName: "1",
				Address:  "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MarketKeeper(t)
	market.InitGenesis(ctx, *k, genesisState)
	got := market.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.PoolList, got.PoolList)
	require.ElementsMatch(t, genesisState.ProviderList, got.ProviderList)
	require.ElementsMatch(t, genesisState.LiqProvList, got.LiqProvList)
	// this line is used by starport scaffolding # genesis/test/assert
}
