package upgrades

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BaseAppParamManager defines an interrace that BaseApp is expected to fulfill
// that allows upgrade handlers to modify BaseApp parameters.
type BaseAppParamManager interface {
	GetConsensusParams(ctx sdk.Context) *tmproto.ConsensusParams
	StoreConsensusParams(ctx sdk.Context, cp *tmproto.ConsensusParams)
}

type Upgrade struct {
	// Upgrade version name, for the upgrade handler, e.g. `v7`
	UpgradeName string

	// Store upgrades, should be used for any new modules introduced, new modules deleted, or store names renamed.
	StoreUpgrades store.StoreUpgrades
}
