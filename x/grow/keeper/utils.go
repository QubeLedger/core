package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CheckDepositAmount(msgAmountIn string, pair types.Pair) error {
	msgAmountInCoins, err := sdk.ParseCoinsNormalized(msgAmountIn)
	if err != nil {
		return err
	}

	pairMinAmountInCoins, err := sdk.ParseCoinsNormalized(pair.MinAmountIn)
	if err != nil {
		return err
	}
	if !msgAmountInCoins.AmountOf(pair.AmountInMetadata.Base).GT(pairMinAmountInCoins.AmountOf(pair.AmountInMetadata.Base)) {
		return types.ErrAmountInGTEminAmountIn
	}

	return nil
}

func (k Keeper) CheckWithdrawalAmount(msgAmountIn string, pair types.Pair) error {
	msgAmountOutCoins, err := sdk.ParseCoinsNormalized(msgAmountIn)
	if err != nil {
		return err
	}

	pairMinAmountoutCoins, err := sdk.ParseCoinsNormalized(pair.MinAmountOut)
	if err != nil {
		return err
	}
	if !msgAmountOutCoins.AmountOf(pair.AmountOutMetadata.Base).GT(pairMinAmountoutCoins.AmountOf(pair.AmountOutMetadata.Base)) {
		return types.ErrAmountOutGTEminAmountOut
	}

	return nil
}
