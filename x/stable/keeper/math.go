package keeper

import (
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CalculateAmountToMint(ibcBaseTokenDenomAmount sdk.Int, atomPrice sdk.Int, mintingFee sdk.Int) sdk.Int {
	return (ibcBaseTokenDenomAmount.Mul(atomPrice).Sub((ibcBaseTokenDenomAmount.Mul(atomPrice).Mul(mintingFee)).Quo(types.MintUsqMultiplier))).Quo(types.Multiplier)
}

func (k Keeper) CalculateMintingFeeForBurningFund(ibcBaseTokenDenomAmount sdk.Int, atomPrice sdk.Int, mintingFee sdk.Int) sdk.Int {
	return (((ibcBaseTokenDenomAmount.Mul(atomPrice)).Mul(mintingFee)).Quo(types.MintUsqMultiplier)).Quo(types.Multiplier)
}

func (k Keeper) CalculateAmountToSend(qAssetTokenAmount sdk.Int, atomPrice sdk.Int, burningFee sdk.Int) sdk.Int {
	fee := (((qAssetTokenAmount.Mul(types.BurnUsqMultiplier)).Quo(atomPrice)).Mul(burningFee)).Quo(types.MintUsqMultiplier)
	return (((qAssetTokenAmount.Mul(types.BurnUsqMultiplier)).Quo(atomPrice)).Sub(fee)).Quo(types.FeeMultiplier)
}

func (k Keeper) CalculateBurningFeeForBurningFund(qAssetTokenAmount sdk.Int, atomPrice sdk.Int, burningFee sdk.Int) sdk.Int {
	return ((((qAssetTokenAmount.Mul(types.BurnUsqMultiplier)).Quo(atomPrice)).Mul(burningFee)).Quo(types.MintUsqMultiplier)).Quo(types.FeeMultiplier)
}
