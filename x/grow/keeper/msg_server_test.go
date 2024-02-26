package keeper_test

import (
	"fmt"

	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *GrowKeeperTestSuite) TestGrowDepositDeactivate() {

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("ATOM")
	suite.RegisterValidator()

	config := s.GetNormalConfig()

	suite.Run(fmt.Sprintf("Grow Deposit Deactivated"), func() {
		suite.AddTestCoins(config.sendTokenAmount, config.sendTokenDenom)
		msg := types.NewMsgGrowDeposit(
			suite.Address.String(),
			sdk.NewInt(config.sendTokenAmount).String()+config.sendTokenDenom,
			s.GetNormalGTokenPair(0).GTokenMetadata.Base,
		)
		ctx := sdk.WrapSDKContext(suite.ctx)
		_, err := suite.app.GrowKeeper.GrowDeposit(ctx, msg)
		suite.Require().Error(err, types.ErrDepositNotActivated)
	})

}

func (suite *GrowKeeperTestSuite) TestGrowCollateralDeactivate() {

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("ATOM")
	suite.RegisterValidator()

	config := s.GetNormalConfig()

	suite.Run(fmt.Sprintf("Grow Collateral Deactivated"), func() {
		suite.AddTestCoins(config.sendTokenAmount, config.sendTokenDenom)
		msg := types.NewMsgCreateLend(
			suite.Address.String(),
			sdk.NewInt(config.sendTokenAmount).String()+config.sendTokenDenom,
		)
		ctx := sdk.WrapSDKContext(suite.ctx)
		_, err := suite.app.GrowKeeper.CreateLend(ctx, msg)
		suite.Require().Error(err, types.ErrCollateralNotActivated)
	})

}

func (suite *GrowKeeperTestSuite) TestGrowBorrowDeactivate() {

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("ATOM")
	suite.RegisterValidator()

	config := s.GetNormalConfig()

	suite.Run(fmt.Sprintf("Grow Borrow Deactivated"), func() {
		suite.AddTestCoins(config.sendTokenAmount, config.sendTokenDenom)
		msg := types.NewMsgCreateBorrow(
			suite.Address.String(),
			config.sendTokenDenom,
			sdk.NewInt(config.lendTokenAmount).String(),
		)
		ctx := sdk.WrapSDKContext(suite.ctx)
		_, err := suite.app.GrowKeeper.CreateBorrow(ctx, msg)
		suite.Require().Error(err, types.ErrBorrowNotActivated)
	})

}
