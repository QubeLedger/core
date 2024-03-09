package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/* #nosec */
func (k Keeper) ExecuteCreateLend(ctx sdk.Context, msg *types.MsgCreateLend, Asset types.Asset) (error, string) {
	AmountInCoins, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil {
		return err, ""
	}

	if k.AddressEmptyCheck(ctx) {
		return types.ErrReserveAddressEmpty, ""
	}

	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return err, ""
	}

	if err := CheckCoinsLen(AmountInCoins, 1); err != nil {
		return err, ""
	}

	DenomIn := AmountInCoins.GetDenomByIndex(0)

	PositionId := k.CalculateDepositId(depositor.String())

	if _, found := k.GetPositionByPositionId(ctx, PositionId); !found {
		position := types.Position{
			Creator:             depositor.String(),
			DepositId:           k.CalculateDepositId(depositor.String()),
			LendId:              []string{},
			LendAmountInUSD:     0,
			BorrowedAmountInUSD: 0,
			LoanId:              []string{},
		}

		k.AppendPosition(ctx, position)
	}

	price, err := k.GetPriceByDenom(ctx, Asset.OracleAssetId)
	if err != nil {
		return err, ""
	}

	position, _ := k.GetPositionByPositionId(ctx, PositionId)
	found := k.CheckLendIdInPosition(ctx, k.CalculateLendId(depositor.String(), DenomIn, PositionId), position)
	if !found {
		lend := types.Lend{
			LendId:         k.CalculateLendId(depositor.String(), DenomIn, PositionId),
			Borrower:       depositor.String(),
			AmountIn:       msg.AmountIn,
			AmountInAmount: sdk.NewDecFromInt(AmountInCoins.AmountOf(DenomIn)),
			AmountInDenom:  DenomIn,
			StartTime:      uint64(ctx.BlockTime().Unix()),
			OracleTicker:   Asset.OracleAssetId,
		}
		k.AppendLend(ctx, lend)
		position = k.PushLendToPosition(ctx, lend.LendId, position)
	} else {
		old_lend, _ := k.GetLendByLendId(ctx, k.CalculateLendId(depositor.String(), DenomIn, PositionId))
		lend := types.Lend{
			LendId:         k.CalculateLendId(depositor.String(), DenomIn, PositionId),
			Borrower:       depositor.String(),
			AmountIn:       msg.AmountIn,
			AmountInAmount: old_lend.AmountInAmount.Add(sdk.NewDecFromInt(AmountInCoins.AmountOf(DenomIn))),
			AmountInDenom:  DenomIn,
			StartTime:      uint64(ctx.BlockTime().Unix()),
			OracleTicker:   Asset.OracleAssetId,
		}
		k.SetLend(ctx, lend)
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, depositor, types.ModuleName, AmountInCoins)
	if err != nil {
		return err, ""
	}

	position.LendAmountInUSD += k.CalculateAmountByPriceAndAmountIn(AmountInCoins.AmountOf(DenomIn), price).Uint64()
	k.SetPosition(ctx, position)

	Asset.ProvideValue += (AmountInCoins.AmountOf(DenomIn)).Uint64()
	k.SetAsset(ctx, Asset)

	return nil, position.DepositId
}

/* #nosec */
func (k Keeper) ExecuteWithdrawalLend(ctx sdk.Context, msg *types.MsgWithdrawalLend, Asset types.Asset) (error, sdk.Coin) {
	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return err, sdk.Coin{}
	}

	if k.AddressEmptyCheck(ctx) {
		return types.ErrReserveAddressEmpty, sdk.Coin{}
	}

	amountInCoins, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil {
		return err, sdk.Coin{}
	}

	price, err := k.GetPriceByDenom(ctx, Asset.OracleAssetId)
	if err != nil {
		return err, sdk.Coin{}
	}

	DenomIn := amountInCoins.GetDenomByIndex(0)
	amountInInt := amountInCoins.AmountOf(DenomIn)

	position, found := k.GetPositionByPositionId(ctx, k.CalculateDepositId(depositor.String()))
	if !found {
		return types.ErrPositionNotFound, sdk.Coin{}
	}

	lend, found := k.GetLendByLendId(ctx, k.CalculateLendId(depositor.String(), DenomIn, position.DepositId))
	if !found {
		return types.ErrLendNotFound, sdk.Coin{}
	}

	if amountInInt.GTE(lend.AmountInAmount.RoundInt()) {
		if position.BorrowedAmountInUSD != 0 {
			return types.ErrRiskRatioMustBeZero, sdk.Coin{}
		}
		k.RemoveLend(ctx, lend.Id)
		position = k.RemoveLendInPosition(ctx, lend.LendId, position)
	} else {
		new_lend := types.Lend{
			LendId:         lend.LendId,
			Borrower:       depositor.String(),
			AmountIn:       lend.AmountIn,
			AmountInAmount: lend.AmountInAmount.Sub(sdk.NewDecFromInt(amountInCoins.AmountOf(DenomIn))),
			AmountInDenom:  DenomIn,
			StartTime:      uint64(ctx.BlockTime().Unix()),
			OracleTicker:   Asset.OracleAssetId,
		}
		k.SetLend(ctx, new_lend)
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositor, amountInCoins)
	if err != nil {
		return err, sdk.Coin{}
	}

	reduceLendAmountInUSD := k.CalculateAmountByPriceAndAmountIn(amountInCoins.AmountOf(DenomIn), price).Uint64()

	if reduceLendAmountInUSD >= position.LendAmountInUSD {
		position.LendAmountInUSD = uint64(0)
	} else {
		position.LendAmountInUSD -= reduceLendAmountInUSD
	}
	k.SetPosition(ctx, position)

	reduceProvideValue := (amountInCoins.AmountOf(DenomIn)).Uint64()
	if reduceProvideValue >= Asset.ProvideValue {
		Asset.ProvideValue = uint64(0)
	} else {
		Asset.ProvideValue -= (amountInCoins.AmountOf(DenomIn)).Uint64()
	}
	k.SetAsset(ctx, Asset)

	return nil, sdk.NewCoin(DenomIn, amountInCoins.AmountOf(DenomIn))
}
