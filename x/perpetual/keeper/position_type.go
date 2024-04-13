package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/QuadrateOrg/core/x/perpetual/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k Keeper) GetPositionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PositionCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetPositionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PositionCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendPosition(
	ctx sdk.Context,
	Position types.TradePosition,
) uint64 {
	count := k.GetPositionCount(ctx)
	Position.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionKey))
	appendedValue := k.cdc.MustMarshal(&Position)
	store.Set(GetPositionIDBytes(Position.Id), appendedValue)
	k.SetPositionCount(ctx, count+1)
	return count
}

func (k Keeper) SetPosition(ctx sdk.Context, Position types.TradePosition) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionKey))
	b := k.cdc.MustMarshal(&Position)
	store.Set(GetPositionIDBytes(Position.Id), b)
}

func (k Keeper) RemovePosition(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionKey))
	store.Delete(GetPositionIDBytes(id))
}

func (k Keeper) GetAllPositions(ctx sdk.Context) (list []types.TradePosition) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var val types.TradePosition
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

func GetPositionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func GetPositionIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) GenerateTraderPositionId(creator string, denom string, trading_asset_id string, trade_type types.PerpetualTradeType, leverage sdk.Dec) string {
	return fmt.Sprintf("%x", crypto.Sha256(append([]byte(creator+denom+trading_asset_id+trade_type.String()+leverage.String()))))
}

func (k Keeper) GetPositionByPositionId(ctx sdk.Context, position_id string) (val types.TradePosition, found bool) {
	allPair := k.GetAllPositions(ctx)
	for _, v := range allPair {
		if v.TradePositionId == position_id {
			return v, true
		}
	}
	return val, false
}
