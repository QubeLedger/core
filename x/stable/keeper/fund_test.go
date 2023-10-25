package keeper_test

import "fmt"

func (suite *StableKeeperTestSuite) TestSetFundsAddress() {
	s.Setup()
	bf := suite.app.StableKeeper.GetBurningFundAddress(suite.ctx)
	rf := suite.app.StableKeeper.GetReserveFundAddress(suite.ctx)
	fmt.Printf("%s\n%s\n", rf.String(), bf.String())

}
