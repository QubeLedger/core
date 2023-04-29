package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/feeburner/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) TotalBurnedNeutronsAmount(goCtx context.Context, _ *types.QueryTotalBurnedNeutronsAmountRequest) (*types.QueryTotalBurnedNeutronsAmountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	totalBurnedNeutronsAmount := k.GetTotalBurnedNeutronsAmount(ctx)

	return &types.QueryTotalBurnedNeutronsAmountResponse{TotalBurnedNeutronsAmount: totalBurnedNeutronsAmount}, nil
}
