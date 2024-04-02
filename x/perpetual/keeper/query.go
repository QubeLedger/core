package keeper

import (
	"github.com/QuadrateOrg/core/x/perpetual/types"
)

var _ types.QueryServer = Keeper{}
