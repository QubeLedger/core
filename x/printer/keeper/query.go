package keeper

import (
	"github.com/QuadrateOrg/core/x/printer/types"
)

var _ types.QueryServer = Keeper{}
