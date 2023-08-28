package keeper

import (
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	AtomPrice   sdk.Int
	TestingMode bool = false
)

func (k Keeper) UpdateAtomPrice(ctx sdk.Context, pair types.Pair) error {
	if TestingMode {
		return nil
	}
	if AtomPrice.IsNil() {
		AtomPrice = sdk.NewInt(0)
	}
	price, err := k.oracleKeeper.GetExchangeRate(ctx, pair.AmountInMetadata.Base)
	if err != nil {
		return err
	}
	if price.IsNil() {
		return types.ErrAtomPriceNil
	}
	AtomPrice = price.MulInt64(10000).RoundInt()
	return nil
}

func (k Keeper) UpdateAtomPriceTesting(ctx sdk.Context, price sdk.Int) error {
	if !TestingMode {
		return nil
	}
	AtomPrice = price
	return nil
}

func (k Keeper) GetAtomPrice(ctx sdk.Context, pair types.Pair) (sdk.Int, error) {
	err := k.UpdateAtomPrice(ctx, pair)
	if err != nil {
		return sdk.Int{}, err
	}
	return AtomPrice, err
}

func (k Keeper) SetTestingMode(value bool) {
	TestingMode = value
}
