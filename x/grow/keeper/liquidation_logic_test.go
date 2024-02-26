package keeper_test

import (
	"fmt"
	"time"

	"github.com/QuadrateOrg/core/app/apptesting"
	"github.com/QuadrateOrg/core/x/grow"
	"github.com/QuadrateOrg/core/x/grow/types"
	stabletypes "github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *GrowKeeperTestSuite) TestExecuteOpenLiqPosition() {
	testCases := []struct {
		name            string
		qStablePair     stabletypes.Pair
		gTokenPair      types.GTokenPair
		Asset           types.Asset
		sendTokenDenom  string
		sendTokenAmount int64
		premium         string
		err             bool
		errString       string
	}{
		{
			"ok-create-liq-position",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalAsset(0),
			"uusdc",
			500 * 1000000,
			"5",
			false,
			"",
		},
		{
			"ok-create-liq-position",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalAsset(0),
			"uusdc",
			800 * 1000000,
			"5",
			false,
			"",
		},
		{
			"false-send wrong premuim",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalAsset(0),
			"uusdc",
			500 * 1000000,
			"a",
			true,
			"ErrWrongPremium err",
		},
		{
			"false-send wrong premuim",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalAsset(0),
			"ueur",
			500 * 1000000,
			"5",
			true,
			"ErrAssetNotFound err",
		},
	}

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("WBTC")
	suite.RegisterValidator()
	suite.app.GrowKeeper.ChangeDepositMethodStatus(s.ctx)
	suite.app.GrowKeeper.ChangeCollateralMethodStatus(s.ctx)
	suite.app.GrowKeeper.ChangeBorrowMethodStatus(s.ctx)
	s.ctx = s.ctx.WithBlockTime(time.Now())
	for _, tc := range testCases {

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		suite.app.GrowKeeper.AppendPair(s.ctx, tc.gTokenPair)
		suite.app.GrowKeeper.AppendAsset(s.ctx, tc.Asset)
		suite.app.GrowKeeper.AppendAsset(s.ctx, s.GetSecondStableNormalAsset(0))

		suite.OracleAggregateExchangeRateFromInput("0.5" + tc.Asset.AssetMetadata.Name)

		suite.AddTestCoinsToCustomAccount(sdk.NewInt(100000*1000000), tc.qStablePair.AmountOutMetadata.Base, s.app.GrowKeeper.GetGrowStakingReserveAddress(s.ctx))

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.AddTestCoins(tc.sendTokenAmount, tc.sendTokenDenom)
			msg := types.NewMsgOpenLiquidationPosition(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				tc.Asset.OracleAssetId,
				tc.premium,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.app.GrowKeeper.OpenLiquidationPosition(ctx, msg)
			if !tc.err {
				suite.Require().NoError(err)

				position, found := s.app.GrowKeeper.GetLiquidatorPositionByLiquidatorPositionId(s.ctx, res.LiquidatorPositionId)
				suite.Require().Equal(found, true)

				suite.Require().Equal(s.GetSecondStableNormalAsset(0).OracleAssetId, position.ProvidedAssetId)
				suite.Require().Equal(tc.Asset.OracleAssetId, position.WantAssetId)

				suite.Require().Equal(position.Liquidator, s.Address.String())
				suite.Require().Equal(position.Amount, sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom)
				premiumInt, _ := s.app.GrowKeeper.ParseAndCheckPremium(tc.premium)
				suite.Require().Equal(position.Premium, premiumInt.Uint64())
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}

func (suite *GrowKeeperTestSuite) TestExecuteCloseLiqPosition() {
	testCases := []struct {
		name                 string
		qStablePair          stabletypes.Pair
		gTokenPair           types.GTokenPair
		Asset                types.Asset
		sendTokenDenom       string
		sendTokenAmount      int64
		premium              string
		LiquidatorPositionId string
		err                  bool
		errString            string
	}{
		{
			"ok-create-liq-position",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalAsset(0),
			"uusdc",
			500 * 1000000,
			"5",
			"",
			false,
			"",
		},
		{
			"ok-wrong liquidatorPositionId",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalAsset(0),
			"uusdc",
			500 * 1000000,
			"5",
			"test",
			true,
			"ErrLiqPositionNotFound err",
		},
	}

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("WBTC")
	suite.RegisterValidator()
	suite.app.GrowKeeper.ChangeDepositMethodStatus(s.ctx)
	suite.app.GrowKeeper.ChangeCollateralMethodStatus(s.ctx)
	suite.app.GrowKeeper.ChangeBorrowMethodStatus(s.ctx)
	s.ctx = s.ctx.WithBlockTime(time.Now())

	suite.app.StableKeeper.AppendPair(s.ctx, s.GetNormalQStablePair(0))
	suite.app.GrowKeeper.AppendPair(s.ctx, s.GetNormalGTokenPair(0))
	suite.app.GrowKeeper.AppendAsset(s.ctx, s.GetNormalAsset(0))
	suite.app.GrowKeeper.AppendAsset(s.ctx, s.GetSecondStableNormalAsset(0))

	suite.OracleAggregateExchangeRateFromInput("0.5" + s.GetNormalAsset(0).AssetMetadata.Name)

	suite.AddTestCoinsToCustomAccount(sdk.NewInt(100000*1000000), s.GetNormalQStablePair(0).AmountOutMetadata.Base, s.app.GrowKeeper.GetGrowStakingReserveAddress(s.ctx))

	suite.AddTestCoins(500*1000000, "uusdc")
	msg := types.NewMsgOpenLiquidationPosition(
		suite.Address.String(),
		sdk.NewInt(500*1000000).String()+"uusdc",
		s.GetNormalAsset(0).AssetMetadata.Name,
		"5",
	)
	ctx := sdk.WrapSDKContext(suite.ctx)
	res, err := suite.app.GrowKeeper.OpenLiquidationPosition(ctx, msg)
	suite.Require().NoError(err)

	for _, tc := range testCases {
		if !tc.err {
			tc.LiquidatorPositionId = res.LiquidatorPositionId
		}
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {

			msg := types.NewMsgCloseLiquidationPosition(
				suite.Address.String(),
				tc.LiquidatorPositionId,
			)
			ctx = sdk.WrapSDKContext(suite.ctx)
			res, err := suite.app.GrowKeeper.CloseLiquidationPosition(ctx, msg)
			if !tc.err {
				suite.Require().NoError(err)

				_, found := s.app.GrowKeeper.GetLiquidatorPositionByLiquidatorPositionId(s.ctx, tc.LiquidatorPositionId)
				suite.Require().Equal(found, false)

				amountOut, _ := sdk.ParseCoinsNormalized(res.AmountOut)

				suite.Require().Equal(sdk.NewInt(500*1000000), amountOut.AmountOf("uusdc"))
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}

func (suite *GrowKeeperTestSuite) TestLiquidatePositionFull() {
	main_address := apptesting.CreateRandomAccounts(1)[0]

	testCases := []struct {
		name                 string
		qStablePair          stabletypes.Pair
		gTokenPair           types.GTokenPair
		Asset                []Asset_type
		Price_change         []Asset_type
		lends                []Lend_type_with_address
		borrows              []Borrow_type_with_address
		liqudation_positions []Liquidation_msg_data
		err                  bool
		errString            string
	}{
		{
			"ok-liqudation-full-loan-1-token",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "0.1",
				},
				{
					Asset: s.GetSecondStableNormalAsset(0),
					Price: "1",
				},
			},
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "1",
				},
				{
					Asset: s.GetSecondStableNormalAsset(0),
					Price: "1",
				},
			},
			[]Lend_type_with_address{
				{
					amount:      100 * 1000000,
					Address:     main_address,
					denom:       "uusdc",
					OracleDenom: "USDC",
				},
				{
					amount:      20000 * 1000000,
					Address:     apptesting.CreateRandomAccounts(1)[0],
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			[]Borrow_type_with_address{
				{
					amount:      60 * 1000000,
					Address:     main_address,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			[]Liquidation_msg_data{
				{
					sendAmount: 10 * 1000000,
					address:    apptesting.CreateRandomAccounts(1)[0],
					sendDenom:  "uwbtc",
					asset:      "USDC",
					premium:    "3",
				},
				{
					sendAmount: 1 * 1000000,
					address:    apptesting.CreateRandomAccounts(1)[0],
					sendDenom:  "uwbtc",
					asset:      "USDC",
					premium:    "5",
				},
			},
			false,
			"",
		},
		{
			"ok-liqudation-full-loan-2-token",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "0.1",
				},
				{
					Asset: s.GetSecondStableNormalAsset(0),
					Price: "1",
				},
				{
					Asset: s.GetThirdVolatileNormalAsset(0),
					Price: "0.1",
				},
			},
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "1",
				},
				{
					Asset: s.GetSecondStableNormalAsset(0),
					Price: "1",
				},
				{
					Asset: s.GetThirdVolatileNormalAsset(0),
					Price: "1",
				},
			},
			[]Lend_type_with_address{
				{
					amount:      100 * 1000000,
					Address:     main_address,
					denom:       "uusdc",
					OracleDenom: "USDC",
				},
				{
					amount:      20000 * 1000000,
					Address:     apptesting.CreateRandomAccounts(1)[0],
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
				{
					amount:      20000 * 1000000,
					Address:     apptesting.CreateRandomAccounts(1)[0],
					denom:       "uosmo",
					OracleDenom: "OSMO",
				},
			},
			[]Borrow_type_with_address{
				{
					amount:      30 * 1000000,
					Address:     main_address,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
				{
					amount:      30 * 1000000,
					Address:     main_address,
					denom:       "uosmo",
					OracleDenom: "OSMO",
				},
			},
			[]Liquidation_msg_data{
				{
					sendAmount: 10 * 1000000,
					address:    apptesting.CreateRandomAccounts(1)[0],
					sendDenom:  "uwbtc",
					asset:      "USDC",
					premium:    "3",
				},
				{
					sendAmount: 10 * 1000000,
					address:    apptesting.CreateRandomAccounts(1)[0],
					sendDenom:  "uosmo",
					asset:      "USDC",
					premium:    "4",
				},
				{
					sendAmount: 1 * 1000000,
					address:    apptesting.CreateRandomAccounts(1)[0],
					sendDenom:  "uwbtc",
					asset:      "USDC",
					premium:    "5",
				},
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {

		suite.Setup()
		suite.Commit()
		suite.RegisterValidator()
		suite.app.GrowKeeper.ChangeCollateralMethodStatus(s.ctx)
		suite.app.GrowKeeper.ChangeBorrowMethodStatus(s.ctx)
		s.ctx = s.ctx.WithBlockTime(time.Now())

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		suite.app.GrowKeeper.AppendPair(s.ctx, tc.gTokenPair)

		params := s.app.GrowKeeper.GetParams(s.ctx)
		params.LastTimeUpdateReserve = uint64(s.ctx.BlockTime().Unix())
		params.UStaticVolatile = 60
		params.MaxRateVolatile = 300
		params.Slope = 7
		s.app.GrowKeeper.SetParams(s.ctx, params)

		oracle_denom := ""
		for i, asset_type := range tc.Asset {
			suite.app.GrowKeeper.AppendAsset(s.ctx, asset_type.Asset)
			suite.SetupOracleKeeper(asset_type.Asset.OracleAssetId)
			oracle_denom += asset_type.Price + asset_type.Asset.OracleAssetId
			if i != len(tc.Asset)-1 {
				oracle_denom += ","
			}
		}
		suite.OracleAggregateExchangeRateFromInput(oracle_denom)

		for _, lend_type := range tc.lends {
			suite.AddTestCoinsToCustomAccount(sdk.NewInt(lend_type.amount), lend_type.denom, lend_type.Address)
			suite.AddTestCoins(lend_type.amount, lend_type.denom)
		}

		for _, liq_pos := range tc.liqudation_positions {
			suite.AddTestCoinsToCustomAccount(sdk.NewInt(liq_pos.sendAmount), liq_pos.sendDenom, liq_pos.address)
		}

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {

			for _, lend_type := range tc.lends {
				msg := types.NewMsgCreateLend(
					lend_type.Address.String(),
					sdk.NewInt(lend_type.amount).String()+lend_type.denom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				res, err := suite.app.GrowKeeper.CreateLend(ctx, msg)
				suite.Require().NoError(err)

				position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
				suite.Require().Equal(true, found)
				suite.Require().Equal(res.PositionId, position.DepositId)
			}

			for _, borrow_type := range tc.borrows {
				msg := types.NewMsgCreateBorrow(
					borrow_type.Address.String(),
					borrow_type.denom,
					sdk.NewInt(borrow_type.amount).String()+borrow_type.denom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				res, err := suite.app.GrowKeeper.CreateBorrow(ctx, msg)
				suite.Require().NoError(err)
				loan, found := s.app.GrowKeeper.GetLoadByLoanId(s.ctx, res.LoanId)
				suite.Require().Equal(found, true)
				suite.Require().Equal(loan.Borrower, borrow_type.Address.String())
				suite.Require().Equal(loan.StartTime, uint64(s.ctx.BlockTime().Unix()))
			}

			liq_pos_ids := []string{}

			for _, liq_pos := range tc.liqudation_positions {
				msg := types.NewMsgOpenLiquidationPosition(
					liq_pos.address.String(),
					sdk.NewInt(liq_pos.sendAmount).String()+liq_pos.sendDenom,
					liq_pos.asset,
					liq_pos.premium,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				res, err := suite.app.GrowKeeper.OpenLiquidationPosition(ctx, msg)
				suite.Require().NoError(err)
				position, found := s.app.GrowKeeper.GetLiquidatorPositionByLiquidatorPositionId(s.ctx, res.LiquidatorPositionId)
				suite.Require().Equal(found, true)
				suite.Require().Equal(liq_pos.asset, position.WantAssetId)

				liq_pos_ids = append(liq_pos_ids, res.LiquidatorPositionId)
			}

			new_oracle_denom := ""
			for i, price_change := range tc.Price_change {
				new_oracle_denom += price_change.Price + price_change.Asset.OracleAssetId
				if i != len(tc.Asset)-1 {
					new_oracle_denom += ","
				}
			}
			suite.OracleAggregateExchangeRateFromInput(new_oracle_denom)

			err := grow.EndBlocker(s.ctx, s.app.GrowKeeper)
			suite.Require().NoError(err)

			err = grow.EndBlocker(s.ctx, s.app.GrowKeeper)
			suite.Require().NoError(err)

			position, _ := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, s.app.GrowKeeper.CalculateDepositId(main_address.String()))
			rr, _ := s.app.GrowKeeper.CalculateRiskRate(sdk.NewIntFromUint64(position.LendAmountInUSD), sdk.NewIntFromUint64(position.BorrowedAmountInUSD))
			_ = rr
			/*
				fmt.Printf("Position LendAmountInUSD:: %v\n", position.LendAmountInUSD)
				fmt.Printf("Position BorrowedAmountInUSD:: %v\n", position.BorrowedAmountInUSD)
				fmt.Printf("Position RR: %v\n\n", rr)

				for _, liq_pos_id := range liq_pos_ids {
					liq_pos, _ := s.app.GrowKeeper.GetLiquidatorPositionByLiquidatorPositionId(s.ctx, liq_pos_id)
					asset, _ := s.app.GrowKeeper.GetAssetByOracleAssetId(s.ctx, liq_pos.WantAssetId)
					liqBalance1 := s.app.BankKeeper.GetBalance(s.ctx, sdk.AccAddress(liq_pos.Liquidator), asset.AssetMetadata.Base)
					fmt.Printf("Liquidator Position Ampunt: %s\n", liq_pos.Amount)
					fmt.Printf("Liquidator Balance: %s\n\n", liqBalance1.String())
				}
			*/
		})
	}

}
