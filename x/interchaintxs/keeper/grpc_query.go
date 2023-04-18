package keeper

import (
	"github.com/QuadrateOrg/core/x/interchaintxs/types"
)

var _ types.QueryServer = Keeper{}
