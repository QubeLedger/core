package v5rc02

import (
	"fmt"

	growmodulekeeper "github.com/QuadrateOrg/core/x/grow/keeper"
	growmoduletypes "github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

/* #nosec */
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	growkeeper growmodulekeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		migrations, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, err
		}
		// update already created asset

		// asset_id already created asset
		asset_id := growkeeper.GenerateAssetIdHash("factory/qube1t2ydw7r4asmk74ymuvykjshdzun8dxye0az5wz/wbtc")

		asset, found := growkeeper.GetAssetByAssetId(ctx, asset_id)
		if !found {
			return nil, growmoduletypes.ErrAssetNotFound
		}

		asset.ProvideValue = uint64(0)
		growkeeper.SetAsset(ctx, asset)

		ctx.Logger().Info(fmt.Sprintf("qLabs: set asset in x/grow with AssetId: %v", asset.AssetId))

		new_asset, new_found := growkeeper.GetAssetByAssetId(ctx, asset_id)
		if !new_found {
			return nil, growmoduletypes.ErrAssetNotFound
		}

		if new_asset.ProvideValue != uint64(0) {
			return nil, growmoduletypes.ErrIntNegativeOrZero
		}

		return migrations, nil
	}
}
