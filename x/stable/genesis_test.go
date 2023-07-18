package stable_test

import (
	"testing"

	"github.com/QuadrateOrg/core/x/stable"
	"github.com/QuadrateOrg/core/x/stable/types"
	"github.com/stretchr/testify/suite"

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

func (s *StableGenesisTestSuite) Setup() {
	s.app = quadrateapptest.Setup(s.T(), "quadrate_5120-1", false, 1)

}

func (s *StableGenesisTestSuite) TestGenesis() {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	stable.InitGenesis(s.ctx, s.app.StableKeeper, genesisState)
	got := stable.ExportGenesis(s.ctx, s.app.StableKeeper)
	s.Require().NotNil(got)
	s.Require().Equal(genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(StableGenesisTestSuite))
}
