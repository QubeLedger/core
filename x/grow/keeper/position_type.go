package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/QuadrateOrg/core/x/grow/types"
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
	Position types.Position,
) uint64 {
	count := k.GetPositionCount(ctx)
	Position.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionKey))
	appendedValue := k.cdc.MustMarshal(&Position)
	store.Set(GetPositionIDBytes(Position.Id), appendedValue)
	k.SetPositionCount(ctx, count+1)
	return count
}

func (k Keeper) SetPosition(ctx sdk.Context, Position types.Position) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionKey))
	b := k.cdc.MustMarshal(&Position)
	store.Set(GetPositionIDBytes(Position.Id), b)
}

func (k Keeper) RemovePosition(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionKey))
	store.Delete(GetPositionIDBytes(id))
}

func (k Keeper) GetAllPosition(ctx sdk.Context) (list []types.Position) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Position
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

func (k Keeper) GetPositionByPositionId(ctx sdk.Context, PositionId string) (val types.Position, found bool) {
	allPosition := k.GetAllPosition(ctx)
	for _, v := range allPosition {
		if v.DepositId == PositionId {
			return v, true
		}
	}
	return val, false
}

func (k Keeper) GetPositionByID(ctx sdk.Context, id uint64) (val types.Position, found bool) {
	allPosition := k.GetAllPosition(ctx)
	for _, v := range allPosition {
		if v.Id == id {
			return v, true
		}
	}
	return val, false
}

func (k Keeper) CheckIfPositionAlredyCreate(ctx sdk.Context, depositor string, denom string) error {
	allPosition := k.GetAllPosition(ctx)
	for _, v := range allPosition {
		if v.Creator == depositor && v.DepositId == k.CalculateDepositId(depositor, denom) {
			return types.ErrUserAlredyDepositCollateral
		}
	}
	return nil
}

//nolint:all
func (k Keeper) CalculateDepositId(address string, denom string) string {
	return fmt.Sprintf("%x", crypto.Sha256(append([]byte(address+denom))))
}

func (k Keeper) CheckLoanIdInPosition(ctx sdk.Context, loanId string, position types.Position) bool {
	for _, v := range position.LoanIds {
		if v == loanId {
			return true
		}
	}
	return false
}

func (k Keeper) PushLoanToPosition(ctx sdk.Context, loanId string, position types.Position) types.Position {
	position.LoanIds = append(position.LoanIds, loanId)
	return position
}

func (k Keeper) RemoveLoanInPosition(ctx sdk.Context, loanId string, position types.Position) types.Position {
	for i, lid := range position.LoanIds {
		if lid == loanId {
			position.LoanIds = append(position.LoanIds[:i], position.LoanIds[i+1:]...)
		}
	}
	return position
}

func (k Keeper) IncreaseBorrowedAmountInUSDInPosition(ctx sdk.Context, position types.Position, amt sdk.Int) types.Position {
	position.BorrowedAmountInUSD = position.BorrowedAmountInUSD + amt.Uint64()
	return position
}

func (k Keeper) ReduceBorrowedAmountInUSDInPosition(ctx sdk.Context, position types.Position, amt sdk.Int) types.Position {
	position.BorrowedAmountInUSD = position.BorrowedAmountInUSD - amt.Uint64()
	return position
}

func (k Keeper) ReduceCollateralInPosition(ctx sdk.Context, position types.Position, amt sdk.Int) types.Position {
	oldAmt, denom, _ := k.GetAmountIntFromCoins(position.Collateral)
	position.Collateral = k.FastCoins(denom, oldAmt.Sub(amt)).String()
	return position
}
