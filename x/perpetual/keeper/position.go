package keeper

import (
	"github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateNewPosition(ctx sdk.Context, msg *types.MsgOpen, vault types.Vault) error {

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

	position := types.TradePosition{
		TradePositionId:  k.GenerateTraderPositionId(msg.Creator, vault.AmountXMetadata.Base, msg.TradingAsset, msg.TradeType),
		Creator:          msg.Creator,
		TradeType:        msg.TradeType,
		Leverage:         msg.Leverage,
		TradingAsset:     msg.TradingAsset,
		CollateralAmount: collateral_coins.AmountOf(vault.AmountXMetadata.Base),
		CollateralDenom:  vault.AmountXMetadata.Base,
		ReturnAmount:     return_amount,
		ReturnDenom:      vault.AmountYMetadata.Base,
	}

	_ = k.AppendPosition(ctx, position)

	/*switch msg.TradeType {
	case types.PerpetualTradeType_PERPETUAL_LONG_POSITION:
		vault.LongPosition = append(vault.LongPosition, position)
	case types.PerpetualTradeType_PERPETUAL_SHORT_POSITION:
		vault.ShortPosition = append(vault.ShortPosition, position)
	}*/

	k.SetVault(ctx, vault)

	return nil
}

func (k Keeper) CloseOrDecreasePosition(ctx sdk.Context, msg *types.MsgClose, vault types.Vault, position types.TradePosition) error {

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}

	if msg.Amount.GT(position.ReturnAmount) {
		return types.ErrAmountGreaterThanPositionSize
	}

	vault, return_amount, err := k.UpdateVaultByTradeTypeWhenClosePosition(ctx, vault, msg.Amount, position)
	if err != nil {
		return err
	}

	if return_amount.IsNil() {
		return types.ErrInCalculationUpdateVault
	}

	if msg.Amount.Equal(position.ReturnAmount) {

		/*switch position.TradeType {
		case types.PerpetualTradeType_PERPETUAL_LONG_POSITION:
			vault = k.RemoveLongFromVault(ctx, position.TradePositionId, vault)
		case types.PerpetualTradeType_PERPETUAL_SHORT_POSITION:
			vault = k.RemoveShortFromVault(ctx, position.TradePositionId, vault)
		}*/

		return_coins := sdk.NewCoins(
			sdk.NewCoin(
				position.CollateralDenom,
				return_amount.Quo(position.Leverage.RoundInt()),
			),
		)

		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, return_coins)
		if err != nil {
			return err
		}

		k.RemovePosition(ctx, position.Id)
	} else {
		position.ReturnAmount = position.ReturnAmount.Sub(msg.Amount)
		position.CollateralAmount = position.CollateralAmount.Add(return_amount.Sub(position.CollateralAmount))
		k.SetPosition(ctx, position)
	}

	k.SetVault(ctx, vault)
	return nil
}
