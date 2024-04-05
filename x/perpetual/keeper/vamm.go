package keeper

import (
	"github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) UpdateVaultByTradeTypeWhenOpenPosition(ctx sdk.Context, vault types.Vault, leverage sdk.Dec, amount sdk.Int, trade_type types.PerpetualTradeType) (types.Vault, sdk.Int, error) {
	switch trade_type {
	case types.PerpetualTradeType_PERPETUAL_LONG_POSITION:
		vault, return_amount := k.UpdateVaultWhenOpenLong(ctx, vault, leverage, amount)
		return vault, return_amount, nil
	case types.PerpetualTradeType_PERPETUAL_SHORT_POSITION:
		vault, return_amount := k.UpdateVaultWhenOpenShort(ctx, vault, leverage, amount)
		return vault, return_amount, nil
	}
	return vault, sdk.Int{}, types.ErrSample
}

func (k Keeper) UpdateVaultWhenOpenLong(ctx sdk.Context, vault types.Vault, leverage sdk.Dec, amount sdk.Int) (types.Vault, sdk.Int) {
	user_amount_with_leverage := amount.Mul(leverage.RoundInt())
	vault.X = vault.X.Add(user_amount_with_leverage)
	return_amount := k.CalculateLongReturnAmountByVaultAndInput(ctx, vault, user_amount_with_leverage)
	vault.Y = vault.Y.Sub(return_amount)
	return vault, return_amount
}

func (k Keeper) UpdateVaultWhenOpenShort(ctx sdk.Context, vault types.Vault, leverage sdk.Dec, amount sdk.Int) (types.Vault, sdk.Int) {
	user_amount_with_leverage := amount.Mul(leverage.RoundInt())
	vault.X = vault.X.Sub(user_amount_with_leverage)
	return_amount := k.CalculateShortReturnAmountByVaultAndInput(ctx, vault, user_amount_with_leverage)
	vault.Y = vault.Y.Add(return_amount)
	return vault, return_amount
}

func (k Keeper) UpdateVaultByTradeTypeWhenClosePosition(ctx sdk.Context, vault types.Vault, amount sdk.Int, position types.TradePosition) (types.Vault, sdk.Int, error) {
	switch position.TradeType {
	case types.PerpetualTradeType_PERPETUAL_LONG_POSITION:
		vault, return_amount := k.UpdateVaultWhenCloseLong(ctx, vault, amount, position)
		return vault, return_amount, nil
	case types.PerpetualTradeType_PERPETUAL_SHORT_POSITION:
		vault, return_amount := k.UpdateVaultWhenCloseShort(ctx, vault, amount, position)
		return vault, return_amount, nil
	}
	return vault, sdk.Int{}, types.ErrSample
}

func (k Keeper) UpdateVaultWhenCloseLong(ctx sdk.Context, vault types.Vault, amount sdk.Int, position types.TradePosition) (types.Vault, sdk.Int) {
	vault.Y = vault.Y.Add(position.ReturnAmount)
	return_amount := k.CalculateLongCloseAmountByAmount(ctx, vault, amount)
	vault.X = vault.X.Sub(return_amount)
	return vault, return_amount
}

func (k Keeper) UpdateVaultWhenCloseShort(ctx sdk.Context, vault types.Vault, amount sdk.Int, position types.TradePosition) (types.Vault, sdk.Int) {
	vault.Y = vault.Y.Sub(position.ReturnAmount)
	return_amount := k.CalculateShortCloseAmountByAmount(ctx, vault, amount)
	vault.X = vault.X.Add(return_amount)
	return vault, return_amount
}

func (k Keeper) CalculateLongReturnAmountByVaultAndInput(ctx sdk.Context, vault types.Vault, amount sdk.Int) sdk.Int {

	// (x + amount) * (y - return_amount) = k
	// return_amount = (k / (x + amount)) - y * -1

	x_value := vault.X
	y_value := vault.Y
	k_value := vault.K

	return_amount := ((k_value.Quo(x_value)).Sub(y_value)).MulRaw(-1)
	return return_amount
}

func (k Keeper) CalculateShortReturnAmountByVaultAndInput(ctx sdk.Context, vault types.Vault, amount sdk.Int) sdk.Int {

	// (x + amount) * (y - return_amount) = k
	// return_amount = (k / (x + amount)) - y * -1

	x_value := vault.X
	y_value := vault.Y
	k_value := vault.K

	return_amount := ((k_value.Quo(x_value)).Sub(y_value))
	return return_amount
}

func (k Keeper) CalculateLongCloseAmountByAmount(ctx sdk.Context, vault types.Vault, amount sdk.Int) sdk.Int {

	// (x - return_amount) * (y + amount) = k
	// (k / (y + amount)) - x * -1

	x_value := vault.X
	y_value := vault.Y
	k_value := vault.K

	return_amount := (k_value.Quo(y_value).Sub(x_value)).MulRaw(-1)
	return return_amount
}

func (k Keeper) CalculateShortCloseAmountByAmount(ctx sdk.Context, vault types.Vault, amount sdk.Int) sdk.Int {

	// (x + return_amount) * (y - amount) = k
	// (k / (y + amount)) - x * -1

	x_value := vault.X
	y_value := vault.Y
	k_value := vault.K

	return_amount := (k_value.Quo(y_value).Sub(x_value)).MulRaw(-1)
	return return_amount
}
