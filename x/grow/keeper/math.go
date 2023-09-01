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
