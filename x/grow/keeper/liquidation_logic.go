package keeper

import (
	"sort"

	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/* #nosec */
func (k Keeper) ExecuteLiquidation(ctx sdk.Context, liquidateLendPositionList []string) error {

	for _, position_id := range liquidateLendPositionList {
		position, _ := k.GetPositionByPositionId(ctx, position_id)

		amount_need_liquidate := k.CalculateAmountLiquidate(ctx, int64(position.LendAmountInUSD), int64(position.BorrowedAmountInUSD))

		temp_position_lend := position.LendId
		temp_position_loan := position.LoanId

		sort.SliceStable(temp_position_lend, func(i, j int) bool {

			lend_i, _ := k.GetLendByLendId(ctx, temp_position_lend[i])
			price_1, _ := k.GetPriceByDenom(ctx, lend_i.OracleTicker)
			lend_j, _ := k.GetLendByLendId(ctx, temp_position_lend[j])
			price_j, _ := k.GetPriceByDenom(ctx, lend_j.OracleTicker)

			return lend_i.AmountInAmount.RoundInt().Mul(price_1).Quo(types.Multiplier).Int64() > lend_j.AmountInAmount.RoundInt().Mul(price_j).Quo(types.Multiplier).Int64()
		})

		sort.SliceStable(temp_position_loan, func(i, j int) bool {

			loan_i, _ := k.GetLoadByLoanId(ctx, temp_position_loan[i])
			price_1, _ := k.GetPriceByDenom(ctx, loan_i.OracleTicker)
			loan_j, _ := k.GetLoadByLoanId(ctx, temp_position_loan[j])
			price_j, _ := k.GetPriceByDenom(ctx, loan_j.OracleTicker)

			return loan_i.AmountOutAmount.RoundInt().Mul(price_1).Quo(types.Multiplier).Int64() > loan_j.AmountOutAmount.RoundInt().Mul(price_j).Quo(types.Multiplier).Int64()
		})

		for i, lend_id := range temp_position_lend {
			lend, _ := k.GetLendByLendId(ctx, lend_id)
			price_lend, _ := k.GetPriceByDenom(ctx, lend.OracleTicker)

			loan, _ := k.GetLoadByLoanId(ctx, temp_position_loan[i])
			price_loan, _ := k.GetPriceByDenom(ctx, loan.OracleTicker)

			lend_amount_usd := lend.AmountInAmount.RoundInt().Mul(price_lend).Quo(types.Multiplier)
			loan_amount_usd := loan.AmountOutAmount.RoundInt().Mul(price_loan).Quo(types.Multiplier)

			if lend_amount_usd.GTE(loan_amount_usd) {
				if loan_amount_usd.GTE(amount_need_liquidate) {
					liqudated, err := k.LiquidatePositionByLiquidators(ctx, lend.OracleTicker, loan.OracleTicker, amount_need_liquidate.Mul(types.Multiplier).Quo(price_loan), price_loan, price_lend)
					if err != nil {
						return err
					}
					amount_need_liquidate = amount_need_liquidate.Sub(liqudated.Mul(price_lend).Quo(types.Multiplier))
					lend.AmountInAmount = lend.AmountInAmount.Sub(liqudated.ToDec())
					loan.AmountOutAmount = loan.AmountOutAmount.Sub((((liqudated.Mul(price_lend).Quo(types.Multiplier)).Mul(types.Multiplier)).Quo(price_loan)).ToDec())
				} else {
					liqudated, err := k.LiquidatePositionByLiquidators(ctx, lend.OracleTicker, loan.OracleTicker, amount_need_liquidate.Mul(types.Multiplier).Quo(price_loan), price_loan, price_lend)
					if err != nil {
						return err
					}
					amount_need_liquidate = amount_need_liquidate.Sub(liqudated.Mul(price_lend).Quo(types.Multiplier))
					lend.AmountInAmount = lend.AmountInAmount.Sub(liqudated.ToDec())
					position = k.RemoveLoanInPosition(ctx, loan.LoanId, position)
					k.RemoveLoan(ctx, loan.Id)
				}
			} else {
				if lend_amount_usd.GTE(amount_need_liquidate) {
					liqudated, err := k.LiquidatePositionByLiquidators(ctx, lend.OracleTicker, loan.OracleTicker, amount_need_liquidate.Mul(types.Multiplier).Quo(price_lend), price_lend, price_loan)
					if err != nil {
						return err
					}
					amount_need_liquidate = amount_need_liquidate.Sub(liqudated.Mul(price_loan).Quo(types.Multiplier))
					lend.AmountInAmount = lend.AmountInAmount.Sub((((liqudated.Mul(price_loan).Quo(types.Multiplier)).Mul(types.Multiplier)).Quo(price_lend)).ToDec())
					loan.AmountOutAmount = loan.AmountOutAmount.Sub(liqudated.ToDec())
				} else {
					liqudated, err := k.LiquidatePositionByLiquidators(ctx, lend.OracleTicker, loan.OracleTicker, amount_need_liquidate.Mul(types.Multiplier).Quo(price_lend), price_lend, price_loan)
					if err != nil {
						return err
					}
					amount_need_liquidate = amount_need_liquidate.Sub(liqudated.Mul(price_loan).Quo(types.Multiplier))
					loan.AmountOutAmount = loan.AmountOutAmount.Sub(liqudated.ToDec())
					position = k.RemoveLendInPosition(ctx, lend.LendId, position)
					k.RemoveLend(ctx, lend.Id)
				}
			}

			k.SetLend(ctx, lend)
			k.SetLoan(ctx, loan)
		}
		k.SetPosition(ctx, position)
		k.ReCalculateLendLoanAmountsInUsd(ctx, position)
	}
	return nil
}

/* #nosec */
func (k Keeper) LiquidatePositionByLiquidators(ctx sdk.Context, lendOracleId string, borrowOracleId string, amount sdk.Int, price sdk.Int, price1 sdk.Int) (sdk.Int, error) {
	allLiquidated := sdk.NewInt(0)
	liquidator_position := k.GetLiquidatorPositionsByAssetAndDenom(ctx, lendOracleId, borrowOracleId)
	liquidator_position = k.SortLiquidatorPositionsByPremium(ctx, liquidator_position)
	for _, lp := range liquidator_position {
		if allLiquidated.GTE(amount) {
			return allLiquidated, nil
		}

		lp_amount, lp_denom, err := k.GetAmountIntFromCoins(lp.Amount)
		if err != nil {
			return sdk.Int{}, err
		}

		asset_lend, err := k.GetAssetByOracleAssetId(ctx, lendOracleId)
		if err != nil {
			return sdk.Int{}, err
		}

		if lp_amount.GTE(amount) {
			return_amount_in_price1_denom, return_amount_in_price_denom := k.CalculatePremiumAmount(ctx, amount, int64(lp.Premium), price, price1)
			err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(lp.Liquidator), k.FastCoins(asset_lend.AssetMetadata.Base, return_amount_in_price1_denom))
			if err != nil {
				return sdk.Int{}, err
			}

			allLiquidated = allLiquidated.Add(return_amount_in_price1_denom)
			lp.Amount = k.FastCoins(lp_denom, lp_amount.Sub(return_amount_in_price_denom)).String()
			k.SetLiquidatorPosition(ctx, lp)
		} else {
			return_amount_in_price1_denom, _ := k.CalculatePremiumAmount(ctx, lp_amount, int64(lp.Premium), price, price1)
			err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(lp.Liquidator), k.FastCoins(asset_lend.AssetMetadata.Base, return_amount_in_price1_denom))
			if err != nil {
				return sdk.Int{}, err
			}
			allLiquidated = allLiquidated.Add(return_amount_in_price1_denom)
			k.RemoveLiquidatorPosition(ctx, lp.Id)
		}
	}
	return sdk.Int{}, types.ErrLiquidationMechanismError
}
