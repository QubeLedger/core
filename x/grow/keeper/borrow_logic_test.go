package keeper_test

import (
	"fmt"
	"time"

	"github.com/QuadrateOrg/core/app/apptesting"
	"github.com/QuadrateOrg/core/x/grow/types"
	stabletypes "github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *GrowKeeperTestSuite) TestExecuteCreateBorrow() {
	testCases := []struct {
		name        string
		qStablePair stabletypes.Pair
		gTokenPair  types.GTokenPair
		Asset       []Asset_type
		lends       []Lend_type
		borrows     []Borrow_type
		err         bool
		errString   string
	}{
		{
			"ok-create-borrow",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "45000",
				},
			},
			[]Lend_type{
				{
					amount:      1 * 1000000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			[]Borrow_type{
				{
					amount:      1 * 100000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			false,
			"",
		},
		{
			"ok-create-2-borrow",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "45000",
				},
			},
			[]Lend_type{
				{
					amount:      1 * 1000000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			[]Borrow_type{
				{
					amount:      1 * 100000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
				{
					amount:      2 * 100000,
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
		suite.app.GrowKeeper.ChangeDepositMethodStatus(s.ctx)
		suite.app.GrowKeeper.ChangeCollateralMethodStatus(s.ctx)
		suite.app.GrowKeeper.ChangeBorrowMethodStatus(s.ctx)
		s.ctx = s.ctx.WithBlockTime(time.Now())

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		suite.app.GrowKeeper.AppendPair(s.ctx, tc.gTokenPair)

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
			suite.AddTestCoinsToCustomAccount(sdk.NewInt(lend_type.amount), lend_type.denom, s.Address)
			suite.AddTestCoins(lend_type.amount, lend_type.denom)
		}

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {

			for _, lend_type := range tc.lends {
				msg := types.NewMsgCreateLend(
					suite.Address.String(),
					sdk.NewInt(lend_type.amount).String()+lend_type.denom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				res, err := suite.app.GrowKeeper.CreateLend(ctx, msg)
				suite.Require().NoError(err)

				position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
				suite.Require().Equal(true, found)
				suite.Require().Equal(res.PositionId, position.DepositId)

				if len(tc.lends) == 1 {
					lend, found := s.app.GrowKeeper.GetLendByLendId(s.ctx, s.app.GrowKeeper.CalculateLendId(s.Address.String(), lend_type.denom, position.DepositId))
					suite.Require().Equal(true, found)
					suite.Require().Equal(sdk.NewInt(lend_type.amount).String()+lend_type.denom, lend.AmountIn)
					suite.Require().Equal(sdk.NewInt(lend_type.amount), lend.AmountInAmount.RoundInt())
				} else {
					lend, found := s.app.GrowKeeper.GetLendByLendId(s.ctx, s.app.GrowKeeper.CalculateLendId(s.Address.String(), lend_type.denom, position.DepositId))
					suite.Require().Equal(true, found)
					suite.Require().GreaterOrEqual(lend.AmountInAmount.RoundInt().Int64(), lend_type.amount)
				}
			}

			position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, s.app.GrowKeeper.CalculateDepositId(s.Address.String()))
			suite.Require().Equal(true, found)
			LendAmountInUsd := 0
			ProvideValue := 0

			for _, lend_type := range tc.lends {
				price, _ := s.app.GrowKeeper.GetPriceByDenom(s.ctx, lend_type.OracleDenom)
				LendAmountInUsd += (int(lend_type.amount) * int(price.Int64())) / 10000
				ProvideValue += int(lend_type.amount)
			}
			suite.Require().Equal(LendAmountInUsd, int(position.LendAmountInUSD))

			asset, _ := s.app.GrowKeeper.GetAssetByAssetId(s.ctx, s.GetNormalAsset(0).AssetId)
			suite.Require().Equal(ProvideValue, int(asset.ProvideValue))

			for _, borrow_type := range tc.borrows {
				msg := types.NewMsgCreateBorrow(
					s.Address.String(),
					borrow_type.denom,
					sdk.NewInt(borrow_type.amount).String()+borrow_type.denom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				res, err := suite.app.GrowKeeper.CreateBorrow(ctx, msg)
				if !tc.err {
					suite.Require().NoError(err)
					suite.Require().Equal(res.AmountOut, sdk.NewInt(borrow_type.amount).String()+borrow_type.denom)
					loan, found := s.app.GrowKeeper.GetLoadByLoanId(s.ctx, res.LoanId)
					suite.Require().Equal(found, true)
					suite.Require().Equal(loan.Borrower, s.Address.String())
					suite.Require().Equal(loan.StartTime, uint64(s.ctx.BlockTime().Unix()))
				} else {
					suite.Require().Error(err, tc.errString)
				}
			}

			if !tc.err {
				position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, s.app.GrowKeeper.CalculateDepositId(s.Address.String()))
				suite.Require().Equal(true, found)
				BorrowAmountInUsd := 0
				BorrowValue := 0

				for _, borrow_type := range tc.borrows {
					price, _ := s.app.GrowKeeper.GetPriceByDenom(s.ctx, borrow_type.OracleDenom)
					BorrowAmountInUsd += (int(borrow_type.amount) * int(price.Int64())) / 10000
					BorrowValue += int(borrow_type.amount)
				}
				suite.Require().Equal(BorrowAmountInUsd, int(position.BorrowedAmountInUSD))

				asset, _ := s.app.GrowKeeper.GetAssetByAssetId(s.ctx, s.GetNormalAsset(0).AssetId)
				suite.Require().Equal(BorrowValue, int(asset.CollectivelyBorrowValue))
			}
		})
	}
}

func (suite *GrowKeeperTestSuite) TestExecuteDeleteBorrow() {
	testCases := []struct {
		name        string
		qStablePair stabletypes.Pair
		gTokenPair  types.GTokenPair
		Asset       []Asset_type
		lends       []Lend_type
		borrows     []Borrow_type
		withdrawals []Withdrawal_Borrow_type
		err         bool
		errString   string
	}{
		{
			"ok-delete-1-borrow",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "45000",
				},
			},
			[]Lend_type{
				{
					amount:      1 * 1000000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			[]Borrow_type{
				{
					amount:      1 * 100000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			[]Withdrawal_Borrow_type{
				{
					amount:      1 * 100000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			false,
			"",
		},
		{
			"ok-delete-part-borrow",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "45000",
				},
			},
			[]Lend_type{
				{
					amount:      1 * 1000000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			[]Borrow_type{
				{
					amount:      1 * 100000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			[]Withdrawal_Borrow_type{
				{
					amount:      5 * 10000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			false,
			"",
		},
		{
			"ok-delete-2-borrow",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "45000",
				},
			},
			[]Lend_type{
				{
					amount:      1 * 1000000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			[]Borrow_type{
				{
					amount:      1 * 100000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			[]Withdrawal_Borrow_type{
				{
					amount:      5 * 10000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
				{
					amount:      5 * 10000,
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
		suite.app.GrowKeeper.ChangeDepositMethodStatus(s.ctx)
		suite.app.GrowKeeper.ChangeCollateralMethodStatus(s.ctx)
		suite.app.GrowKeeper.ChangeBorrowMethodStatus(s.ctx)
		s.ctx = s.ctx.WithBlockTime(time.Now())

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		suite.app.GrowKeeper.AppendPair(s.ctx, tc.gTokenPair)

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
			suite.AddTestCoinsToCustomAccount(sdk.NewInt(lend_type.amount), lend_type.denom, s.Address)
			suite.AddTestCoins(lend_type.amount, lend_type.denom)
		}

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {

			for _, lend_type := range tc.lends {
				msg := types.NewMsgCreateLend(
					suite.Address.String(),
					sdk.NewInt(lend_type.amount).String()+lend_type.denom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				res, err := suite.app.GrowKeeper.CreateLend(ctx, msg)
				suite.Require().NoError(err)

				position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
				suite.Require().Equal(true, found)
				suite.Require().Equal(res.PositionId, position.DepositId)

				if len(tc.lends) == 1 {
					lend, found := s.app.GrowKeeper.GetLendByLendId(s.ctx, s.app.GrowKeeper.CalculateLendId(s.Address.String(), lend_type.denom, position.DepositId))
					suite.Require().Equal(true, found)
					suite.Require().Equal(sdk.NewInt(lend_type.amount).String()+lend_type.denom, lend.AmountIn)
					suite.Require().Equal(sdk.NewInt(lend_type.amount), lend.AmountInAmount.RoundInt())
				} else {
					lend, found := s.app.GrowKeeper.GetLendByLendId(s.ctx, s.app.GrowKeeper.CalculateLendId(s.Address.String(), lend_type.denom, position.DepositId))
					suite.Require().Equal(true, found)
					suite.Require().GreaterOrEqual(lend.AmountInAmount.RoundInt().Int64(), lend_type.amount)
				}
			}

			position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, s.app.GrowKeeper.CalculateDepositId(s.Address.String()))
			suite.Require().Equal(true, found)
			LendAmountInUsd := 0
			ProvideValue := 0

			for _, lend_type := range tc.lends {
				price, _ := s.app.GrowKeeper.GetPriceByDenom(s.ctx, lend_type.OracleDenom)
				LendAmountInUsd += (int(lend_type.amount) * int(price.Int64())) / 10000
				ProvideValue += int(lend_type.amount)
			}
			suite.Require().Equal(LendAmountInUsd, int(position.LendAmountInUSD))

			asset, _ := s.app.GrowKeeper.GetAssetByAssetId(s.ctx, s.GetNormalAsset(0).AssetId)
			suite.Require().Equal(ProvideValue, int(asset.ProvideValue))

			for _, borrow_type := range tc.borrows {
				msg := types.NewMsgCreateBorrow(
					s.Address.String(),
					borrow_type.denom,
					sdk.NewInt(borrow_type.amount).String()+borrow_type.denom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				res, err := suite.app.GrowKeeper.CreateBorrow(ctx, msg)
				if !tc.err {
					suite.Require().NoError(err)
					suite.Require().Equal(res.AmountOut, sdk.NewInt(borrow_type.amount).String()+borrow_type.denom)
					loan, found := s.app.GrowKeeper.GetLoadByLoanId(s.ctx, res.LoanId)
					suite.Require().Equal(found, true)
					suite.Require().Equal(loan.Borrower, s.Address.String())
					suite.Require().Equal(loan.StartTime, uint64(s.ctx.BlockTime().Unix()))
				} else {
					suite.Require().Error(err, tc.errString)
				}
			}

			BorrowAmountInUsd := 0
			BorrowValue := 0

			if !tc.err {
				position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, s.app.GrowKeeper.CalculateDepositId(s.Address.String()))
				suite.Require().Equal(true, found)

				for _, borrow_type := range tc.borrows {
					price, _ := s.app.GrowKeeper.GetPriceByDenom(s.ctx, borrow_type.OracleDenom)
					BorrowAmountInUsd += (int(borrow_type.amount) * int(price.Int64())) / 10000
					BorrowValue += int(borrow_type.amount)
				}
				suite.Require().Equal(BorrowAmountInUsd, int(position.BorrowedAmountInUSD))

				asset, _ := s.app.GrowKeeper.GetAssetByAssetId(s.ctx, s.GetNormalAsset(0).AssetId)
				suite.Require().Equal(BorrowValue, int(asset.CollectivelyBorrowValue))
			}

			for _, w_type := range tc.withdrawals {
				loan_id := s.app.GrowKeeper.GenerateLoanIdHash(w_type.denom, s.Address.String())
				old_loan, _ := s.app.GrowKeeper.GetLoadByLoanId(s.ctx, loan_id)

				msg := types.NewMsgDeleteBorrow(
					suite.Address.String(),
					w_type.denom,
					sdk.NewInt(w_type.amount).String()+w_type.denom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.GrowKeeper.DeleteBorrow(ctx, msg)
				if !tc.err {
					suite.Require().NoError(err)

					if w_type.amount == old_loan.AmountOutAmount.RoundInt().Int64() {
						_, found := s.app.GrowKeeper.GetLoadByLoanId(s.ctx, loan_id)
						suite.Require().Equal(false, found)
					} else {
						loan, found := s.app.GrowKeeper.GetLoadByLoanId(s.ctx, loan_id)
						suite.Require().Equal(true, found)
						suite.Require().Equal(old_loan.AmountOutAmount.RoundInt().Sub(sdk.NewInt(w_type.amount)), loan.AmountOutAmount.RoundInt())
					}

				} else {
					suite.Require().Error(err, tc.errString)
				}
			}

			if !tc.err {
				position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, s.app.GrowKeeper.CalculateDepositId(s.Address.String()))
				suite.Require().Equal(true, found)

				for _, w_type := range tc.withdrawals {
					price, _ := s.app.GrowKeeper.GetPriceByDenom(s.ctx, w_type.OracleDenom)
					BorrowAmountInUsd -= (int(w_type.amount) * int(price.Int64())) / 10000
					BorrowValue -= int(w_type.amount)
				}

				suite.Require().Equal(BorrowAmountInUsd, int(position.BorrowedAmountInUSD))

				asset, _ := s.app.GrowKeeper.GetAssetByAssetId(s.ctx, s.GetNormalAsset(0).AssetId)
				suite.Require().Equal(BorrowValue, int(asset.CollectivelyBorrowValue))
			}
		})
	}
}

func (suite *GrowKeeperTestSuite) TestExecuteCreateBorrowInAnotherToken() {

	main_address := apptesting.CreateRandomAccounts(1)[0]

	testCases := []struct {
		name        string
		qStablePair stabletypes.Pair
		gTokenPair  types.GTokenPair
		Asset       []Asset_type
		lends       []Lend_type_with_address
		borrows     []Borrow_type_with_address
		err         bool
		errString   string
	}{
		{
			"ok-create-borrow",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "45000",
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
					amount:      2 * 1000000,
					Address:     apptesting.CreateRandomAccounts(1)[0],
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			[]Borrow_type_with_address{
				{
					amount:      1 * 100000,
					Address:     main_address,
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
		suite.app.GrowKeeper.ChangeDepositMethodStatus(s.ctx)
		suite.app.GrowKeeper.ChangeCollateralMethodStatus(s.ctx)
		suite.app.GrowKeeper.ChangeBorrowMethodStatus(s.ctx)
		s.ctx = s.ctx.WithBlockTime(time.Now())

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		suite.app.GrowKeeper.AppendPair(s.ctx, tc.gTokenPair)

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

				if len(tc.lends) == 1 {
					lend, found := s.app.GrowKeeper.GetLendByLendId(s.ctx, s.app.GrowKeeper.CalculateLendId(lend_type.Address.String(), lend_type.denom, position.DepositId))
					suite.Require().Equal(true, found)
					suite.Require().Equal(sdk.NewInt(lend_type.amount).String()+lend_type.denom, lend.AmountIn)
					suite.Require().Equal(sdk.NewInt(lend_type.amount), lend.AmountInAmount.RoundInt())
				} else {
					lend, found := s.app.GrowKeeper.GetLendByLendId(s.ctx, s.app.GrowKeeper.CalculateLendId(lend_type.Address.String(), lend_type.denom, position.DepositId))
					suite.Require().Equal(true, found)
					suite.Require().GreaterOrEqual(lend.AmountInAmount.RoundInt().Int64(), lend_type.amount)
				}
			}

			for _, borrow_type := range tc.borrows {
				msg := types.NewMsgCreateBorrow(
					borrow_type.Address.String(),
					borrow_type.denom,
					sdk.NewInt(borrow_type.amount).String()+borrow_type.denom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				res, err := suite.app.GrowKeeper.CreateBorrow(ctx, msg)
				if !tc.err {
					suite.Require().NoError(err)
					suite.Require().Equal(res.AmountOut, sdk.NewInt(borrow_type.amount).String()+borrow_type.denom)
					loan, found := s.app.GrowKeeper.GetLoadByLoanId(s.ctx, res.LoanId)
					suite.Require().Equal(found, true)
					suite.Require().Equal(loan.Borrower, borrow_type.Address.String())
					suite.Require().Equal(loan.StartTime, uint64(s.ctx.BlockTime().Unix()))
				} else {
					suite.Require().Error(err, tc.errString)
				}
			}

			if !tc.err {
				position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, s.app.GrowKeeper.CalculateDepositId(main_address.String()))
				suite.Require().Equal(true, found)
				BorrowAmountInUsd := 0
				BorrowValue := 0

				for _, borrow_type := range tc.borrows {
					price, _ := s.app.GrowKeeper.GetPriceByDenom(s.ctx, borrow_type.OracleDenom)
					BorrowAmountInUsd += (int(borrow_type.amount) * int(price.Int64())) / 10000
					BorrowValue += int(borrow_type.amount)
				}
				suite.Require().Equal(BorrowAmountInUsd, int(position.BorrowedAmountInUSD))

				asset, _ := s.app.GrowKeeper.GetAssetByAssetId(s.ctx, s.GetNormalAsset(0).AssetId)
				suite.Require().Equal(BorrowValue, int(asset.CollectivelyBorrowValue))
			}
		})
	}
}
