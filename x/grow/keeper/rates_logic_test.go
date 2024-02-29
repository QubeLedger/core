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
	third_address := apptesting.CreateRandomAccounts(1)[0]

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
					amount:      10000 * 1000000,
					Address:     main_address,
					denom:       "uusdc",
					OracleDenom: "USDC",
				},
				{
					amount:      10000 * 1000000,
					Address:     second_address,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
				{
					amount:      10000 * 1000000,
					Address:     third_address,
					denom:       "uusdc",
					OracleDenom: "USDC",
				},
			},
			[]Borrow_type_with_address{
				{
					amount:      5000 * 1000000,
					Address:     main_address,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
				{
					amount:      5000 * 1000000,
					Address:     second_address,
					denom:       "uusdc",
					OracleDenom: "USDC",
				},
				{
					amount:      1000 * 1000000,
					Address:     third_address,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
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
		params.MaxRateVolatile = 200
		params.UStaticStable = 80
		params.MaxRateStable = 100
		params.Slope_1 = 1
		params.Slope_2 = 8
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

			for _, asset := range tc.Asset {
				asset, _ := s.app.GrowKeeper.GetAssetByAssetId(s.ctx, asset.Asset.AssetId)
				utilization_rate := (float64(asset.CollectivelyBorrowValue) / float64(asset.ProvideValue))
				bir, sir, _ := s.app.GrowKeeper.GetRatesByUtilizationRate(s.ctx, utilization_rate, asset)
				fmt.Printf("Asset Ticker:  %v\n", asset.OracleAssetId)
				fmt.Printf("UR:  %f\n", utilization_rate)
				fmt.Printf("Borrow Interest Rate:  %f\n", bir)
				fmt.Printf("Supply Interest Rate: %f\n", sir)
				fmt.Printf("\n")
			}

			s.ctx = s.ctx.WithBlockTime(time.Unix((s.ctx.BlockTime().Unix() + (2419200)), 0))

			err := grow.EndBlocker(s.ctx, s.app.GrowKeeper)
			s.Require().NoError(err)

			position, _ := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, s.app.GrowKeeper.CalculateDepositId(main_address.String()))
			fmt.Printf("position.LendAmountInUSD: %v\n", position.LendAmountInUSD)
			fmt.Printf("position.BorrowedAmountInUSD: %v\n", position.BorrowedAmountInUSD)
			fmt.Printf("\n")

		})
	}
}
