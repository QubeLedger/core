package keeper

import (
	"github.com/QubeLedger/core/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
