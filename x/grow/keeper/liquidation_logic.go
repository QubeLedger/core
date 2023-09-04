package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ExecuteLiquidation(ctx sdk.Context, liquidatorList []types.LiquidatorPosition, pos types.Position) error {
	for i := 1; i < 100; i++ {
		for _, lp := range liquidatorList {
			if lp.Premium == uint64(i) {
				//k.LiquidatePosition()
			}
		}
	}
	return nil
}

func (k Keeper) LiquidatePosition(ctx sdk.Context, liqPosition types.LiquidatorPosition, pos types.Position) bool {
	amountPositionCoins, _ := sdk.ParseCoinsNormalized(pos.Amount)
	amountPositionInt := amountPositionCoins.AmountOf(types.DefaultDenom)
	price, _ := k.GetPriceByDenom(ctx, pos.OracleTicker)
	liquidateAmount := k.CalculateAmountWhichLiquidatorSend(ctx, pos.BorrowedAmountInUSD, amountPositionInt, price, liqPosition.Premium)
	if liquidateAmount.GT(amountPositionInt) {
		return false
	}
	return true
}
