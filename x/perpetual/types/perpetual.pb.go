// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/perpetual/v1beta1/perpetual.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PerpetualTradeType int32

const (
	PerpetualTradeType_PERPETUAL_LONG_POSITION  PerpetualTradeType = 0
	PerpetualTradeType_PERPETUAL_SHORT_POSITION PerpetualTradeType = 1
)

var PerpetualTradeType_name = map[int32]string{
	0: "PERPETUAL_LONG_POSITION",
	1: "PERPETUAL_SHORT_POSITION",
}

var PerpetualTradeType_value = map[string]int32{
	"PERPETUAL_LONG_POSITION":  0,
	"PERPETUAL_SHORT_POSITION": 1,
}

func (x PerpetualTradeType) String() string {
	return proto.EnumName(PerpetualTradeType_name, int32(x))
}

func (PerpetualTradeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17f41e0ca14bd547, []int{0}
}

type TradePosition struct {
	Id               uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TradePositionId  string                                 `protobuf:"bytes,2,opt,name=trade_position_id,json=tradePositionId,proto3" json:"trade_position_id,omitempty"`
	Creator          string                                 `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	TradeType        PerpetualTradeType                     `protobuf:"varint,4,opt,name=trade_type,json=tradeType,proto3,enum=core.perpetual.v1beta1.PerpetualTradeType" json:"trade_type,omitempty"`
	Leverage         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=leverage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"leverage"`
	TradingAsset     string                                 `protobuf:"bytes,6,opt,name=trading_asset,json=tradingAsset,proto3" json:"trading_asset,omitempty"`
	CollateralAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=collateral_amount,json=collateralAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"collateral_amount"`
	CollateralDenom  string                                 `protobuf:"bytes,8,opt,name=collateral_denom,json=collateralDenom,proto3" json:"collateral_denom,omitempty"`
	ReturnAmount     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=return_amount,json=returnAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"return_amount"`
	ReturnDenom      string                                 `protobuf:"bytes,10,opt,name=return_denom,json=returnDenom,proto3" json:"return_denom,omitempty"`
	ProfitAmount     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=profit_amount,json=profitAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"profit_amount"`
}

func (m *TradePosition) Reset()         { *m = TradePosition{} }
func (m *TradePosition) String() string { return proto.CompactTextString(m) }
func (*TradePosition) ProtoMessage()    {}
func (*TradePosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_17f41e0ca14bd547, []int{0}
}
func (m *TradePosition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TradePosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TradePosition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TradePosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradePosition.Merge(m, src)
}
func (m *TradePosition) XXX_Size() int {
	return m.Size()
}
func (m *TradePosition) XXX_DiscardUnknown() {
	xxx_messageInfo_TradePosition.DiscardUnknown(m)
}

var xxx_messageInfo_TradePosition proto.InternalMessageInfo

func (m *TradePosition) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TradePosition) GetTradePositionId() string {
	if m != nil {
		return m.TradePositionId
	}
	return ""
}

func (m *TradePosition) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *TradePosition) GetTradeType() PerpetualTradeType {
	if m != nil {
		return m.TradeType
	}
	return PerpetualTradeType_PERPETUAL_LONG_POSITION
}

func (m *TradePosition) GetTradingAsset() string {
	if m != nil {
		return m.TradingAsset
	}
	return ""
}

func (m *TradePosition) GetCollateralDenom() string {
	if m != nil {
		return m.CollateralDenom
	}
	return ""
}

func (m *TradePosition) GetReturnDenom() string {
	if m != nil {
		return m.ReturnDenom
	}
	return ""
}

type Vault struct {
	Id              uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	VaultId         string                                 `protobuf:"bytes,2,opt,name=vault_id,json=vaultId,proto3" json:"vault_id,omitempty"`
	AmountXMetadata types.Metadata                         `protobuf:"bytes,3,opt,name=amountXMetadata,proto3" json:"amountXMetadata" yaml:"amountXMetadata"`
	AmountYMetadata types.Metadata                         `protobuf:"bytes,4,opt,name=amountYMetadata,proto3" json:"amountYMetadata" yaml:"amountYMetadata"`
	X               github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=x,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"x"`
	Y               github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=y,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"y"`
	K               github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=k,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"k"`
	OracleAssetId   string                                 `protobuf:"bytes,8,opt,name=OracleAssetId,proto3" json:"OracleAssetId,omitempty"`
	LongPosition    []TradePosition                        `protobuf:"bytes,9,rep,name=long_position,json=longPosition,proto3" json:"long_position"`
	ShortPosition   []TradePosition                        `protobuf:"bytes,10,rep,name=short_position,json=shortPosition,proto3" json:"short_position"`
}

func (m *Vault) Reset()         { *m = Vault{} }
func (m *Vault) String() string { return proto.CompactTextString(m) }
func (*Vault) ProtoMessage()    {}
func (*Vault) Descriptor() ([]byte, []int) {
	return fileDescriptor_17f41e0ca14bd547, []int{1}
}
func (m *Vault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vault.Merge(m, src)
}
func (m *Vault) XXX_Size() int {
	return m.Size()
}
func (m *Vault) XXX_DiscardUnknown() {
	xxx_messageInfo_Vault.DiscardUnknown(m)
}

var xxx_messageInfo_Vault proto.InternalMessageInfo

func (m *Vault) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Vault) GetVaultId() string {
	if m != nil {
		return m.VaultId
	}
	return ""
}

func (m *Vault) GetAmountXMetadata() types.Metadata {
	if m != nil {
		return m.AmountXMetadata
	}
	return types.Metadata{}
}

func (m *Vault) GetAmountYMetadata() types.Metadata {
	if m != nil {
		return m.AmountYMetadata
	}
	return types.Metadata{}
}

func (m *Vault) GetOracleAssetId() string {
	if m != nil {
		return m.OracleAssetId
	}
	return ""
}

func (m *Vault) GetLongPosition() []TradePosition {
	if m != nil {
		return m.LongPosition
	}
	return nil
}

func (m *Vault) GetShortPosition() []TradePosition {
	if m != nil {
		return m.ShortPosition
	}
	return nil
}

func init() {
	proto.RegisterEnum("core.perpetual.v1beta1.PerpetualTradeType", PerpetualTradeType_name, PerpetualTradeType_value)
	proto.RegisterType((*TradePosition)(nil), "core.perpetual.v1beta1.TradePosition")
	proto.RegisterType((*Vault)(nil), "core.perpetual.v1beta1.Vault")
}

func init() {
	proto.RegisterFile("core/perpetual/v1beta1/perpetual.proto", fileDescriptor_17f41e0ca14bd547)
}

var fileDescriptor_17f41e0ca14bd547 = []byte{
	// 658 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x4f, 0xd4, 0x40,
	0x18, 0xde, 0xc2, 0xf2, 0xb1, 0xb3, 0x2c, 0x1f, 0x13, 0x83, 0x15, 0xb5, 0xac, 0xab, 0x92, 0x95,
	0x84, 0x36, 0xe0, 0xcd, 0x78, 0x61, 0x03, 0xc1, 0x1a, 0xa4, 0x6b, 0x59, 0x8d, 0xe8, 0xa1, 0x99,
	0x6d, 0xc7, 0xd2, 0x6c, 0xb7, 0xd3, 0x4c, 0x67, 0x09, 0xfb, 0x17, 0x3c, 0xf9, 0xb3, 0x38, 0x72,
	0x34, 0x1e, 0x88, 0x81, 0x7f, 0xe0, 0xc1, 0xb3, 0x99, 0x99, 0x7e, 0x2c, 0xa0, 0x07, 0xf6, 0xd4,
	0xce, 0x33, 0xcf, 0xfb, 0x3c, 0x6f, 0xdf, 0x3c, 0x7d, 0xc1, 0x9a, 0x4b, 0x28, 0x36, 0x62, 0x4c,
	0x63, 0xcc, 0x06, 0x28, 0x34, 0x4e, 0x36, 0xbb, 0x98, 0xa1, 0xcd, 0x02, 0xd1, 0x63, 0x4a, 0x18,
	0x81, 0xcb, 0x9c, 0xa7, 0x17, 0x68, 0xca, 0x5b, 0xb9, 0xe7, 0x13, 0x9f, 0x08, 0x8a, 0xc1, 0xdf,
	0x24, 0x7b, 0x45, 0x73, 0x49, 0xd2, 0x27, 0x89, 0xd1, 0x45, 0x51, 0x2f, 0x97, 0xe4, 0x07, 0x79,
	0xdf, 0xf8, 0x53, 0x06, 0xb5, 0x0e, 0x45, 0x1e, 0x6e, 0x93, 0x24, 0x60, 0x01, 0x89, 0xe0, 0x3c,
	0x98, 0x08, 0x3c, 0x55, 0xa9, 0x2b, 0xcd, 0xb2, 0x3d, 0x11, 0x78, 0x70, 0x1d, 0x2c, 0x31, 0x4e,
	0x70, 0xe2, 0x94, 0xe1, 0x04, 0x9e, 0x3a, 0x51, 0x57, 0x9a, 0x15, 0x7b, 0x81, 0x8d, 0x56, 0x9a,
	0x1e, 0x54, 0xc1, 0x8c, 0x4b, 0x31, 0x62, 0x84, 0xaa, 0x93, 0x82, 0x91, 0x1d, 0xa1, 0x09, 0x80,
	0x54, 0x61, 0xc3, 0x18, 0xab, 0xe5, 0xba, 0xd2, 0x9c, 0xdf, 0x5a, 0xd7, 0xff, 0xfd, 0x29, 0x7a,
	0x3b, 0x43, 0x44, 0x67, 0x9d, 0x61, 0x8c, 0xed, 0x0a, 0xcb, 0x5e, 0xe1, 0x5b, 0x30, 0x1b, 0xe2,
	0x13, 0x4c, 0x91, 0x8f, 0xd5, 0x29, 0xee, 0xd2, 0xd2, 0xcf, 0x2e, 0x56, 0x4b, 0x3f, 0x2f, 0x56,
	0xd7, 0xfc, 0x80, 0x1d, 0x0f, 0xba, 0xba, 0x4b, 0xfa, 0x46, 0xfa, 0xdd, 0xf2, 0xb1, 0x91, 0x78,
	0x3d, 0x83, 0x3b, 0x27, 0xfa, 0x0e, 0x76, 0xed, 0xbc, 0x1e, 0x3e, 0x05, 0x35, 0x2e, 0x1c, 0x44,
	0xbe, 0x83, 0x92, 0x04, 0x33, 0x75, 0x5a, 0xb4, 0x3d, 0x97, 0x82, 0xdb, 0x1c, 0x83, 0x5f, 0xc0,
	0x92, 0x4b, 0xc2, 0x10, 0x31, 0x4c, 0x51, 0xe8, 0xa0, 0x3e, 0x19, 0x44, 0x4c, 0x9d, 0xb9, 0xb3,
	0xb3, 0x19, 0x31, 0x7b, 0xb1, 0x10, 0xda, 0x16, 0x3a, 0xf0, 0x05, 0x18, 0xc1, 0x1c, 0x0f, 0x47,
	0xa4, 0xaf, 0xce, 0xca, 0xe9, 0x16, 0xf8, 0x0e, 0x87, 0xe1, 0x21, 0xa8, 0x51, 0xcc, 0x06, 0x34,
	0xca, 0x7a, 0xa8, 0x8c, 0xd5, 0xc3, 0x9c, 0x14, 0x49, 0xfd, 0x9f, 0x80, 0xf4, 0x9c, 0x7a, 0x03,
	0xe1, 0x5d, 0x95, 0x58, 0xee, 0x1b, 0x53, 0xf2, 0x35, 0x60, 0x99, 0x6f, 0x75, 0x3c, 0x5f, 0x29,
	0x22, 0x7d, 0x1b, 0xdf, 0xa6, 0xc0, 0xd4, 0x47, 0x34, 0x08, 0xd9, 0xad, 0xc0, 0x3d, 0x00, 0xb3,
	0x27, 0xfc, 0xa2, 0xc8, 0xd9, 0x8c, 0x38, 0x9b, 0x1e, 0xf4, 0xc1, 0x82, 0x6c, 0xe1, 0xd3, 0x3b,
	0xcc, 0x90, 0x87, 0x18, 0x12, 0x39, 0xab, 0x6e, 0x3d, 0xd6, 0xa5, 0xa5, 0x2e, 0xa2, 0x9d, 0xe5,
	0x28, 0x23, 0xb5, 0x34, 0xde, 0xea, 0xef, 0x8b, 0xd5, 0xe5, 0x21, 0xea, 0x87, 0xaf, 0x1a, 0x37,
	0x34, 0x1a, 0xf6, 0x4d, 0xd5, 0xc2, 0xe8, 0x28, 0x37, 0x2a, 0x8f, 0x6d, 0x74, 0x74, 0xcb, 0x28,
	0x47, 0xe0, 0x6b, 0xa0, 0x9c, 0x8e, 0x91, 0x62, 0x3e, 0x4f, 0xe5, 0x94, 0x57, 0x0f, 0x65, 0x64,
	0xef, 0x5e, 0x3d, 0xe4, 0xd5, 0xbd, 0x31, 0x73, 0xac, 0xf4, 0xe0, 0x33, 0x50, 0xb3, 0x28, 0x72,
	0x43, 0x2c, 0x7e, 0x12, 0xd3, 0x4b, 0x53, 0x7b, 0x1d, 0x84, 0x6d, 0x50, 0x0b, 0x49, 0xe4, 0xe7,
	0xcb, 0x43, 0xad, 0xd4, 0x27, 0x9b, 0xd5, 0xad, 0xe7, 0xff, 0xfb, 0xf5, 0xaf, 0xed, 0xa2, 0x56,
	0x99, 0xb7, 0x65, 0xcf, 0x71, 0x85, 0x7c, 0x3f, 0xd9, 0x60, 0x3e, 0x39, 0x26, 0x94, 0x15, 0x92,
	0xe0, 0xee, 0x92, 0x35, 0x21, 0x91, 0x81, 0xeb, 0x16, 0x80, 0xb7, 0x77, 0x0e, 0x7c, 0x08, 0xee,
	0xb7, 0x77, 0xed, 0xf6, 0x6e, 0xe7, 0xc3, 0xf6, 0xbe, 0xb3, 0x6f, 0x1d, 0xec, 0x39, 0x6d, 0xeb,
	0xd0, 0xec, 0x98, 0xd6, 0xc1, 0x62, 0x09, 0x3e, 0x02, 0x6a, 0x71, 0x79, 0xf8, 0xc6, 0xb2, 0x3b,
	0xc5, 0xad, 0xd2, 0xda, 0x3b, 0xbb, 0xd4, 0x94, 0xf3, 0x4b, 0x4d, 0xf9, 0x75, 0xa9, 0x29, 0xdf,
	0xaf, 0xb4, 0xd2, 0xf9, 0x95, 0x56, 0xfa, 0x71, 0xa5, 0x95, 0x3e, 0x6f, 0x8c, 0x4c, 0xf8, 0xfd,
	0x00, 0x79, 0x14, 0x31, 0x6c, 0x51, 0xdf, 0x10, 0xdb, 0xff, 0x74, 0x64, 0xff, 0x8b, 0x61, 0x77,
	0xa7, 0xc5, 0x9a, 0x7e, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x95, 0xf0, 0xdf, 0x1e, 0x06,
	0x00, 0x00,
}

func (m *TradePosition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TradePosition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TradePosition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l

	{
		size := m.ProfitAmount.Size()
		i -= size
		if _, err := m.ProfitAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetual(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a

	if len(m.ReturnDenom) > 0 {
		i -= len(m.ReturnDenom)
		copy(dAtA[i:], m.ReturnDenom)
		i = encodeVarintPerpetual(dAtA, i, uint64(len(m.ReturnDenom)))
		i--
		dAtA[i] = 0x52
	}

	{
		size := m.ReturnAmount.Size()
		i -= size
		if _, err := m.ReturnAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetual(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a

	if len(m.CollateralDenom) > 0 {
		i -= len(m.CollateralDenom)
		copy(dAtA[i:], m.CollateralDenom)
		i = encodeVarintPerpetual(dAtA, i, uint64(len(m.CollateralDenom)))
		i--
		dAtA[i] = 0x42
	}
	{
		size := m.CollateralAmount.Size()
		i -= size
		if _, err := m.CollateralAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetual(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.TradingAsset) > 0 {
		i -= len(m.TradingAsset)
		copy(dAtA[i:], m.TradingAsset)
		i = encodeVarintPerpetual(dAtA, i, uint64(len(m.TradingAsset)))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.Leverage.Size()
		i -= size
		if _, err := m.Leverage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetual(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.TradeType != 0 {
		i = encodeVarintPerpetual(dAtA, i, uint64(m.TradeType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPerpetual(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TradePositionId) > 0 {
		i -= len(m.TradePositionId)
		copy(dAtA[i:], m.TradePositionId)
		i = encodeVarintPerpetual(dAtA, i, uint64(len(m.TradePositionId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPerpetual(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Vault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ShortPosition) > 0 {
		for iNdEx := len(m.ShortPosition) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ShortPosition[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPerpetual(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.LongPosition) > 0 {
		for iNdEx := len(m.LongPosition) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LongPosition[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPerpetual(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.OracleAssetId) > 0 {
		i -= len(m.OracleAssetId)
		copy(dAtA[i:], m.OracleAssetId)
		i = encodeVarintPerpetual(dAtA, i, uint64(len(m.OracleAssetId)))
		i--
		dAtA[i] = 0x42
	}
	{
		size := m.K.Size()
		i -= size
		if _, err := m.K.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetual(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.Y.Size()
		i -= size
		if _, err := m.Y.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetual(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.X.Size()
		i -= size
		if _, err := m.X.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetual(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.AmountYMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPerpetual(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.AmountXMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPerpetual(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.VaultId) > 0 {
		i -= len(m.VaultId)
		copy(dAtA[i:], m.VaultId)
		i = encodeVarintPerpetual(dAtA, i, uint64(len(m.VaultId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPerpetual(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPerpetual(dAtA []byte, offset int, v uint64) int {
	offset -= sovPerpetual(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TradePosition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPerpetual(uint64(m.Id))
	}
	l = len(m.TradePositionId)
	if l > 0 {
		n += 1 + l + sovPerpetual(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPerpetual(uint64(l))
	}
	if m.TradeType != 0 {
		n += 1 + sovPerpetual(uint64(m.TradeType))
	}
	l = m.Leverage.Size()
	n += 1 + l + sovPerpetual(uint64(l))
	l = len(m.TradingAsset)
	if l > 0 {
		n += 1 + l + sovPerpetual(uint64(l))
	}
	l = m.CollateralAmount.Size()
	n += 1 + l + sovPerpetual(uint64(l))
	l = len(m.CollateralDenom)
	if l > 0 {
		n += 1 + l + sovPerpetual(uint64(l))
	}
	l = m.ReturnAmount.Size()
	n += 1 + l + sovPerpetual(uint64(l))
	l = len(m.ReturnDenom)
	if l > 0 {
		n += 1 + l + sovPerpetual(uint64(l))
	}
	l = m.ProfitAmount.Size()
	n += 1 + l + sovPerpetual(uint64(l))
	return n
}

func (m *Vault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPerpetual(uint64(m.Id))
	}
	l = len(m.VaultId)
	if l > 0 {
		n += 1 + l + sovPerpetual(uint64(l))
	}
	l = m.AmountXMetadata.Size()
	n += 1 + l + sovPerpetual(uint64(l))
	l = m.AmountYMetadata.Size()
	n += 1 + l + sovPerpetual(uint64(l))
	l = m.X.Size()
	n += 1 + l + sovPerpetual(uint64(l))
	l = m.Y.Size()
	n += 1 + l + sovPerpetual(uint64(l))
	l = m.K.Size()
	n += 1 + l + sovPerpetual(uint64(l))
	l = len(m.OracleAssetId)
	if l > 0 {
		n += 1 + l + sovPerpetual(uint64(l))
	}
	if len(m.LongPosition) > 0 {
		for _, e := range m.LongPosition {
			l = e.Size()
			n += 1 + l + sovPerpetual(uint64(l))
		}
	}
	if len(m.ShortPosition) > 0 {
		for _, e := range m.ShortPosition {
			l = e.Size()
			n += 1 + l + sovPerpetual(uint64(l))
		}
	}
	return n
}

func sovPerpetual(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPerpetual(x uint64) (n int) {
	return sovPerpetual(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TradePosition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPerpetual
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TradePosition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TradePosition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradePositionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradePositionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeType", wireType)
			}
			m.TradeType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TradeType |= PerpetualTradeType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leverage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Leverage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradingAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradingAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReturnAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReturnAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReturnDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReturnDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfitAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ProfitAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPerpetual(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPerpetual
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Vault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPerpetual
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Vault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VaultId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VaultId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountXMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountXMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountYMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountYMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.X.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Y.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field K", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.K.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleAssetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleAssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LongPosition", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LongPosition = append(m.LongPosition, TradePosition{})
			if err := m.LongPosition[len(m.LongPosition)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShortPosition", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPerpetual
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetual
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShortPosition = append(m.ShortPosition, TradePosition{})
			if err := m.ShortPosition[len(m.ShortPosition)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPerpetual(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPerpetual
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPerpetual(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPerpetual
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPerpetual
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPerpetual
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPerpetual
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPerpetual
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPerpetual        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPerpetual          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPerpetual = fmt.Errorf("proto: unexpected end of group")
)
