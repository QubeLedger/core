package app_helpers

import (
	"encoding/json"

	quadrateapp "github.com/QuadrateOrg/core/app"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

type EmptyAppOptions struct{}

func (EmptyAppOptions) Get(o string) interface{} { return nil }

func SetupTestingApp() (testApp ibctesting.TestingApp, genesisState map[string]json.RawMessage) {
	db := dbm.NewMemDB()
	encCdc := quadrateapp.MakeEncodingConfig()
	app := quadrateapp.NewQuadrateApp(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		quadrateapp.DefaultNodeHome,
		0,
		encCdc,
		EmptyAppOptions{},
	)
	return app, quadrateapp.NewDefaultGenesisState()
}
