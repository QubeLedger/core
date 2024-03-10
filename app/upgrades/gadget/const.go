package gadget

import (
	"github.com/QuadrateOrg/core/app/upgrades"

	store "github.com/cosmos/cosmos-sdk/store/types"
)

const (
	UpgradeName = "gadget"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName: UpgradeName,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	},
}
