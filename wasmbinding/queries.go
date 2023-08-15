package wasmbinding

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/QubeLedger/core/wasmbinding/bindings"
	oraclekeeper "github.com/QubeLedger/core/x/oracle/keeper"
	tokenfactorykeeper "github.com/QubeLedger/core/x/tokenfactory/keeper"
)

type QueryPlugin struct {
	tokenFactoryKeeper *tokenfactorykeeper.Keeper
	oracleKeeper       *oraclekeeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(tfk *tokenfactorykeeper.Keeper, oracle *oraclekeeper.Keeper) *QueryPlugin {
	return &QueryPlugin{
		tokenFactoryKeeper: tfk,
		oracleKeeper:       oracle,
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
