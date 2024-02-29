package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/QuadrateOrg/core/x/grow/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k Keeper) GetLendCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LendCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetLendCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LendCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendLend(
	ctx sdk.Context,
	Lend types.Lend,
) uint64 {
	// Create the Lend
	count := k.GetLendCount(ctx)

	// Set the ID of the appended value
	Lend.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LendKey))
	appendedValue := k.cdc.MustMarshal(&Lend)
	store.Set(GetLendIDBytes(Lend.Id), appendedValue)

	// Update Lend count
	k.SetLendCount(ctx, count+1)

	return count
}

// SetLend set a specific Lend in the store
func (k Keeper) SetLend(ctx sdk.Context, Lend types.Lend) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LendKey))
	b := k.cdc.MustMarshal(&Lend)
	store.Set(GetLendIDBytes(Lend.Id), b)
}

// GetLend returns a Lend from its id
func (k Keeper) GetLend(ctx sdk.Context, id uint64) (val types.Lend, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LendKey))
	b := store.Get(GetLendIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLend removes a Lend from the store
func (k Keeper) RemoveLend(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LendKey))
	store.Delete(GetLendIDBytes(id))
}

// GetAllLend returns all Lend
func (k Keeper) GetAllLend(ctx sdk.Context) (list []types.Lend) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LendKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Lend
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetLendIDBytes returns the byte representation of the ID
func GetLendIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetLendIDFromBytes returns ID in uint64 format from a byte array
func GetLendIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

//nolint:all
func (k Keeper) GenerateLendIdHash(denom1 string, amount string, borrower string, time string) string {
	return fmt.Sprintf("%x", crypto.Sha256(append([]byte(denom1+amount+borrower+time))))
}

func (k Keeper) GetLendByLendId(ctx sdk.Context, lend_id string) (val types.Lend, found bool) {
	allLend := k.GetAllLend(ctx)
	for _, v := range allLend {
		if v.LendId == lend_id {
			return v, true
		}
	}
	return val, false
}
