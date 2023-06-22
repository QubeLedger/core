package evm

import (
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
)

const (
	UpgradeName = "TF"
)

var (
	AddModules = []string{evmtypes.ModuleName, feemarkettypes.ModuleName}
)
