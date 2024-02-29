package keeper

import (
	"math"

	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
Deposit Math Helpers
*/

func (k Keeper) CalculateGTokenAmountOut(amount sdk.Int, price sdk.Int) sdk.Int {
	return ((amount.Mul(types.CoinMultiplier)).Quo(price))
}

func (k Keeper) CalculateReturnQubeStableAmountOut(amount sdk.Int, price sdk.Int) sdk.Int {
	return ((amount.Mul(price)).Quo(types.CoinMultiplier))
}

/* #nosec */
func (k Keeper) CalculateGTokenAPY(lastAmount sdk.Int, growRate sdk.Int, day sdk.Int) sdk.Int {
	lastAmountInt := lastAmount.Int64()
	growRateInt := growRate.Int64()
	dayInt := day.Int64()

	res := float64(lastAmountInt) * (math.Pow((1 + (float64(growRateInt)/1000)/365), (float64(dayInt) - 1)))
	return sdk.NewInt(int64(res))
}

/*
Lend Math Helpers
*/

func (k Keeper) CalculateAmountByPriceAndAmountIn(amount sdk.Int, price sdk.Int) sdk.Int {
	return ((amount.Mul(price)).Quo(types.Multiplier))
}

func (k Keeper) CalculateDeleteLendAmountOut(amount sdk.Int, price sdk.Int) sdk.Int {
	return ((amount.Mul(types.Multiplier)).Quo(price))
}

func (k Keeper) CalculateNeedAmountToGet(borrow_amount sdk.Int, borrow_time sdk.Int) sdk.Int {
	return (borrow_amount.Add(((borrow_amount.MulRaw(15).Mul(borrow_time)).QuoRaw(100)).QuoRaw(31536000)))
}

func (k Keeper) CalculateAmountForRemoveFromCollateral(amt sdk.Int, price sdk.Int) sdk.Int {
	return ((amt.Mul(types.Multiplier)).Quo(price))
}

/*
RR Math Logic
*/

/* #nosec */
func (k Keeper) CalculateRiskRate(collateral sdk.Int, borrow sdk.Int) (sdk.Int, error) {
	mul := float64(1) / float64(60)
	riskRatio := (((float64(borrow.Int64()) / float64(collateral.Int64())) * mul) * 10000)
	return sdk.NewInt(int64(riskRatio)), nil
}

func (k Keeper) CheckRiskRate(collateral sdk.Int, borrow sdk.Int, desiredAmount sdk.Int) error {
	rr, err := k.CalculateRiskRate(collateral, borrow)
	if err != nil {
		return err
	}
	if rr.GT(sdk.NewInt(95)) {
		return types.ErrRiskRateIsGreaterThenShouldBe
	}

	rr, err = k.CalculateRiskRate(collateral, borrow.Add(desiredAmount))
	if err != nil {
		return err
	}

	if rr.GT(sdk.NewInt(95)) {
		return types.ErrRiskRateIsGreaterThenShouldBe
	}
	return nil
}

/*
Liquidator Math Logic
*/

/* #nosec */
func (k Keeper) CalculateAmountLiquidate(ctx sdk.Context, collateral int64, borrow int64) sdk.Int {
	return sdk.NewInt(int64(((2.85 * float64(collateral)) - float64(5*borrow)) / float64(-2.15)))

}

func (k Keeper) CalculatePremiumAmount(ctx sdk.Context, amount sdk.Int, premium int64, price sdk.Int, price1 sdk.Int) (sdk.Int, sdk.Int) {
	amountInt := amount.Int64()
	priceInt := price.Int64()
	price1Int := price1.Int64()

	plusAmount := amountInt + ((amountInt * premium) / 100)
	return sdk.NewInt((((plusAmount * priceInt) / 10000) * 10000) / price1Int), sdk.NewInt(plusAmount)
}
