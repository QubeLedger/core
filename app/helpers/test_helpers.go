package helpers

import (
	"encoding/json"
	"testing"
	"time"

	quadrateapp "github.com/QuadrateOrg/core/app"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

// SimAppChainID hardcoded chainID for simulation
const (
	SimAppChainID = "qube-1"
)

// DefaultConsensusParams defines the default Tendermint consensus params used
var DefaultConsensusParams = &abci.ConsensusParams{
	Block: &abci.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

type EmptyAppOptions struct{}

func (EmptyAppOptions) Get(o string) interface{} { return nil }

var defaultGenesisBz []byte

func getDefaultGenesisStateBytes() []byte {
	if len(defaultGenesisBz) == 0 {
		genesisState := quadrateapp.NewDefaultGenesisState()
		stateBytes, _ := json.MarshalIndent(genesisState, "", " ")
		/*if err != nil {
			panic(err)
		}*/
		defaultGenesisBz = stateBytes
	}
	return defaultGenesisBz
}

func Setup(t *testing.T, chainId string, isCheckTx bool, invCheckPeriod uint) *quadrateapp.QuadrateApp {
	//t.Helper()

	db := dbm.NewMemDB()
	encCdc := quadrateapp.MakeEncodingConfig()
	app := quadrateapp.NewQuadrateApp(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		quadrateapp.DefaultNodeHome,
		invCheckPeriod,
		encCdc,
		EmptyAppOptions{},
	)

	if !isCheckTx {
		app.InitChain(
			abci.RequestInitChain{
				ChainId:         chainId,
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: DefaultConsensusParams,
				AppStateBytes:   getDefaultGenesisStateBytes(),
			},
		)
	}

	return app
}
