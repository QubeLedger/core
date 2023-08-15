package keeper_test

import (
	"github.com/QubeLedger/core/x/stable/types"
)

func (s *StableKeeperTestSuite) TestGetParams() {
	s.Setup()
	params := types.DefaultParams()

	s.app.StableKeeper.SetParams(s.ctx, params)

	s.Require().EqualValues(params, s.app.StableKeeper.GetParams(s.ctx))
}
