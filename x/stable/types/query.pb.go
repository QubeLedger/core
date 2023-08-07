// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/stable/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	types "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dec94a25afa6fc54, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dec94a25afa6fc54, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

type PairByPairIdRequest struct {
	PairId string `protobuf:"bytes,1,opt,name=pairId,proto3" json:"pairId,omitempty"`
}

func (m *PairByPairIdRequest) Reset()         { *m = PairByPairIdRequest{} }
func (m *PairByPairIdRequest) String() string { return proto.CompactTextString(m) }
func (*PairByPairIdRequest) ProtoMessage()    {}
func (*PairByPairIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dec94a25afa6fc54, []int{2}
}
func (m *PairByPairIdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PairByPairIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PairByPairIdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PairByPairIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairByPairIdRequest.Merge(m, src)
}
func (m *PairByPairIdRequest) XXX_Size() int {
	return m.Size()
}
func (m *PairByPairIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PairByPairIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PairByPairIdRequest proto.InternalMessageInfo

func (m *PairByPairIdRequest) GetPairId() string {
	if m != nil {
		return m.PairId
	}
	return ""
}

type PairByIdRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *PairByIdRequest) Reset()         { *m = PairByIdRequest{} }
func (m *PairByIdRequest) String() string { return proto.CompactTextString(m) }
func (*PairByIdRequest) ProtoMessage()    {}
func (*PairByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dec94a25afa6fc54, []int{3}
}
func (m *PairByIdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PairByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PairByIdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PairByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairByIdRequest.Merge(m, src)
}
func (m *PairByIdRequest) XXX_Size() int {
	return m.Size()
}
func (m *PairByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PairByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PairByIdRequest proto.InternalMessageInfo

func (m *PairByIdRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PairRequestResponse struct {
	PairId            string                                 `protobuf:"bytes,1,opt,name=pairId,proto3" json:"pairId,omitempty"`
	AmountInMetadata  types.Metadata                         `protobuf:"bytes,2,opt,name=amountInMetadata,proto3" json:"amountInMetadata" yaml:"amountInMetadata"`
	AmountOutMetadata types.Metadata                         `protobuf:"bytes,3,opt,name=amountOutMetadata,proto3" json:"amountOutMetadata" yaml:"amountOutMetadata"`
	Qm                github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=qm,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"qm" yaml:"qm"`
	Ar                github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=ar,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"ar" yaml:"ar"`
	MinAmountIn      string                                 `protobuf:"bytes,6,opt,name=minAmountInt,proto3" json:"minAmountInt,omitempty"`
}

func (m *PairRequestResponse) Reset()         { *m = PairRequestResponse{} }
func (m *PairRequestResponse) String() string { return proto.CompactTextString(m) }
func (*PairRequestResponse) ProtoMessage()    {}
func (*PairRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dec94a25afa6fc54, []int{4}
}
func (m *PairRequestResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PairRequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PairRequestResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PairRequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairRequestResponse.Merge(m, src)
}
func (m *PairRequestResponse) XXX_Size() int {
	return m.Size()
}
func (m *PairRequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PairRequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PairRequestResponse proto.InternalMessageInfo

func (m *PairRequestResponse) GetPairId() string {
	if m != nil {
		return m.PairId
	}
	return ""
}

func (m *PairRequestResponse) GetAmountInMetadata() types.Metadata {
	if m != nil {
		return m.AmountInMetadata
	}
	return types.Metadata{}
}

func (m *PairRequestResponse) GetAmountOutMetadata() types.Metadata {
	if m != nil {
		return m.AmountOutMetadata
	}
	return types.Metadata{}
}

func (m *PairRequestResponse) GetMinAmountIn() string {
	if m != nil {
		return m.MinAmountIn
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "core.stable.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "core.stable.v1beta1.QueryParamsResponse")
	proto.RegisterType((*PairByPairIdRequest)(nil), "core.stable.v1beta1.PairByPairIdRequest")
	proto.RegisterType((*PairByIdRequest)(nil), "core.stable.v1beta1.PairByIdRequest")
	proto.RegisterType((*PairRequestResponse)(nil), "core.stable.v1beta1.PairRequestResponse")
}

func init() { proto.RegisterFile("core/stable/query.proto", fileDescriptor_dec94a25afa6fc54) }

var fileDescriptor_dec94a25afa6fc54 = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x9b, 0x74, 0xab, 0x98, 0x99, 0xf8, 0xe3, 0x4e, 0x2c, 0xea, 0x58, 0x5a, 0xcc, 0x34,
	0xaa, 0x4a, 0x8b, 0xb5, 0xed, 0x06, 0x27, 0x2a, 0x2e, 0x05, 0xa1, 0x6d, 0x39, 0x72, 0x73, 0x1b,
	0x2b, 0x84, 0x35, 0x71, 0xea, 0x38, 0x88, 0x0a, 0x4d, 0x42, 0xfb, 0x00, 0x08, 0x89, 0x2f, 0xc0,
	0xa7, 0x41, 0x3b, 0x4e, 0xe2, 0x82, 0x38, 0x54, 0xa8, 0xe5, 0x13, 0xec, 0x13, 0xa0, 0xd8, 0x4e,
	0x97, 0xad, 0xeb, 0x40, 0x70, 0x69, 0x1b, 0xbf, 0x8f, 0x7f, 0xef, 0xf3, 0x3a, 0x8f, 0x0b, 0x56,
	0x7b, 0x8c, 0x53, 0x9c, 0x08, 0xd2, 0xed, 0x53, 0x3c, 0x48, 0x29, 0x1f, 0x3a, 0x31, 0x67, 0x82,
	0xc1, 0x6a, 0x56, 0x70, 0x54, 0xc1, 0x79, 0xbb, 0xdd, 0xa5, 0x82, 0x6c, 0xd7, 0x56, 0x7c, 0xe6,
	0x33, 0x59, 0xc7, 0xd9, 0x2f, 0x25, 0xad, 0xdd, 0xf7, 0x19, 0xf3, 0xfb, 0x14, 0x93, 0x38, 0xc0,
	0x24, 0x8a, 0x98, 0x20, 0x22, 0x60, 0x51, 0xa2, 0xab, 0xad, 0x1e, 0x4b, 0x42, 0x96, 0xe0, 0x2e,
	0x49, 0x74, 0x07, 0xac, 0x71, 0x38, 0x26, 0x7e, 0x10, 0x49, 0xb1, 0xd6, 0x5a, 0x45, 0x37, 0x31,
	0xe1, 0x24, 0xcc, 0x29, 0xf6, 0x94, 0x12, 0x1d, 0x4e, 0xf7, 0x67, 0x0f, 0xaa, 0x8e, 0x56, 0x00,
	0x3c, 0xc8, 0xd8, 0xfb, 0x72, 0x93, 0x4b, 0x07, 0x29, 0x4d, 0x04, 0x7a, 0x0e, 0xaa, 0x17, 0x56,
	0x93, 0x98, 0x45, 0x09, 0x85, 0xbb, 0xa0, 0xa2, 0xe0, 0x96, 0xd1, 0x30, 0x9a, 0x37, 0x77, 0xd6,
	0x9c, 0x2b, 0x86, 0x75, 0xf4, 0x26, 0x2d, 0x45, 0x5b, 0xa0, 0xba, 0x4f, 0x02, 0xde, 0x1e, 0x66,
	0x9f, 0x1d, 0x4f, 0xb7, 0x80, 0xf7, 0x32, 0x56, 0xb6, 0x20, 0x59, 0x4b, 0xae, 0x7e, 0x42, 0x0f,
	0xc0, 0x6d, 0x25, 0x3f, 0x97, 0xde, 0x02, 0x66, 0xa0, 0x64, 0x0b, 0xae, 0x19, 0x78, 0xe8, 0x6b,
	0x59, 0x21, 0x75, 0x7d, 0x6a, 0x6f, 0x0e, 0x12, 0xbe, 0x01, 0x77, 0x48, 0xc8, 0xd2, 0x48, 0x74,
	0xa2, 0x97, 0x54, 0x10, 0x8f, 0x08, 0x62, 0x99, 0x72, 0x80, 0x75, 0x47, 0x1d, 0x8f, 0x23, 0x4f,
	0x24, 0x1f, 0x20, 0x17, 0xb5, 0xeb, 0x27, 0xa3, 0x7a, 0xe9, 0x6c, 0x54, 0x5f, 0x1d, 0x92, 0xb0,
	0xff, 0x18, 0x5d, 0x86, 0x20, 0x77, 0x86, 0x0b, 0x43, 0x70, 0x57, 0xad, 0xed, 0xa5, 0x62, 0xda,
	0xac, 0xfc, 0x37, 0xcd, 0x1a, 0xba, 0x99, 0x55, 0x6c, 0x56, 0xa0, 0x20, 0x77, 0x96, 0x0c, 0x5f,
	0x00, 0x73, 0x10, 0x5a, 0x0b, 0xd9, 0xb8, 0xed, 0x27, 0x19, 0xe0, 0xc7, 0xa8, 0xbe, 0xe9, 0x07,
	0xe2, 0x75, 0xda, 0x75, 0x7a, 0x2c, 0xc4, 0xfa, 0xed, 0xab, 0xaf, 0xad, 0xc4, 0x3b, 0xc4, 0x62,
	0x18, 0xd3, 0xc4, 0xe9, 0x44, 0xe2, 0x6c, 0x54, 0x5f, 0x52, 0xad, 0x06, 0x21, 0x72, 0xcd, 0x41,
	0x98, 0xc1, 0x08, 0xb7, 0x16, 0xff, 0x0f, 0x46, 0x38, 0x72, 0x4d, 0xc2, 0x21, 0x02, 0xcb, 0x61,
	0x10, 0x3d, 0xd5, 0xe7, 0x23, 0xac, 0x8a, 0x7c, 0x25, 0x17, 0xd6, 0x76, 0xbe, 0x94, 0xc1, 0xa2,
	0xcc, 0x19, 0xfc, 0x60, 0x80, 0x8a, 0xca, 0x0d, 0x7c, 0x74, 0x65, 0xa8, 0x66, 0x43, 0x5a, 0x6b,
	0xfe, 0x59, 0xa8, 0x82, 0x81, 0x1e, 0x1e, 0x7f, 0xfb, 0xf5, 0xd9, 0x5c, 0x87, 0x6b, 0xb8, 0x78,
	0x4f, 0xce, 0x6f, 0x93, 0xec, 0xfb, 0xd1, 0x00, 0xcb, 0xc5, 0xa0, 0xc2, 0xe6, 0x9c, 0x74, 0xcf,
	0x64, 0xb9, 0x36, 0x5f, 0x79, 0x29, 0xa2, 0xa8, 0x25, 0x9d, 0x6c, 0x40, 0x34, 0xc7, 0x49, 0xc0,
	0xf1, 0x7b, 0x95, 0xda, 0x23, 0x78, 0x6c, 0x80, 0x1b, 0xf9, 0x55, 0x80, 0x1b, 0xd7, 0x98, 0xf9,
	0x17, 0x23, 0x9b, 0xd2, 0x48, 0x03, 0xda, 0xd7, 0x18, 0x09, 0xbc, 0xa3, 0xf6, 0xb3, 0x93, 0xb1,
	0x6d, 0x9c, 0x8e, 0x6d, 0xe3, 0xe7, 0xd8, 0x36, 0x3e, 0x4d, 0xec, 0xd2, 0xe9, 0xc4, 0x2e, 0x7d,
	0x9f, 0xd8, 0xa5, 0x57, 0xad, 0x42, 0x32, 0x0e, 0x52, 0xe2, 0x71, 0x22, 0xe8, 0x1e, 0xf7, 0x15,
	0xef, 0x5d, 0x4e, 0x94, 0x09, 0xe9, 0x56, 0xe4, 0x9f, 0xcd, 0xee, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xbf, 0xe5, 0xd3, 0x86, 0x36, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	PairByPairId(ctx context.Context, in *PairByPairIdRequest, opts ...grpc.CallOption) (*PairRequestResponse, error)
	PairById(ctx context.Context, in *PairByIdRequest, opts ...grpc.CallOption) (*PairRequestResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/core.stable.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PairByPairId(ctx context.Context, in *PairByPairIdRequest, opts ...grpc.CallOption) (*PairRequestResponse, error) {
	out := new(PairRequestResponse)
	err := c.cc.Invoke(ctx, "/core.stable.v1beta1.Query/PairByPairId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PairById(ctx context.Context, in *PairByIdRequest, opts ...grpc.CallOption) (*PairRequestResponse, error) {
	out := new(PairRequestResponse)
	err := c.cc.Invoke(ctx, "/core.stable.v1beta1.Query/PairById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	PairByPairId(context.Context, *PairByPairIdRequest) (*PairRequestResponse, error)
	PairById(context.Context, *PairByIdRequest) (*PairRequestResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) PairByPairId(ctx context.Context, req *PairByPairIdRequest) (*PairRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PairByPairId not implemented")
}
func (*UnimplementedQueryServer) PairById(ctx context.Context, req *PairByIdRequest) (*PairRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PairById not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.stable.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PairByPairId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PairByPairIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PairByPairId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.stable.v1beta1.Query/PairByPairId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PairByPairId(ctx, req.(*PairByPairIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PairById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PairByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PairById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.stable.v1beta1.Query/PairById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PairById(ctx, req.(*PairByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "core.stable.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "PairByPairId",
			Handler:    _Query_PairByPairId_Handler,
		},
		{
			MethodName: "PairById",
			Handler:    _Query_PairById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/stable/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PairByPairIdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PairByPairIdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PairByPairIdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PairId) > 0 {
		i -= len(m.PairId)
		copy(dAtA[i:], m.PairId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PairId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PairByIdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PairByIdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PairByIdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PairRequestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PairRequestResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PairRequestResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinAmountIn) > 0 {
		i -= len(m.MinAmountIn)
		copy(dAtA[i:], m.MinAmountIn)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.MinAmountIn)))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.Ar.Size()
		i -= size
		if _, err := m.Ar.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Qm.Size()
		i -= size
		if _, err := m.Qm.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.AmountOutMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.AmountInMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.PairId) > 0 {
		i -= len(m.PairId)
		copy(dAtA[i:], m.PairId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PairId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *PairByPairIdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PairId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *PairByIdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *PairRequestResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PairId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = m.AmountInMetadata.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.AmountOutMetadata.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.Qm.Size()
	n += 1 + l + sovPair(uint64(l))
	l = m.Ar.Size()
	n += 1 + l + sovPair(uint64(l))
	l = len(m.MinAmountIn)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *PairByPairIdRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: PairByPairIdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PairByPairIdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *PairByIdRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: PairByIdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PairByIdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *PairRequestResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: PairRequestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PairRequestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountInMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountInMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountOutMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountOutMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Qm", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Qm.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Ar.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAmountIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinAmountIn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
