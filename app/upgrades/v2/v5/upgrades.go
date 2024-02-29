package v5

import (
	"fmt"

	growmodulekeeper "github.com/QuadrateOrg/core/x/grow/keeper"
	growmoduletypes "github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

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

		params := growkeeper.GetParams(ctx)

		params.LastTimeUpdateReserve = uint64(ctx.BlockTime().Unix())

		params.UStaticVolatile = uint64(60)
		params.UStaticStable = uint64(80)

		params.MaxRateVolatile = uint64(300)
		params.MaxRateStable = uint64(100)

		params.Slope_1 = uint64(1)
		params.Slope_2 = uint64(8)

		params.CollateralMethodStatus = true
		params.BorrowMethodStatus = true

		growkeeper.SetParams(ctx, params)
		if err != nil {
			return nil, err
		}

		ctx.Logger().Info(fmt.Sprintf("qLabs: set x/grow params: params: %v -- value: %v", "LastTimeUpdateReserve", params.LastTimeUpdateReserve))
		ctx.Logger().Info(fmt.Sprintf("qLabs: set x/grow params: params: %v -- value: %v", "UStaticVolatile", params.UStaticVolatile))
		ctx.Logger().Info(fmt.Sprintf("qLabs: set x/grow params: params: %v -- value: %v", "UStaticStable", params.UStaticStable))
		ctx.Logger().Info(fmt.Sprintf("qLabs: set x/grow params: params: %v -- value: %v", "MaxRateVolatile", params.MaxRateVolatile))
		ctx.Logger().Info(fmt.Sprintf("qLabs: set x/grow params: params: %v -- value: %v", "Slope_1", params.Slope_1))
		ctx.Logger().Info(fmt.Sprintf("qLabs: set x/grow params: params: %v -- value: %v", "Slope_2", params.Slope_2))
		ctx.Logger().Info(fmt.Sprintf("qLabs: set x/grow params: params: %v -- value: %v", "CollateralMethodStatus", params.CollateralMethodStatus))
		ctx.Logger().Info(fmt.Sprintf("qLabs: set x/grow params: params: %v -- value: %v", "BorrowMethodStatus", params.BorrowMethodStatus))

		// update already created asset

		// asset_id already created asset
		asset_id := "c5b4376538178084416cab617a2cace5a17db8ec762fed02a0ad35ef1a156e29"

		asset, found := growkeeper.GetAssetByAssetId(ctx, asset_id)
		if !found {
			return nil, growmoduletypes.ErrAssetNotFound
		}

		asset.ProvideValue = uint64(0)
		asset.CollectivelyBorrowValue = uint64(0)
		asset.Type = "volatile"

		ctx.Logger().Info(fmt.Sprintf("qLabs: set asset in x/grow with AssetId: %v", asset.AssetId))

		growkeeper.SetAsset(ctx, asset)

		return migrations, nil
	}
}
