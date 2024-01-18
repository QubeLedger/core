package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgGrowDeposit{}, "grow/GrowDeposit", nil)
	cdc.RegisterConcrete(&MsgGrowWithdrawal{}, "grow/GrowWithdrawal", nil)
	cdc.RegisterConcrete(&MsgCreateLend{}, "grow/CreateLend", nil)
	cdc.RegisterConcrete(&MsgDeleteLend{}, "grow/DeleteLend", nil)
	cdc.RegisterConcrete(&MsgDepositCollateral{}, "grow/DepositCollateral", nil)
	cdc.RegisterConcrete(&MsgWithdrawalCollateral{}, "grow/WithdrawalCollateral", nil)
	cdc.RegisterConcrete(&MsgOpenLiquidationPosition{}, "grow/CreateLiquidationPosition", nil)
	cdc.RegisterConcrete(&MsgCloseLiquidationPosition{}, "grow/CloseLiquidationPosition", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGrowDeposit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGrowWithdrawal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateLend{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteLend{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDepositCollateral{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawalCollateral{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOpenLiquidationPosition{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCloseLiquidationPosition{},
	)
	registry.RegisterImplementations(
		(*gov.Content)(nil),
		&RegisterLendAssetProposal{},
		&RegisterGTokenPairProposal{},
		&RegisterChangeGrowYieldReserveAddressProposal{},
		&RegisterChangeUSQReserveAddressProposal{},
		&RegisterChangeGrowStakingReserveAddressProposal{},
		&RegisterChangeRealRateProposal{},
		&RegisterChangeBorrowRateProposal{},
		&RegisterActivateGrowModuleProposal{},
		&RegisterRemoveLendAssetProposal{},
		&RegisterRemoveGTokenPairProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
