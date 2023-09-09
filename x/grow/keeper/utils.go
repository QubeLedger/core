package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	stabletypes "github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
Coins Helpers
*/
func CheckCoinsLen(coins sdk.Coins, amt int) error {
	if coins.Len() != amt {
		return types.ErrCoinsLen
	}
	return nil
}

func CheckCoinDenom(coins sdk.Coins, denom string) error {
	if err := CheckCoinsLen(coins, 1); err != nil {
		return err
	}
	if coins.GetDenomByIndex(0) != denom {
		return types.ErrDenomsNotEqual
	}
	return nil
}

func (k Keeper) GetAmountIntFromCoins(coins string) (sdk.Int, string, error) {
	amountPositionCoins, err := sdk.ParseCoinsNormalized(coins)
	if err != nil {
		return sdk.Int{}, "", err
	}
	amountPositionInt := amountPositionCoins.AmountOf(amountPositionCoins.GetDenomByIndex(0))
	return amountPositionInt, amountPositionCoins.GetDenomByIndex(0), nil
}

func (k Keeper) FastCoins(denom string, amt sdk.Int) sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin(denom, amt))
}

/*
Deposit Helpers
*/
func (k Keeper) CheckDepositAmount(ctx sdk.Context, msgAmountIn string, pair types.GTokenPair) error {
	msgAmountInCoins, err := sdk.ParseCoinsNormalized(msgAmountIn)
	if err != nil {
		return err
	}

	qStablePair, found := k.stableKeeper.GetPairByPairID(ctx, pair.QStablePairId)
	if !found {
		return stabletypes.ErrPairNotFound
	}

	pairMinAmountInCoins, err := sdk.ParseCoinsNormalized(pair.MinAmountIn)
	if err != nil {
		return err
	}

	if msgAmountInCoins.AmountOf(qStablePair.AmountOutMetadata.Base).LT(pairMinAmountInCoins.AmountOf(qStablePair.AmountOutMetadata.Base)) {
		return types.ErrAmountInGTEminAmountIn
	}

	return nil
}

func (k Keeper) CheckWithdrawalAmount(msgAmountIn string, pair types.GTokenPair) error {
	msgAmountOutCoins, err := sdk.ParseCoinsNormalized(msgAmountIn)
	if err != nil {
		return err
	}

	pairMinAmountoutCoins, err := sdk.ParseCoinsNormalized(pair.MinAmountOut)
	if err != nil {
		return err
	}
	if !msgAmountOutCoins.AmountOf(pair.GTokenMetadata.Base).GT(pairMinAmountoutCoins.AmountOf(pair.GTokenMetadata.Base)) {
		return types.ErrAmountOutGTEminAmountOut
	}

	return nil
}

/*
Lend Helpers
*/
func (k Keeper) CheckOracleAssetId(ctx sdk.Context, LendAsset types.LendAsset) error {
	denomList := k.oracleKeeper.Whitelist(ctx)
	for _, dl := range denomList {
		if dl.Name == LendAsset.OracleAssetId {
			return nil
		}
	}
	return types.ErrOracleAssetIdNotFound
}

/*
EndBlocker Helpers
*/
func (k Keeper) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) GetPairByPairID(ctx sdk.Context, id string) (stabletypes.Pair, bool) {
	return k.stableKeeper.GetPairByPairID(ctx, id)
}

/*
Liquidations Helpers
*/

func (k Keeper) ParseAndCheckPremium(amount string) (sdk.Int, error) {
	amtInt, suc := sdk.NewIntFromString(amount)
	if !suc {
		return sdk.Int{}, types.ErrSdkIntError
	}
	if amtInt.IsZero() || amtInt.IsNil() || amtInt.IsNegative() {
		return sdk.Int{}, types.ErrIntNegativeOrZero
	}
	return amtInt, nil

}

func (k Keeper) CheckLiquidator(liq sdk.Address, pos types.LiquidatorPosition) error {
	if pos.Liquidator != liq.String() {
		return types.ErrLiquidatorAddresesNotEqual
	}
	liqPosId := k.GenerateLiquidatorPositionId(liq.String(), types.DefaultDenom, pos.Amount)
	if pos.LiquidatorPositionId != liqPosId {
		return types.ErrLiquidatorPositionIdNotEqual
	}
	return nil
}

/*
Price Helpers
*/
func (k Keeper) GetPriceByDenom(ctx sdk.Context, denom string) (sdk.Int, error) {
	price, err := k.oracleKeeper.GetExchangeRate(ctx, denom)
	if err != nil {
		return sdk.Int{}, err
	}
	if price.IsNil() {
		return sdk.Int{}, types.ErrPriceNil
	}
	priceInt := price.MulInt64(10000).RoundInt()
	return priceInt, nil
}

/*
Reserve Helpers
*/

func (k Keeper) IncreaseGrowStakingReserve(ctx sdk.Context, amountIn sdk.Coins, gTokenPair types.GTokenPair, qStablePair stabletypes.Pair) (types.GTokenPair, error) {
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, k.GetGrowStakingReserveAddress(ctx), amountIn)
	if err != nil {
		return gTokenPair, err
	}

	gTokenPair.St = gTokenPair.St.Add(amountIn.AmountOf(qStablePair.AmountOutMetadata.Base))

	return gTokenPair, nil
}

func (k Keeper) ReduceGrowStakingReserve(ctx sdk.Context, amountIn sdk.Coins, gTokenPair types.GTokenPair) (types.GTokenPair, error) {
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, k.GetGrowStakingReserveAddress(ctx), types.ModuleName, amountIn)
	if err != nil {
		return gTokenPair, err
	}

	gTokenPair.St = gTokenPair.St.Sub(amountIn.AmountOf(gTokenPair.GTokenMetadata.Base))

	return gTokenPair, nil
}
