package v1

import (
	"github.com/QuadrateOrg/core/app"
	growtypes "github.com/QuadrateOrg/core/x/grow/types"
	stabletypes "github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	app *app.QuadrateApp,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		migrations, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, err
		}

		app.StableKeeper.SetParams(ctx, stabletypes.DefaultParams())
		app.GrowKeeper.SetParams(ctx, growtypes.DefaultParams())

		pair, _ := app.GrowKeeper.GetPairByDenomID(ctx, app.GrowKeeper.GenerateDenomIdHash("uusd"))
		pair.GTokenLastPrice = sdk.NewInt(1 * 1000000)
		app.GrowKeeper.SetPair(ctx, pair)

		return migrations, nil
	}
}
