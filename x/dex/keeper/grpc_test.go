package keeper_test

import (
	"testing"
	"time"

	"github.com/QuadrateOrg/core/app"
	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	quadrateapptest "github.com/QuadrateOrg/core/app/helpers"
	"github.com/QuadrateOrg/core/x/dex/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type GrpcDexTestSuite struct {
	suite.Suite
	ctx         sdk.Context
	app         *app.QuadrateApp
	genesis     types.GenesisState
	Address     sdk.AccAddress
	ValPubKeys  []cryptotypes.PubKey
	PoolAddress sdk.AccAddress
}

var s *GrpcDexTestSuite

func (suite *GrpcDexTestSuite) Commit() {
	header := suite.ctx.BlockHeader()
	suite.ctx = suite.app.BaseApp.NewContext(false, header)
}

func (s *GrpcDexTestSuite) Setup() {
	s.app = quadrateapptest.Setup(s.T(), "qube-1", false, 1)
	s.Address = apptesting.CreateRandomAccounts(1)[0]
	s.PoolAddress = apptesting.CreateRandomAccounts(1)[0]
	s.ValPubKeys = simapp.CreateTestPubKeys(1)
	s.ctx = s.ctx.WithBlockTime(time.Now())
}

func TestGrpcDexTestSuite(t *testing.T) {
	s = new(GrpcDexTestSuite)
	suite.Run(t, s)
	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keeper Suite")
}
