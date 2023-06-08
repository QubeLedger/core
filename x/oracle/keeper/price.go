package keeper

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/QuadrateOrg/core/x/oracle/types"
	jsonparser "github.com/buger/jsonparser"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetPriceCount get the total number of price
func (k Keeper) GetPriceCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PriceCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPriceCount set the total number of price
func (k Keeper) SetPriceCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PriceCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPrice appends a price in the store with a new id and update the count
func (k Keeper) AppendPrice(
	ctx sdk.Context,
	price types.Price,
) uint64 {
	// Create the price
	count := k.GetPriceCount(ctx)

	// Set the ID of the appended value
	price.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceKey))
	appendedValue := k.cdc.MustMarshal(&price)
	store.Set(GetPriceIDBytes(price.Id), appendedValue)

	// Update price count
	k.SetPriceCount(ctx, count+1)

	return count
}

// SetPrice set a specific price in the store
func (k Keeper) SetPrice(ctx sdk.Context, price types.Price) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceKey))
	b := k.cdc.MustMarshal(&price)
	store.Set(GetPriceIDBytes(price.Id), b)
}

// TODO
// Sometimes the osmosis api does not return the stATOM price
// Set a more reliable stATOM price source
func (k Keeper) GetTokensActualPrice(ctx sdk.Context) (string, string, error) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	res, err := client.Get("https://api.coinbase.com/v2/exchange-rates?currency=ATOM")
	if err != nil {
		return "", "", err
	}
	body, _ := ioutil.ReadAll(res.Body)

	var atomPrice string
	if value, err := jsonparser.GetString(body, "data", "rates", "USD"); err == nil {
		atomPrice = value
	} else {
		return "", "", err
	}

	res1, err := client.Get("https://api.osmosis.zone/tokens/v2/price/statom")
	if err != nil {
		return atomPrice, "", err
	}
	body1, _ := ioutil.ReadAll(res1.Body)

	var statomPrice string
	if value, err := jsonparser.GetFloat(body1, "price"); err == nil {
		statomPrice = fmt.Sprintf("%v", value)
	} else {
		return "", "", err
	}
	return atomPrice, statomPrice, nil
}

// GetPrice returns a price from its id
func (k Keeper) GetPrice(ctx sdk.Context, id uint64) (val types.Price, found bool) {
	atom, statom, err := k.GetTokensActualPrice(ctx)
	if err != nil {
		return val, false
	}
	var prices = types.Price{
		Id:          0,
		AtomPrice:   fmt.Sprintf("%v", atom),
		StatomPrice: fmt.Sprintf("%v", statom),
	}
	k.AppendPrice(ctx, prices)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceKey))
	b := store.Get(GetPriceIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePrice removes a price from the store
func (k Keeper) RemovePrice(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceKey))
	store.Delete(GetPriceIDBytes(id))
}

// GetAllPrice returns all price
func (k Keeper) GetAllPrice(ctx sdk.Context) (list []types.Price) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Price
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPriceIDBytes returns the byte representation of the ID
func GetPriceIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetPriceIDFromBytes returns ID in uint64 format from a byte array
func GetPriceIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
