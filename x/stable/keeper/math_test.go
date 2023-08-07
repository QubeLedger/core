package keeper_test

import sdk "github.com/cosmos/cosmos-sdk/types"

func (suite *StableKeeperTestSuite) TestCalculateAmountToMint() {
	suite.Setup()
	atom := sdk.NewInt(1000).Mul(sdk.NewInt(1000000))
	price := sdk.NewInt(95000)
	fee := sdk.NewInt(3)
	res := suite.app.StableKeeper.CalculateAmountToMint(atom, price, fee)
	suite.Assert().Equal(res, sdk.NewInt(9471500000))
}

func (suite *StableKeeperTestSuite) TestCalculateMintingFeeForBurningFund() {
	suite.Setup()
	atom := sdk.NewInt(1000).Mul(sdk.NewInt(1000000))
	price := sdk.NewInt(95000)
	fee := sdk.NewInt(3)
	res := suite.app.StableKeeper.CalculateMintingFeeForBurningFund(atom, price, fee)
	suite.Assert().Equal(res, sdk.NewInt(28500000))
}

func (suite *StableKeeperTestSuite) TestCalculateAmountToSend() {
	suite.Setup()
	stable := sdk.NewInt(100).Mul(sdk.NewInt(1000000)) // 100 stable * 1*10**6
	price := sdk.NewInt(95000)
	fee := sdk.NewInt(3)
	res := suite.app.StableKeeper.CalculateAmountToSend(stable, price, fee)
	suite.Assert().Equal(res, sdk.NewInt(10494736))
}

func (suite *StableKeeperTestSuite) TestCalculateBurningFeeForBurningFund() {
	suite.Setup()
	stable := sdk.NewInt(100).Mul(sdk.NewInt(1000000))
	price := sdk.NewInt(95000)
	fee := sdk.NewInt(3)
	res := suite.app.StableKeeper.CalculateBurningFeeForBurningFund(stable, price, fee)
	suite.Assert().Equal(res, sdk.NewInt(31578))
}
