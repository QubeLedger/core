package v2

import (
	"errors"

	stablemodulekeeper "github.com/QuadrateOrg/core/x/stable/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	stablekeepers stablemodulekeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		migrations, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, err
		}

		pair, found := stablekeepers.GetPairByPairID(ctx, stablekeepers.GeneratePairIdHash("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", "uusd"))
		if !found {
			return nil, errors.New("pair not found")
		}

		pair.Model = "gmb"
		pair.MinAmountIn = "20ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2"
		stablekeepers.SetPair(ctx, pair)

		return migrations, nil
	}
}
