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
