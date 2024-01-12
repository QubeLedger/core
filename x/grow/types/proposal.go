package types

import (
	fmt "fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
)

// constants
const (
	ProposalTypeRegisterLendAssetProposal                       string = "RegisterLendAssetProposal"
	ProposalTypeRegisterGTokenPairProposal                      string = "RegisterGTokenPairProposal"
	ProposalTypeRegisterChangeGrowYieldReserveAddressProposal   string = "RegisterChangeGrowYieldReserveAddressProposal"
	ProposalTypeRegisterChangeUSQReserveAddressProposal         string = "RegisterChangeUSQReserveAddressProposal"
	ProposalTypeRegisterChangeGrowStakingReserveAddressProposal string = "RegisterChangeGrowStakingReserveAddressProposal"
	ProposalTypeRegisterChangeRealRateProposal                  string = "RegisterChangeRealRateProposal"
	ProposalTypeRegisterChangeBorrowRateProposal                string = "RegisterChangeBorrowRateProposal"
	ProposalTypeRegisterChangeLendRateProposal                  string = "RegisterChangeLendRateProposal"
	ProposalTypeRegisterActivateGrowModuleProposal              string = "RegisterActivateGrowModuleProposal"
	ProposalTypeRegisterRemoveLendAssetProposal                 string = "RegisterRemoveLendAssetProposal"
	ProposalTypeRegisterRemoveGTokenPairProposal                string = "RegisterRemoveGTokenPairProposal"
)

// Implements Proposal Interface
var (
	_ govtypes.Content = &RegisterLendAssetProposal{}
	_ govtypes.Content = &RegisterGTokenPairProposal{}
	_ govtypes.Content = &RegisterChangeGrowYieldReserveAddressProposal{}
	_ govtypes.Content = &RegisterChangeUSQReserveAddressProposal{}
	_ govtypes.Content = &RegisterChangeGrowStakingReserveAddressProposal{}
	_ govtypes.Content = &RegisterChangeRealRateProposal{}
	_ govtypes.Content = &RegisterChangeBorrowRateProposal{}
	_ govtypes.Content = &RegisterChangeLendRateProposal{}
	_ govtypes.Content = &RegisterActivateGrowModuleProposal{}
	_ govtypes.Content = &RegisterRemoveLendAssetProposal{}
	_ govtypes.Content = &RegisterRemoveGTokenPairProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeRegisterLendAssetProposal)
	govtypes.RegisterProposalTypeCodec(&RegisterLendAssetProposal{}, "grow/RegisterLendAssetProposal")

	govtypes.RegisterProposalType(ProposalTypeRegisterGTokenPairProposal)
	govtypes.RegisterProposalTypeCodec(&RegisterGTokenPairProposal{}, "grow/RegisterGTokenPairProposal")

	govtypes.RegisterProposalType(ProposalTypeRegisterChangeGrowYieldReserveAddressProposal)
	govtypes.RegisterProposalTypeCodec(&RegisterChangeGrowYieldReserveAddressProposal{}, "grow/RegisterChangeGrowYieldReserveAddressProposal")

	govtypes.RegisterProposalType(ProposalTypeRegisterChangeUSQReserveAddressProposal)
	govtypes.RegisterProposalTypeCodec(&RegisterChangeUSQReserveAddressProposal{}, "grow/RegisterChangeUSQReserveAddressProposal")

	govtypes.RegisterProposalType(ProposalTypeRegisterChangeGrowStakingReserveAddressProposal)
	govtypes.RegisterProposalTypeCodec(&RegisterChangeGrowStakingReserveAddressProposal{}, "grow/RegisterChangeGrowStakingReserveAddressProposal")

	govtypes.RegisterProposalType(ProposalTypeRegisterChangeRealRateProposal)
	govtypes.RegisterProposalTypeCodec(&RegisterChangeRealRateProposal{}, "grow/RegisterChangeRealRateProposal")

	govtypes.RegisterProposalType(ProposalTypeRegisterChangeBorrowRateProposal)
	govtypes.RegisterProposalTypeCodec(&RegisterChangeBorrowRateProposal{}, "grow/RegisterChangeBorrowRateProposal")

	govtypes.RegisterProposalType(ProposalTypeRegisterChangeLendRateProposal)
	govtypes.RegisterProposalTypeCodec(&RegisterChangeLendRateProposal{}, "grow/RegisterChangeLendRateProposal")

	govtypes.RegisterProposalType(ProposalTypeRegisterActivateGrowModuleProposal)
	govtypes.RegisterProposalTypeCodec(&RegisterActivateGrowModuleProposal{}, "grow/RegisterActivateGrowModuleProposal")

	govtypes.RegisterProposalType(ProposalTypeRegisterRemoveLendAssetProposal)
	govtypes.RegisterProposalTypeCodec(&RegisterRemoveLendAssetProposal{}, "grow/RegisterRemoveLendAssetProposal")

	govtypes.RegisterProposalType(ProposalTypeRegisterRemoveGTokenPairProposal)
	govtypes.RegisterProposalTypeCodec(&RegisterRemoveGTokenPairProposal{}, "grow/RegisterRemoveGTokenPairProposal")
}

/*
RegisterLendAssetProposal
*/

func NewRegisterLendAssetProposal(title, description string, assetMetadata banktypes.Metadata, oracleId string) govtypes.Content {
	return &RegisterLendAssetProposal{
		Title:         title,
		Description:   description,
		AssetMetadata: assetMetadata,
		OracleAssetId: oracleId,
	}
}

func (*RegisterLendAssetProposal) ProposalRoute() string { return RouterKey }

func (*RegisterLendAssetProposal) ProposalType() string {
	return ProposalTypeRegisterLendAssetProposal
}

func (rtbp *RegisterLendAssetProposal) ValidateBasic() error {
	{
		if strings.TrimSpace(rtbp.AssetMetadata.Name) == "" {
			return fmt.Errorf("name field cannot be blank")
		}

		if strings.TrimSpace(rtbp.AssetMetadata.Symbol) == "" {
			return fmt.Errorf("symbol field cannot be blank")
		}

		if err := sdk.ValidateDenom(rtbp.AssetMetadata.Base); err != nil {
			return fmt.Errorf("invalid metadata base denom: %w", err)
		}

		if err := sdk.ValidateDenom(rtbp.AssetMetadata.Display); err != nil {
			return fmt.Errorf("invalid metadata display denom: %w", err)
		}
		if err := ibctransfertypes.ValidateIBCDenom(rtbp.AssetMetadata.Base); err != nil {
			return err
		}
	}
	{
		if len(rtbp.OracleAssetId) == 0 {
			return ErrInvalidLength
		}
	}

	return nil
}

/*
RegisterGTokenPairProposal
*/

func NewRegisterGTokenPairProposal(title, description string, metadata banktypes.Metadata, qStableId string, minAmountIn string, minAmountOut string) govtypes.Content {
	return &RegisterGTokenPairProposal{
		Title:          title,
		Description:    description,
		GTokenMetadata: metadata,
		QStablePairId:  qStableId,
		MinAmountIn:    minAmountIn,
		MinAmountOut:   minAmountOut,
	}
}

func (*RegisterGTokenPairProposal) ProposalRoute() string { return RouterKey }

func (*RegisterGTokenPairProposal) ProposalType() string {
	return ProposalTypeRegisterGTokenPairProposal
}

func (rtbp *RegisterGTokenPairProposal) ValidateBasic() error {
	{
		if strings.TrimSpace(rtbp.GTokenMetadata.Name) == "" {
			return fmt.Errorf("name field cannot be blank")
		}

		if strings.TrimSpace(rtbp.GTokenMetadata.Symbol) == "" {
			return fmt.Errorf("symbol field cannot be blank")
		}

		if err := sdk.ValidateDenom(rtbp.GTokenMetadata.Base); err != nil {
			return fmt.Errorf("invalid metadata base denom: %w", err)
		}

		if err := sdk.ValidateDenom(rtbp.GTokenMetadata.Display); err != nil {
			return fmt.Errorf("invalid metadata display denom: %w", err)
		}
		if err := ibctransfertypes.ValidateIBCDenom(rtbp.GTokenMetadata.Base); err != nil {
			return err
		}
	}
	{
		if len(rtbp.QStablePairId) == 0 {
			return ErrInvalidLength
		}
	}
	{
		val, err := sdk.ParseCoinsNormalized(rtbp.MinAmountIn)
		if err != nil || val.String() == "" {
			return sdk.ErrInvalidLengthCoin
		}
		val, err = sdk.ParseCoinsNormalized(rtbp.MinAmountOut)
		if err != nil || val.String() == "" {
			return sdk.ErrInvalidLengthCoin
		}
	}

	return nil
}

/*
RegisterChangeGrowYieldReserveAddressProposal
*/

func NewRegisterChangeGrowYieldReserveAddressProposal(title, description string, address string) govtypes.Content {
	return &RegisterChangeGrowYieldReserveAddressProposal{
		Title:       title,
		Description: description,
		Address:     address,
	}
}

func (*RegisterChangeGrowYieldReserveAddressProposal) ProposalRoute() string { return RouterKey }

func (*RegisterChangeGrowYieldReserveAddressProposal) ProposalType() string {
	return ProposalTypeRegisterChangeGrowYieldReserveAddressProposal
}

func (rtbp *RegisterChangeGrowYieldReserveAddressProposal) ValidateBasic() error {
	{
		if len(rtbp.Address) == 0 {
			return ErrInvalidLength
		}
		_, err := sdk.AccAddressFromBech32(rtbp.Address)
		if err != nil {
			return err
		}
		return nil
	}
}

/*
RegisterChangeUSQReserveAddressProposal
*/

func NewRegisterChangeUSQReserveAddressProposal(title, description string, address string) govtypes.Content {
	return &RegisterChangeUSQReserveAddressProposal{
		Title:       title,
		Description: description,
		Address:     address,
	}
}

func (*RegisterChangeUSQReserveAddressProposal) ProposalRoute() string { return RouterKey }

func (*RegisterChangeUSQReserveAddressProposal) ProposalType() string {
	return ProposalTypeRegisterChangeUSQReserveAddressProposal
}

func (rtbp *RegisterChangeUSQReserveAddressProposal) ValidateBasic() error {
	{
		if len(rtbp.Address) == 0 {
			return ErrInvalidLength
		}
		_, err := sdk.AccAddressFromBech32(rtbp.Address)
		if err != nil {
			return err
		}
		return nil
	}
}

/*
RegisterChangeGrowStakingReserveAddressProposal
*/

func NewRegisterChangeGrowStakingReserveAddressProposal(title, description string, address string) govtypes.Content {
	return &RegisterChangeUSQReserveAddressProposal{
		Title:       title,
		Description: description,
		Address:     address,
	}
}

func (*RegisterChangeGrowStakingReserveAddressProposal) ProposalRoute() string { return RouterKey }

func (*RegisterChangeGrowStakingReserveAddressProposal) ProposalType() string {
	return ProposalTypeRegisterChangeGrowStakingReserveAddressProposal
}

func (rtbp *RegisterChangeGrowStakingReserveAddressProposal) ValidateBasic() error {
	{
		if len(rtbp.Address) == 0 {
			return ErrInvalidLength
		}
		_, err := sdk.AccAddressFromBech32(rtbp.Address)
		if err != nil {
			return err
		}
		return nil
	}
}

/*
RegisterChangeRealRateProposal
*/

func NewRegisterChangeRealRateProposal(title, description string, rate uint64) govtypes.Content {
	return &RegisterChangeRealRateProposal{
		Title:       title,
		Description: description,
		Rate:        rate,
	}
}

func (*RegisterChangeRealRateProposal) ProposalRoute() string { return RouterKey }

func (*RegisterChangeRealRateProposal) ProposalType() string {
	return ProposalTypeRegisterChangeRealRateProposal
}

/* #nosec */
func (rtbp *RegisterChangeRealRateProposal) ValidateBasic() error {
	{
		if rtbp.Rate == uint64(0) {
			return ErrIntNegativeOrZero
		}
		value := sdk.NewInt(int64(rtbp.Rate))
		if value.IsNegative() || value.IsNil() || value.IsZero() {
			return ErrIntNegativeOrZero
		}
	}
	return nil
}

/*
RegisterChangeBorrowRateProposal
*/

func NewRegisterChangeBorrowRateProposal(title, description string, rate uint64) govtypes.Content {
	return &RegisterChangeBorrowRateProposal{
		Title:       title,
		Description: description,
		Rate:        rate,
	}
}

func (*RegisterChangeBorrowRateProposal) ProposalRoute() string { return RouterKey }

func (*RegisterChangeBorrowRateProposal) ProposalType() string {
	return ProposalTypeRegisterChangeBorrowRateProposal
}

/* #nosec */
func (rtbp *RegisterChangeBorrowRateProposal) ValidateBasic() error {
	{
		if rtbp.Rate == uint64(0) {
			return ErrIntNegativeOrZero
		}
		value := sdk.NewInt(int64(rtbp.Rate))
		if value.IsNegative() || value.IsNil() || value.IsZero() {
			return ErrIntNegativeOrZero
		}
	}
	return nil
}

/*
RegisterChangeLendRateProposal
*/

func NewRegisterChangeLendRateProposal(title, description string, rate uint64, id string) govtypes.Content {
	return &RegisterChangeLendRateProposal{
		Title:       title,
		Description: description,
		Rate:        rate,
		Id:          id,
	}
}

func (*RegisterChangeLendRateProposal) ProposalRoute() string { return RouterKey }

func (*RegisterChangeLendRateProposal) ProposalType() string {
	return ProposalTypeRegisterChangeLendRateProposal
}

/* #nosec */
func (rtbp *RegisterChangeLendRateProposal) ValidateBasic() error {
	{
		if rtbp.Rate == uint64(0) {
			return ErrIntNegativeOrZero
		}

		value := sdk.NewInt(int64(rtbp.Rate))
		if value.IsNegative() || value.IsNil() || value.IsZero() {
			return ErrIntNegativeOrZero
		}

		if len(rtbp.Id) == 0 {
			return ErrInvalidLength
		}
	}
	return nil
}

/*
RegisterActivateGrowModule
*/

func NewRegisterActivateGrowModuleProposal(title, description string) govtypes.Content {
	return &RegisterActivateGrowModuleProposal{
		Title:       title,
		Description: description,
	}
}

func (*RegisterActivateGrowModuleProposal) ProposalRoute() string { return RouterKey }

func (*RegisterActivateGrowModuleProposal) ProposalType() string {
	return ProposalTypeRegisterActivateGrowModuleProposal
}

/* #nosec */
func (rtbp *RegisterActivateGrowModuleProposal) ValidateBasic() error {
	return nil
}

/*
RegisterRemoveLendAssetProposal
*/

func NewRegisterRemoveLendAssetProposal(title, description string, id string) govtypes.Content {
	return &RegisterRemoveLendAssetProposal{
		Title:       title,
		Description: description,
		LendAssetId: id,
	}
}

func (*RegisterRemoveLendAssetProposal) ProposalRoute() string { return RouterKey }

func (*RegisterRemoveLendAssetProposal) ProposalType() string {
	return ProposalTypeRegisterRemoveLendAssetProposal
}

/* #nosec */
func (rtbp *RegisterRemoveLendAssetProposal) ValidateBasic() error {
	if len(rtbp.LendAssetId) == 0 {
		return ErrInvalidLength
	}
	return nil
}

/*
RegisterRemoveGTokenPairProposal
*/

func NewRegisterRemoveGTokenPairProposal(title, description string, id string) govtypes.Content {
	return &RegisterRemoveGTokenPairProposal{
		Title:        title,
		Description:  description,
		GTokenPairID: id,
	}
}

func (*RegisterRemoveGTokenPairProposal) ProposalRoute() string { return RouterKey }

func (*RegisterRemoveGTokenPairProposal) ProposalType() string {
	return ProposalTypeRegisterRemoveGTokenPairProposal
}

/* #nosec */
func (rtbp *RegisterRemoveGTokenPairProposal) ValidateBasic() error {
	if len(rtbp.GTokenPairID) == 0 {
		return ErrInvalidLength
	}
	return nil
}
