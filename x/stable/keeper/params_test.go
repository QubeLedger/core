package keeper_test

import (
	"github.com/QuadrateOrg/core/x/stable/types"
)

func (s *StableKeeperTestSuite) TestGetParams() {
	s.Setup()
	s.Commit()
	params := types.DefaultParams()

	s.app.StableKeeper.SetParams(s.ctx, params)

	s.Require().EqualValues(params, s.app.StableKeeper.GetParams(s.ctx))
}
