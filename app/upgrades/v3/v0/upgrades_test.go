package v0_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"time"

	"github.com/QuadrateOrg/core/app"
	quadrateapptest "github.com/QuadrateOrg/core/app/helpers"
	apptypes "github.com/QuadrateOrg/core/types"
	lsmtypes "github.com/QuadrateOrg/core/x/liquidstakeibc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/tendermint/tendermint/abci/types"
)

const (
	UpgradeHeight = 15
)

type UpgradeTestSuite struct {
	suite.Suite
	App *app.QuadrateApp
	ctx sdk.Context
}

func (s *UpgradeTestSuite) SetupTest() {
	s.Setup()
}

var s *UpgradeTestSuite

func (s *UpgradeTestSuite) Setup() {
	apptypes.SetConfig()
	s.App = quadrateapptest.Setup(s.T(), "qube-1", false, 1)
	s.ctx = s.App.BaseApp.NewContext(false, s.ctx.BlockHeader())
	s.ctx = s.ctx.WithBlockHeight(0)
}

func TestUpgradeTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (s *UpgradeTestSuite) runV_0_3_0_Upgrade() {
	s.ctx = s.ctx.WithBlockHeight(UpgradeHeight - 1)
	plan := upgradetypes.Plan{Name: "v0.3.0", Height: UpgradeHeight}
	err := s.App.UpgradeKeeper.ScheduleUpgrade(s.ctx, plan)
	s.Require().NoError(err)
	_, exists := s.App.UpgradeKeeper.GetUpgradePlan(s.ctx)
	s.Require().True(exists)

	s.ctx = s.ctx.WithBlockHeight(UpgradeHeight)
}

func (s *UpgradeTestSuite) TestUpgrade() {
	s.Setup()

	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(time.Hour * 24))
	s.App.BeginBlocker(s.ctx, types.RequestBeginBlock{})

	s.runV_0_3_0_Upgrade()
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(time.Hour * 24 * 7))
	s.App.BeginBlocker(s.ctx, types.RequestBeginBlock{})

	params := s.App.LiquidStakeIBCKeeper.GetParams(s.ctx)
	s.Require().Equal(params, lsmtypes.DefaultParams())
}
