package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	chain "github.com/QuadrateOrg/core/app"
	"github.com/QuadrateOrg/core/app/apptesting"
	testhelpers "github.com/QuadrateOrg/core/app/helpers"
	apptypes "github.com/QuadrateOrg/core/types"
	"github.com/QuadrateOrg/core/x/epochs/types"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	app         *chain.QuadrateApp
	queryClient types.QueryClient
	Ctx         sdk.Context
}

func (s *KeeperTestSuite) Setup() {
	apptypes.SetConfig()
	s.app = testhelpers.Setup(s.T(), "qube-1", false, 1)
	s.Ctx = s.app.BaseApp.NewContext(false, s.Ctx.BlockHeader())
	s.Ctx = s.Ctx.WithBlockGasMeter(sdk.NewGasMeter(uint64(1000000)))
	s.queryClient = types.NewQueryClient(s.QueryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
