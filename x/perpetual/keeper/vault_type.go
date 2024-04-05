package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/QuadrateOrg/core/x/perpetual/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k Keeper) GetVaultCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VaultCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetVaultCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VaultCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendVault(
	ctx sdk.Context,
	Vault types.Vault,
) uint64 {
	count := k.GetVaultCount(ctx)
	Vault.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	appendedValue := k.cdc.MustMarshal(&Vault)
	store.Set(GetVaultIDBytes(Vault.Id), appendedValue)
	k.SetVaultCount(ctx, count+1)
	return count
}

func (k Keeper) SetVault(ctx sdk.Context, Vault types.Vault) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	b := k.cdc.MustMarshal(&Vault)
	store.Set(GetVaultIDBytes(Vault.Id), b)
}

func (k Keeper) RemoveVault(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	store.Delete(GetVaultIDBytes(id))
}

func (k Keeper) GetAllVault(ctx sdk.Context) (list []types.Vault) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var val types.Vault
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

func GetVaultIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func GetVaultIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) GenerateVaultIdHash(denom1 string, denom2 string) string {
	return fmt.Sprintf("%x", crypto.Sha256(append([]byte(denom1+denom2))))
}

func (k Keeper) GetVaultByVaultId(ctx sdk.Context, vault_id string) (val types.Vault, found bool) {
	allPair := k.GetAllVault(ctx)
	for _, v := range allPair {
		if v.VaultId == vault_id {
			return v, true
		}
	}
	return val, false
}

func (k Keeper) RemoveLongFromVault(ctx sdk.Context, pos_id string, vault types.Vault) types.Vault {
	for i, lid := range vault.LongPosition {
		if lid.TradePositionId == pos_id {
			vault.LongPosition = append(vault.LongPosition[:i], vault.LongPosition[i+1:]...)
		}
	}
	return vault
}

func (k Keeper) RemoveShortFromVault(ctx sdk.Context, pos_id string, vault types.Vault) types.Vault {
	for i, lid := range vault.ShortPosition {
		if lid.TradePositionId == pos_id {
			vault.ShortPosition = append(vault.ShortPosition[:i], vault.ShortPosition[i+1:]...)
		}
	}
	return vault
}
