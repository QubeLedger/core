// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/stable/proposal.proto

package types

import (
	fmt "fmt"
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

type RegisterPairProposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// description of the proposal
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// metadata slice of the native Cosmos coins
	AmountInMetadata  types.Metadata `protobuf:"bytes,3,opt,name=amountInMetadata,proto3" json:"amountInMetadata" yaml:"amountInMetadata"`
	AmountOutMetadata types.Metadata `protobuf:"bytes,4,opt,name=amountOutMetadata,proto3" json:"amountOutMetadata" yaml:"amountOutMetadata"`
	MinAmountIn       string         `protobuf:"bytes,5,opt,name=minAmountIn,proto3" json:"minAmountIn,omitempty"`
	MinAmountOut      string         `protobuf:"bytes,6,opt,name=minAmountOut,proto3" json:"minAmountOut,omitempty"`
}

func (m *RegisterPairProposal) Reset()         { *m = RegisterPairProposal{} }
func (m *RegisterPairProposal) String() string { return proto.CompactTextString(m) }
func (*RegisterPairProposal) ProtoMessage()    {}
func (*RegisterPairProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_65b2ec78985a1875, []int{0}
}
func (m *RegisterPairProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterPairProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterPairProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterPairProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterPairProposal.Merge(m, src)
}
func (m *RegisterPairProposal) XXX_Size() int {
	return m.Size()
}
func (m *RegisterPairProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterPairProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterPairProposal proto.InternalMessageInfo

func (m *RegisterPairProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RegisterPairProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RegisterPairProposal) GetAmountInMetadata() types.Metadata {
	if m != nil {
		return m.AmountInMetadata
	}
	return types.Metadata{}
}

func (m *RegisterPairProposal) GetAmountOutMetadata() types.Metadata {
	if m != nil {
		return m.AmountOutMetadata
	}
	return types.Metadata{}
}

func (m *RegisterPairProposal) GetMinAmountIn() string {
	if m != nil {
		return m.MinAmountIn
	}
	return ""
}

func (m *RegisterPairProposal) GetMinAmountOut() string {
	if m != nil {
		return m.MinAmountOut
	}
	return ""
}

type ProposalMetadata struct {
	AmountInMetadata  types.Metadata `protobuf:"bytes,1,opt,name=amountInMetadata,proto3" json:"amountInMetadata" yaml:"amountInMetadata"`
	AmountOutMetadata types.Metadata `protobuf:"bytes,2,opt,name=amountOutMetadata,proto3" json:"amountOutMetadata" yaml:"amountOutMetadata"`
	MinAmountIn       string         `protobuf:"bytes,3,opt,name=minAmountIn,proto3" json:"minAmountIn,omitempty"`
	MinAmountOut      string         `protobuf:"bytes,4,opt,name=minAmountOut,proto3" json:"minAmountOut,omitempty"`
}

func (m *ProposalMetadata) Reset()         { *m = ProposalMetadata{} }
func (m *ProposalMetadata) String() string { return proto.CompactTextString(m) }
func (*ProposalMetadata) ProtoMessage()    {}
func (*ProposalMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_65b2ec78985a1875, []int{1}
}
func (m *ProposalMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalMetadata.Merge(m, src)
}
func (m *ProposalMetadata) XXX_Size() int {
	return m.Size()
}
func (m *ProposalMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalMetadata proto.InternalMessageInfo

func (m *ProposalMetadata) GetAmountInMetadata() types.Metadata {
	if m != nil {
		return m.AmountInMetadata
	}
	return types.Metadata{}
}

func (m *ProposalMetadata) GetAmountOutMetadata() types.Metadata {
	if m != nil {
		return m.AmountOutMetadata
	}
	return types.Metadata{}
}

func (m *ProposalMetadata) GetMinAmountIn() string {
	if m != nil {
		return m.MinAmountIn
	}
	return ""
}

func (m *ProposalMetadata) GetMinAmountOut() string {
	if m != nil {
		return m.MinAmountOut
	}
	return ""
}

type RegisterChangeBurningFundAddressProposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// description of the proposal
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *RegisterChangeBurningFundAddressProposal) Reset() {
	*m = RegisterChangeBurningFundAddressProposal{}
}
func (m *RegisterChangeBurningFundAddressProposal) String() string { return proto.CompactTextString(m) }
func (*RegisterChangeBurningFundAddressProposal) ProtoMessage()    {}
func (*RegisterChangeBurningFundAddressProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_65b2ec78985a1875, []int{2}
}
func (m *RegisterChangeBurningFundAddressProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterChangeBurningFundAddressProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterChangeBurningFundAddressProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterChangeBurningFundAddressProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterChangeBurningFundAddressProposal.Merge(m, src)
}
func (m *RegisterChangeBurningFundAddressProposal) XXX_Size() int {
	return m.Size()
}
func (m *RegisterChangeBurningFundAddressProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterChangeBurningFundAddressProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterChangeBurningFundAddressProposal proto.InternalMessageInfo

func (m *RegisterChangeBurningFundAddressProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RegisterChangeBurningFundAddressProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RegisterChangeBurningFundAddressProposal) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterPairProposal)(nil), "core.stable.v1beta1.RegisterPairProposal")
	proto.RegisterType((*ProposalMetadata)(nil), "core.stable.v1beta1.ProposalMetadata")
	proto.RegisterType((*RegisterChangeBurningFundAddressProposal)(nil), "core.stable.v1beta1.RegisterChangeBurningFundAddressProposal")
}

func init() { proto.RegisterFile("core/stable/proposal.proto", fileDescriptor_65b2ec78985a1875) }

var fileDescriptor_65b2ec78985a1875 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x3d, 0x8b, 0xdb, 0x40,
	0x10, 0xd5, 0xca, 0x1f, 0x21, 0xeb, 0x14, 0x8e, 0x62, 0x88, 0x30, 0x44, 0x12, 0xaa, 0x4c, 0x0a,
	0x09, 0x27, 0x9d, 0x3b, 0x3b, 0x21, 0x90, 0x22, 0xd8, 0x51, 0x99, 0x6e, 0x25, 0x2d, 0xf2, 0x26,
	0xd2, 0xae, 0xd8, 0x5d, 0x85, 0xb8, 0xcd, 0x2f, 0xc8, 0x4f, 0xc8, 0x7f, 0xc8, 0x9f, 0x70, 0xe9,
	0x32, 0x95, 0x39, 0xec, 0xe6, 0xea, 0x83, 0xeb, 0x0f, 0x6b, 0x25, 0xa3, 0x3b, 0x1f, 0xdc, 0xc1,
	0x81, 0x3b, 0xcd, 0xbc, 0xe1, 0xbd, 0xb7, 0x4f, 0x33, 0x70, 0x18, 0x31, 0x8e, 0x7d, 0x21, 0x51,
	0x98, 0x62, 0x3f, 0xe7, 0x2c, 0x67, 0x02, 0xa5, 0x5e, 0xce, 0x99, 0x64, 0xc6, 0xab, 0x03, 0xe6,
	0x29, 0xcc, 0xfb, 0x39, 0x0e, 0xb1, 0x44, 0xe3, 0xe1, 0x20, 0x61, 0x09, 0x2b, 0x71, 0xff, 0xf0,
	0xa5, 0x46, 0x87, 0x56, 0xc4, 0x44, 0xc6, 0x84, 0x1f, 0x22, 0xfa, 0xc3, 0xaf, 0x46, 0xcb, 0x42,
	0xe1, 0xee, 0xb5, 0x0e, 0x07, 0x01, 0x4e, 0x88, 0x90, 0x98, 0x2f, 0x10, 0xe1, 0x8b, 0x4a, 0xc9,
	0x18, 0xc0, 0x8e, 0x24, 0x32, 0xc5, 0x26, 0x70, 0xc0, 0xe8, 0x79, 0xa0, 0x0a, 0xc3, 0x81, 0xbd,
	0x18, 0x8b, 0x88, 0x93, 0x5c, 0x12, 0x46, 0x4d, 0xbd, 0xc4, 0x9a, 0x2d, 0xe3, 0x3b, 0xec, 0xa3,
	0x8c, 0x15, 0x54, 0x7e, 0xa6, 0x5f, 0xb0, 0x44, 0x31, 0x92, 0xc8, 0x6c, 0x39, 0x60, 0xd4, 0x7b,
	0xf7, 0xc6, 0x53, 0x5e, 0xbc, 0x52, 0xbe, 0xf2, 0xe2, 0xd5, 0x43, 0x33, 0x7b, 0xbd, 0xb5, 0xb5,
	0xab, 0xad, 0xfd, 0x7a, 0x85, 0xb2, 0x74, 0xe2, 0xde, 0x25, 0x71, 0x83, 0x13, 0x5e, 0x23, 0x83,
	0x2f, 0x55, 0x6f, 0x5e, 0xc8, 0xa3, 0x58, 0xfb, 0x31, 0x62, 0x4e, 0x25, 0x66, 0x36, 0xc5, 0x1a,
	0x2c, 0x6e, 0x70, 0xca, 0x7c, 0x78, 0x7c, 0x46, 0xe8, 0xb4, 0x72, 0x61, 0x76, 0xd4, 0xe3, 0x1b,
	0x2d, 0xc3, 0x85, 0x2f, 0x8e, 0xe5, 0xbc, 0x90, 0x66, 0xb7, 0x1c, 0xb9, 0xd5, 0x9b, 0xb4, 0x2f,
	0xff, 0xda, 0x9a, 0xfb, 0x4f, 0x87, 0xfd, 0x3a, 0xeb, 0xa3, 0xc0, 0x7d, 0xd9, 0x81, 0x73, 0x66,
	0xa7, 0x9f, 0x2b, 0xbb, 0xd6, 0xc3, 0xd9, 0xb5, 0x4f, 0xb3, 0x73, 0x7f, 0x03, 0x38, 0xaa, 0xb7,
	0xf5, 0xc3, 0x12, 0xd1, 0x04, 0xcf, 0x0a, 0x4e, 0x09, 0x4d, 0x3e, 0x15, 0x34, 0x9e, 0xc6, 0x31,
	0xc7, 0x42, 0x3c, 0x79, 0x83, 0x4d, 0xf8, 0x0c, 0x29, 0xaa, 0xca, 0x66, 0x5d, 0xaa, 0x5f, 0x37,
	0xfb, 0xb8, 0xde, 0x59, 0x60, 0xb3, 0xb3, 0xc0, 0xc5, 0xce, 0x02, 0x7f, 0xf6, 0x96, 0xb6, 0xd9,
	0x5b, 0xda, 0xff, 0xbd, 0xa5, 0x7d, 0x7b, 0x9b, 0x10, 0xb9, 0x2c, 0x42, 0x2f, 0x62, 0x99, 0xff,
	0xb5, 0x40, 0x31, 0x47, 0x12, 0xcf, 0x79, 0xe2, 0x97, 0xa7, 0xfc, 0xab, 0x3e, 0x66, 0xb9, 0xca,
	0xb1, 0x08, 0xbb, 0xe5, 0xfd, 0xbd, 0xbf, 0x09, 0x00, 0x00, 0xff, 0xff, 0x11, 0x2c, 0x18, 0xe6,
	0xe8, 0x03, 0x00, 0x00,
}

func (m *RegisterPairProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterPairProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterPairProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinAmountOut) > 0 {
		i -= len(m.MinAmountOut)
		copy(dAtA[i:], m.MinAmountOut)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.MinAmountOut)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.MinAmountIn) > 0 {
		i -= len(m.MinAmountIn)
		copy(dAtA[i:], m.MinAmountIn)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.MinAmountIn)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.AmountOutMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.AmountInMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProposalMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinAmountOut) > 0 {
		i -= len(m.MinAmountOut)
		copy(dAtA[i:], m.MinAmountOut)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.MinAmountOut)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MinAmountIn) > 0 {
		i -= len(m.MinAmountIn)
		copy(dAtA[i:], m.MinAmountIn)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.MinAmountIn)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.AmountOutMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.AmountInMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *RegisterChangeBurningFundAddressProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterChangeBurningFundAddressProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterChangeBurningFundAddressProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegisterPairProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = m.AmountInMetadata.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = m.AmountOutMetadata.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = len(m.MinAmountIn)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.MinAmountOut)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func (m *ProposalMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AmountInMetadata.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = m.AmountOutMetadata.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = len(m.MinAmountIn)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.MinAmountOut)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func (m *RegisterChangeBurningFundAddressProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisterPairProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: RegisterPairProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterPairProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountInMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountInMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountOutMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountOutMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAmountIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinAmountIn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAmountOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinAmountOut = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *ProposalMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: ProposalMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountInMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountInMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountOutMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountOutMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAmountIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinAmountIn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAmountOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinAmountOut = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *RegisterChangeBurningFundAddressProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: RegisterChangeBurningFundAddressProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterChangeBurningFundAddressProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
