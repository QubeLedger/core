package keeper

import (
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	AtomPrice   sdk.Int
	TestingMode bool = false
)

func (k Keeper) UpdateAtomPrice(ctx sdk.Context) error {
	if TestingMode {
		return nil
	}
	if AtomPrice.IsNil() {
		AtomPrice = sdk.NewInt(0)
	}
	atomPrice, _, err := k.oracleKeeper.GetTokensActualPriceInt(ctx)
	if err != nil {
		return err
	}
	if atomPrice.IsNil() {
		return types.ErrAtomPriceNil
	}
	AtomPrice = atomPrice
	return nil
}

func (k Keeper) UpdateAtomPriceTesting(ctx sdk.Context, price sdk.Int) error {
	if !TestingMode {
		return nil
	}
	AtomPrice = price
	return nil
}

func (k Keeper) GetAtomPrice(ctx sdk.Context) sdk.Int {
	return AtomPrice
}

func (k Keeper) SetTestingMode(value bool) {
	TestingMode = value
}
