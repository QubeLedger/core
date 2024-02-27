package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *GrowKeeperTestSuite) TestCalculateGTokenAmountOut() {
	s.Setup()
	amt := sdk.NewInt(100 * 1000000)
	price := sdk.NewInt(5 * 1000000)
	res := s.app.GrowKeeper.CalculateGTokenAmountOut(amt, price)
	s.Require().Equal(res.Int64(), int64(20*1000000))
}

func (s *GrowKeeperTestSuite) TestCalculateReturnQubeStableAmountOut() {
	s.Setup()
	amt := sdk.NewInt(100 * 1000000)
	price := sdk.NewInt(5 * 1000000)
	res := s.app.GrowKeeper.CalculateReturnQubeStableAmountOut(amt, price)
	s.Require().Equal(res.Int64(), int64(500*1000000))
}

func (s *GrowKeeperTestSuite) TestCalculateGTokenAPY() {
	s.Setup()
	lastAmount := sdk.NewInt(1 * 1000000)
	growRate := sdk.NewInt(150)
	day := sdk.NewInt(150)
	res := s.app.GrowKeeper.CalculateGTokenAPY(lastAmount, growRate, day)
	s.Require().Equal(res.Int64(), int64(1063133))
	day = sdk.NewInt(365)
	res = s.app.GrowKeeper.CalculateGTokenAPY(lastAmount, growRate, day)
	s.Require().Equal(res.Int64(), int64(1161321))
}

func (s *GrowKeeperTestSuite) TestCalculateAmountByPriceAndAmountIn() {
	s.Setup()
	amt := sdk.NewInt(100 * 1000000)
	price := sdk.NewInt(5 * 10000)
	res := s.app.GrowKeeper.CalculateAmountByPriceAndAmountIn(amt, price)
	s.Require().Equal(res.Int64(), int64(500*1000000))
}
func (s *GrowKeeperTestSuite) TestCalculateDeleteLendAmountOut() {
	s.Setup()
	amt := sdk.NewInt(100 * 1000000)
	price := sdk.NewInt(5 * 10000)
	res := s.app.GrowKeeper.CalculateDeleteLendAmountOut(amt, price)
	s.Require().Equal(res.Int64(), int64(20*1000000))
}
func (s *GrowKeeperTestSuite) TestCalculateNeedAmountToGet() {
	s.Setup()
	amt := sdk.NewInt(10 * 1000000)
	time := sdk.NewInt(31536000)
	res := s.app.GrowKeeper.CalculateNeedAmountToGet(amt, time)
	s.Require().Equal(res.Int64(), int64(11500000))
}
func (s *GrowKeeperTestSuite) TestCalculateRiskRate() {
	s.Setup()
	collateral := sdk.NewInt(100 * 1000000)
	borrow := sdk.NewInt(60 * 1000000)
	res, err := s.app.GrowKeeper.CalculateRiskRate(collateral, borrow)
	s.Require().NoError(err)
	s.Require().Equal(int64(100), res.Int64())
}
func (s *GrowKeeperTestSuite) TestCheckRiskRate() {
	s.Setup()
	collateral := sdk.NewInt(100 * 1000000)
	borrow := sdk.NewInt(20 * 1000000)
	desired := sdk.NewInt(5 * 1000000)
	err := s.app.GrowKeeper.CheckRiskRate(collateral, borrow, desired)
	s.Require().NoError(err)

	borrow = sdk.NewInt(50 * 1000000)
	desired = sdk.NewInt(20 * 1000000)
	err = s.app.GrowKeeper.CheckRiskRate(collateral, borrow, desired)
	s.Require().Error(err)
}
func (s *GrowKeeperTestSuite) TestCalculateAmountLiquidate() {
	s.Setup()
	collateral := sdk.NewInt(100 * 1000000)
	borrow := sdk.NewInt(60 * 1000000)
	res := s.app.GrowKeeper.CalculateAmountLiquidate(s.ctx, collateral.Int64(), borrow.Int64())
	s.Require().Equal(res.Int64(), int64(6976744))
}
func (s *GrowKeeperTestSuite) TestCalculatePremiumAmount() {
	s.Setup()
	amt := sdk.NewInt(100 * 1000000)
	price := sdk.NewInt(5 * 10000)
	price1 := sdk.NewInt(5 * 10000)
	premium := int64(3)
	res1, _ := s.app.GrowKeeper.CalculatePremiumAmount(s.ctx, amt, premium, price, price1)
	s.Require().Equal(res1.Int64(), int64(103000000))
}
