package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	"github.com/QuadrateOrg/core/x/tokenfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type KeeperTestSuite struct {
	suite.Suite
	apptesting.KeeperTestHelper

	queryClient types.QueryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()

	// Fund every TestAcc with 100 denom creation fees.
	fundAccsAmount := sdk.NewCoins(
		sdk.NewCoin(
			//types.DefaultParams().DenomCreationFee[0].Denom,
			//types.DefaultParams().DenomCreationFee[0].Amount.MulRaw(100)
			"uqube",
			sdk.NewInt(20_000_000_000),
		),
	)
	for _, acc := range s.TestAccs {
		s.FundAcc(acc, fundAccsAmount)
	}

	s.SetupTokenFactory()

	s.queryClient = types.NewQueryClient(s.QueryHelper)
}
