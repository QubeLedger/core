package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func CreateCoins(denom string, amount sdk.Int) sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin(denom, amount))
}
