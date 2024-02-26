package keeper_test

import (
	"fmt"
	"sort"
)

type Lns struct {
	amount int64
	denom  string
}

type Lnd struct {
	amount int64
	denom  string
}

type Pos struct {
	loans []Lns
	lends []Lnd
}

func (s *GrowKeeperTestSuite) TestArray() {
	s.Setup()

	positions := Pos{
		loans: []Lns{
			{
				amount: 10,
				denom:  "btc",
			},
			{
				amount: 2,
				denom:  "btc",
			},
			{
				amount: 45,
				denom:  "btc",
			},
			{
				amount: 15,
				denom:  "btc",
			},
			{
				amount: 5,
				denom:  "btc",
			},
		},
		lends: []Lnd{
			{
				amount: 10,
				denom:  "btc",
			},
			{
				amount: 2,
				denom:  "btc",
			},
			{
				amount: 45,
				denom:  "btc",
			},
			{
				amount: 15,
				denom:  "btc",
			},
			{
				amount: 5,
				denom:  "btc",
			},
		},
	}

	sort.SliceStable(positions.lends, func(i, j int) bool { return positions.lends[i].amount > positions.lends[j].amount })
	fmt.Printf("%v", positions.lends)
}
