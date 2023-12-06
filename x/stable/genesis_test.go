package stable_test

import (
	"fmt"
	"testing"

	apptypes "github.com/QuadrateOrg/core/types"
	"github.com/QuadrateOrg/core/x/stable"
	"github.com/QuadrateOrg/core/x/stable/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto"

	"github.com/QuadrateOrg/core/app"
	quadrateapptest "github.com/QuadrateOrg/core/app/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StableGenesisTestSuite struct {
	suite.Suite
	ctx     sdk.Context
	app     *app.QuadrateApp
	genesis types.GenesisState
}

func (suite *StableGenesisTestSuite) Commit() {
	header := suite.ctx.BlockHeader()
	suite.ctx = suite.app.BaseApp.NewContext(false, header)
}

func (s *StableGenesisTestSuite) Setup() {
	apptypes.SetConfig()
	s.app = quadrateapptest.Setup(s.T(), "qube-1", false, 1)
}

func (s *StableGenesisTestSuite) TestInitGenesis() {
	testCases := []struct {
		name         string
		genesisState types.GenesisState
	}{
		{
			"default genesis",
			*types.DefaultGenesis(),
		},
		{
			"custom genesis",
			types.NewGenesisState(
				types.DefaultParams(),
				"stable",
				[]types.Pair{
					{
						Id:     0,
						PairId: fmt.Sprintf("%x", crypto.Sha256(append([]byte("uatom"+"uusd")))),
						AmountInMetadata: banktypes.Metadata{
							Description: "",
							DenomUnits: []*banktypes.DenomUnit{
								{Denom: "uatom", Exponent: uint32(0), Aliases: []string{"microatom"}},
							},
							Base:    "uatom",
							Display: "atom",
							Name:    "ATOM",
							Symbol:  "ATOM",
						},
						AmountOutMetadata: banktypes.Metadata{
							Description: "",
							DenomUnits: []*banktypes.DenomUnit{
								{Denom: "uusd", Exponent: uint32(0), Aliases: []string{"microusd"}},
							},
							Base:    "uusd",
							Display: "usd",
							Name:    "USQ",
							Symbol:  "USQ",
						},
						Qm:           sdk.NewInt(0),
						Ar:           sdk.NewInt(0),
						MinAmountIn:  "20uatom",
						MinAmountOut: "20uusd",
					},
				},
			),
		},
	}
	for _, tc := range testCases {
		s.Setup()
		s.Commit()
		s.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			s.Require().NotPanics(func() {
				stable.InitGenesis(s.ctx, s.app.StableKeeper, tc.genesisState)
			})
			params := s.app.StableKeeper.GetParams(s.ctx)

			pairs := s.app.StableKeeper.GetAllPair(s.ctx)
			s.Require().Equal(tc.genesisState.Params, params)
			if len(pairs) > 0 {
				s.Require().Equal(tc.genesisState.Pairs, pairs)
			} else {
				s.Require().Len(tc.genesisState.Pairs, 0)
			}

			reserveFundAddress := s.app.StableKeeper.GetReserveFundAddress(s.ctx)
			rf, err := sdk.AccAddressFromBech32(params.ReserveFundAddress)
			s.Require().NoError(err, tc.name)
			s.Require().Equal(rf, reserveFundAddress)

			burningFundBalance := s.app.StableKeeper.GetBurningFundAddress(s.ctx)
			bf, err := sdk.AccAddressFromBech32(params.BurningFundAddress)
			s.Require().NoError(err, tc.name)
			s.Require().Equal(bf, burningFundBalance)
		})
	}

}

func TestStableGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(StableGenesisTestSuite))
}
