package trinity

import (
	"github.com/QuadrateOrg/core/app/upgrades"

	store "github.com/cosmos/cosmos-sdk/store/types"
)

const (
	UpgradeName = "trinity"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName: UpgradeName,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	},
}
