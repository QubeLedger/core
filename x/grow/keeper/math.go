package keeper

import (
	"math"

	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
Deposit Helpers
*/
func (k Keeper) CalculateGTokenAmountOut(amount sdk.Int, price sdk.Int) sdk.Int {
	return ((amount.Mul(types.DepositMultiplier)).Quo(price))
}

func (k Keeper) CalculateReturnQubeStableAmountOut(amount sdk.Int, price sdk.Int) sdk.Int {
	return ((amount.Mul(price)).Quo(types.DepositMultiplier))
}

func (k Keeper) CalculateGTokenAPY(lastAmount sdk.Int, growRate sdk.Int, day sdk.Int) sdk.Int {
	lastAmountInt := lastAmount.Int64()
	growRateInt := growRate.Int64()
	dayInt := day.Int64()

	res := float64(lastAmountInt) * (math.Pow((1 + (float64(growRateInt)/1000)/365), (float64(dayInt) - 1)))
	return sdk.NewInt(int64(res))
}

/*
Lend Helpers
*/

func (k Keeper) CalculateCreateLendAmountOut(amount sdk.Int, price sdk.Int) sdk.Int {
	return ((amount.Mul(price)).Quo(types.Multiplier))
}

func (k Keeper) CalculateDeleteLendAmountOut(amount sdk.Int, price sdk.Int) sdk.Int {
	return ((amount.Mul(types.Multiplier)).Quo(price))
}

func (k Keeper) CalculateNeedAmountToGet(borrow_amount sdk.Int, borrow_time sdk.Int) sdk.Int {
	return (borrow_amount.Add(((borrow_amount.MulRaw(15).Mul(borrow_time)).QuoRaw(100)).QuoRaw(31536000)))
}

/*
RR Logic
*/

func (k Keeper) CalculateRiskRate(collateral sdk.Int, price sdk.Int, borrow sdk.Int) (sdk.Int, error) {
	amtCollateral := (collateral.Mul(price)).QuoRaw(10000)
	riskRatio := uint64((float64((borrow.Quo(amtCollateral)).Int64()) * float64(1/60)) * 10000)
	return sdk.NewIntFromUint64(riskRatio), nil
}

func (k Keeper) CheckRiskRate(collateral sdk.Int, price sdk.Int, borrow sdk.Int, desiredAmount sdk.Int) error {
	rr, err := k.CalculateRiskRate(collateral, price, borrow)
	if err != nil {
		return err
	}
	if rr.GT(sdk.NewInt(95)) {
		return types.ErrRiskRateIsGreaterThenShouldBe
	}

	rr, err = k.CalculateRiskRate(collateral, price, borrow.Add(desiredAmount))
	if err != nil {
		return err
	}

	if rr.GT(sdk.NewInt(95)) {
		return types.ErrRiskRateIsGreaterThenShouldBe
	}
	return nil
}
