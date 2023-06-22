package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PriceList: []Price{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in price
	priceIdMap := make(map[uint64]bool)
	priceCount := gs.GetPriceCount()
	for _, elem := range gs.PriceList {
		if _, ok := priceIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for price")
		}
		if elem.Id >= priceCount {
			return fmt.Errorf("price id should be lower or equal than the last id")
		}
		priceIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
