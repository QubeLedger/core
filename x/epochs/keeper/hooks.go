package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AfterEpochEnd gets called at the end of the epoch, end of epoch is the timestamp of first block produced after epoch duration.
func (k Keeper) AfterEpochEnd(ctx sdk.Context, identifier string, epochNumber int64) error {
	// Error is not handled as AfterEpochEnd Hooks use osmoutils.ApplyFuncIfNoError()
	err := k.hooks.AfterEpochEnd(ctx, identifier, epochNumber)
	return err
}

// BeforeEpochStart new epoch is next block of epoch end block
func (k Keeper) BeforeEpochStart(ctx sdk.Context, identifier string, epochNumber int64) error {
	// Error is not handled as BeforeEpochStart Hooks use osmoutils.ApplyFuncIfNoError()
	err := k.hooks.BeforeEpochStart(ctx, identifier, epochNumber)
	return err
}
