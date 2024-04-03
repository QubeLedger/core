// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/perpetual/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	grpc1 "github.com/gogo/protobuf/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgOpen struct {
	Creator         string                                 `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	TradeType       PerpetualTradeType                     `protobuf:"varint,2,opt,name=trade_type,json=tradeType,proto3,enum=core.perpetual.v1beta1.PerpetualTradeType" json:"trade_type,omitempty"`
	Leverage        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=leverage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"leverage"`
	TradingAsset    string                                 `protobuf:"bytes,4,opt,name=trading_asset,json=tradingAsset,proto3" json:"trading_asset,omitempty"`
	Collateral      string                                 `protobuf:"bytes,5,opt,name=collateral,proto3" json:"collateral,omitempty"`
	TakeProfitPrice github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=take_profit_price,json=takeProfitPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"take_profit_price"`
}

func (m *MsgOpen) Reset()         { *m = MsgOpen{} }
func (m *MsgOpen) String() string { return proto.CompactTextString(m) }
func (*MsgOpen) ProtoMessage()    {}
func (*MsgOpen) Descriptor() ([]byte, []int) {
	return fileDescriptor_815e94f45385e249, []int{0}
}
func (m *MsgOpen) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgOpen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgOpen.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgOpen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgOpen.Merge(m, src)
}
func (m *MsgOpen) XXX_Size() int {
	return m.Size()
}
func (m *MsgOpen) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgOpen.DiscardUnknown(m)
}

var xxx_messageInfo_MsgOpen proto.InternalMessageInfo

func (m *MsgOpen) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgOpen) GetTradeType() PerpetualTradeType {
	if m != nil {
		return m.TradeType
	}
	return PerpetualTradeType_PERPETUAL_LONG_POSITION
}

func (m *MsgOpen) GetTradingAsset() string {
	if m != nil {
		return m.TradingAsset
	}
	return ""
}

func (m *MsgOpen) GetCollateral() string {
	if m != nil {
		return m.Collateral
	}
	return ""
}

type MsgOpenResponse struct {
}

func (m *MsgOpenResponse) Reset()         { *m = MsgOpenResponse{} }
func (m *MsgOpenResponse) String() string { return proto.CompactTextString(m) }
func (*MsgOpenResponse) ProtoMessage()    {}
func (*MsgOpenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_815e94f45385e249, []int{1}
}
func (m *MsgOpenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgOpenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgOpenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgOpenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgOpenResponse.Merge(m, src)
}
func (m *MsgOpenResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgOpenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgOpenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgOpenResponse proto.InternalMessageInfo

type MsgClose struct {
	Creator string                                 `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64                                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Amount  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *MsgClose) Reset()         { *m = MsgClose{} }
func (m *MsgClose) String() string { return proto.CompactTextString(m) }
func (*MsgClose) ProtoMessage()    {}
func (*MsgClose) Descriptor() ([]byte, []int) {
	return fileDescriptor_815e94f45385e249, []int{2}
}
func (m *MsgClose) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClose) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClose.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClose) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClose.Merge(m, src)
}
func (m *MsgClose) XXX_Size() int {
	return m.Size()
}
func (m *MsgClose) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClose.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClose proto.InternalMessageInfo

func (m *MsgClose) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgClose) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgCloseResponse struct {
}

func (m *MsgCloseResponse) Reset()         { *m = MsgCloseResponse{} }
func (m *MsgCloseResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCloseResponse) ProtoMessage()    {}
func (*MsgCloseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_815e94f45385e249, []int{3}
}
func (m *MsgCloseResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCloseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCloseResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCloseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCloseResponse.Merge(m, src)
}
func (m *MsgCloseResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCloseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCloseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCloseResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgOpen)(nil), "core.perpetual.v1beta1.MsgOpen")
	proto.RegisterType((*MsgOpenResponse)(nil), "core.perpetual.v1beta1.MsgOpenResponse")
	proto.RegisterType((*MsgClose)(nil), "core.perpetual.v1beta1.MsgClose")
	proto.RegisterType((*MsgCloseResponse)(nil), "core.perpetual.v1beta1.MsgCloseResponse")
}

func init() { proto.RegisterFile("core/perpetual/v1beta1/tx.proto", fileDescriptor_815e94f45385e249) }

var fileDescriptor_815e94f45385e249 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0x2d, 0xc7, 0x71, 0x92, 0xa1, 0x4d, 0x9a, 0xa5, 0x14, 0xe1, 0x83, 0x6c, 0x5c, 0x48,
	0x4d, 0x21, 0x12, 0x49, 0x9f, 0xa0, 0x69, 0x69, 0x49, 0xc1, 0xc4, 0x55, 0x73, 0xca, 0xc5, 0xac,
	0xa5, 0xe9, 0x56, 0x44, 0xd6, 0x2e, 0xbb, 0xa3, 0x90, 0x40, 0x1f, 0xa2, 0x0f, 0xd2, 0x07, 0xc9,
	0x31, 0xc7, 0xd2, 0x42, 0x28, 0xf6, 0x8b, 0x94, 0xdd, 0x48, 0xa9, 0x0f, 0xf9, 0x03, 0x3e, 0x69,
	0x35, 0xfc, 0xbe, 0x6f, 0x66, 0x3f, 0x69, 0xa0, 0x9b, 0x48, 0x8d, 0x91, 0x42, 0xad, 0x90, 0x4a,
	0x9e, 0x47, 0x67, 0x7b, 0x13, 0x24, 0xbe, 0x17, 0xd1, 0x79, 0xa8, 0xb4, 0x24, 0xc9, 0x5e, 0x58,
	0x20, 0xbc, 0x05, 0xc2, 0x0a, 0xe8, 0x3c, 0x17, 0x52, 0x48, 0x87, 0x44, 0xf6, 0x74, 0x43, 0x77,
	0x76, 0xee, 0xb1, 0xfb, 0xaf, 0x77, 0x5c, 0xff, 0x4f, 0x13, 0xd6, 0x86, 0x46, 0x1c, 0x29, 0x2c,
	0x98, 0x0f, 0x6b, 0x89, 0x46, 0x4e, 0x52, 0xfb, 0x5e, 0xcf, 0x1b, 0x6c, 0xc4, 0xf5, 0x2b, 0x3b,
	0x04, 0x20, 0xcd, 0x53, 0x1c, 0xd3, 0x85, 0x42, 0xbf, 0xd9, 0xf3, 0x06, 0x9b, 0xfb, 0xaf, 0xc3,
	0xbb, 0x07, 0x0a, 0x47, 0x75, 0xe5, 0xd8, 0x4a, 0x8e, 0x2f, 0x14, 0xc6, 0x1b, 0x54, 0x1f, 0xd9,
	0x27, 0x58, 0xcf, 0xf1, 0x0c, 0x35, 0x17, 0xe8, 0xaf, 0xd8, 0x2e, 0x07, 0xe1, 0xe5, 0x75, 0xb7,
	0xf1, 0xfb, 0xba, 0xbb, 0x23, 0x32, 0xfa, 0x56, 0x4e, 0xc2, 0x44, 0x4e, 0xa3, 0x44, 0x9a, 0xa9,
	0x34, 0xd5, 0x63, 0xd7, 0xa4, 0xa7, 0x91, 0xed, 0x6c, 0xc2, 0xf7, 0x98, 0xc4, 0xb7, 0x7a, 0xf6,
	0x12, 0x9e, 0x5a, 0xe3, 0xac, 0x10, 0x63, 0x6e, 0x0c, 0x92, 0xdf, 0x72, 0x63, 0x3f, 0xa9, 0x8a,
	0x6f, 0x6d, 0x8d, 0x05, 0x00, 0x89, 0xcc, 0x73, 0x4e, 0xa8, 0x79, 0xee, 0xaf, 0x3a, 0x62, 0xa1,
	0xc2, 0x4e, 0x60, 0x9b, 0xf8, 0x29, 0x8e, 0x95, 0x96, 0x5f, 0x33, 0x1a, 0x2b, 0x9d, 0x25, 0xe8,
	0xb7, 0x97, 0x9a, 0x6c, 0xcb, 0x1a, 0x8d, 0x9c, 0xcf, 0xc8, 0xda, 0xf4, 0xb7, 0x61, 0xab, 0x0a,
	0x37, 0x46, 0xa3, 0x64, 0x61, 0xb0, 0xff, 0x1d, 0xd6, 0x87, 0x46, 0xbc, 0xcb, 0xa5, 0xc1, 0x07,
	0x02, 0xdf, 0x84, 0x66, 0x96, 0xba, 0xa0, 0x5b, 0x71, 0x33, 0x4b, 0xd9, 0x07, 0x68, 0xf3, 0xa9,
	0x2c, 0x0b, 0x5a, 0x22, 0xb3, 0xc3, 0x82, 0xe2, 0x4a, 0xdd, 0x67, 0xf0, 0xac, 0xee, 0x5e, 0x4f,
	0xb4, 0xff, 0xd3, 0x83, 0x95, 0xa1, 0x11, 0x6c, 0x04, 0x2d, 0xf7, 0x1b, 0x74, 0xef, 0xfb, 0xb0,
	0xd5, 0x55, 0x3a, 0xaf, 0x1e, 0x01, 0x6a, 0x67, 0xf6, 0x05, 0x56, 0x6f, 0x2e, 0xda, 0x7b, 0x40,
	0xe1, 0x88, 0xce, 0xe0, 0x31, 0xa2, 0x36, 0x3d, 0xf8, 0x78, 0x39, 0x0b, 0xbc, 0xab, 0x59, 0xe0,
	0xfd, 0x9d, 0x05, 0xde, 0x8f, 0x79, 0xd0, 0xb8, 0x9a, 0x07, 0x8d, 0x5f, 0xf3, 0xa0, 0x71, 0xb2,
	0xbb, 0x10, 0xc6, 0xe7, 0x92, 0xa7, 0x9a, 0x13, 0x1e, 0x69, 0x11, 0xb9, 0x55, 0x38, 0x5f, 0x58,
	0x06, 0x97, 0xcb, 0xa4, 0xed, 0x36, 0xe0, 0xcd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0xf9,
	0xee, 0xb8, 0x7a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Open(ctx context.Context, in *MsgOpen, opts ...grpc.CallOption) (*MsgOpenResponse, error)
	Close(ctx context.Context, in *MsgClose, opts ...grpc.CallOption) (*MsgCloseResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Open(ctx context.Context, in *MsgOpen, opts ...grpc.CallOption) (*MsgOpenResponse, error) {
	out := new(MsgOpenResponse)
	err := c.cc.Invoke(ctx, "/core.perpetual.v1beta1.Msg/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Close(ctx context.Context, in *MsgClose, opts ...grpc.CallOption) (*MsgCloseResponse, error) {
	out := new(MsgCloseResponse)
	err := c.cc.Invoke(ctx, "/core.perpetual.v1beta1.Msg/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Open(context.Context, *MsgOpen) (*MsgOpenResponse, error)
	Close(context.Context, *MsgClose) (*MsgCloseResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Open(ctx context.Context, req *MsgOpen) (*MsgOpenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (*UnimplementedMsgServer) Close(ctx context.Context, req *MsgClose) (*MsgCloseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgOpen)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.perpetual.v1beta1.Msg/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Open(ctx, req.(*MsgOpen))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClose)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.perpetual.v1beta1.Msg/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Close(ctx, req.(*MsgClose))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "core.perpetual.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Open",
			Handler:    _Msg_Open_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Msg_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/perpetual/v1beta1/tx.proto",
}

func (m *MsgOpen) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgOpen) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgOpen) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TakeProfitPrice.Size()
		i -= size
		if _, err := m.TakeProfitPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Collateral) > 0 {
		i -= len(m.Collateral)
		copy(dAtA[i:], m.Collateral)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Collateral)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TradingAsset) > 0 {
		i -= len(m.TradingAsset)
		copy(dAtA[i:], m.TradingAsset)
		i = encodeVarintTx(dAtA, i, uint64(len(m.TradingAsset)))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.Leverage.Size()
		i -= size
		if _, err := m.Leverage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.TradeType != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TradeType))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgOpenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgOpenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgOpenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgClose) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClose) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClose) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCloseResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCloseResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCloseResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgOpen) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TradeType != 0 {
		n += 1 + sovTx(uint64(m.TradeType))
	}
	l = m.Leverage.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.TradingAsset)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Collateral)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.TakeProfitPrice.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgOpenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgClose) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgCloseResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgOpen) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgOpen: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgOpen: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeType", wireType)
			}
			m.TradeType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leverage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Leverage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradingAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradingAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collateral = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakeProfitPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TakeProfitPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgOpenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgOpenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgOpenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgClose) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClose: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClose: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCloseResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCloseResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCloseResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
