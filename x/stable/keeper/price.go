package keeper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/QuadrateOrg/core/x/stable/types"
	"github.com/buger/jsonparser"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	AtomPrice   sdk.Int
	TestingMode bool = false
)

func (k Keeper) UpdateAtomPrice(ctx sdk.Context, pair types.Pair) error {
	if TestingMode {
		return nil
	}
	if AtomPrice.IsNil() {
		AtomPrice = sdk.NewInt(0)
	}
	price, err := GetPrice() //k.oracleKeeper.GetExchangeRate(ctx, pair.AmountInMetadata.Base)
	if err != nil {
		return err
	}
	if price.IsNil() {
		return types.ErrAtomPriceNil
	}
	AtomPrice = price.MulInt64(10000).RoundInt()
	return nil
}

func GetPrice() (sdk.Dec, error) {
	var atomPriceString string

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	res, err := client.Get("https://api.coinbase.com/v2/exchange-rates?currency=ATOM")
	if err != nil {
		return sdk.ZeroDec(), err
	}
	body, _ := ioutil.ReadAll(res.Body)

	if value, err := jsonparser.GetString(body, "data", "rates", "USD"); err == nil {
		atomPriceString = fmt.Sprintf("%v", value)
	} else {
		return sdk.ZeroDec(), err
	}

	val, err := sdk.NewDecFromStr(atomPriceString)

	return val, nil
}

func (k Keeper) UpdateAtomPriceTesting(ctx sdk.Context, price sdk.Int) error {
	if !TestingMode {
		return nil
	}
	AtomPrice = price
	return nil
}

func (k Keeper) GetAtomPrice(ctx sdk.Context, pair types.Pair) (sdk.Int, error) {
	err := k.UpdateAtomPrice(ctx, pair)
	if err != nil {
		return sdk.Int{}, err
	}
	return AtomPrice, err
}

func (k Keeper) SetTestingMode(value bool) {
	TestingMode = value
}
