package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ExecuteDeposit(ctx sdk.Context, msg *types.MsgDeposit, pair types.Pair) (error, sdk.Coin) {
	return nil, sdk.Coin{}
}

func (k Keeper) ExecuteWithdrawal(ctx sdk.Context, msg *types.MsgWithdrawal, pair types.Pair) (error, sdk.Coin) {
	return nil, sdk.Coin{}
}
