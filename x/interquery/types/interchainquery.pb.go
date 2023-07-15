// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/interquery/v1beta1/interchainquery.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Query struct {
	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ConnectionId string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	ChainId      string `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	QueryType    string `protobuf:"bytes,4,opt,name=query_type,json=queryType,proto3" json:"query_type,omitempty"`
	Request      []byte `protobuf:"bytes,5,opt,name=request,proto3" json:"request,omitempty"`
	// change these to uint64 in v0.5.0
	Period       int64                                  `protobuf:"varint,6,opt,name=period,proto3" json:"period,omitempty"`
	LastHeight   int64                                  `protobuf:"varint,7,opt,name=last_height,json=lastHeight,proto3" json:"last_height,omitempty"`
	CallbackId   string                                 `protobuf:"bytes,8,opt,name=callback_id,json=callbackId,proto3" json:"callback_id,omitempty"`
	Ttl          uint64                                 `protobuf:"varint,9,opt,name=ttl,proto3" json:"ttl,omitempty"`
	LastEmission github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=last_emission,json=lastEmission,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"last_emission"`
	FromAddress  string                                 `protobuf:"bytes,11,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_984fb9cf8c413fcb, []int{0}
}
func (m *Query) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Query.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return m.Size()
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Query) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *Query) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *Query) GetQueryType() string {
	if m != nil {
		return m.QueryType
	}
	return ""
}

func (m *Query) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Query) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *Query) GetLastHeight() int64 {
	if m != nil {
		return m.LastHeight
	}
	return 0
}

func (m *Query) GetCallbackId() string {
	if m != nil {
		return m.CallbackId
	}
	return ""
}

func (m *Query) GetTtl() uint64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *Query) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

type DataPoint struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// change these to uint64 in v0.5.0
	RemoteHeight github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=remote_height,json=remoteHeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"remote_height"`
	LocalHeight  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=local_height,json=localHeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"local_height"`
	Value        []byte                                 `protobuf:"bytes,4,opt,name=value,proto3" json:"result,omitempty"`
}

func (m *DataPoint) Reset()         { *m = DataPoint{} }
func (m *DataPoint) String() string { return proto.CompactTextString(m) }
func (*DataPoint) ProtoMessage()    {}
func (*DataPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_984fb9cf8c413fcb, []int{1}
}
func (m *DataPoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataPoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataPoint.Merge(m, src)
}
func (m *DataPoint) XXX_Size() int {
	return m.Size()
}
func (m *DataPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_DataPoint.DiscardUnknown(m)
}

var xxx_messageInfo_DataPoint proto.InternalMessageInfo

func (m *DataPoint) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataPoint) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Query)(nil), "core.interquery.v1beta1.Query")
	proto.RegisterType((*DataPoint)(nil), "core.interquery.v1beta1.DataPoint")
}

func init() {
	proto.RegisterFile("core/interquery/v1beta1/interchainquery.proto", fileDescriptor_984fb9cf8c413fcb)
}

var fileDescriptor_984fb9cf8c413fcb = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x93, 0x26, 0x69, 0x26, 0x0e, 0xaa, 0xac, 0x08, 0xb6, 0x95, 0x70, 0xa2, 0x22, 0xa1,
	0x08, 0x11, 0x47, 0x15, 0x47, 0xb8, 0x10, 0x81, 0xd4, 0x9c, 0xa0, 0x86, 0x13, 0x17, 0x6b, 0xe3,
	0x5d, 0x9c, 0x55, 0x6d, 0x6f, 0xba, 0x3b, 0xae, 0xc8, 0x37, 0x70, 0xe1, 0x03, 0xf8, 0x8c, 0x7e,
	0x44, 0x8f, 0x55, 0x4f, 0x88, 0x43, 0x84, 0x92, 0x1b, 0x5f, 0x81, 0x76, 0xed, 0xa8, 0x95, 0xb8,
	0xf6, 0x94, 0x9d, 0xf7, 0xde, 0xbc, 0x79, 0x19, 0xef, 0xc2, 0x38, 0x96, 0x8a, 0x4f, 0x44, 0x8e,
	0x5c, 0x5d, 0x14, 0x5c, 0xad, 0x26, 0x97, 0x27, 0x73, 0x8e, 0xf4, 0xa4, 0x84, 0xe2, 0x05, 0x15,
	0xb9, 0xc5, 0x83, 0xa5, 0x92, 0x28, 0xbd, 0x27, 0x46, 0x1e, 0xdc, 0xc9, 0x83, 0x4a, 0x7e, 0x74,
	0x18, 0x4b, 0x9d, 0x49, 0x1d, 0x59, 0xd9, 0xa4, 0x2c, 0xca, 0x9e, 0xa3, 0x7e, 0x22, 0x13, 0x59,
	0xe2, 0xe6, 0x54, 0xa2, 0xc7, 0x3f, 0x1b, 0xd0, 0x3c, 0x33, 0x16, 0xde, 0x23, 0xa8, 0x0b, 0x46,
	0x9c, 0xa1, 0x33, 0xea, 0x84, 0x75, 0xc1, 0xbc, 0x67, 0xd0, 0x8b, 0x65, 0x9e, 0xf3, 0x18, 0x85,
	0xcc, 0x23, 0xc1, 0x48, 0xdd, 0x52, 0xee, 0x1d, 0x38, 0x63, 0xde, 0x21, 0xec, 0xdb, 0x70, 0x86,
	0x6f, 0x58, 0xbe, 0x6d, 0xeb, 0x19, 0xf3, 0x9e, 0x02, 0xd8, 0x6c, 0x11, 0xae, 0x96, 0x9c, 0xec,
	0x59, 0xb2, 0x63, 0x91, 0xcf, 0xab, 0x25, 0xf7, 0x08, 0xb4, 0x15, 0xbf, 0x28, 0xb8, 0x46, 0xd2,
	0x1c, 0x3a, 0x23, 0x37, 0xdc, 0x95, 0xde, 0x63, 0x68, 0x2d, 0xb9, 0x12, 0x92, 0x91, 0xd6, 0xd0,
	0x19, 0x35, 0xc2, 0xaa, 0xf2, 0x06, 0xd0, 0x4d, 0xa9, 0xc6, 0x68, 0xc1, 0x45, 0xb2, 0x40, 0xd2,
	0xb6, 0x24, 0x18, 0xe8, 0xd4, 0x22, 0x46, 0x10, 0xd3, 0x34, 0x9d, 0xd3, 0xf8, 0xdc, 0xe4, 0xd9,
	0xb7, 0x23, 0x61, 0x07, 0xcd, 0x98, 0x77, 0x00, 0x0d, 0xc4, 0x94, 0x74, 0x86, 0xce, 0x68, 0x2f,
	0x34, 0x47, 0x8f, 0x42, 0xcf, 0x7a, 0xf2, 0x4c, 0x68, 0x2d, 0x64, 0x4e, 0xc0, 0x34, 0x4d, 0xdf,
	0x5c, 0xaf, 0x07, 0xb5, 0xdf, 0xeb, 0xc1, 0xf3, 0x44, 0xe0, 0xa2, 0x98, 0x07, 0xb1, 0xcc, 0xaa,
	0x65, 0x56, 0x3f, 0x63, 0xcd, 0xce, 0x27, 0xe6, 0x8f, 0xe9, 0x60, 0x96, 0xe3, 0xed, 0xd5, 0x18,
	0xaa, 0x5d, 0xcf, 0x72, 0x0c, 0x5d, 0x63, 0xf9, 0xbe, 0x72, 0xf4, 0x5e, 0x83, 0xfb, 0x55, 0xc9,
	0x2c, 0xa2, 0x8c, 0x29, 0xae, 0x35, 0xe9, 0xda, 0x09, 0xe4, 0xf6, 0x6a, 0xdc, 0xaf, 0x7a, 0xde,
	0x96, 0xcc, 0x27, 0x54, 0x22, 0x4f, 0xc2, 0xae, 0x51, 0x57, 0xd0, 0xf1, 0xf7, 0x3a, 0x74, 0xde,
	0x51, 0xa4, 0x1f, 0xa5, 0xc8, 0xf1, 0xbf, 0x4f, 0x44, 0xa1, 0xa7, 0x78, 0x26, 0x91, 0xef, 0x76,
	0x52, 0x7f, 0x88, 0xf4, 0xa5, 0x65, 0xb5, 0xd3, 0x08, 0xdc, 0x54, 0xc6, 0x34, 0xdd, 0x4d, 0x68,
	0x3c, 0xc0, 0x84, 0xae, 0x75, 0xac, 0x06, 0xbc, 0x80, 0xe6, 0x25, 0x4d, 0x8b, 0xf2, 0x86, 0xb8,
	0xd3, 0xfe, 0xdf, 0xf5, 0xe0, 0x40, 0x71, 0x5d, 0xa4, 0xf8, 0x52, 0x66, 0x02, 0x79, 0xb6, 0xc4,
	0x55, 0x58, 0x4a, 0xa6, 0xa7, 0xd7, 0x1b, 0xdf, 0xb9, 0xd9, 0xf8, 0xce, 0x9f, 0x8d, 0xef, 0xfc,
	0xd8, 0xfa, 0xb5, 0x9b, 0xad, 0x5f, 0xfb, 0xb5, 0xf5, 0x6b, 0x5f, 0x82, 0x7b, 0x41, 0xce, 0x0a,
	0xca, 0x14, 0x45, 0xfe, 0x41, 0x25, 0x13, 0xfb, 0xac, 0xbe, 0xdd, 0x7f, 0x58, 0x36, 0xd4, 0xbc,
	0x65, 0x6f, 0xff, 0xab, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x68, 0x2f, 0xfb, 0x78, 0x03,
	0x00, 0x00,
}

func (m *Query) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Query) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Query) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x5a
	}
	{
		size := m.LastEmission.Size()
		i -= size
		if _, err := m.LastEmission.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInterchainquery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.Ttl != 0 {
		i = encodeVarintInterchainquery(dAtA, i, uint64(m.Ttl))
		i--
		dAtA[i] = 0x48
	}
	if len(m.CallbackId) > 0 {
		i -= len(m.CallbackId)
		copy(dAtA[i:], m.CallbackId)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.CallbackId)))
		i--
		dAtA[i] = 0x42
	}
	if m.LastHeight != 0 {
		i = encodeVarintInterchainquery(dAtA, i, uint64(m.LastHeight))
		i--
		dAtA[i] = 0x38
	}
	if m.Period != 0 {
		i = encodeVarintInterchainquery(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Request) > 0 {
		i -= len(m.Request)
		copy(dAtA[i:], m.Request)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.Request)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.QueryType) > 0 {
		i -= len(m.QueryType)
		copy(dAtA[i:], m.QueryType)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.QueryType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DataPoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataPoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataPoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.LocalHeight.Size()
		i -= size
		if _, err := m.LocalHeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInterchainquery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.RemoteHeight.Size()
		i -= size
		if _, err := m.RemoteHeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInterchainquery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintInterchainquery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInterchainquery(dAtA []byte, offset int, v uint64) int {
	offset -= sovInterchainquery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Query) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	l = len(m.QueryType)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	l = len(m.Request)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	if m.Period != 0 {
		n += 1 + sovInterchainquery(uint64(m.Period))
	}
	if m.LastHeight != 0 {
		n += 1 + sovInterchainquery(uint64(m.LastHeight))
	}
	l = len(m.CallbackId)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	if m.Ttl != 0 {
		n += 1 + sovInterchainquery(uint64(m.Ttl))
	}
	l = m.LastEmission.Size()
	n += 1 + l + sovInterchainquery(uint64(l))
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	return n
}

func (m *DataPoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	l = m.RemoteHeight.Size()
	n += 1 + l + sovInterchainquery(uint64(l))
	l = m.LocalHeight.Size()
	n += 1 + l + sovInterchainquery(uint64(l))
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovInterchainquery(uint64(l))
	}
	return n
}

func sovInterchainquery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInterchainquery(x uint64) (n int) {
	return sovInterchainquery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Query) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterchainquery
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
			return fmt.Errorf("proto: Query: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Query: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Request = append(m.Request[:0], dAtA[iNdEx:postIndex]...)
			if m.Request == nil {
				m.Request = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastHeight", wireType)
			}
			m.LastHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallbackId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ttl", wireType)
			}
			m.Ttl = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ttl |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastEmission", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastEmission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInterchainquery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterchainquery
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
func (m *DataPoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterchainquery
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
			return fmt.Errorf("proto: DataPoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataPoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemoteHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
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
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LocalHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterchainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInterchainquery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInterchainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInterchainquery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterchainquery
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
func skipInterchainquery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInterchainquery
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
					return 0, ErrIntOverflowInterchainquery
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
					return 0, ErrIntOverflowInterchainquery
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
				return 0, ErrInvalidLengthInterchainquery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInterchainquery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInterchainquery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInterchainquery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInterchainquery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInterchainquery = fmt.Errorf("proto: unexpected end of group")
)
