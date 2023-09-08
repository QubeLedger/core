package keeper

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/QuadrateOrg/core/x/grow/types"
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
	pair types.GTokenPair,
) uint64 {
	count := k.GetPairCount(ctx)
	pair.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairKey))
	appendedValue := k.cdc.MustMarshal(&pair)
	store.Set(GetPairIDBytes(pair.Id), appendedValue)
	k.SetPairCount(ctx, count+1)
	return count
}

func (k Keeper) SetPair(ctx sdk.Context, Pair types.GTokenPair) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairKey))
	b := k.cdc.MustMarshal(&Pair)
	store.Set(GetPairIDBytes(Pair.Id), b)
}

func (k Keeper) RemovePair(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairKey))
	store.Delete(GetPairIDBytes(id))
}

func (k Keeper) GetAllPair(ctx sdk.Context) (list []types.GTokenPair) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GTokenPair
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

func (k Keeper) GetPairByDenomID(ctx sdk.Context, denomID string) (val types.GTokenPair, found bool) {
	allPair := k.GetAllPair(ctx)
	for _, v := range allPair {
		if v.DenomID == denomID {
			return v, true
		}
	}
	return val, false
}

func (k Keeper) GetPairByID(ctx sdk.Context, id uint64) (val types.GTokenPair, found bool) {
	allPair := k.GetAllPair(ctx)
	for _, v := range allPair {
		if v.Id == id {
			return v, true
		}
	}
	return val, false
}

//nolint:all
func (k Keeper) GenerateDenomIdHash(denom1 string) string {
	return fmt.Sprintf("%x", crypto.Sha256(append([]byte(denom1))))
}

func (k Keeper) GetDenomIdDeposit(denom string) (string, error) {
	return k.GenerateDenomIdHash(denom), nil
}

func (k Keeper) GetDenomIdWithdrawal(amountIn string) (string, error) {
	msgAmountInCoins, err := sdk.ParseCoinsNormalized(amountIn)
	if err != nil {
		return "", err
	}

	return k.GenerateDenomIdHash(msgAmountInCoins.GetDenomByIndex(0)), nil
}

func (k Keeper) RegisterPair(ctx sdk.Context, p types.GTokenPair) error {
	pair := types.GTokenPair{
		DenomID:                     k.GenerateDenomIdHash(p.GTokenMetadata.Base),
		QStablePairId:               p.QStablePairId,
		GTokenMetadata:              p.GTokenMetadata,
		MinAmountIn:                 p.MinAmountIn,
		MinAmountOut:                p.MinAmountOut,
		GTokenLastPrice:             sdk.Int{},
		GTokenLatestPriceUpdateTime: uint64(time.Now().Unix()),
	}
	_ = k.AppendPair(ctx, pair)
	return nil
}
