package grow

import (
	"time"

	"github.com/QuadrateOrg/core/x/grow/keeper"
	"github.com/QuadrateOrg/core/x/grow/types"
	stabletypes "github.com/QuadrateOrg/core/x/stable/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker is called at the end of every block
/* #nosec */
func EndBlocker(ctx sdk.Context, k keeper.Keeper) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)
	allGTokenPair := k.GetAllPair(ctx)
	for _, gp := range allGTokenPair {
		action, value, err := k.CheckYieldRate(ctx, gp)
		if err != nil {
			return err
		}

		if value.IsNil() {
			return types.ErrIntNegativeOrZero
		}

		if action == types.SendToReserveAction {
			SendToReserveAction(k, ctx, value, gp)
		}

		if action == types.SendFromReserveAction {
			SendFromReserveAction(k, ctx, value, gp)
		}

		err = k.UpdateGTokenPrice(ctx, gp)
		if err != nil {
			return err
		}
	}

	return nil
}

func SendToReserveAction(k keeper.Keeper, ctx sdk.Context, value sdk.Int, gTokenPair types.GTokenPair) error {
	qStablePair, found := k.GetPairByPairID(ctx, gTokenPair.QStablePairId)
	if !found {
		return stabletypes.ErrPairNotFound
	}

	amt := sdk.NewCoins(sdk.NewCoin(qStablePair.AmountOutMetadata.Base, value))

	err := k.SendCoinsFromAccountToModule(ctx, k.GetUSQStakingReserveAddress(ctx), types.ModuleName, amt)
	if err != nil {
		return err
	}

	err = k.SendCoinsFromModuleToAccount(ctx, types.ModuleName, k.GetUSQReserveAddress(ctx), amt)
	if err != nil {
		return err
	}

	return nil
}

func SendFromReserveAction(k keeper.Keeper, ctx sdk.Context, value sdk.Int, gTokenPair types.GTokenPair) error {
	qStablePair, found := k.GetPairByPairID(ctx, gTokenPair.QStablePairId)
	if !found {
		return stabletypes.ErrPairNotFound
	}

	amt := sdk.NewCoins(sdk.NewCoin(qStablePair.AmountOutMetadata.Base, value))

	err := k.SendCoinsFromAccountToModule(ctx, k.GetUSQReserveAddress(ctx), types.ModuleName, amt)
	if err != nil {
		return err
	}

	err = k.SendCoinsFromModuleToAccount(ctx, types.ModuleName, k.GetUSQStakingReserveAddress(ctx), amt)
	if err != nil {
		return err
	}

	return nil
}
