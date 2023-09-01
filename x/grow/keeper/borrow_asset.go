package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/QuadrateOrg/core/x/grow/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k Keeper) GetBorrowAssetCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BorrowAssetCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetBorrowAssetCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BorrowAssetCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendBorrowAsset(
	ctx sdk.Context,
	borrowAsset types.BorrowAsset,
) uint64 {
	count := k.GetBorrowAssetCount(ctx)
	borrowAsset.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAssetKey))
	appendedValue := k.cdc.MustMarshal(&borrowAsset)
	store.Set(GetBorrowAssetIDBytes(borrowAsset.Id), appendedValue)
	k.SetBorrowAssetCount(ctx, count+1)
	return count
}

func (k Keeper) SetBorrowAsset(ctx sdk.Context, BorrowAsset types.BorrowAsset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAssetKey))
	b := k.cdc.MustMarshal(&BorrowAsset)
	store.Set(GetBorrowAssetIDBytes(BorrowAsset.Id), b)
}

func (k Keeper) RemoveBorrowAsset(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAssetKey))
	store.Delete(GetBorrowAssetIDBytes(id))
}

func (k Keeper) GetAllBorrowAsset(ctx sdk.Context) (list []types.BorrowAsset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAssetKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BorrowAsset
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func GetBorrowAssetIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func GetBorrowAssetIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) GetBorrowAssetByBorrowAssetId(ctx sdk.Context, borrowAssetId string) (val types.BorrowAsset, found bool) {
	allBorrowAsset := k.GetAllBorrowAsset(ctx)
	for _, v := range allBorrowAsset {
		if v.BorrowAssetId == borrowAssetId {
			return v, true
		}
	}
	return val, false
}

func (k Keeper) GetBorrowAssetByID(ctx sdk.Context, id uint64) (val types.BorrowAsset, found bool) {
	allBorrowAsset := k.GetAllBorrowAsset(ctx)
	for _, v := range allBorrowAsset {
		if v.Id == id {
			return v, true
		}
	}
	return val, false
}

//nolint:all
func (k Keeper) GenerateBorrowAssetIdHash(denom1 string, denom2 string) string {
	return fmt.Sprintf("%x", crypto.Sha256(append([]byte(denom1+denom2))))
}

func (k Keeper) RegisterBorrowAsset(ctx sdk.Context, p types.BorrowAsset) error {
	borrowAsset := types.BorrowAsset{
		BorrowAssetId:          k.GenerateBorrowAssetIdHash(p.AmountInAssetMetadata.Base, p.AmountOutAssetMetadata.Base),
		AmountInAssetMetadata:  p.AmountInAssetMetadata,
		AmountOutAssetMetadata: p.AmountOutAssetMetadata,
		OracleAssetId:          p.OracleAssetId,
	}
	_ = k.AppendBorrowAsset(ctx, borrowAsset)
	return nil
}

func (k Keeper) GetBorrowAssetIdCreateLend(amountIn string, denom2 string) (string, error) {
	amountInCoins, err := sdk.ParseCoinsNormalized(amountIn)
	if err != nil {
		return "", err
	}
	return k.GenerateBorrowAssetIdHash(amountInCoins.GetDenomByIndex(0), denom2), nil
}

func (k Keeper) GetBorrowAssetIdDeleteLend(amountIn string, denom2 string) (string, error) {
	amountInCoins, err := sdk.ParseCoinsNormalized(amountIn)
	if err != nil {
		return "", err
	}
	return k.GenerateBorrowAssetIdHash(denom2, amountInCoins.GetDenomByIndex(0)), nil
}
