package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Data struct {
	lp types.LiquidatorPosition
	id string
}

/* #nosec */
func (k Keeper) ExecuteLiquidation(ctx sdk.Context, liquidateLendPositionList []string) error {
	liquidatorList := []Data{}
	allLiquidator := k.GetAllLiquidatorPosition(ctx)
	for _, llp := range liquidateLendPositionList {
		pos, _ := k.GetPositionByPositionId(ctx, llp)
		for _, liquidator := range allLiquidator {
			if pos.OracleTicker == liquidator.BorrowAssetId {
				liquidatorList = append(liquidatorList, Data{
					lp: liquidator,
					id: pos.DepositId,
				})
			}
		}
	}
	for i := 1; i < 100; i++ {
		for _, data := range liquidatorList {
			if data.lp.Premium == uint64(i) {
				res, err := k.LiquidatePosition(ctx, data.lp, data.id)
				if res {
					return nil
				}
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

/* #nosec */
func (k Keeper) LiquidatePosition(ctx sdk.Context, liqPosition types.LiquidatorPosition, posId string) (bool, error) {
	pos, _ := k.GetPositionByPositionId(ctx, posId)
	collateralAmountInt, collateralDenom, err := k.GetAmountIntFromCoins(pos.Collateral)
	if err != nil {
		return false, err
	}

	liqPositioAmountInt, liqPosDenom, err := k.GetAmountIntFromCoins(liqPosition.Amount)
	if err != nil {
		return false, err
	}

	price, err := k.GetPriceByDenom(ctx, pos.OracleTicker)
	if err != nil {
		return false, err
	}

	rr, err := k.CalculateRiskRate(collateralAmountInt, price, sdk.NewIntFromUint64(pos.BorrowedAmountInUSD))
	//fmt.Printf("Risk Ratio: %d\n", rr.Int64())
	if err != nil {
		return false, err
	}

	if rr.LTE(sdk.NewInt(95)) {
		return true, nil
	}

	liquidator, err := sdk.AccAddressFromBech32(liqPosition.Liquidator)
	if err != nil {
		return false, err
	}

	collateralAmountInUsq := (collateralAmountInt.Mul(price)).QuoRaw(10000)
	amtLiquidate := k.CalculateAmountLiquidate(ctx, collateralAmountInUsq, sdk.NewIntFromUint64(pos.BorrowedAmountInUSD))
	if amtLiquidate.LTE(liqPositioAmountInt) {
		usdAmount, assetAmount := k.CalculatePremiumAmount(ctx, amtLiquidate, price, int64(liqPosition.Premium))
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, liquidator, k.FastCoins(collateralDenom, assetAmount))
		if err != nil {
			return false, err
		}
		pos.BorrowedAmountInUSD = pos.BorrowedAmountInUSD - usdAmount.Uint64()
		pos.Collateral = k.FastCoins(collateralDenom, collateralAmountInt.Sub(assetAmount)).String()
		liqPosition.Amount = k.FastCoins(liqPosDenom, liqPositioAmountInt.Sub(usdAmount)).String()
	} else if amtLiquidate.GT(liqPositioAmountInt) {
		usdAmount, assetAmount := k.CalculatePremiumAmount(ctx, liqPositioAmountInt, price, int64(liqPosition.Premium))
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, liquidator, k.FastCoins(collateralDenom, assetAmount))
		if err != nil {
			return false, err
		}
		pos.BorrowedAmountInUSD = pos.BorrowedAmountInUSD - usdAmount.Uint64()
		pos.Collateral = k.FastCoins(collateralDenom, collateralAmountInt.Sub(assetAmount)).String()
		k.RemoveLiquidatorPosition(ctx, liqPosition.Id)
	}
	k.SetPosition(ctx, pos)
	return false, nil
}
