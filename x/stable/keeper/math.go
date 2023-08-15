package keeper

import (
	"github.com/QubeLedger/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CalculateAmountToMint(ibcBaseTokenDenomAmount sdk.Int, atomPrice sdk.Int, mintingFee sdk.Int) sdk.Int {
	return (ibcBaseTokenDenomAmount.Mul(atomPrice).Sub((ibcBaseTokenDenomAmount.Mul(atomPrice).Mul(mintingFee)).Quo(types.MintUsqMultiplier))).Quo(types.Multiplier)
}

func (k Keeper) CalculateMintingFeeForBurningFund(ibcBaseTokenDenomAmount sdk.Int, atomPrice sdk.Int, mintingFee sdk.Int) sdk.Int {
	return (((ibcBaseTokenDenomAmount.Mul(atomPrice)).Mul(mintingFee)).Quo(types.MintUsqMultiplier)).Quo(types.Multiplier)
}

func (k Keeper) CalculateAmountToSend(uusdTokenAmount sdk.Int, atomPrice sdk.Int, burningFee sdk.Int) sdk.Int {
	fee := (((uusdTokenAmount.Mul(types.BurnUsqMultiplier)).Quo(atomPrice)).Mul(burningFee)).Quo(types.MintUsqMultiplier)
	return (((uusdTokenAmount.Mul(types.BurnUsqMultiplier)).Quo(atomPrice)).Sub(fee)).Quo(types.FeeMultiplier)
}

func (k Keeper) CalculateBurningFeeForBurningFund(uusdTokenAmount sdk.Int, atomPrice sdk.Int, burningFee sdk.Int) sdk.Int {
	return ((((uusdTokenAmount.Mul(types.BurnUsqMultiplier)).Quo(atomPrice)).Mul(burningFee)).Quo(types.MintUsqMultiplier)).Quo(types.FeeMultiplier)
}
