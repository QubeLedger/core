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

	vault, return_amount, err := k.UpdateVaultByTradeType(ctx, vault, msg.Leverage, collateral_coins.AmountOf(vault.AmountXMetadata.Base), msg.TradeType)
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

	switch msg.TradeType {
	case types.PerpetualTradeType_PERPETUAL_LONG_POSITION:
		vault.LongPosition = append(vault.LongPosition, position)
	case types.PerpetualTradeType_PERPETUAL_SHORT_POSITION:
		vault.ShortPosition = append(vault.ShortPosition, position)
	}

	k.SetVault(ctx, vault)

	return nil
}
