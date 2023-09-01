package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetPriceByDenom(ctx sdk.Context, denom string) (sdk.Int, error) {
	price, err := k.oracleKeeper.GetExchangeRate(ctx, denom)
	if err != nil {
		return sdk.Int{}, err
	}
	if price.IsNil() {
		return sdk.Int{}, types.ErrPriceNil
	}
	priceInt := price.MulInt64(10000).RoundInt()
	return priceInt, nil
}
