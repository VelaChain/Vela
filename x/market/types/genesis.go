package types

import (
	"fmt"
	host "github.com/cosmos/ibc-go/v2/modules/core/24-host"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:      PortID,
		PoolList:    []Pool{},
		LiqProvList: []LiqProv{},
		FeeMapList:  []FeeMap{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated index in pool
	poolIndexMap := make(map[string]struct{})

	for _, elem := range gs.PoolList {
		index := string(PoolKey(elem.DenomA, elem.DenomB))
		if _, ok := poolIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pool")
		}
		poolIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in liqProv
	liqProvIndexMap := make(map[string]struct{})

	for _, elem := range gs.LiqProvList {
		index := string(LiqProvKey(elem.PoolName, elem.Address))
		if _, ok := liqProvIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for liqProv")
		}
		liqProvIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in feeMap
	feeMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.FeeMapList {
		index := string(FeeMapKey(elem.PoolName))
		if _, ok := feeMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for feeMap")
		}
		feeMapIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
