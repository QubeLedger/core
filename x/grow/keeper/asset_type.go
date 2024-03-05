package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/QuadrateOrg/core/x/grow/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k Keeper) GetAssetCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AssetCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetAssetCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AssetCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendAsset(
	ctx sdk.Context,
	Asset types.Asset,
) uint64 {
	count := k.GetAssetCount(ctx)
	Asset.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	appendedValue := k.cdc.MustMarshal(&Asset)
	store.Set(GetAssetIDBytes(Asset.Id), appendedValue)
	k.SetAssetCount(ctx, count+1)
	return count
}

func (k Keeper) SetAsset(ctx sdk.Context, Asset types.Asset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	b := k.cdc.MustMarshal(&Asset)
	store.Set(GetAssetIDBytes(Asset.Id), b)
}

func (k Keeper) RemoveAsset(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	store.Delete(GetAssetIDBytes(id))
}

func (k Keeper) GetAllAsset(ctx sdk.Context) (list []types.Asset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Asset
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func GetAssetIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func GetAssetIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

//nolint:all
func (k Keeper) GenerateAssetIdHash(denom1 string) string {
	return fmt.Sprintf("%x", crypto.Sha256(append([]byte(denom1))))
}

// func for gov proposal
func (k Keeper) RegisterAsset(ctx sdk.Context, p types.Asset) error {
	Asset := types.Asset{
		AssetId:                 k.GenerateAssetIdHash(p.AssetMetadata.Base),
		AssetMetadata:           p.AssetMetadata,
		OracleAssetId:           p.OracleAssetId,
		ProvideValue:            p.ProvideValue,
		CollectivelyBorrowValue: p.CollectivelyBorrowValue,
		Type:                    p.Type,
	}
	_ = k.AppendAsset(ctx, Asset)
	return nil
}

/*
functions for get lend asset
*/

func (k Keeper) GetAssetIdByCoins(amountIn string) (string, error) {
	amountInCoins, err := sdk.ParseCoinsNormalized(amountIn)
	if err != nil {
		return "", err
	}
	return k.GenerateAssetIdHash(amountInCoins.GetDenomByIndex(0)), nil
}

func (k Keeper) GetAssetByDenom(ctx sdk.Context, denom string) (types.Asset, bool) {
	assetId := k.GenerateAssetIdHash(denom)
	asset, found := k.GetAssetByAssetId(ctx, assetId)
	return asset, found
}

func (k Keeper) GetAssetByOracleAssetId(ctx sdk.Context, oracleAssetId string) (val types.Asset, err error) {
	allAsset := k.GetAllAsset(ctx)
	for _, v := range allAsset {
		if v.OracleAssetId == oracleAssetId {
			return v, nil
		}
	}
	return val, types.ErrAssetNotFound
}

func (k Keeper) GetAssetByAssetId(ctx sdk.Context, AssetId string) (val types.Asset, found bool) {
	allAsset := k.GetAllAsset(ctx)
	for _, v := range allAsset {
		if v.AssetId == AssetId {
			return v, true
		}
	}
	return val, false
}

func (k Keeper) GetAssetByID(ctx sdk.Context, id uint64) (val types.Asset, found bool) {
	allAsset := k.GetAllAsset(ctx)
	for _, v := range allAsset {
		if v.Id == id {
			return v, true
		}
	}
	return val, false
}
