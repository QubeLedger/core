package keeper

import (
	"github.com/QuadrateOrg/core/x/dex/types"
)

var _ types.QueryServer = Keeper{}
