package keeper

import (
	"github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateNewPosition(ctx sdk.Context, msg *types.MsgOpen, vault types.Vault) error {

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}

	collateral_coins, err := sdk.ParseCoinsNormalized(msg.Collateral)
	if err != nil {
		return err
	}

	vault, return_amount, err := k.UpdateVaultByTradeTypeWhenOpenPosition(ctx, vault, msg.Leverage, collateral_coins.AmountOf(vault.AmountXMetadata.Base), msg.TradeType)
	if err != nil {
		return err
	}

	if return_amount.IsNil() {
		return types.ErrInCalculationUpdateVault
	}

	err = k.BankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, collateral_coins)
	if err != nil {
		return err
	}

	position := types.TradePosition{
		TradePositionId:  k.GenerateTraderPositionId(msg.Creator, vault.AmountXMetadata.Base, msg.TradingAsset, msg.TradeType),
		Creator:          msg.Creator,
		TradeType:        msg.TradeType,
		Leverage:         msg.Leverage,
		TradingAsset:     msg.TradingAsset,
		CollateralAmount: (collateral_coins.AmountOf(vault.AmountXMetadata.Base)).ToDec(),
		CollateralDenom:  vault.AmountXMetadata.Base,
		ReturnAmount:     return_amount.ToDec(),
		ReturnDenom:      vault.AmountYMetadata.Base,
	}

	_ = k.AppendPosition(ctx, position)

	switch msg.TradeType {
	case types.PerpetualTradeType_PERPETUAL_LONG_POSITION:
		vault.LongPositionId = append(vault.LongPositionId, position.TradePositionId)
	case types.PerpetualTradeType_PERPETUAL_SHORT_POSITION:
		vault.ShortPositionId = append(vault.ShortPositionId, position.TradePositionId)
	}

	k.SetVault(ctx, vault)

	return nil
}

func (k Keeper) CloseOrDecreasePosition(ctx sdk.Context, msg *types.MsgClose, vault types.Vault, position types.TradePosition) error {

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}

	if (msg.Amount.ToDec()).GT(position.ReturnAmount) {
		return types.ErrAmountGreaterThanPositionSize
	}

	vault, return_amount, err := k.UpdateVaultByTradeTypeWhenClosePosition(ctx, vault, msg.Amount, position)
	if err != nil {
		return err
	}

	if return_amount.IsNil() {
		return types.ErrInCalculationUpdateVault
	}

	position.ProfitAmount = position.ProfitAmount.Add(return_amount.ToDec())

	if (msg.Amount.ToDec()).Equal(position.ReturnAmount) {

		if position.TradeType == types.PerpetualTradeType_PERPETUAL_LONG_POSITION {
			if position.ProfitAmount.GT(position.CollateralAmount) {
				return_amount_plus := position.ProfitAmount.Sub(position.CollateralAmount.Mul(position.Leverage))
				return_amount = (position.CollateralAmount.Add(return_amount_plus)).RoundInt()
			} else {
				return_amount = (position.ProfitAmount.Quo(position.Leverage)).RoundInt()
			}

			return_coins := sdk.NewCoins(
				sdk.NewCoin(
					position.CollateralDenom,
					return_amount,
				),
			)

			err := k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, return_coins)
			if err != nil {
				return err
			}

			vault = k.RemoveLongFromVault(ctx, position.TradePositionId, vault)
			k.RemovePosition(ctx, position.Id)
		} else if position.TradeType == types.PerpetualTradeType_PERPETUAL_SHORT_POSITION {

			if position.ProfitAmount.GT(position.CollateralAmount) {
				return_amount_plus := position.ProfitAmount.Sub(position.CollateralAmount.Mul(position.Leverage))
				return_amount = (position.CollateralAmount.Sub(return_amount_plus)).RoundInt()
			} else {

				return_amount_plus := (position.CollateralAmount.Mul(position.Leverage)).Sub(position.ProfitAmount)
				return_amount = (position.CollateralAmount.Add(return_amount_plus)).RoundInt()
			}

			return_coins := sdk.NewCoins(
				sdk.NewCoin(
					position.CollateralDenom,
					return_amount,
				),
			)

			err := k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, return_coins)
			if err != nil {
				return err
			}

			vault = k.RemoveShortFromVault(ctx, position.TradePositionId, vault)
			k.RemovePosition(ctx, position.Id)
		}
	} else {
		position.ReturnAmount = position.ReturnAmount.Sub(msg.Amount.ToDec())
		k.SetPosition(ctx, position)
	}

	k.SetVault(ctx, vault)
	return nil
}
