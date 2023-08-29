package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	stabletypes "github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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

	if !msgAmountInCoins.AmountOf(qStablePair.AmountOutMetadata.Base).GT(pairMinAmountInCoins.AmountOf(qStablePair.AmountOutMetadata.Base)) {
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
