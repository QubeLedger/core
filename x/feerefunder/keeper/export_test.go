package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/QuadrateOrg/core/x/feerefunder/types"
)

func (k Keeper) CheckFees(ctx sdk.Context, fees types.Fee) error {
	return k.checkFees(ctx, fees)
}
