package evm

import (
	"encoding/json"

	tokenfactorykeeper "github.com/QuadrateOrg/core/x/tokenfactory/keeper"
	tokenfactorytypes "github.com/QuadrateOrg/core/x/tokenfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	ek evmkeeper.Keeper,
	fk feemarketkeeper.Keeper,
	tf tokenfactorykeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		fromVM[evmtypes.ModuleName] = mm.Modules[evmtypes.ModuleName].ConsensusVersion()
		fromVM[feemarkettypes.ModuleName] = mm.Modules[feemarkettypes.ModuleName].ConsensusVersion()
		fromVM[tokenfactorytypes.ModuleName] = mm.Modules[tokenfactorytypes.ModuleName].ConsensusVersion()

		var evmparams EvmUpgradeParams
		var tfparams TfUpgradeParams
		err := json.Unmarshal([]byte(plan.Info), &evmparams)
		if err != nil {
			panic(err)
		}

		ek.SetParams(ctx, evmparams.Evm)
		fk.SetParams(ctx, evmparams.FeeMarket)
		tf.SetParams(ctx, tfparams.Tf)

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
