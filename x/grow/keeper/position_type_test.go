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
		lendLen int
	}{
		{
			"1 loan & 1 lend",
			types.Position{
				Id:        0,
				Creator:   "test",
				DepositId: "test",
				LendId: []string{
					"lendid1",
				},
				LoanId: []string{
					"loanid1",
				},
			},
			0,
			0,
		},
	}
	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			new_pos := s.app.GrowKeeper.RemoveLoanInPosition(s.ctx, tc.pos.LoanId[0], tc.pos)
			s.Require().Equal(len(new_pos.LoanId), tc.loanLen)

			new_pos_1 := s.app.GrowKeeper.RemoveLendInPosition(s.ctx, tc.pos.LendId[0], tc.pos)
			s.Require().Equal(len(new_pos_1.LendId), tc.lendLen)
		})
	}
}
