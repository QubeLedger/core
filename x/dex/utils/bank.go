package utils

import (
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SanitizeCoins takes an unsorted list of coins and sorts them, removes coins with amount zero and combines duplicate coins
func SanitizeCoins(coins []sdk.Coin) sdk.Coins {
	sort.SliceStable(coins, func(i, j int) bool {
		return coins[i].Denom < coins[j].Denom
	})
	cleanCoins := sdk.Coins{}
	lastDenom := ""
	for _, coin := range coins {
		if coin.IsZero() {
			continue
		}
		if lastDenom != coin.Denom {
			cleanCoins = append(cleanCoins, coin)
		} else {
			cleanCoins[len(cleanCoins)-1].Add(coin)
		}
		lastDenom = coin.Denom
	}
	return cleanCoins
}
