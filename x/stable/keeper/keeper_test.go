package keeper_test

import (
	"testing"

	"github.com/QuadrateOrg/core/app"
	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	quadrateapptest "github.com/QuadrateOrg/core/app/helpers"
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
