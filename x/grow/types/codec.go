package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDeposit{}, "grow/Deposit", nil)
	cdc.RegisterConcrete(&MsgWithdrawal{}, "grow/Withdrawal", nil)
	cdc.RegisterConcrete(&MsgCreateLend{}, "grow/CreateLend", nil)
	cdc.RegisterConcrete(&MsgDeleteLend{}, "grow/DeleteLend", nil)
	cdc.RegisterConcrete(&MsgDepositCollateral{}, "grow/DepositCollateral", nil)
	cdc.RegisterConcrete(&MsgWithdrawalCollateral{}, "grow/WithdrawalCollateral", nil)
	cdc.RegisterConcrete(&MsgOpenLiquidationPosition{}, "grow/CreateLiquidationPosition", nil)
	cdc.RegisterConcrete(&MsgCloseLiquidationPosition{}, "grow/CloseLiquidationPosition", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeposit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawal{},
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

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
