// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/stable/pair.proto

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

type Pair struct {
	Id                uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PairId            string                                 `protobuf:"bytes,2,opt,name=pairId,proto3" json:"pairId,omitempty"`
	AmountInMetadata  types.Metadata                         `protobuf:"bytes,3,opt,name=amountInMetadata,proto3" json:"amountInMetadata" yaml:"amountInMetadata"`
	AmountOutMetadata types.Metadata                         `protobuf:"bytes,4,opt,name=amountOutMetadata,proto3" json:"amountOutMetadata" yaml:"amountOutMetadata"`
	Qm                github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=qm,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"qm" yaml:"qm"`
	Ar                github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=ar,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"ar" yaml:"ar"`
	MinAmountInt      string                                 `protobuf:"bytes,7,opt,name=minAmountInt,proto3" json:"minAmountInt,omitempty"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_25f7dcf26ad59df9, []int{0}
}
func (m *Pair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return m.Size()
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pair) GetPairId() string {
	if m != nil {
		return m.PairId
	}
	return ""
}

func (m *Pair) GetAmountInMetadata() types.Metadata {
	if m != nil {
		return m.AmountInMetadata
	}
	return types.Metadata{}
}

func (m *Pair) GetAmountOutMetadata() types.Metadata {
	if m != nil {
		return m.AmountOutMetadata
	}
	return types.Metadata{}
}

func (m *Pair) GetMinAmountInt() string {
	if m != nil {
		return m.MinAmountInt
	}
	return ""
}

func init() {
	proto.RegisterType((*Pair)(nil), "core.stable.v1beta1.Pair")
}

func init() { proto.RegisterFile("core/stable/pair.proto", fileDescriptor_25f7dcf26ad59df9) }

var fileDescriptor_25f7dcf26ad59df9 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x93, 0xb4, 0xb7, 0x97, 0xce, 0xbd, 0x88, 0x46, 0xa9, 0xa1, 0x60, 0x12, 0xb2, 0x90,
	0x22, 0x38, 0x43, 0x75, 0xa7, 0x2b, 0x8b, 0x9b, 0x22, 0x52, 0xcd, 0xd2, 0xdd, 0x49, 0x13, 0x62,
	0x6c, 0x27, 0xd3, 0x4e, 0x26, 0x62, 0xdf, 0xc2, 0xb7, 0xb2, 0xcb, 0x2e, 0xc5, 0x45, 0x90, 0xf6,
	0x0d, 0xfa, 0x04, 0x92, 0x4c, 0x2a, 0xd5, 0x6e, 0x04, 0x57, 0xc9, 0xf9, 0xcf, 0xe1, 0xfb, 0xc8,
	0xc9, 0x41, 0x8d, 0x3e, 0xe3, 0x01, 0x49, 0x04, 0x78, 0xc3, 0x80, 0x8c, 0x20, 0xe2, 0x78, 0xc4,
	0x99, 0x60, 0xfa, 0x6e, 0x9e, 0x63, 0x99, 0xe3, 0xc7, 0xb6, 0x17, 0x08, 0x68, 0x37, 0xf7, 0x42,
	0x16, 0xb2, 0xa2, 0x4f, 0xf2, 0x37, 0x39, 0xda, 0x34, 0xfb, 0x2c, 0xa1, 0x2c, 0x21, 0x1e, 0xc4,
	0x03, 0x52, 0x8e, 0x16, 0x85, 0xec, 0x3b, 0x2f, 0x15, 0x54, 0xbd, 0x81, 0x88, 0xeb, 0x5b, 0x48,
	0x8b, 0x7c, 0x43, 0xb5, 0xd5, 0x56, 0xd5, 0xd5, 0x22, 0x5f, 0x6f, 0xa0, 0x5a, 0x6e, 0xec, 0xfa,
	0x86, 0x66, 0xab, 0xad, 0xba, 0x5b, 0x56, 0xfa, 0x03, 0xda, 0x06, 0xca, 0xd2, 0x58, 0x74, 0xe3,
	0xeb, 0x40, 0x80, 0x0f, 0x02, 0x8c, 0x8a, 0xad, 0xb6, 0xfe, 0x9d, 0x1c, 0x60, 0xe9, 0xc2, 0x05,
	0xbe, 0x74, 0xe1, 0xd5, 0x50, 0xc7, 0x9a, 0x66, 0x96, 0xb2, 0xcc, 0xac, 0xfd, 0x09, 0xd0, 0xe1,
	0x99, 0xf3, 0x1d, 0xe2, 0xb8, 0x1b, 0x5c, 0x9d, 0xa2, 0x1d, 0x99, 0xf5, 0x52, 0xf1, 0x29, 0xab,
	0xfe, 0x44, 0x66, 0x97, 0x32, 0x63, 0x5d, 0xb6, 0x46, 0x71, 0xdc, 0x4d, 0xb2, 0x7e, 0x85, 0xb4,
	0x31, 0x35, 0xfe, 0xe4, 0x9f, 0xdb, 0x39, 0xcf, 0x01, 0x6f, 0x99, 0x75, 0x18, 0x46, 0xe2, 0x3e,
	0xf5, 0x70, 0x9f, 0x51, 0x52, 0xae, 0x52, 0x3e, 0x8e, 0x13, 0x7f, 0x40, 0xc4, 0x64, 0x14, 0x24,
	0xb8, 0x1b, 0x8b, 0x65, 0x66, 0xd5, 0xa5, 0x6a, 0x4c, 0x1d, 0x57, 0x1b, 0xd3, 0x1c, 0x06, 0xdc,
	0xa8, 0xfd, 0x0e, 0x06, 0xdc, 0x71, 0x35, 0xe0, 0xba, 0x83, 0xfe, 0xd3, 0x28, 0xbe, 0x28, 0xf7,
	0x23, 0x8c, 0xbf, 0xc5, 0x2f, 0xf9, 0x92, 0x75, 0x2e, 0xa7, 0x73, 0x53, 0x9d, 0xcd, 0x4d, 0xf5,
	0x7d, 0x6e, 0xaa, 0xcf, 0x0b, 0x53, 0x99, 0x2d, 0x4c, 0xe5, 0x75, 0x61, 0x2a, 0x77, 0x47, 0x6b,
	0xda, 0xdb, 0x14, 0x7c, 0x0e, 0x22, 0xe8, 0xf1, 0x90, 0x14, 0xd7, 0xf5, 0xb4, 0xba, 0xaf, 0x42,
	0xef, 0xd5, 0x8a, 0xb3, 0x38, 0xfd, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x00, 0xca, 0xf7, 0x7b,
	0x02, 0x00, 0x00,
}

func (m *Pair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinAmountInt) > 0 {
		i -= len(m.MinAmountInt)
		copy(dAtA[i:], m.MinAmountInt)
		i = encodeVarintPair(dAtA, i, uint64(len(m.MinAmountInt)))
		i--
		dAtA[i] = 0x3a
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
	dAtA[i] = 0x32
	{
		size := m.Qm.Size()
		i -= size
		if _, err := m.Qm.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.AmountOutMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPair(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.AmountInMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPair(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.PairId) > 0 {
		i -= len(m.PairId)
		copy(dAtA[i:], m.PairId)
		i = encodeVarintPair(dAtA, i, uint64(len(m.PairId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPair(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPair(dAtA []byte, offset int, v uint64) int {
	offset -= sovPair(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPair(uint64(m.Id))
	}
	l = len(m.PairId)
	if l > 0 {
		n += 1 + l + sovPair(uint64(l))
	}
	l = m.AmountInMetadata.Size()
	n += 1 + l + sovPair(uint64(l))
	l = m.AmountOutMetadata.Size()
	n += 1 + l + sovPair(uint64(l))
	l = m.Qm.Size()
	n += 1 + l + sovPair(uint64(l))
	l = m.Ar.Size()
	n += 1 + l + sovPair(uint64(l))
	l = len(m.MinAmountInt)
	if l > 0 {
		n += 1 + l + sovPair(uint64(l))
	}
	return n
}

func sovPair(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPair(x uint64) (n int) {
	return sovPair(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPair
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
			return fmt.Errorf("proto: Pair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPair
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
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPair
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
				return ErrInvalidLengthPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountInMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPair
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
				return ErrInvalidLengthPair
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPair
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
					return ErrIntOverflowPair
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
				return ErrInvalidLengthPair
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPair
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
				return fmt.Errorf("proto: wrong wireType = %d for field Qm", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPair
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
				return ErrInvalidLengthPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Qm.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPair
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
				return ErrInvalidLengthPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Ar.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAmountInt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPair
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
				return ErrInvalidLengthPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinAmountInt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPair
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
func skipPair(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPair
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
					return 0, ErrIntOverflowPair
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
					return 0, ErrIntOverflowPair
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
				return 0, ErrInvalidLengthPair
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPair
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPair
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPair        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPair          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPair = fmt.Errorf("proto: unexpected end of group")
)
