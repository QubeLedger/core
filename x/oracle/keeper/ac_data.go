package keeper

import (
	"encoding/binary"

	"github.com/QuadrateOrg/core/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAcDataCount get the total number of acData
func (k Keeper) GetAcDataCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AcDataCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAcDataCount set the total number of acData
func (k Keeper) SetAcDataCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AcDataCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAcData appends a acData in the store with a new id and update the count
func (k Keeper) AppendAcData(
	ctx sdk.Context,
	acData types.AcData,
) uint64 {
	// Create the acData
	count := k.GetAcDataCount(ctx)

	// Set the ID of the appended value
	acData.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AcDataKey))
	appendedValue := k.cdc.MustMarshal(&acData)
	store.Set(GetAcDataIDBytes(acData.Id), appendedValue)

	// Update acData count
	k.SetAcDataCount(ctx, count+1)

	return count
}

// SetAcData set a specific acData in the store
func (k Keeper) SetAcData(ctx sdk.Context, acData types.AcData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AcDataKey))
	b := k.cdc.MustMarshal(&acData)
	store.Set(GetAcDataIDBytes(acData.Id), b)
}

// GetAcData returns a acData from its id
func (k Keeper) GetAcData(ctx sdk.Context, id uint64) (val types.AcData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AcDataKey))
	b := store.Get(GetAcDataIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAcData removes a acData from the store
func (k Keeper) RemoveAcData(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AcDataKey))
	store.Delete(GetAcDataIDBytes(id))
}

// GetAllAcData returns all acData
func (k Keeper) GetAllAcData(ctx sdk.Context) (list []types.AcData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AcDataKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AcData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAcDataIDBytes returns the byte representation of the ID
func GetAcDataIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAcDataIDFromBytes returns ID in uint64 format from a byte array
func GetAcDataIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
