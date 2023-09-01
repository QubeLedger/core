package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ExecuteLend(ctx sdk.Context, msg *types.MsgCreateLend, borrowAsser types.BorrowAsset) (error, sdk.Coin, string) {
	amountInCoins, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	borrower, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	_ = borrower
	_ = amountInCoins
	return nil, sdk.Coin{}, ""
}

func (k Keeper) ExecuteDeleteLend(ctx sdk.Context, msg *types.MsgDeleteLend, borrowAsser types.BorrowAsset) (error, sdk.Coin, string) {
	amountInCoins, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	borrower, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	_ = borrower
	_ = amountInCoins
	return nil, sdk.Coin{}, ""
}
