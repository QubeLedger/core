package upgrades

import (
	tokenfactorytypes "github.com/QubeLedger/core/x/tokenfactory/types"
)

type TfUpgradeParams struct {
	Tf tokenfactorytypes.Params `json:"tokenfactory,omitempty"`
}
