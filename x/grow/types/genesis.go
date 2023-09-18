package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const DefaultIndex uint64 = 1

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:                    DefaultParams(),
		GTokenPairList:            []GTokenPair{},
		RealRate:                  1,
		BorrowRate:                1,
		GrowStakingReserveAddress: "qube1xpurqvpsxqcrqvpsxqcrqvpsxqcrqvpsxqcrqv",
		USQReserveAddress:         "qube1xpurqvpsxqcrqvpsxqcrqvpsxqcrqvpsxqcrqv",
	}
}

/* #nosec */
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

	if sdk.NewInt(int64(gs.RealRate)).IsNegative() || sdk.NewInt(int64(gs.RealRate)).IsZero() {
		return fmt.Errorf("RealRate negative or zero")
	}

	if sdk.NewInt(int64(gs.BorrowRate)).IsNegative() || sdk.NewInt(int64(gs.BorrowRate)).IsZero() {
		return fmt.Errorf("BorrowRate negative or zero")
	}

	return gs.Params.Validate()
}
