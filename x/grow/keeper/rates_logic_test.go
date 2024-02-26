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

func (suite *GrowKeeperTestSuite) TestExecuteCalculateUtilizationRate() {
	main_address := apptesting.CreateRandomAccounts(1)[0]
	second_address := apptesting.CreateRandomAccounts(1)[0]

	testCases := []struct {
		name         string
		qStablePair  stabletypes.Pair
		gTokenPair   types.GTokenPair
		Asset        []Asset_type
		Price_change []Asset_type
		lends        []Lend_type_with_address
		borrows      []Borrow_type_with_address
		err          bool
		errString    string
	}{
		{
			"ok-calculate-ur",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
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
					amount:      1000 * 1000000,
					Address:     second_address,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			[]Borrow_type_with_address{
				{
					amount:      1 * 1000000,
					Address:     main_address,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
				{
					amount:      1 * 1000000,
					Address:     second_address,
					denom:       "uusdc",
					OracleDenom: "USDC",
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

			s.ctx = s.ctx.WithBlockTime(time.Unix((s.ctx.BlockTime().Unix() + 31536000), 0))

			err := grow.EndBlocker(s.ctx, s.app.GrowKeeper)
			s.Require().NoError(err)

			position, _ := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, s.app.GrowKeeper.CalculateDepositId(main_address.String()))
			fmt.Printf("position.LendAmountInUSD: %v\n", position.LendAmountInUSD)
			fmt.Printf("position.BorrowedAmountInUSD: %v\n", position.BorrowedAmountInUSD)
		})
	}
}
