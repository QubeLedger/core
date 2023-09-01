package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ExecuteLend(ctx sdk.Context, msg *types.MsgCreateLend, borrowAsset types.BorrowAsset) (error, sdk.Coin, string) {
	amountInCoins, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	borrower, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	amountInInt := amountInCoins.AmountOf(borrowAsset.AmountInAssetMetadata.Base)

	price, err := k.GetPriceByDenom(ctx, borrowAsset.OracleAssetId)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	amountOutInt := k.CalculateCreateLendAmountOut(amountInInt, price)
	amountOutCoins := sdk.NewCoins(sdk.NewCoin(borrowAsset.AmountOutAssetMetadata.Base, amountOutInt))

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, amountInCoins)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, k.GetGrowStakingReserveAddress(ctx), types.ModuleName, amountOutCoins)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, amountOutCoins)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	newLoan := types.Loan{
		LoanId:       k.GenerateLoadIdHash(amountInCoins.GetDenomByIndex(0), msg.DenomOut, msg.AmountIn, msg.Borrower, ctx.BlockTime().Format("")),
		Amount:       msg.AmountIn,
		Borrower:     msg.Borrower,
		StartTime:    uint64(ctx.BlockTime().Unix()),
		OracleTicker: borrowAsset.OracleAssetId,
		AmountOut:    sdk.NewCoin(borrowAsset.AmountOutAssetMetadata.Base, amountOutInt).String(),
		Liquidation:  false,
		Liquidator:   "",
		Hf:           0,
	}

	k.AppendLoan(ctx, newLoan)

	return nil, sdk.Coin{}, newLoan.LoanId
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
