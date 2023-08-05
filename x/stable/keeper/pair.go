package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/QuadrateOrg/core/x/stable/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k Keeper) GetPairCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PairCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetPairCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PairCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendPair(
	ctx sdk.Context,
	pair types.Pair,
) uint64 {
	count := k.GetPairCount(ctx)
	pair.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairKey))
	appendedValue := k.cdc.MustMarshal(&pair)
	store.Set(GetPairIDBytes(pair.Id), appendedValue)
	k.SetPairCount(ctx, count+1)
	return count
}

func (k Keeper) SetPair(ctx sdk.Context, Pair types.Pair) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairKey))
	b := k.cdc.MustMarshal(&Pair)
	store.Set(GetPairIDBytes(Pair.Id), b)
}

func (k Keeper) RemovePair(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairKey))
	store.Delete(GetPairIDBytes(id))
}

func (k Keeper) GetAllPair(ctx sdk.Context) (list []types.Pair) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var val types.Pair
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

func GetPairIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func GetPairIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) GetPairByPairID(ctx sdk.Context, pairId string) (val types.Pair, found bool) {
	allPair := k.GetAllPair(ctx)
	for _, v := range allPair {
		if v.PairId == pairId {
			return v, true
		}
	}
	return val, false
}

func (k Keeper) GeneratePairIdHash(denom1 string, denom2 string) string {
	return fmt.Sprintf("%x", crypto.Sha256(append([]byte(denom1+denom2))))
}

func (k Keeper) GetPairIdMint(amountInt string, denom2 string) (string, error) {
	amountIntCoins, err := sdk.ParseCoinsNormalized(amountInt)
	if err != nil {
		return "", err
	}
	return k.GeneratePairIdHash(amountIntCoins.GetDenomByIndex(0), denom2), nil
}

func (k Keeper) GetPairIdBurn(amountInt string, denom2 string) (string, error) {
	amountIntCoins, err := sdk.ParseCoinsNormalized(amountInt)
	if err != nil {
		return "", err
	}
	return k.GeneratePairIdHash(denom2, amountIntCoins.GetDenomByIndex(0)), nil
}
