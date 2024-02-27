package keeper

import (
	"encoding/binary"
	"fmt"
	"sort"

	"github.com/QuadrateOrg/core/x/grow/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k Keeper) GetLiquidatorPositionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LiquidatorPositionCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetLiquidatorPositionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LiquidatorPositionCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendLiquidatorPosition(
	ctx sdk.Context,
	LiquidatorPosition types.LiquidatorPosition,
) uint64 {
	count := k.GetLiquidatorPositionCount(ctx)
	LiquidatorPosition.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LiquidatorPositionKey))
	appendedValue := k.cdc.MustMarshal(&LiquidatorPosition)
	store.Set(GetLiquidatorPositionIDBytes(LiquidatorPosition.Id), appendedValue)
	k.SetLiquidatorPositionCount(ctx, count+1)
	return count
}

func (k Keeper) SetLiquidatorPosition(ctx sdk.Context, LiquidatorPosition types.LiquidatorPosition) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LiquidatorPositionKey))
	b := k.cdc.MustMarshal(&LiquidatorPosition)
	store.Set(GetLiquidatorPositionIDBytes(LiquidatorPosition.Id), b)
}

func (k Keeper) RemoveLiquidatorPosition(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LiquidatorPositionKey))
	store.Delete(GetLiquidatorPositionIDBytes(id))
}

func (k Keeper) GetAllLiquidatorPosition(ctx sdk.Context) (list []types.LiquidatorPosition) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LiquidatorPositionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LiquidatorPosition
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func GetLiquidatorPositionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func GetLiquidatorPositionIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

//nolint:all
func (k Keeper) GenerateLiquidatorPositionId(address string, denom1 string, asset1 string, amount string, premium string) string {
	return fmt.Sprintf("%x", crypto.Sha256(append([]byte(address+denom1+asset1+amount+premium))))
}

func (k Keeper) GetLiquidatorPositionByLiquidatorPositionId(ctx sdk.Context, LiquidatorPositionId string) (val types.LiquidatorPosition, found bool) {
	allLiquidatorPosition := k.GetAllLiquidatorPosition(ctx)
	for _, v := range allLiquidatorPosition {
		if v.LiquidatorPositionId == LiquidatorPositionId {
			return v, true
		}
	}
	return val, false
}

func (k Keeper) GetLiquidatorPositionByID(ctx sdk.Context, id uint64) (val types.LiquidatorPosition, found bool) {
	allLiquidatorPosition := k.GetAllLiquidatorPosition(ctx)
	for _, v := range allLiquidatorPosition {
		if v.Id == id {
			return v, true
		}
	}
	return val, false
}

func (k Keeper) CheckIfLiquidatorPositionAlredyCreate(ctx sdk.Context, depositor string, denom string) error {
	allLiquidatorPosition := k.GetAllLiquidatorPosition(ctx)
	for _, v := range allLiquidatorPosition {
		if v.Liquidator == depositor && v.LiquidatorPositionId == k.CalculateDepositId(depositor) {
			return types.ErrUserAlredyDepositCollateral
		}
	}
	return nil
}

/* #nosec */
func (k Keeper) GetLiquidatorPositionsByAssetAndDenom(ctx sdk.Context, wantAssetId string, provideAssetId string) []types.LiquidatorPosition {
	allLiquidatorPosition := k.GetAllLiquidatorPosition(ctx)
	res := []types.LiquidatorPosition{}
	for _, v := range allLiquidatorPosition {
		if v.WantAssetId == wantAssetId && v.ProvidedAssetId == provideAssetId {
			res = append(res, v)
		}
	}
	return res
}

func (k Keeper) SortLiquidatorPositionsByPremium(ctx sdk.Context, lps []types.LiquidatorPosition) []types.LiquidatorPosition {
	sort.SliceStable(lps, func(i, j int) bool {
		return int64(lps[i].Premium) < int64(lps[j].Premium)
	})
	return lps
}
