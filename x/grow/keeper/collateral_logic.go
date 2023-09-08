package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ExecuteDepositCollateral(ctx sdk.Context, msg *types.MsgDepositCollateral, LendAsset types.LendAsset) (error, string) {
	amountInCoins, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil {
		return err, ""
	}

	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return err, ""
	}

	if err := CheckCoinsLen(amountInCoins, 1); err != nil {
		return err, ""
	}

	denom := amountInCoins.GetDenomByIndex(0)

	if err := k.CheckIfPositionAlredyCreate(ctx, depositor.String(), denom); err != nil {
		return err, ""
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, depositor, types.ModuleName, amountInCoins)
	if err != nil {
		return err, ""
	}

	position := types.Position{
		Creator:             depositor.String(),
		DepositId:           k.CalculateDepositId(depositor.String(), denom),
		Collateral:          msg.AmountIn,
		OracleTicker:        LendAsset.OracleAssetId,
		BorrowedAmountInUSD: 0,
		LoanIds:             []string{},
	}

	k.AppendPosition(ctx, position)

	return nil, position.DepositId
}

func (k Keeper) ExecuteWithdrawalCollateral(ctx sdk.Context, msg *types.MsgWithdrawalCollateral, LendAsset types.LendAsset) (error, sdk.Coin) {
	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return err, sdk.Coin{}
	}

	denom := msg.Denom

	position, found := k.GetPositionByPositionId(ctx, k.CalculateDepositId(depositor.String(), denom))
	if !found {
		return types.ErrPositionNotFound, sdk.Coin{}
	}

	if !sdk.NewIntFromUint64(position.BorrowedAmountInUSD).IsZero() {
		return types.ErrRiskRatioMustBeZero, sdk.Coin{}
	}

	amountOut, err := sdk.ParseCoinsNormalized(position.Collateral)
	if err != nil {
		return err, sdk.Coin{}
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositor, amountOut)
	if err != nil {
		return err, sdk.Coin{}
	}

	k.RemovePosition(ctx, position.Id)

	return nil, sdk.NewCoin(denom, amountOut.AmountOf(denom))
}
