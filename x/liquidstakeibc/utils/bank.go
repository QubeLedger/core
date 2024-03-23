package utils

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// UnmarshalBalanceCompat unmarshal balance amount from storage, it's backward-compatible with the legacy format.
func UnmarshalBalanceCompat(cdc codec.BinaryCodec, bz []byte, denom string) (sdk.Coin, error) {
	if err := sdk.ValidateDenom(denom); err != nil {
		return sdk.Coin{}, err
	}

	amount := sdk.ZeroInt()
	if bz == nil {
		return sdk.NewCoin(denom, amount), nil
	}

	if err := amount.Unmarshal(bz); err != nil {
		// try to unmarshal with the legacy format.
		var balance sdk.Coin
		if cdc.Unmarshal(bz, &balance) != nil {
			// return with the original error
			return sdk.Coin{}, err
		}
		return balance, nil
	}

	return sdk.NewCoin(denom, amount), nil
}
