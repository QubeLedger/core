package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ExecuteLend(ctx sdk.Context, msg *types.MsgCreateLend, LendAsset types.LendAsset) (error, sdk.Coin, string) {

	borrower, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	if err := k.CheckIfPositionAlredyCreate(ctx, borrower.String(), msg.DenomIn); err == nil {
		return err, sdk.Coin{}, ""
	}

	position, found := k.GetPositionByPositionId(ctx, k.CalculateDepositId(borrower.String(), msg.DenomIn))
	if !found {
		return types.ErrPositionNotFound, sdk.Coin{}, ""
	}

	amountPositionCoins, err := sdk.ParseCoinsNormalized(position.Amount)
	if err != nil {
		return err, sdk.Coin{}, ""
	}
	amountPositionInt := amountPositionCoins.AmountOf(msg.DenomIn)

	desiredAmountInt, b := sdk.NewIntFromString(msg.DesiredAmount)
	if !b {
		return types.ErrSdkIntError, sdk.Coin{}, ""
	}
	desiredAmountCoin := sdk.NewCoin(types.DefaultDenom, desiredAmountInt)

	desiredAmountCoins := sdk.NewCoins(desiredAmountCoin)

	price, err := k.GetPriceByDenom(ctx, position.OracleTicker)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	if !sdk.NewIntFromUint64(position.BorrowedAmountInUSD).IsZero() {
		collateral := (amountPositionInt.Mul(price)).QuoRaw(10000)
		err = k.CheckRiskRate(collateral, price, sdk.NewIntFromUint64(position.BorrowedAmountInUSD), desiredAmountInt)
		if err != nil {
			return err, sdk.Coin{}, ""
		}
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, k.GetGrowStakingReserveAddress(ctx), types.ModuleName, desiredAmountCoins)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, desiredAmountCoins)
	if err != nil {
		return err, sdk.Coin{}, ""
	}

	loanId := k.GenerateLoadIdHash(msg.DenomIn, types.DefaultDenom, desiredAmountCoins.String(), borrower.String(), ctx.BlockTime().Format(""))

	loan := types.Loan{
		LoanId:       k.GenerateLoadIdHash(msg.DenomIn, types.DefaultDenom, desiredAmountCoins.String(), borrower.String(), ctx.BlockTime().Format("")),
		Borrower:     borrower.String(),
		AmountOut:    desiredAmountCoins.String(),
		StartTime:    uint64(ctx.BlockTime().Unix()),
		OracleTicker: position.OracleTicker,
	}

	k.AppendLoan(ctx, loan)
	position = k.PushLoanToPosition(ctx, loanId, position)
	position = k.IncreaseBorrowedAmountInUSDInPosition(ctx, position, desiredAmountInt)
	k.SetPosition(ctx, position)

	return nil, desiredAmountCoin, loanId
}

func (k Keeper) ExecuteDeleteLend(ctx sdk.Context, msg *types.MsgDeleteLend, LendAsset types.LendAsset) (error, string) {
	amountInCoins, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil {
		return err, ""
	}

	borrower, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		return err, ""
	}

	if err := CheckCoinsLen(amountInCoins, 1); err != nil {
		return err, ""
	}

	if amountInCoins.GetDenomByIndex(0) != types.DefaultDenom {
		return types.ErrNeedSendUSQ, ""
	}

	amountInInt := amountInCoins.AmountOf(types.DefaultDenom)

	loan, found := k.GetLoadByLoadId(ctx, msg.LoanId)
	if !found {
		return types.ErrLoanNotFound, ""
	}

	position, found := k.GetPositionByPositionId(ctx, k.CalculateDepositId(borrower.String(), msg.DenomOut))
	if !found {
		return types.ErrPositionNotFound, ""
	}

	found = k.CheckLoanIdInPosition(ctx, loan.LoanId, position)
	if !found {
		return types.ErrLoanNotFoundInPosition, ""
	}

	borrowAmountCoins, err := sdk.ParseCoinsNormalized(loan.AmountOut)
	if err != nil {
		return err, ""
	}

	borrowTime := sdk.NewInt(ctx.BlockTime().Unix() - int64(loan.StartTime))
	borrowAmountInt := borrowAmountCoins.AmountOf(types.DefaultDenom)

	rightAmount := k.CalculateNeedAmountToGet(borrowAmountInt, borrowTime)

	if !amountInInt.GTE(rightAmount) {
		return types.ErrNotEnoughAmountIn, ""
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, amountInCoins)
	if err != nil {
		return err, ""
	}

	amtToReserves := (rightAmount.Sub(borrowAmountInt)).QuoRaw(2)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, k.GetUSQReserveAddress(ctx), sdk.NewCoins(sdk.NewCoin(types.DefaultDenom, amtToReserves)))
	if err != nil {
		return err, ""
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, k.stableKeeper.GetBurningFundAddress(ctx), sdk.NewCoins(sdk.NewCoin(types.DefaultDenom, amtToReserves)))
	if err != nil {
		return err, ""
	}

	k.RemoveLoan(ctx, loan.Id)
	position = k.RemoveLoanInPosition(ctx, loan.LoanId, position)
	position = k.ReduceBorrowedAmountInUSDInPosition(ctx, position, borrowAmountInt)
	k.SetPosition(ctx, position)

	return nil, loan.LoanId
}
