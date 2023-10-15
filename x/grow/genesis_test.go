package grow_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/QuadrateOrg/core/app"
	"github.com/QuadrateOrg/core/app/apptesting"
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
				RealRate:                  15,
				BorrowRate:                15,
				GrowStakingReserveAddress: apptesting.CreateRandomAccounts(1)[0].String(),
				USQReserveAddress:         apptesting.CreateRandomAccounts(1)[0].String(),
			},
			valid: true,
		},
		{
			name: "address null",
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
						GTokenLastPrice:             sdk.NewInt(1 * 1000000),
						GTokenLatestPriceUpdateTime: uint64(time.Now().Unix()),
					},
				},
				RealRate:                  15,
				BorrowRate:                15,
				GrowStakingReserveAddress: "",
				USQReserveAddress:         "",
			},
			valid: false,
		},
		{
			name: "percent null",
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
						GTokenLastPrice:             sdk.NewInt(1 * 1000000),
						GTokenLatestPriceUpdateTime: uint64(time.Now().Unix()),
					},
				},
				RealRate:                  0,
				BorrowRate:                0,
				GrowStakingReserveAddress: "",
				USQReserveAddress:         "",
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

				uf, err := sdk.AccAddressFromBech32(tc.genesisState.USQReserveAddress)
				s.Require().NoError(err, tc.name)
				gf, err := sdk.AccAddressFromBech32(tc.genesisState.GrowStakingReserveAddress)
				s.Require().NoError(err, tc.name)

				s.Require().Equal(uf, s.app.GrowKeeper.GetUSQReserveAddress(s.ctx))
				s.Require().Equal(gf, s.app.GrowKeeper.GetGrowStakingReserveAddress(s.ctx))
			}
		})
	}

}

func TestGrowGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GrowGenesisTestSuite))
}
