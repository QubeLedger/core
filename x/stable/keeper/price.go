package keeper

import (
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	AtomPrice sdk.Int
)

func (k Keeper) SetAtomPriceFromOracle(ctx sdk.Context) error {
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

func (k Keeper) SetAtomPriceForTest(ctx sdk.Context, price sdk.Int) {
	AtomPrice = price
}

func (k Keeper) GetAtomPrice(ctx sdk.Context) sdk.Int {
	return AtomPrice
}
