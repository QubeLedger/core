// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/grow/v1beta1/params.proto

package types

import (
	fmt "fmt"
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

// Params defines the parameters for the module.
type Params struct {
	LastTimeUpdateReserve     uint64 `protobuf:"varint,1,opt,name=LastTimeUpdateReserve,proto3" json:"LastTimeUpdateReserve,omitempty"`
	GrowStakingReserveAddress string `protobuf:"bytes,2,opt,name=GrowStakingReserveAddress,proto3" json:"GrowStakingReserveAddress,omitempty"`
	USQReserveAddress         string `protobuf:"bytes,3,opt,name=USQReserveAddress,proto3" json:"USQReserveAddress,omitempty"`
	GrowYieldReserveAddress   string `protobuf:"bytes,4,opt,name=GrowYieldReserveAddress,proto3" json:"GrowYieldReserveAddress,omitempty"`
	DepositMethodStatus       bool   `protobuf:"varint,5,opt,name=DepositMethodStatus,proto3" json:"DepositMethodStatus,omitempty"`
	CollateralMethodStatus    bool   `protobuf:"varint,6,opt,name=CollateralMethodStatus,proto3" json:"CollateralMethodStatus,omitempty"`
	BorrowMethodStatus        bool   `protobuf:"varint,7,opt,name=BorrowMethodStatus,proto3" json:"BorrowMethodStatus,omitempty"`
	UStaticVolatile           uint64 `protobuf:"varint,8,opt,name=u_static_volatile,json=uStaticVolatile,proto3" json:"u_static_volatile,omitempty"`
	UStaticStable             uint64 `protobuf:"varint,9,opt,name=u_static_stable,json=uStaticStable,proto3" json:"u_static_stable,omitempty"`
	MaxRateVolatile           uint64 `protobuf:"varint,10,opt,name=max_rate_volatile,json=maxRateVolatile,proto3" json:"max_rate_volatile,omitempty"`
	MaxRateStable             uint64 `protobuf:"varint,11,opt,name=max_rate_stable,json=maxRateStable,proto3" json:"max_rate_stable,omitempty"`
	Slope                     uint64 `protobuf:"varint,12,opt,name=slope,proto3" json:"slope,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_28929114d7d669ed, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetLastTimeUpdateReserve() uint64 {
	if m != nil {
		return m.LastTimeUpdateReserve
	}
	return 0
}

func (m *Params) GetGrowStakingReserveAddress() string {
	if m != nil {
		return m.GrowStakingReserveAddress
	}
	return ""
}

func (m *Params) GetUSQReserveAddress() string {
	if m != nil {
		return m.USQReserveAddress
	}
	return ""
}

func (m *Params) GetGrowYieldReserveAddress() string {
	if m != nil {
		return m.GrowYieldReserveAddress
	}
	return ""
}

func (m *Params) GetDepositMethodStatus() bool {
	if m != nil {
		return m.DepositMethodStatus
	}
	return false
}

func (m *Params) GetCollateralMethodStatus() bool {
	if m != nil {
		return m.CollateralMethodStatus
	}
	return false
}

func (m *Params) GetBorrowMethodStatus() bool {
	if m != nil {
		return m.BorrowMethodStatus
	}
	return false
}

func (m *Params) GetUStaticVolatile() uint64 {
	if m != nil {
		return m.UStaticVolatile
	}
	return 0
}

func (m *Params) GetUStaticStable() uint64 {
	if m != nil {
		return m.UStaticStable
	}
	return 0
}

func (m *Params) GetMaxRateVolatile() uint64 {
	if m != nil {
		return m.MaxRateVolatile
	}
	return 0
}

func (m *Params) GetMaxRateStable() uint64 {
	if m != nil {
		return m.MaxRateStable
	}
	return 0
}

func (m *Params) GetSlope() uint64 {
	if m != nil {
		return m.Slope
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "core.grow.v1beta1.Params")
}

func init() { proto.RegisterFile("core/grow/v1beta1/params.proto", fileDescriptor_28929114d7d669ed) }

var fileDescriptor_28929114d7d669ed = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x13, 0xcd, 0xad, 0xf7, 0x8e, 0xca, 0xa5, 0xe3, 0x55, 0xa3, 0x8b, 0x58, 0x5c, 0x48,
	0x10, 0x49, 0x2c, 0x8a, 0x88, 0xb8, 0xb1, 0x0a, 0x6e, 0x14, 0x6d, 0x62, 0x05, 0xdd, 0x94, 0x93,
	0xe6, 0x90, 0x06, 0x27, 0x4e, 0x98, 0x39, 0x69, 0xeb, 0x5b, 0xb8, 0x74, 0xe9, 0xb3, 0xb8, 0x72,
	0xd9, 0xa5, 0x4b, 0x69, 0x5f, 0x44, 0x32, 0x09, 0xd5, 0xd6, 0x76, 0x97, 0xfc, 0xdf, 0xff, 0xfd,
	0x81, 0x70, 0x98, 0x37, 0x91, 0x0a, 0xc3, 0x4c, 0xc9, 0x79, 0x38, 0xeb, 0x27, 0x48, 0xd0, 0x0f,
	0x4b, 0x50, 0x50, 0xe8, 0xa0, 0x54, 0x92, 0x24, 0xef, 0xd6, 0x3c, 0xa8, 0x79, 0xd0, 0xf2, 0x9b,
	0x67, 0x99, 0xcc, 0xa4, 0xa1, 0x61, 0xfd, 0xd4, 0x14, 0x6f, 0xff, 0x70, 0x58, 0xe7, 0xad, 0x31,
	0xf9, 0x43, 0x76, 0xf5, 0x15, 0x68, 0x7a, 0x97, 0x17, 0x38, 0x2a, 0x53, 0x20, 0x8c, 0x50, 0xa3,
	0x9a, 0xa1, 0x6b, 0xf7, 0x6c, 0xdf, 0x89, 0xf6, 0x43, 0xfe, 0x94, 0xdd, 0x78, 0xa9, 0xe4, 0x3c,
	0x26, 0xf8, 0x94, 0x7f, 0xce, 0xda, 0xf4, 0x59, 0x9a, 0x2a, 0xd4, 0xda, 0x3d, 0xd7, 0xb3, 0xfd,
	0x93, 0xe8, 0x70, 0x81, 0xdf, 0x63, 0xdd, 0x51, 0x3c, 0xdc, 0xb1, 0xce, 0x1b, 0xeb, 0x7f, 0xc0,
	0x1f, 0xb3, 0xeb, 0xf5, 0xd4, 0x87, 0x1c, 0x45, 0xba, 0xe3, 0x38, 0xc6, 0x39, 0x84, 0xf9, 0x7d,
	0x76, 0xe5, 0x05, 0x96, 0x52, 0xe7, 0xf4, 0x1a, 0x69, 0x2a, 0xd3, 0x98, 0x80, 0x2a, 0xed, 0x1e,
	0xf5, 0x6c, 0xff, 0x38, 0xda, 0x87, 0xf8, 0x23, 0x76, 0xed, 0xb9, 0x14, 0x02, 0x08, 0x15, 0x88,
	0x2d, 0xa9, 0x63, 0xa4, 0x03, 0x94, 0x07, 0x8c, 0x0f, 0xa4, 0x52, 0x72, 0xbe, 0xe5, 0x5c, 0x30,
	0xce, 0x1e, 0xc2, 0xef, 0xb2, 0x6e, 0x35, 0xd6, 0x04, 0x94, 0x4f, 0xc6, 0x33, 0x29, 0x80, 0x72,
	0x81, 0xee, 0xb1, 0xf9, 0xe3, 0xa7, 0x55, 0x6c, 0xf2, 0xf7, 0x6d, 0xcc, 0xef, 0xb0, 0xd3, 0x4d,
	0x57, 0x13, 0x24, 0x02, 0xdd, 0x13, 0xd3, 0xbc, 0xdc, 0x36, 0x63, 0x13, 0xd6, 0x9b, 0x05, 0x2c,
	0xc6, 0x0a, 0x08, 0xff, 0x6e, 0xb2, 0x66, 0xb3, 0x80, 0x45, 0x04, 0x84, 0xff, 0x6e, 0x6e, 0xba,
	0xed, 0xe6, 0xc5, 0x66, 0xb3, 0x6d, 0xb6, 0x9b, 0x67, 0xec, 0x48, 0x0b, 0x59, 0xa2, 0x7b, 0xc9,
	0xd0, 0xe6, 0xe5, 0x89, 0xf3, 0xed, 0xfb, 0x2d, 0x6b, 0x30, 0xf8, 0xb9, 0xf2, 0xec, 0xe5, 0xca,
	0xb3, 0x7f, 0xaf, 0x3c, 0xfb, 0xeb, 0xda, 0xb3, 0x96, 0x6b, 0xcf, 0xfa, 0xb5, 0xf6, 0xac, 0x8f,
	0x7e, 0x96, 0xd3, 0xb4, 0x4a, 0x82, 0x89, 0x2c, 0xc2, 0x61, 0x05, 0x69, 0xfd, 0x99, 0x37, 0x2a,
	0x0b, 0xcd, 0xf9, 0x2e, 0x9a, 0x03, 0xa6, 0x2f, 0x25, 0xea, 0xa4, 0x63, 0xee, 0xf1, 0xc1, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x23, 0x25, 0x21, 0x4e, 0xda, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Slope != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.Slope))
		i--
		dAtA[i] = 0x60
	}
	if m.MaxRateStable != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxRateStable))
		i--
		dAtA[i] = 0x58
	}
	if m.MaxRateVolatile != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxRateVolatile))
		i--
		dAtA[i] = 0x50
	}
	if m.UStaticStable != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UStaticStable))
		i--
		dAtA[i] = 0x48
	}
	if m.UStaticVolatile != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UStaticVolatile))
		i--
		dAtA[i] = 0x40
	}
	if m.BorrowMethodStatus {
		i--
		if m.BorrowMethodStatus {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.CollateralMethodStatus {
		i--
		if m.CollateralMethodStatus {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.DepositMethodStatus {
		i--
		if m.DepositMethodStatus {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.GrowYieldReserveAddress) > 0 {
		i -= len(m.GrowYieldReserveAddress)
		copy(dAtA[i:], m.GrowYieldReserveAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.GrowYieldReserveAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.USQReserveAddress) > 0 {
		i -= len(m.USQReserveAddress)
		copy(dAtA[i:], m.USQReserveAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.USQReserveAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.GrowStakingReserveAddress) > 0 {
		i -= len(m.GrowStakingReserveAddress)
		copy(dAtA[i:], m.GrowStakingReserveAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.GrowStakingReserveAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.LastTimeUpdateReserve != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.LastTimeUpdateReserve))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LastTimeUpdateReserve != 0 {
		n += 1 + sovParams(uint64(m.LastTimeUpdateReserve))
	}
	l = len(m.GrowStakingReserveAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.USQReserveAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.GrowYieldReserveAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.DepositMethodStatus {
		n += 2
	}
	if m.CollateralMethodStatus {
		n += 2
	}
	if m.BorrowMethodStatus {
		n += 2
	}
	if m.UStaticVolatile != 0 {
		n += 1 + sovParams(uint64(m.UStaticVolatile))
	}
	if m.UStaticStable != 0 {
		n += 1 + sovParams(uint64(m.UStaticStable))
	}
	if m.MaxRateVolatile != 0 {
		n += 1 + sovParams(uint64(m.MaxRateVolatile))
	}
	if m.MaxRateStable != 0 {
		n += 1 + sovParams(uint64(m.MaxRateStable))
	}
	if m.Slope != 0 {
		n += 1 + sovParams(uint64(m.Slope))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTimeUpdateReserve", wireType)
			}
			m.LastTimeUpdateReserve = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastTimeUpdateReserve |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrowStakingReserveAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GrowStakingReserveAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field USQReserveAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.USQReserveAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrowYieldReserveAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GrowYieldReserveAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositMethodStatus", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DepositMethodStatus = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralMethodStatus", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CollateralMethodStatus = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowMethodStatus", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BorrowMethodStatus = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UStaticVolatile", wireType)
			}
			m.UStaticVolatile = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UStaticVolatile |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UStaticStable", wireType)
			}
			m.UStaticStable = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UStaticStable |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRateVolatile", wireType)
			}
			m.MaxRateVolatile = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRateVolatile |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRateStable", wireType)
			}
			m.MaxRateStable = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRateStable |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slope", wireType)
			}
			m.Slope = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Slope |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
