package evm

import (
	tokenfactorytypes "github.com/QuadrateOrg/core/x/tokenfactory/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
)

type EvmUpgradeParams struct {
	Evm       evmtypes.Params       `json:"evm,omitempty"`
	FeeMarket feemarkettypes.Params `json:"fee_market,omitempty"`
}

type TfUpgradeParams struct {
	Tf tokenfactorytypes.Params `json:"tokenfactory,omitempty"`
}
