package types

import (
	"fmt"
)

const DefaultIndex uint64 = 1

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:         DefaultParams(),
		GTokenPairList: []GTokenPair{},
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
	return gs.Params.Validate()
}
