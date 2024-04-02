package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgPerpetualDeposit{}, "perpetual/PerpetualDeposit", nil)
	cdc.RegisterConcrete(&MsgPerpetualWithdraw{}, "perpetual/PerpetualWithdraw", nil)
	cdc.RegisterConcrete(&MsgCreatePosition{}, "perpetual/CreatePosition", nil)
	cdc.RegisterConcrete(&MsgClosePosition{}, "perpetual/ClosePosition", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPerpetualDeposit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPerpetualWithdraw{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePosition{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClosePosition{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
