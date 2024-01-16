package keeper

import (
	dexmoduletypes "github.com/QuadrateOrg/core/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

func NewACoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenA", amt)
}

func NewBCoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenB", amt)
}

func FundAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, dexmoduletypes.ModuleName, amounts); err != nil {
		return err
	}

	return bankKeeper.SendCoinsFromModuleToAccount(ctx, dexmoduletypes.ModuleName, addr, amounts)
}
