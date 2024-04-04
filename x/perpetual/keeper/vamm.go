package keeper

import (
	"github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) UpdateVaultByTradeType(ctx sdk.Context, vault types.Vault, leverage sdk.Dec, amount sdk.Int, trade_type types.PerpetualTradeType) (types.Vault, sdk.Int, error) {
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

func (k Keeper) CalculateLongReturnAmountByVaultAndInput(ctx sdk.Context, vault types.Vault, amount sdk.Int) sdk.Int {

	// (x + amount) * (y - return_amount) = k
	// return_amount = (k / (x + amount)) - y * -1

	x_value := vault.X
	y_value := vault.Y
	k_value := vault.K

	return_amount := ((k_value.Quo(x_value)).Sub(y_value)).Mul(sdk.NewInt(-1))
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
