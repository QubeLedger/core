package keeper_test

import (
	"fmt"

	"github.com/QuadrateOrg/core/x/grow/types"
)

func (s *GrowKeeperTestSuite) TestRemoveLoanInPosition() {
	s.Setup()

	testCases := []struct {
		name    string
		pos     types.Position
		loanLen int
	}{
		{
			"1 loan",
			types.Position{
				Id:           0,
				Creator:      "test",
				DepositId:    "test",
				Collateral:   "test",
				OracleTicker: "BTC",
				LoanIds: []string{
					"loanid1",
				},
			},
			0,
		},
		{
			"2 loan",
			types.Position{
				Id:           0,
				Creator:      "test",
				DepositId:    "test",
				Collateral:   "test",
				OracleTicker: "BTC",
				LoanIds: []string{
					"loanid1",
					"loanid2",
				},
			},
			1,
		},
	}
	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			new_pos := s.app.GrowKeeper.RemoveLoanInPosition(s.ctx, tc.pos.LoanIds[0], tc.pos)
			s.Require().Equal(len(new_pos.LoanIds), tc.loanLen)
		})
	}
}
