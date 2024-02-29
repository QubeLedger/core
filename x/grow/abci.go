package grow

import (
	"fmt"
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

	params := k.GetParams(ctx)

	err := k.CheckDepositMethodStatus(ctx)
	if err == nil {
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
	}

	if k.CheckCollateralMethodStatus(ctx) == nil && k.CheckBorrowMethodStatus(ctx) == nil {
		allPosition := k.GetAllPosition(ctx)
		for _, position := range allPosition {
			NewLendAmountInUSD := 0.0
			NewBorrowAmountInUSD := 0.0

			for _, lend_id := range position.LendId {
				lend, found := k.GetLendByLendId(ctx, lend_id)
				if !found {
					return types.ErrLendNotFound
				}

				price, err := k.GetPriceByDenom(ctx, lend.OracleTicker)
				if err != nil {
					return err
				}
				if err != nil && price.IsNil() {
					return types.ErrIntNegativeOrZero
				}

				collateral_amount, err := lend.AmountInAmount.Float64()
				if err != nil {
					return err
				}

				asset, err := k.GetAssetByOracleAssetId(ctx, lend.OracleTicker)
				if err != nil {
					return types.ErrAssetNotFound
				}

				sir := 0.0

				utilization_rate := (float64(asset.CollectivelyBorrowValue) / float64(asset.ProvideValue))
				if utilization_rate > 0 {
					_, sir_temp, err := k.GetRatesByUtilizationRate(ctx, utilization_rate, asset)
					if err != nil {
						return types.ErrCalculateBIROrSIR
					}
					sir = sir_temp
				}

				time := ctx.BlockTime().Unix() - int64(params.LastTimeUpdateReserve)
				result := collateral_amount + ((collateral_amount * sir * float64(time)) / 31536000)
				if sir <= 0 {
					lend.AmountInAmount = sdk.MustNewDecFromStr(fmt.Sprintf("%f", collateral_amount))
					NewLendAmountInUSD += result * (float64(price.Int64()) / 10000)
				} else {
					lend.AmountInAmount = sdk.MustNewDecFromStr(fmt.Sprintf("%f", result))
					NewLendAmountInUSD += result * (float64(price.Int64()) / 10000)
				}
				k.SetLend(ctx, lend)
			}
			for _, loan_id := range position.LoanId {
				loan, found := k.GetLoadByLoanId(ctx, loan_id)
				if !found {
					return types.ErrLoanNotFound
				}

				price, err := k.GetPriceByDenom(ctx, loan.OracleTicker)
				if err != nil {
					return err
				}
				if err != nil && price.IsNil() {
					return types.ErrIntNegativeOrZero
				}

				borrow_amount, err := loan.AmountOutAmount.Float64()
				if err != nil {
					return err
				}

				asset, err := k.GetAssetByOracleAssetId(ctx, loan.OracleTicker)
				if err != nil {
					return types.ErrAssetNotFound
				}

				bir := 0.0

				utilization_rate := (float64(asset.CollectivelyBorrowValue) / float64(asset.ProvideValue))
				if utilization_rate > 0 {
					bir_temp, _, err := k.GetRatesByUtilizationRate(ctx, utilization_rate, asset)
					if err != nil {
						return types.ErrCalculateBIROrSIR
					}
					bir = bir_temp
				}

				time := ctx.BlockTime().Unix() - int64(params.LastTimeUpdateReserve)
				result := borrow_amount + ((borrow_amount * bir * float64(time)) / 31536000)
				if bir <= 0 {
					loan.AmountOutAmount = sdk.MustNewDecFromStr(fmt.Sprintf("%f", borrow_amount))
					NewBorrowAmountInUSD += borrow_amount * (float64(price.Int64()) / 10000)
				} else {
					loan.AmountOutAmount = sdk.MustNewDecFromStr(fmt.Sprintf("%f", result))
					NewBorrowAmountInUSD += result * (float64(price.Int64()) / 10000)
				}
				k.SetLoan(ctx, loan)
			}

			position.BorrowedAmountInUSD = uint64(NewBorrowAmountInUSD)
			position.LendAmountInUSD = uint64(NewLendAmountInUSD)
			k.SetPosition(ctx, position)
		}
	}

	if k.CheckCollateralMethodStatus(ctx) == nil && k.CheckBorrowMethodStatus(ctx) == nil {
		allPosition := k.GetAllPosition(ctx)
		liquidateLendPositionList := []string{}
		for _, pos := range allPosition {
			if !sdk.NewIntFromUint64(pos.BorrowedAmountInUSD).IsZero() {
				rr, err := k.CalculateRiskRate(sdk.NewIntFromUint64(pos.LendAmountInUSD), sdk.NewIntFromUint64(pos.BorrowedAmountInUSD))
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
			k.ReCalculateLendLoanAmountsInUsd(ctx, pos)
		}
		if len(liquidateLendPositionList) != 0 {
			err := k.ExecuteLiquidation(ctx, liquidateLendPositionList)
			if err != nil {
				return err
			}
		}
	}

	params.LastTimeUpdateReserve = uint64(ctx.BlockTime().Unix())

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
