package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/QuadrateOrg/core/x/grow/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k Keeper) GetLendAssetCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LendAssetCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetLendAssetCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LendAssetCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendLendAsset(
	ctx sdk.Context,
	LendAsset types.LendAsset,
) uint64 {
	count := k.GetLendAssetCount(ctx)
	LendAsset.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LendAssetKey))
	appendedValue := k.cdc.MustMarshal(&LendAsset)
	store.Set(GetLendAssetIDBytes(LendAsset.Id), appendedValue)
	k.SetLendAssetCount(ctx, count+1)
	return count
}

func (k Keeper) SetLendAsset(ctx sdk.Context, LendAsset types.LendAsset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LendAssetKey))
	b := k.cdc.MustMarshal(&LendAsset)
	store.Set(GetLendAssetIDBytes(LendAsset.Id), b)
}

func (k Keeper) RemoveLendAsset(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LendAssetKey))
	store.Delete(GetLendAssetIDBytes(id))
}

func (k Keeper) GetAllLendAsset(ctx sdk.Context) (list []types.LendAsset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LendAssetKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LendAsset
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func GetLendAssetIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func GetLendAssetIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

//nolint:all
func (k Keeper) GenerateLendAssetIdHash(denom1 string) string {
	return fmt.Sprintf("%x", crypto.Sha256(append([]byte(denom1))))
}

// func for gov proposal
func (k Keeper) RegisterLendAsset(ctx sdk.Context, p types.LendAsset) error {
	LendAsset := types.LendAsset{
		LendAssetId:   k.GenerateLendAssetIdHash(p.AssetMetadata.Base),
		AssetMetadata: p.AssetMetadata,
		OracleAssetId: p.OracleAssetId,
	}
	_ = k.AppendLendAsset(ctx, LendAsset)
	return nil
}

/*
functions for get lend asset
*/

func (k Keeper) GetLendAssetIdByCoins(amountIn string) (string, error) {
	amountInCoins, err := sdk.ParseCoinsNormalized(amountIn)
	if err != nil {
		return "", err
	}
	return k.GenerateLendAssetIdHash(amountInCoins.GetDenomByIndex(0)), nil
}

func (k Keeper) GetLendAssetByOracleAssetId(ctx sdk.Context, oracleAssetId string) (val types.LendAsset, err error) {
	allLendAsset := k.GetAllLendAsset(ctx)
	for _, v := range allLendAsset {
		if v.OracleAssetId == oracleAssetId {
			return v, nil
		}
	}
	return val, types.ErrLendAssetNotFound
}

func (k Keeper) GetLendAssetByLendAssetId(ctx sdk.Context, LendAssetId string) (val types.LendAsset, found bool) {
	allLendAsset := k.GetAllLendAsset(ctx)
	for _, v := range allLendAsset {
		if v.LendAssetId == LendAssetId {
			return v, true
		}
	}
	return val, false
}

func (k Keeper) GetLendAssetByID(ctx sdk.Context, id uint64) (val types.LendAsset, found bool) {
	allLendAsset := k.GetAllLendAsset(ctx)
	for _, v := range allLendAsset {
		if v.Id == id {
			return v, true
		}
	}
	return val, false
}
