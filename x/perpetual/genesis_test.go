package perpetual_test

import (
	"testing"

	"github.com/QuadrateOrg/core/app"
	quadrateapptest "github.com/QuadrateOrg/core/app/helpers"
	apptypes "github.com/QuadrateOrg/core/types"
	"github.com/QuadrateOrg/core/x/perpetual"
	"github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type PerpetualGenesisTestSuite struct {
	suite.Suite
	ctx     sdk.Context
	app     *app.QuadrateApp
	genesis types.GenesisState
}

func (suite *PerpetualGenesisTestSuite) Commit() {
	header := suite.ctx.BlockHeader()
	suite.ctx = suite.app.BaseApp.NewContext(false, header)
}

func (s *PerpetualGenesisTestSuite) Setup() {
	apptypes.SetConfig()
	s.app = quadrateapptest.Setup(s.T(), "qube-1", false, 1)
}

func (s *PerpetualGenesisTestSuite) TestGenesis(t *testing.T) {
	s.Setup()
	s.Commit()
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}
	keeper := s.app.PerpetualKeeper

	perpetual.InitGenesis(s.ctx, keeper, genesisState)
	got := perpetual.ExportGenesis(s.ctx, keeper)
	require.NotNil(t, got)

	s.Require().NotNil(&genesisState)
	s.Require().NotNil(got)
}
