package keeper_test

import sdk "github.com/cosmos/cosmos-sdk/types"

func (suite *GrowKeeperTestSuite) TestSetRealRate() {
	suite.Setup()
	suite.Commit()

	suite.app.GrowKeeper.AppendPair(s.ctx, suite.GetNormalGTokenPair(0))

	s.Require().Equal(s.app.GrowKeeper.GetRealRate(s.ctx, suite.GetNormalGTokenPair(0).DenomID).Int64(), int64(1))
	s.app.GrowKeeper.SetRealRate(s.ctx, sdk.NewInt(15), suite.GetNormalGTokenPair(0).DenomID)
	s.Require().Equal(s.app.GrowKeeper.GetRealRate(s.ctx, suite.GetNormalGTokenPair(0).DenomID).Int64(), int64(15))
}

func (suite *GrowKeeperTestSuite) TestSetBorrowRate() {
	suite.Setup()
	suite.Commit()

	suite.app.GrowKeeper.AppendPair(s.ctx, suite.GetNormalGTokenPair(0))

	s.Require().Equal(s.app.GrowKeeper.GetBorrowRate(s.ctx, suite.GetNormalGTokenPair(0).DenomID).Int64(), int64(1))
	s.app.GrowKeeper.SetBorrowRate(s.ctx, sdk.NewInt(25), suite.GetNormalGTokenPair(0).DenomID)
	s.Require().Equal(s.app.GrowKeeper.GetBorrowRate(s.ctx, suite.GetNormalGTokenPair(0).DenomID).Int64(), int64(25))
}

func (suite *GrowKeeperTestSuite) TestSetLastTimeUpdateReserve() {
	suite.Setup()
	suite.Commit()

	s.Require().Equal(s.app.GrowKeeper.GetLastTimeUpdateReserve(s.ctx).Int64(), int64(0))
	s.app.GrowKeeper.SetLastTimeUpdateReserve(s.ctx, sdk.NewInt(15))
	s.Require().Equal(s.app.GrowKeeper.GetLastTimeUpdateReserve(s.ctx).Int64(), int64(15))
}
