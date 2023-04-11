package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AcDataList: []AcData{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in acData
	acDataIdMap := make(map[uint64]bool)
	acDataCount := gs.GetAcDataCount()
	for _, elem := range gs.AcDataList {
		if _, ok := acDataIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for acData")
		}
		if elem.Id >= acDataCount {
			return fmt.Errorf("acData id should be lower or equal than the last id")
		}
		acDataIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
