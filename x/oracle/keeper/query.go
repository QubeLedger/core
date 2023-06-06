package keeper

import (
	"github.com/QuadrateOrg/core/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
