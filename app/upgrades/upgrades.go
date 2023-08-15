package upgrades

import (
	"encoding/json"

	tokenfactorykeeper "github.com/QubeLedger/core/x/tokenfactory/keeper"
	tokenfactorytypes "github.com/QubeLedger/core/x/tokenfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	tf tokenfactorykeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		fromVM[tokenfactorytypes.ModuleName] = mm.Modules[tokenfactorytypes.ModuleName].ConsensusVersion()

		var tfparams TfUpgradeParams
		err := json.Unmarshal([]byte(plan.Info), &tfparams)
		if err != nil {
			panic(err)
		}

		tf.SetParams(ctx, tfparams.Tf)

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
