package keeper

import (
	"fmt"

	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

var (
	BaseTokenDenom string
	SendTokenDenom string
)

func (k Keeper) ChangeBaseTokenDenom(ctx sdk.Context, coinMetadata banktypes.Metadata) error {
	if err := k.verifyMetadata(ctx, coinMetadata); err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidCoins, "base denomination '%s' cannot have a supply of 0", coinMetadata.Base)
	}
	k.SetBaseTokenDenom(ctx, coinMetadata.Base)
	return nil
}

func (k Keeper) ChangeSendTokenDenom(ctx sdk.Context, coinMetadata banktypes.Metadata) error {
	if err := k.verifyMetadata(ctx, coinMetadata); err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidCoins, "base denomination '%s' cannot have a supply of 0", coinMetadata.Base)
	}
	k.SetSendTokenDenom(ctx, coinMetadata.Base)
	return nil
}

func (k Keeper) SetBaseTokenDenom(ctx sdk.Context, newBaseTokenDenom string) {
	BaseTokenDenom = newBaseTokenDenom
}

func (k Keeper) SetSendTokenDenom(ctx sdk.Context, newSendTokenDenom string) {
	SendTokenDenom = newSendTokenDenom
}

func (k Keeper) GetBaseTokenDenom(ctx sdk.Context) string {
	return BaseTokenDenom
}

func (k Keeper) verifyMetadata(
	ctx sdk.Context,
	coinMetadata banktypes.Metadata,
) error {
	meta, found := k.bankKeeper.GetDenomMetaData(ctx, coinMetadata.Base)
	if !found {
		k.bankKeeper.SetDenomMetaData(ctx, coinMetadata)
		return nil
	}

	// If it already existed, check that is equal to what is stored
	return EqualMetadata(meta, coinMetadata)
}

func EqualMetadata(a, b banktypes.Metadata) error {
	if a.Base == b.Base && a.Description == b.Description && a.Display == b.Display && a.Name == b.Name && a.Symbol == b.Symbol {
		if len(a.DenomUnits) != len(b.DenomUnits) {
			return fmt.Errorf("metadata provided has different denom units from stored, %d ≠ %d", len(a.DenomUnits), len(b.DenomUnits))
		}

		for i, v := range a.DenomUnits {
			if (v.Exponent != b.DenomUnits[i].Exponent) || (v.Denom != b.DenomUnits[i].Denom) || !EqualStringSlice(v.Aliases, b.DenomUnits[i].Aliases) {
				return fmt.Errorf("metadata provided has different denom unit from stored, %s ≠ %s", a.DenomUnits[i], b.DenomUnits[i])
			}
		}

		return nil
	}
	return fmt.Errorf("metadata provided is different from stored")
}

func EqualStringSlice(aliasesA, aliasesB []string) bool {
	if len(aliasesA) != len(aliasesB) {
		return false
	}

	for i := 0; i < len(aliasesA); i++ {
		if aliasesA[i] != aliasesB[i] {
			return false
		}
	}

	return true
}
