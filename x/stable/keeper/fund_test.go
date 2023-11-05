package keeper_test

import (
	"fmt"

	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *StableKeeperTestSuite) TestGetFundsAddress() {
	s.Setup()
	bf := s.app.StableKeeper.GetBurningFundAddress(s.ctx)
	rf := s.app.StableKeeper.GetReserveFundAddress(s.ctx)
	fmt.Printf("%s\n%s\n", rf.String(), bf.String())
}

func (s *StableKeeperTestSuite) TestSetFundsAddress() {
	s.Setup()
	bfold := s.app.StableKeeper.GetBurningFundAddress(s.ctx)
	rfold := s.app.StableKeeper.GetReserveFundAddress(s.ctx)

	bfnew := apptesting.CreateRandomAccounts(1)[0]
	rfnew := apptesting.CreateRandomAccounts(1)[0]

	s.app.StableKeeper.SetBurningFundAddress(s.ctx, bfnew)
	s.app.StableKeeper.SetReserveFundAddress(s.ctx, rfnew)

	bf := s.app.StableKeeper.GetBurningFundAddress(s.ctx)
	rf := s.app.StableKeeper.GetReserveFundAddress(s.ctx)

	s.Require().NotEqual(bfold, bf)
	s.Require().NotEqual(rfold, rf)

	s.Require().Equal(bfnew, bf)
	s.Require().Equal(rfnew, rf)
}

func (s *StableKeeperTestSuite) TestAddressEmptyCheck() {
	s.Setup()
	s.Require().Equal(s.app.StableKeeper.AddressEmptyCheck(s.ctx), false)

	bfnew := sdk.AccAddress("")
	rfnew := sdk.AccAddress("")

	s.app.StableKeeper.SetBurningFundAddress(s.ctx, bfnew)
	s.app.StableKeeper.SetReserveFundAddress(s.ctx, rfnew)

	s.Require().Equal(s.app.StableKeeper.AddressEmptyCheck(s.ctx), true)
}
