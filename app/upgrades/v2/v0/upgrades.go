package v0

import (
	"errors"

	growmodulekeeper "github.com/QuadrateOrg/core/x/grow/keeper"
	growtypes "github.com/QuadrateOrg/core/x/grow/types"
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
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		migrations, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, err
		}

		stablekeepers.SetParams(ctx, stabletypes.DefaultParams())
		growkeepers.SetParams(ctx, growtypes.DefaultParams())

		pair, found := growkeepers.GetPairByDenomID(ctx, growkeepers.GenerateDenomIdHash("ugusd"))
		if !found {
			return nil, errors.New("gTokenPair not found")
		}

		pair.GTokenLastPrice = sdk.NewInt(1 * 1000000)
		growkeepers.SetPair(ctx, pair)

		return migrations, nil
	}
}
