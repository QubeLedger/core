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
		action, rawValue, err := k.CheckYieldRate(ctx, gp)
		if err != nil {
			return err
		}

		if rawValue.IsNil() {
			return types.ErrIntNegativeOrZero
		}

		value, blocked := k.CalculateAddToReserveValue(ctx, rawValue, gp)
		if blocked {
			if action == types.SendToReserveAction {
				err := SendToReserveAction(k, ctx, value, gp)
				if err != nil {
					return err
				}
			}

			if action == types.SendFromReserveAction {
				err := SendFromReserveAction(k, ctx, value, gp)
				if err != nil {
					return err
				}
			}
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

	err := k.SendCoinsFromAccountToModule(ctx, k.GetGrowStakingReserveAddress(ctx), types.ModuleName, amt)
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

	err = k.SendCoinsFromModuleToAccount(ctx, types.ModuleName, k.GetGrowStakingReserveAddress(ctx), amt)
	if err != nil {
		return err
	}

	return nil
}
