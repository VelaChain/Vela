package types_test

import (
	"testing"

	"github.com/VelaChain/vela/x/market/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
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
				FeeMapList: []types.FeeMap{
					{
						PoolName: "0",
					},
					{
						PoolName: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated pool",
			genState: &types.GenesisState{
				PoolList: []types.Pool{
					{
						DenomA: "0",
						DenomB: "0",
					},
					{
						DenomA: "0",
						DenomB: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated provider",
			genState: &types.GenesisState{
				ProviderList: []types.Provider{
					{
						DenomA: "0",
						DenomB: "0",
					},
					{
						DenomA: "0",
						DenomB: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated liqProv",
			genState: &types.GenesisState{
				LiqProvList: []types.LiqProv{
					{
						PoolName: "0",
						Address:  "0",
					},
					{
						PoolName: "0",
						Address:  "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated feeMap",
			genState: &types.GenesisState{
				FeeMapList: []types.FeeMap{
					{
						PoolName: "0",
					},
					{
						PoolName: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
