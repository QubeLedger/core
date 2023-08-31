package types

import (
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	// this line is used by starport scaffolding # genesis/types/import
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:             DefaultParams(),
		PortId:             PortID,
		Pairs:              []Pair{},
		ReserveFundAddress: "",
		BurningFundAddress: "",
	}
}

func NewGenesisState(params Params, portID string, pairs []Pair, reserveFundAddress string, burningFundBalance string) GenesisState {
	return GenesisState{
		Params:             params,
		PortId:             portID,
		Pairs:              pairs,
		ReserveFundAddress: reserveFundAddress,
		BurningFundAddress: burningFundBalance,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}

	return gs.Params.Validate()
}
