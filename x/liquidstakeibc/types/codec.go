package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/liquidstakeibc interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterHostChain{}, "liquidstakeibc/MsgRegisterHostChain", nil)
	cdc.RegisterConcrete(&MsgUpdateHostChain{}, "liquidstakeibc/MsgUpdateHostChain", nil)
	cdc.RegisterConcrete(&MsgLiquidStake{}, "liquidstakeibc/MsgLiquidStake", nil)
	cdc.RegisterConcrete(&MsgLiquidStakeLSM{}, "liquidstakeibc/MsgLiquidStakeLSM", nil)
	cdc.RegisterConcrete(&MsgLiquidUnstake{}, "liquidstakeibc/MsgLiquidUnstake", nil)
	cdc.RegisterConcrete(&MsgRedeem{}, "liquidstakeibc/MsgRedeem", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "liquidstakeibc/MsgUpdateParams", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterHostChain{},
		&MsgUpdateHostChain{},
		&MsgLiquidStake{},
		&MsgLiquidStakeLSM{},
		&MsgLiquidUnstake{},
		&MsgRedeem{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)
}
