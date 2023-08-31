package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const DefaultIndex uint64 = 1

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

func (gs GenesisState) Validate() error {
	gTokenPairListIdMap := make(map[uint64]bool)
	for _, elem := range gs.GTokenPairList {
		if _, ok := gTokenPairListIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for gTokenPair")
		}
		gTokenPairListIdMap[elem.Id] = true
	}

	if sdk.AccAddress(gs.GrowStakingReserveAddress).Empty() {
		return fmt.Errorf("GrowStakingReserveAddress empty")
	}

	if sdk.AccAddress(gs.USQReserveAddress).Empty() {
		return fmt.Errorf("GrowStakingReserveAddress empty")
	}

	if sdk.NewIntFromUint64(gs.RealRate).IsNegative() {
		return fmt.Errorf("RealRate negative")
	}

	return gs.Params.Validate()
}
