package keeper_test

import (
	"testing"

	"github.com/QuadrateOrg/core/app"
	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	quadrateapptest "github.com/QuadrateOrg/core/app/helpers"
	gmd "github.com/QuadrateOrg/core/x/stable/gmb"
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type StableKeeperTestSuite struct {
	suite.Suite
	ctx     sdk.Context
	app     *app.QuadrateApp
	genesis types.GenesisState
	Address sdk.AccAddress
}

var s *StableKeeperTestSuite

func (s *StableKeeperTestSuite) Setup() {
	s.app = quadrateapptest.Setup(s.T(), "qube-1", false, 1)
	s.Address = apptesting.CreateRandomAccounts(1)[0]
}

func TestStableKeeperTestSuite(t *testing.T) {
	s = new(StableKeeperTestSuite)
	suite.Run(t, s)
	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keeper Suite")
}

func (suite *StableKeeperTestSuite) Commit() {
	header := suite.ctx.BlockHeader()
	// update ctx
	suite.ctx = suite.app.BaseApp.NewContext(false, header)
}

func (suite *StableKeeperTestSuite) IncreaseBalance(address sdk.AccAddress, denom string, amount sdk.Int) {
	suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, amount)))
	suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, address, sdk.NewCoins(sdk.NewCoin("uatom", amount)))
}

func (suite *StableKeeperTestSuite) GetBackingRatio() int64 {
	atomPrice := suite.app.StableKeeper.GetAtomPrice(suite.ctx)
	qm := suite.app.StableKeeper.GetStablecoinSupply(suite.ctx)
	ar := suite.app.StableKeeper.GetAtomReserve(suite.ctx)
	backing_ratio := gmd.CalculateBackingRatio(atomPrice, ar, qm)
	return backing_ratio.Int64()
}

func (suite *StableKeeperTestSuite) MintUsq() error {
	suite.app.StableKeeper.SetBaseTokenDenom(suite.ctx, "uatom")
	suite.app.StableKeeper.UpdateAtomPriceTesting(suite.ctx, sdk.NewInt(95000))
	msg := types.NewMsgMintUsq(
		suite.Address.String(),
		sdk.NewInt(10000).String()+"uatom",
	)
	ctx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.app.StableKeeper.MintUsq(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}
