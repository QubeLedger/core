package keeper

import (
	"github.com/QuadrateOrg/core/x/contractmanager/types"
)

var _ types.QueryServer = Keeper{}
