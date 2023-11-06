package TF

import (
	tokenfactorytypes "github.com/QuadrateOrg/core/x/tokenfactory/types"
)

type TfUpgradeParams struct {
	Tf tokenfactorytypes.Params `json:"tokenfactory,omitempty"`
}
