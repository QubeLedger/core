package keeper

import (
	"example/x/example/types"
)

var _ types.QueryServer = Keeper{}
