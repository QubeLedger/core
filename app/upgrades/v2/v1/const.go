package v1

import (
	"github.com/QuadrateOrg/core/app/upgrades"

	dexmoduletypes "github.com/QuadrateOrg/core/x/dex/types"
	ibcswapmoduletypes "github.com/QuadrateOrg/core/x/ibcswap/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

const (
	UpgradeName = "v0.2.1"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName: UpgradeName,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			dexmoduletypes.ModuleName,
			ibcswapmoduletypes.ModuleName,
		},
		Deleted: []string{},
	},
}
