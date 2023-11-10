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

	err := k.CheckGrowStatus(ctx)
	if err != nil {
		return types.ErrGrowNotActivated
	}

	allGTokenPair := k.GetAllPair(ctx)
	for _, gp := range allGTokenPair {
		action, rawValue, err := k.CheckYieldRate(ctx, gp)
		if err != nil {
			return err
		}

		if rawValue.IsNil() || rawValue.IsZero() {
			return types.ErrIntNegativeOrZero
		}

		value, blocked := k.CalculateAddToReserveValue(ctx, rawValue, gp)
		if !blocked && value.IsNil() {
			return types.ErrIntNegativeOrZero
		}

		err = ExecuteReserveAction(k, ctx, value, gp, action, blocked)
		if err != nil {
			return err
		}

		err = k.UpdateGTokenPrice(ctx, gp)
		if err != nil {
			return err
		}
	}

	allPosition := k.GetAllPosition(ctx)
	liquidateLendPositionList := []string{}
	for _, pos := range allPosition {
		if !sdk.NewIntFromUint64(pos.BorrowedAmountInUSD).IsZero() {
			price, err := k.GetPriceByDenom(ctx, pos.OracleTicker)
			if err != nil {
				return err
			}
			if err != nil && price.IsNil() {
				return types.ErrIntNegativeOrZero
			}

			amountPositionInt, _, err := k.GetAmountIntFromCoins(pos.Collateral)
			if err != nil {
				return err
			}
			if err != nil && amountPositionInt.IsNil() {
				return types.ErrIntNegativeOrZero
			}

			rr, err := k.CalculateRiskRate(amountPositionInt, price, sdk.NewIntFromUint64(pos.BorrowedAmountInUSD))
			if err != nil {
				return err
			}

			if err != nil && rr.IsNil() {
				return types.ErrIntNegativeOrZero
			}

			if rr.GTE(sdk.NewInt(95)) {
				liquidateLendPositionList = append(liquidateLendPositionList, pos.DepositId)
			}
		}
	}
	if len(liquidateLendPositionList) != 0 {
		err := k.ExecuteLiquidation(ctx, liquidateLendPositionList)
		if err != nil {
			return err
		}
	}

	return nil
}

func ExecuteReserveAction(k keeper.Keeper, ctx sdk.Context, value sdk.Int, gp types.GTokenPair, action string, blocked bool) error {
	if !blocked {
		switch action {
		case types.SendToReserveAction:
			if err := SendToReserveAction(k, ctx, value, gp); err != nil {
				return err
			}
		case types.SendFromReserveAction:
			if err := SendFromReserveAction(k, ctx, value, gp); err != nil {
				return err
			}

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
