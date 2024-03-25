package v0

import (
	"github.com/QuadrateOrg/core/app/upgrades"

	epochmoduletypes "github.com/QuadrateOrg/core/x/epochs/types"
	ibchookermoduletypes "github.com/QuadrateOrg/core/x/ibchooker/types"
	interchainquerymoduletypes "github.com/QuadrateOrg/core/x/interchainquery/types"
	liquidstakeibcmoduletypes "github.com/QuadrateOrg/core/x/liquidstakeibc/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

const (
	UpgradeName = "v0.3.0"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName: UpgradeName,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			epochmoduletypes.ModuleName,
			interchainquerymoduletypes.ModuleName,
			liquidstakeibcmoduletypes.ModuleName,
			ibchookermoduletypes.ModuleName,
		},
		Deleted: []string{},
	},
}
