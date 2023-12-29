package keeper_test

import "github.com/QuadrateOrg/core/app/apptesting"

func (suite *GrowKeeperTestSuite) TestSetUSQReserveAddress() {
	suite.Setup()
	suite.Commit()

	addr := apptesting.CreateRandomAccounts(1)[0]

	s.app.GrowKeeper.SetUSQReserveAddress(s.ctx, addr)
	s.Require().Equal(s.app.GrowKeeper.GetUSQReserveAddress(s.ctx), addr)
}

func (suite *GrowKeeperTestSuite) TestSetGrowYieldReserveAddress() {
	suite.Setup()
	suite.Commit()

	addr := apptesting.CreateRandomAccounts(1)[0]

	s.app.GrowKeeper.SetGrowYieldReserveAddress(s.ctx, addr)
	s.Require().Equal(s.app.GrowKeeper.GetGrowYieldReserveAddress(s.ctx), addr)
}

func (suite *GrowKeeperTestSuite) TestSetGrowStakingReserveAddress() {
	suite.Setup()
	suite.Commit()

	addr := apptesting.CreateRandomAccounts(1)[0]

	s.app.GrowKeeper.SetGrowStakingReserveAddress(s.ctx, addr)
	s.Require().Equal(s.app.GrowKeeper.GetGrowStakingReserveAddress(s.ctx), addr)
}
