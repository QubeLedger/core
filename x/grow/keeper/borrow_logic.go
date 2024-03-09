package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/* #nosec */
func (k Keeper) ExecuteCreateBorrow(ctx sdk.Context, msg *types.MsgCreateBorrow, Asset types.Asset) (error, sdk.Coin, string) {

	borrower, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	desiredAmountCoins, err := sdk.ParseCoinsNormalized(msg.DesiredAmount)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	DenomIn := desiredAmountCoins.GetDenomByIndex(0)

	if desiredAmountCoins.AmountOf(DenomIn).Uint64() > Asset.ProvideValue {
		return types.ErrNotEnoughProvideValue, sdk.Coin{}, ""
	}

	if k.AddressEmptyCheck(ctx) {
		return types.ErrReserveAddressEmpty, sdk.Coin{}, ""
	}

	position, found := k.GetPositionByPositionId(ctx, k.CalculateDepositId(borrower.String()))
	if !found {
		return types.ErrPositionNotFound, sdk.Coin{}, ""
	}

	price, err := k.GetPriceByDenom(ctx, Asset.OracleAssetId)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	err = k.CheckRiskRate(sdk.NewIntFromUint64(position.LendAmountInUSD), sdk.NewIntFromUint64(position.BorrowedAmountInUSD), desiredAmountCoins.AmountOf(DenomIn).Mul(price).Quo(types.Multiplier))
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	loanId := k.GenerateLoanIdHash(DenomIn, borrower.String())
	found = k.CheckLoanIdInPosition(ctx, loanId, position)
	if !found {
		loan := types.Loan{
			LoanId:          loanId,
			Borrower:        borrower.String(),
			AmountOut:       desiredAmountCoins.String(),
			AmountOutAmount: sdk.NewDecFromInt(desiredAmountCoins.AmountOf(DenomIn)),
			AmountOutDenom:  DenomIn,
			StartTime:       uint64(ctx.BlockTime().Unix()),
			OracleTicker:    Asset.OracleAssetId,
		}
		k.AppendLoan(ctx, loan)
		position = k.PushLoanToPosition(ctx, loan.LoanId, position)
	} else {
		old_loan, _ := k.GetLoadByLoanId(ctx, loanId)
		loan := types.Loan{
			LoanId:          loanId,
			Borrower:        borrower.String(),
			AmountOut:       desiredAmountCoins.String(),
			AmountOutAmount: old_loan.AmountOutAmount.Add(sdk.NewDecFromInt(desiredAmountCoins.AmountOf(DenomIn))),
			AmountOutDenom:  old_loan.AmountOutDenom,
			StartTime:       uint64(ctx.BlockTime().Unix()),
			OracleTicker:    Asset.OracleAssetId,
		}
		k.SetLoan(ctx, loan)
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, desiredAmountCoins)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	Asset.CollectivelyBorrowValue += desiredAmountCoins.AmountOf(DenomIn).Uint64()
	k.SetAsset(ctx, Asset)

	position.BorrowedAmountInUSD += k.CalculateAmountByPriceAndAmountIn(desiredAmountCoins.AmountOf(DenomIn), price).Uint64()
	k.SetPosition(ctx, position)

	return nil, sdk.NewCoin(DenomIn, desiredAmountCoins.AmountOf(DenomIn)), loanId
}

/* #nosec */
func (k Keeper) ExecuteDeleteBorrow(ctx sdk.Context, msg *types.MsgDeleteBorrow, Asset types.Asset) (error, string) {
	amountInCoins, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil {
		return err, ""
	}

	if err := CheckCoinsLen(amountInCoins, 1); err != nil {
		return err, ""
	}

	if k.AddressEmptyCheck(ctx) {
		return types.ErrReserveAddressEmpty, ""
	}

	borrower, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		return err, ""
	}

	price, err := k.GetPriceByDenom(ctx, Asset.OracleAssetId)
	if err != nil {
		return err, ""
	}

	DenomIn := amountInCoins.GetDenomByIndex(0)

	amountInInt := amountInCoins.AmountOf(DenomIn)

	loanId := k.GenerateLoanIdHash(DenomIn, borrower.String())
	loan, found := k.GetLoadByLoanId(ctx, loanId)
	if !found {
		return types.ErrLoanNotFound, ""
	}

	position, found := k.GetPositionByPositionId(ctx, k.CalculateDepositId(borrower.String()))
	if !found {
		return types.ErrPositionNotFound, ""
	}

	found = k.CheckLoanIdInPosition(ctx, loan.LoanId, position)
	if !found {
		return types.ErrLoanNotFoundInPosition, ""
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, amountInCoins)
	if err != nil {
		return err, ""
	}

	if amountInInt.GTE(loan.AmountOutAmount.RoundInt()) {
		k.RemoveLoan(ctx, loan.Id)
		position = k.RemoveLoanInPosition(ctx, loan.LoanId, position)
	} else {
		new_loan := types.Loan{
			LoanId:          loanId,
			Borrower:        borrower.String(),
			AmountOut:       loan.AmountOut,
			AmountOutAmount: loan.AmountOutAmount.Sub(sdk.NewDecFromInt(amountInCoins.AmountOf(DenomIn))),
			AmountOutDenom:  loan.AmountOutDenom,
			StartTime:       uint64(ctx.BlockTime().Unix()),
			OracleTicker:    Asset.OracleAssetId,
		}
		k.SetLoan(ctx, new_loan)
	}

	reduceCollectivelyBorrowValue := amountInCoins.AmountOf(DenomIn).Uint64()

	if reduceCollectivelyBorrowValue >= Asset.CollectivelyBorrowValue {
		Asset.CollectivelyBorrowValue = uint64(0)
	} else {
		Asset.CollectivelyBorrowValue -= reduceCollectivelyBorrowValue
	}
	k.SetAsset(ctx, Asset)

	reduceAmount := k.CalculateAmountByPriceAndAmountIn(amountInCoins.AmountOf(DenomIn), price).Uint64()

	if reduceAmount >= position.BorrowedAmountInUSD {
		position.BorrowedAmountInUSD = uint64(0)
	} else {
		position.BorrowedAmountInUSD -= reduceAmount
	}

	k.SetPosition(ctx, position)

	return nil, loan.LoanId
}
