package v0

import (
	growmodulekeeper "github.com/QuadrateOrg/core/x/grow/keeper"
	growtypes "github.com/QuadrateOrg/core/x/grow/types"
	liquidstakeibckeeper "github.com/QuadrateOrg/core/x/liquidstakeibc/keeper"
	lsmtypes "github.com/QuadrateOrg/core/x/liquidstakeibc/types"
	stablemodulekeeper "github.com/QuadrateOrg/core/x/stable/keeper"
	stabletypes "github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	stablekeepers stablemodulekeeper.Keeper,
	growkeepers growmodulekeeper.Keeper,
	lsmkeepers liquidstakeibckeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		migrations, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, err
		}

		stablekeepers.SetParams(ctx, stabletypes.DefaultParams())
		growkeepers.SetParams(ctx, growtypes.DefaultParams())
		lsmkeepers.SetParams(ctx, lsmtypes.DefaultParams())

		return migrations, nil
	}
}
