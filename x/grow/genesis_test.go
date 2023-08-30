package grow_test

import (
	"fmt"
	"testing"

	"github.com/QuadrateOrg/core/app"
	quadrateapptest "github.com/QuadrateOrg/core/app/helpers"
	"github.com/QuadrateOrg/core/x/grow"
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
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
	s.app = quadrateapptest.Setup(s.T(), "qube-1", false, 1)
}

func (s *GrowGenesisTestSuite) TestInitGenesis() {
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
				LoanList: []types.Loan{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				LoanCount: 2,
			},
			valid: true,
		},
		{
			name: "duplicated loan",
			genesisState: types.GenesisState{
				Params: types.DefaultParams(),
				LoanList: []types.Loan{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			name: "invalid loan count",
			genesisState: types.GenesisState{
				Params: types.DefaultParams(),
				LoanList: []types.Loan{
					{
						Id: 1,
					},
				},
				LoanCount: 0,
			},
			valid: false,
		},
	}
	for _, tc := range testCases {
		s.Setup()
		s.Commit()
		s.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			if tc.valid {
				s.Require().NotPanics(func() {
					grow.InitGenesis(s.ctx, s.app.GrowKeeper, tc.genesisState)
				})
				params := s.app.GrowKeeper.GetParams(s.ctx)

				loans := s.app.GrowKeeper.GetAllLoan(s.ctx)
				s.Require().Equal(tc.genesisState.Params, params)
				if len(loans) > 0 {
					s.Require().Equal(tc.genesisState.LoanList, loans)
				} else {
					s.Require().Len(tc.genesisState.LoanList, 0)
				}
			}
		})
	}

}

func TestGrowGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GrowGenesisTestSuite))
}
