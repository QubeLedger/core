package keeper_test

import (
	"fmt"

	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *StableKeeperTestSuite) TestGetAmountOutByAmountIn() {
	testCases := []struct {
		name   string
		pair   types.Pair
		amount uint64
		action string
	}{
		{
			"ok-mint",
			s.GetNormalPair(0),
			1000,
			"mint",
		},
		{
			"ok-burn",
			s.GetNormalPair(0),
			1000,
			"burn",
		},
	}
	suite.Setup()
	suite.Commit()

	for _, tc := range testCases {
		suite.app.StableKeeper.SetTestingMode(true)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			ctx := sdk.WrapSDKContext(suite.ctx)
			suite.app.StableKeeper.AppendPair(suite.ctx, suite.GetNormalPair(0))
			suite.MintStable(10000000000, s.GetNormalPair(0))

			req := types.GetAmountOutByAmountIn{
				PairId:   tc.pair.PairId,
				AmountIn: tc.amount,
				Action:   tc.action,
			}

			_, err := suite.app.StableKeeper.GetAmountOutByAmountIn(ctx, &req)
			suite.NoError(err)
		})
	}
}
