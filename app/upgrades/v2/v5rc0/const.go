package v5rc0

import (
	"github.com/QuadrateOrg/core/app/upgrades"

	store "github.com/cosmos/cosmos-sdk/store/types"
)

const (
	UpgradeName = "v0.2.5-rc0"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName: UpgradeName,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	},
}
