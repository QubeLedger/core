// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feerefunder/fee.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Fee defines the ICS29 receive, acknowledgement and timeout fees
type Fee struct {
	// the packet receive fee
	RecvFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=recv_fee,json=recvFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"recv_fee" yaml:"recv_fee"`
	// the packet acknowledgement fee
	AckFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=ack_fee,json=ackFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"ack_fee" yaml:"ack_fee"`
	// the packet timeout fee
	TimeoutFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=timeout_fee,json=timeoutFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"timeout_fee" yaml:"timeout_fee"`
}

func (m *Fee) Reset()         { *m = Fee{} }
func (m *Fee) String() string { return proto.CompactTextString(m) }
func (*Fee) ProtoMessage()    {}
func (*Fee) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c6cd4ef4b890305, []int{0}
}
func (m *Fee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fee.Merge(m, src)
}
func (m *Fee) XXX_Size() int {
	return m.Size()
}
func (m *Fee) XXX_DiscardUnknown() {
	xxx_messageInfo_Fee.DiscardUnknown(m)
}

var xxx_messageInfo_Fee proto.InternalMessageInfo

func (m *Fee) GetRecvFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.RecvFee
	}
	return nil
}

func (m *Fee) GetAckFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.AckFee
	}
	return nil
}

func (m *Fee) GetTimeoutFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TimeoutFee
	}
	return nil
}

type PacketID struct {
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	PortId    string `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	Sequence  uint64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *PacketID) Reset()         { *m = PacketID{} }
func (m *PacketID) String() string { return proto.CompactTextString(m) }
func (*PacketID) ProtoMessage()    {}
func (*PacketID) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c6cd4ef4b890305, []int{1}
}
func (m *PacketID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketID.Merge(m, src)
}
func (m *PacketID) XXX_Size() int {
	return m.Size()
}
func (m *PacketID) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketID.DiscardUnknown(m)
}

var xxx_messageInfo_PacketID proto.InternalMessageInfo

func (m *PacketID) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *PacketID) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *PacketID) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*Fee)(nil), "neutron.feerefunder.Fee")
	proto.RegisterType((*PacketID)(nil), "neutron.feerefunder.PacketID")
}

func init() { proto.RegisterFile("feerefunder/fee.proto", fileDescriptor_0c6cd4ef4b890305) }

var fileDescriptor_0c6cd4ef4b890305 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6b, 0xe2, 0x40,
	0x18, 0xc6, 0x33, 0xba, 0xf8, 0x67, 0x84, 0x5d, 0xc8, 0xee, 0xb2, 0xae, 0xb0, 0x51, 0x72, 0xca,
	0xc5, 0x0c, 0xba, 0xb7, 0x3d, 0xea, 0x22, 0x08, 0x3d, 0x14, 0x8f, 0x3d, 0x54, 0x92, 0xc9, 0x9b,
	0x18, 0x62, 0x66, 0x6c, 0x32, 0x91, 0x7a, 0xed, 0x27, 0xe8, 0xe7, 0xe8, 0x27, 0xf1, 0xe8, 0xb1,
	0x27, 0x5b, 0xf4, 0x1b, 0xf4, 0x5e, 0x28, 0x93, 0x8c, 0x25, 0x37, 0xf1, 0x34, 0xef, 0x1f, 0x9e,
	0xe7, 0xf7, 0x0c, 0xbc, 0xf8, 0xa7, 0x0f, 0x90, 0x80, 0x9f, 0x31, 0x0f, 0x12, 0xe2, 0x03, 0xd8,
	0xab, 0x84, 0x0b, 0xae, 0x7f, 0x67, 0x90, 0x89, 0x84, 0x33, 0xbb, 0xb4, 0xee, 0x18, 0x94, 0xa7,
	0x31, 0x4f, 0x89, 0xeb, 0xa4, 0x40, 0xd6, 0x03, 0x17, 0x84, 0x33, 0x20, 0x94, 0x87, 0xac, 0x10,
	0x75, 0x7e, 0x04, 0x3c, 0xe0, 0x79, 0x49, 0x64, 0x55, 0x4c, 0xcd, 0xf7, 0x0a, 0xae, 0x4e, 0x00,
	0xf4, 0x0d, 0x6e, 0x24, 0x40, 0xd7, 0x73, 0x1f, 0xa0, 0x8d, 0x7a, 0x55, 0xab, 0x35, 0xfc, 0x6d,
	0x17, 0x86, 0xb6, 0x34, 0xb4, 0x95, 0xa1, 0x3d, 0xe6, 0x21, 0x1b, 0x8d, 0xb7, 0xfb, 0xae, 0xf6,
	0xb6, 0xef, 0x7e, 0xdb, 0x38, 0xf1, 0xf2, 0x9f, 0x79, 0x12, 0x9a, 0x4f, 0x2f, 0x5d, 0x2b, 0x08,
	0xc5, 0x22, 0x73, 0x6d, 0xca, 0x63, 0xa2, 0x02, 0x15, 0x4f, 0x3f, 0xf5, 0x22, 0x22, 0x36, 0x2b,
	0x48, 0x73, 0x8f, 0x74, 0x56, 0x97, 0x32, 0x89, 0x5e, 0xe3, 0xba, 0x43, 0xa3, 0x9c, 0x5c, 0x39,
	0x47, 0x1e, 0x29, 0xf2, 0xd7, 0x82, 0xac, 0x74, 0x97, 0x81, 0x6b, 0x0e, 0x8d, 0x24, 0xf7, 0x01,
	0xe1, 0x96, 0x08, 0x63, 0xe0, 0x99, 0xc8, 0xe1, 0xd5, 0x73, 0xf0, 0x89, 0x82, 0xeb, 0x05, 0xbc,
	0xa4, 0xbd, 0x2c, 0x00, 0x56, 0xca, 0x09, 0x80, 0x79, 0x8b, 0x1b, 0xd7, 0x0e, 0x8d, 0x40, 0x4c,
	0xff, 0xeb, 0x7f, 0x30, 0xa6, 0x0b, 0x87, 0x31, 0x58, 0xce, 0x43, 0xaf, 0x8d, 0x7a, 0xc8, 0x6a,
	0xce, 0x9a, 0x6a, 0x32, 0xf5, 0xf4, 0x5f, 0xb8, 0xbe, 0xe2, 0x89, 0x90, 0xbb, 0x4a, 0xbe, 0xab,
	0xc9, 0x76, 0xea, 0xe9, 0x1d, 0xdc, 0x48, 0xe1, 0x2e, 0x03, 0x46, 0xe5, 0x27, 0x90, 0xf5, 0x65,
	0xf6, 0xd9, 0x8f, 0xae, 0xb6, 0x07, 0x03, 0xed, 0x0e, 0x06, 0x7a, 0x3d, 0x18, 0xe8, 0xf1, 0x68,
	0x68, 0xbb, 0xa3, 0xa1, 0x3d, 0x1f, 0x0d, 0xed, 0x66, 0x58, 0xca, 0xab, 0xee, 0xa9, 0xcf, 0x93,
	0xe0, 0x54, 0x93, 0x7b, 0x52, 0x3e, 0xbe, 0x3c, 0xbf, 0x5b, 0xcb, 0x8f, 0xe6, 0xef, 0x47, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x63, 0x01, 0x6f, 0xe2, 0x98, 0x02, 0x00, 0x00,
}

func (m *Fee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TimeoutFee) > 0 {
		for iNdEx := len(m.TimeoutFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TimeoutFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AckFee) > 0 {
		for iNdEx := len(m.AckFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AckFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.RecvFee) > 0 {
		for iNdEx := len(m.RecvFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RecvFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PacketID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintFee(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x18
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintFee(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintFee(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Fee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RecvFee) > 0 {
		for _, e := range m.RecvFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	if len(m.AckFee) > 0 {
		for _, e := range m.AckFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	if len(m.TimeoutFee) > 0 {
		for _, e := range m.TimeoutFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	return n
}

func (m *PacketID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovFee(uint64(l))
	}
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovFee(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovFee(uint64(m.Sequence))
	}
	return n
}

func sovFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFee(x uint64) (n int) {
	return sovFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Fee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
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
			return fmt.Errorf("proto: Fee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecvFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecvFee = append(m.RecvFee, types.Coin{})
			if err := m.RecvFee[len(m.RecvFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AckFee = append(m.AckFee, types.Coin{})
			if err := m.AckFee[len(m.AckFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TimeoutFee = append(m.TimeoutFee, types.Coin{})
			if err := m.TimeoutFee[len(m.TimeoutFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
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
func (m *PacketID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
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
			return fmt.Errorf("proto: PacketID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
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
func skipFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFee
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
					return 0, ErrIntOverflowFee
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
					return 0, ErrIntOverflowFee
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
				return 0, ErrInvalidLengthFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFee = fmt.Errorf("proto: unexpected end of group")
)