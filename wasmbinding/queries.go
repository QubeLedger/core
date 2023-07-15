package wasmbinding

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/QuadrateOrg/core/wasmbinding/bindings"

	oraclekeeper "github.com/QuadrateOrg/core/x/oracle/keeper"

	tokenfactorykeeper "github.com/QuadrateOrg/core/x/tokenfactory/keeper"

	interquerykeeper "github.com/QuadrateOrg/core/x/interquery/keeper"
)

type QueryPlugin struct {
	tokenFactoryKeeper *tokenfactorykeeper.Keeper
	oracleKeeper       *oraclekeeper.Keeper
	interqueryKeeper   *interquerykeeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(tfk *tokenfactorykeeper.Keeper, oracle *oraclekeeper.Keeper, interquery *interquerykeeper.Keeper) *QueryPlugin {
	return &QueryPlugin{
		tokenFactoryKeeper: tfk,
		oracleKeeper:       oracle,
		interqueryKeeper:   interquery,
	}
}

// GetDenomAdmin is a query to get denom admin.
func (qp QueryPlugin) GetDenomAdmin(ctx sdk.Context, denom string) (*bindings.DenomAdminResponse, error) {
	metadata, err := qp.tokenFactoryKeeper.GetAuthorityMetadata(ctx, denom)
	if err != nil {
		return nil, fmt.Errorf("failed to get admin for denom: %s", denom)
	}

	return &bindings.DenomAdminResponse{Admin: metadata.Admin}, nil
}

// GetActualProce is a query to get denom admin.
func (qp QueryPlugin) GetActualPrice(ctx sdk.Context) (*bindings.ActualPriceResponse, error) {
	price, err := qp.oracleKeeper.GetPrice(ctx, 0)
	if !err {
		return nil, fmt.Errorf("oracle error")
	}

	return &bindings.ActualPriceResponse{
		Atom:   price.AtomPrice,
		StAtom: price.StatomPrice,
	}, nil
}

func (qp QueryPlugin) GetInterchainQuery(ctx sdk.Context, id string) (*bindings.InterchainQueryResponse, error) {
	res, err := qp.interqueryKeeper.GetQuery(ctx, id)
	if !err {
		return nil, fmt.Errorf("interquery: GetInterchainQuery: error")
	}
	return &bindings.InterchainQueryResponse{
		Id:           res.Id,
		ConnectionId: res.ConnectionId,
		ChainId:      res.ChainId,
		QueryType:    res.QueryType,
		Request:      res.Request,
		Period:       res.Period,
		LastHeight:   res.LastHeight,
		CallbackId:   res.CallbackId,
		Ttl:          res.Ttl,
	}, nil
}

func (qp QueryPlugin) GetInterchainQueryResult(ctx sdk.Context, module, connectionID, chainID, queryType string, request []byte) (*bindings.InterchainQueryResultResponse, error) {
	res, err := qp.interqueryKeeper.GetDatapoint(ctx, module, connectionID, chainID, queryType, request)
	if err != nil {
		return nil, fmt.Errorf("interquery: GetInterchainQueryResult: error")
	}

	return &bindings.InterchainQueryResultResponse{
		Id:           res.Id,
		LocalHeight:  res.LocalHeight.Int64(),
		RemoteHeight: res.RemoteHeight.Int64(),
		Value:        res.Value,
	}, nil
}
