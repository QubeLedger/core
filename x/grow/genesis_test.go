package grow_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/QuadrateOrg/core/app"
	quadrateapptest "github.com/QuadrateOrg/core/app/helpers"
	apptypes "github.com/QuadrateOrg/core/types"
	"github.com/QuadrateOrg/core/x/grow"
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto"
)

type GrowGenesisTestSuite struct {
	suite.Suite
	ctx     sdk.Context
	app     *app.QuadrateApp
	genesis types.GenesisState
}

func (suite *GrowGenesisTestSuite) Commit() {
	header := suite.ctx.BlockHeader()
	suite.ctx = suite.app.BaseApp.NewContext(false, header)
}

func (s *GrowGenesisTestSuite) Setup() {
	apptypes.SetConfig()
	s.app = quadrateapptest.Setup(s.T(), "qube-1", false, 1)
}

func (s *GrowGenesisTestSuite) TestInitGenesis() {
	s.Setup()
	s.Commit()
	testCases := []struct {
		name         string
		genesisState types.GenesisState
		valid        bool
	}{
		{
			name:         "default is valid",
			genesisState: *types.DefaultGenesis(),
			valid:        true,
		},
		{
			name: "valid genesis state",
			genesisState: types.GenesisState{
				Params: types.DefaultParams(),
				GTokenPairList: []types.GTokenPair{
					{
						Id:            0,
						DenomID:       fmt.Sprintf("%x", crypto.Sha256(append([]byte("ugusd")))),
						QStablePairId: fmt.Sprintf("%x", crypto.Sha256(append([]byte("uatom"+"uusd")))),
						GTokenMetadata: banktypes.Metadata{
							Description: "",
							DenomUnits: []*banktypes.DenomUnit{
								{Denom: "ugusd", Exponent: uint32(0), Aliases: []string{"microgusd"}},
							},
							Base:    "ugusd",
							Display: "gusd",
							Name:    "gUSQ",
							Symbol:  "gUSQ",
						},
						MinAmountIn:                 "20uusd",
						MinAmountOut:                "20ugusd",
						GTokenLastPrice:             sdk.NewInt(0),
						GTokenLatestPriceUpdateTime: 0,
						St:                          sdk.NewInt(0),
					},
				},
			},
			valid: true,
		},
		{
			name: "name null",
			genesisState: types.GenesisState{
				Params: types.DefaultParams(),
				GTokenPairList: []types.GTokenPair{
					{
						Id:            0,
						DenomID:       fmt.Sprintf("%x", crypto.Sha256(append([]byte("ugusd")))),
						QStablePairId: fmt.Sprintf("%x", crypto.Sha256(append([]byte("uatom"+"uusd")))),
						GTokenMetadata: banktypes.Metadata{
							Description: "",
							DenomUnits: []*banktypes.DenomUnit{
								{Denom: "ugusd", Exponent: uint32(0), Aliases: []string{"microgusd"}},
							},
							Base:    "ugusd",
							Display: "gusd",
							Name:    "",
							Symbol:  "gUSQ",
						},
						MinAmountIn:                 "20uusd",
						MinAmountOut:                "20ugusd",
						GTokenLastPrice:             sdk.NewInt(1 * 1000000),
						GTokenLatestPriceUpdateTime: uint64(time.Now().Unix()),
					},
				},
			},
			valid: false,
		},
	}
	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			if tc.valid {
				s.Require().NotPanics(func() {
					grow.InitGenesis(s.ctx, s.app.GrowKeeper, tc.genesisState)
				})
				params := s.app.GrowKeeper.GetParams(s.ctx)

				pairs := s.app.GrowKeeper.GetAllPair(s.ctx)
				s.Require().Equal(tc.genesisState.Params, params)
				if len(pairs) > 0 {
					s.Require().Equal(tc.genesisState.GTokenPairList, pairs)
				} else {
					s.Require().Len(tc.genesisState.GTokenPairList, 0)
				}
			}
		})
	}

}

func TestGrowGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GrowGenesisTestSuite))
}
