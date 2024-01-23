package v1_test

import (
	"testing"

	"github.com/QuadrateOrg/core/app/apptesting"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/suite"
)

const (
	v6UpgradeHeight = 15
)

type UpgradeTestSuite struct {
	apptesting.KeeperTestHelper
}

func (s *UpgradeTestSuite) SetupTest() {
	s.Setup()
}

func TestUpgradeTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (s *UpgradeTestSuite) runV_0_2_1_Upgrade() {
	s.Ctx = s.Ctx.WithBlockHeight(v6UpgradeHeight - 1)
	plan := upgradetypes.Plan{Name: "v0.2.1", Height: v6UpgradeHeight}
	err := s.App.UpgradeKeeper.ScheduleUpgrade(s.Ctx, plan)
	s.Require().NoError(err)
	_, exists := s.App.UpgradeKeeper.GetUpgradePlan(s.Ctx)
	s.Require().True(exists)

	s.Ctx = s.Ctx.WithBlockHeight(v6UpgradeHeight)
}

func (s *UpgradeTestSuite) TestUpgrade() {
}
