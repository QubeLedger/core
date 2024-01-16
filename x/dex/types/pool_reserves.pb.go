// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: duality/dex/pool_reserves.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_duality_labs_duality_utils_math "github.com/QuadrateOrg/core/x/dex/utils/math"
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

type PoolReservesKey struct {
	TradePairID           *TradePairID `protobuf:"bytes,1,opt,name=tradePairID,proto3" json:"tradePairID,omitempty"`
	TickIndexTakerToMaker int64        `protobuf:"varint,2,opt,name=TickIndexTakerToMaker,proto3" json:"TickIndexTakerToMaker,omitempty"`
	Fee                   uint64       `protobuf:"varint,3,opt,name=Fee,proto3" json:"Fee,omitempty"`
}

func (m *PoolReservesKey) Reset()         { *m = PoolReservesKey{} }
func (m *PoolReservesKey) String() string { return proto.CompactTextString(m) }
func (*PoolReservesKey) ProtoMessage()    {}
func (*PoolReservesKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_d37077b416662cb1, []int{0}
}
func (m *PoolReservesKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolReservesKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolReservesKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolReservesKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolReservesKey.Merge(m, src)
}
func (m *PoolReservesKey) XXX_Size() int {
	return m.Size()
}
func (m *PoolReservesKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolReservesKey.DiscardUnknown(m)
}

var xxx_messageInfo_PoolReservesKey proto.InternalMessageInfo

func (m *PoolReservesKey) GetTradePairID() *TradePairID {
	if m != nil {
		return m.TradePairID
	}
	return nil
}

func (m *PoolReservesKey) GetTickIndexTakerToMaker() int64 {
	if m != nil {
		return m.TickIndexTakerToMaker
	}
	return 0
}

func (m *PoolReservesKey) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

type PoolReserves struct {
	Key                       *PoolReservesKey                                   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ReservesMakerDenom        github_com_cosmos_cosmos_sdk_types.Int             `protobuf:"bytes,2,opt,name=reservesMakerDenom,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reservesMakerDenom" yaml:"reservesMakerDenom"`
	PriceTakerToMaker         github_com_duality_labs_duality_utils_math.PrecDec `protobuf:"bytes,3,opt,name=priceTakerToMaker,proto3,customtype=github.com/QuadrateOrg/core/x/dex/utils/math.PrecDec" json:"priceTakerToMaker" yaml:"priceTakerToMaker"`
	PriceOppositeTakerToMaker github_com_duality_labs_duality_utils_math.PrecDec `protobuf:"bytes,4,opt,name=priceOppositeTakerToMaker,proto3,customtype=github.com/QuadrateOrg/core/x/dex/utils/math.PrecDec" json:"priceOppositeTakerToMaker" yaml:"priceOppositeTakerToMaker"`
}

func (m *PoolReserves) Reset()         { *m = PoolReserves{} }
func (m *PoolReserves) String() string { return proto.CompactTextString(m) }
func (*PoolReserves) ProtoMessage()    {}
func (*PoolReserves) Descriptor() ([]byte, []int) {
	return fileDescriptor_d37077b416662cb1, []int{1}
}
func (m *PoolReserves) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolReserves) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolReserves.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolReserves) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolReserves.Merge(m, src)
}
func (m *PoolReserves) XXX_Size() int {
	return m.Size()
}
func (m *PoolReserves) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolReserves.DiscardUnknown(m)
}

var xxx_messageInfo_PoolReserves proto.InternalMessageInfo

func (m *PoolReserves) GetKey() *PoolReservesKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*PoolReservesKey)(nil), "duality.dex.PoolReservesKey")
	proto.RegisterType((*PoolReserves)(nil), "duality.dex.PoolReserves")
}

func init() { proto.RegisterFile("duality/dex/pool_reserves.proto", fileDescriptor_d37077b416662cb1) }

var fileDescriptor_d37077b416662cb1 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x8a, 0xd3, 0x40,
	0x18, 0xc7, 0x3b, 0xa6, 0x08, 0x4e, 0x05, 0x75, 0x50, 0xc8, 0x2e, 0x92, 0x94, 0x1c, 0xa4, 0x20,
	0x3b, 0x81, 0xd5, 0xd3, 0x1e, 0x97, 0x2a, 0x14, 0x11, 0xcb, 0xd0, 0x93, 0x97, 0x32, 0x4d, 0x3e,
	0xba, 0x43, 0x92, 0x4e, 0x98, 0x99, 0x4a, 0x83, 0x2f, 0xa1, 0x07, 0x0f, 0xbe, 0x82, 0x6f, 0xe0,
	0x1b, 0xec, 0x71, 0x8f, 0xe2, 0x21, 0x48, 0x7b, 0xdb, 0x63, 0x9f, 0x40, 0x32, 0xcd, 0x62, 0x6a,
	0xb3, 0x78, 0xf0, 0x34, 0x93, 0xf9, 0x7e, 0xf3, 0x9f, 0x5f, 0x66, 0xf8, 0xb0, 0x1f, 0x2f, 0x79,
	0x2a, 0x4c, 0x11, 0xc6, 0xb0, 0x0a, 0x73, 0x29, 0xd3, 0xa9, 0x02, 0x0d, 0xea, 0x03, 0x68, 0x9a,
	0x2b, 0x69, 0x24, 0xe9, 0xd5, 0x00, 0x8d, 0x61, 0x75, 0xfc, 0x78, 0x2e, 0xe7, 0xd2, 0xae, 0x87,
	0xd5, 0x6c, 0x87, 0x1c, 0xef, 0x65, 0x18, 0xc5, 0x63, 0x98, 0xe6, 0x5c, 0xa8, 0xa9, 0x88, 0x77,
	0x40, 0xf0, 0x05, 0xe1, 0x07, 0x63, 0x29, 0x53, 0x56, 0x47, 0xbf, 0x81, 0x82, 0x9c, 0xe1, 0x9e,
	0x45, 0xc7, 0x5c, 0xa8, 0xd1, 0xd0, 0x45, 0x7d, 0x34, 0xe8, 0x9d, 0xba, 0xb4, 0x71, 0x1a, 0x9d,
	0xfc, 0xa9, 0xb3, 0x26, 0x4c, 0x5e, 0xe2, 0x27, 0x13, 0x11, 0x25, 0xa3, 0x45, 0x0c, 0xab, 0x09,
	0x4f, 0x40, 0x4d, 0xe4, 0xdb, 0x6a, 0x70, 0xef, 0xf4, 0xd1, 0xc0, 0x61, 0xed, 0x45, 0xf2, 0x10,
	0x3b, 0xaf, 0x01, 0x5c, 0xa7, 0x8f, 0x06, 0x5d, 0x56, 0x4d, 0x83, 0x6f, 0x5d, 0x7c, 0xbf, 0xe9,
	0x45, 0x28, 0x76, 0x12, 0x28, 0x6a, 0x99, 0xa7, 0x7b, 0x32, 0x7f, 0xf9, 0xb3, 0x0a, 0x24, 0x9f,
	0x11, 0x26, 0x37, 0xf7, 0x65, 0x0f, 0x19, 0xc2, 0x42, 0x66, 0x56, 0xe3, 0xde, 0x39, 0xbf, 0x2c,
	0xfd, 0xce, 0xcf, 0xd2, 0x7f, 0x36, 0x17, 0xe6, 0x62, 0x39, 0xa3, 0x91, 0xcc, 0xc2, 0x48, 0xea,
	0x4c, 0xea, 0x7a, 0x38, 0xd1, 0x71, 0x12, 0x9a, 0x22, 0x07, 0x4d, 0x47, 0x0b, 0x73, 0x5d, 0xfa,
	0x2d, 0x59, 0xdb, 0xd2, 0x3f, 0x2a, 0x78, 0x96, 0x9e, 0x05, 0x87, 0xb5, 0x80, 0xb5, 0x6c, 0x20,
	0x5f, 0x11, 0x7e, 0x94, 0x2b, 0x11, 0xc1, 0xde, 0xcd, 0x38, 0x56, 0x29, 0xa9, 0x95, 0x4e, 0x1b,
	0x4a, 0xf5, 0x4f, 0x9e, 0xa4, 0x7c, 0xa6, 0x6f, 0x3e, 0xc2, 0xa5, 0x11, 0xa9, 0x0e, 0x33, 0x6e,
	0x2e, 0xe8, 0x58, 0x41, 0x34, 0x84, 0xe8, 0xba, 0xf4, 0x0f, 0x63, 0xb7, 0xa5, 0xef, 0xee, 0xec,
	0x0e, 0x4a, 0x01, 0x3b, 0xc4, 0xc9, 0x77, 0x84, 0x8f, 0xec, 0xea, 0xbb, 0x3c, 0x97, 0x5a, 0x98,
	0x7d, 0xc7, 0xae, 0x75, 0xfc, 0xf8, 0x5f, 0x8e, 0xb7, 0xc7, 0x6f, 0x4b, 0xbf, 0xdf, 0x70, 0x6d,
	0x43, 0x02, 0x76, 0xfb, 0xf6, 0xf3, 0x57, 0x97, 0x6b, 0x0f, 0x5d, 0xad, 0x3d, 0xf4, 0x6b, 0xed,
	0xa1, 0x4f, 0x1b, 0xaf, 0x73, 0xb5, 0xf1, 0x3a, 0x3f, 0x36, 0x5e, 0xe7, 0xfd, 0xf3, 0x7f, 0x99,
	0xae, 0x76, 0x9d, 0x51, 0xbd, 0xf4, 0xec, 0xae, 0x6d, 0x89, 0x17, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x92, 0xad, 0x36, 0xbf, 0x79, 0x03, 0x00, 0x00,
}

func (m *PoolReservesKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolReservesKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolReservesKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fee != 0 {
		i = encodeVarintPoolReserves(dAtA, i, uint64(m.Fee))
		i--
		dAtA[i] = 0x18
	}
	if m.TickIndexTakerToMaker != 0 {
		i = encodeVarintPoolReserves(dAtA, i, uint64(m.TickIndexTakerToMaker))
		i--
		dAtA[i] = 0x10
	}
	if m.TradePairID != nil {
		{
			size, err := m.TradePairID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPoolReserves(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PoolReserves) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolReserves) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolReserves) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.PriceOppositeTakerToMaker.Size()
		i -= size
		if _, err := m.PriceOppositeTakerToMaker.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPoolReserves(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.PriceTakerToMaker.Size()
		i -= size
		if _, err := m.PriceTakerToMaker.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPoolReserves(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.ReservesMakerDenom.Size()
		i -= size
		if _, err := m.ReservesMakerDenom.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPoolReserves(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Key != nil {
		{
			size, err := m.Key.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPoolReserves(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPoolReserves(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoolReserves(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolReservesKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TradePairID != nil {
		l = m.TradePairID.Size()
		n += 1 + l + sovPoolReserves(uint64(l))
	}
	if m.TickIndexTakerToMaker != 0 {
		n += 1 + sovPoolReserves(uint64(m.TickIndexTakerToMaker))
	}
	if m.Fee != 0 {
		n += 1 + sovPoolReserves(uint64(m.Fee))
	}
	return n
}

func (m *PoolReserves) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != nil {
		l = m.Key.Size()
		n += 1 + l + sovPoolReserves(uint64(l))
	}
	l = m.ReservesMakerDenom.Size()
	n += 1 + l + sovPoolReserves(uint64(l))
	l = m.PriceTakerToMaker.Size()
	n += 1 + l + sovPoolReserves(uint64(l))
	l = m.PriceOppositeTakerToMaker.Size()
	n += 1 + l + sovPoolReserves(uint64(l))
	return n
}

func sovPoolReserves(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoolReserves(x uint64) (n int) {
	return sovPoolReserves(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolReservesKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoolReserves
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
			return fmt.Errorf("proto: PoolReservesKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolReservesKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradePairID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
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
				return ErrInvalidLengthPoolReserves
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPoolReserves
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TradePairID == nil {
				m.TradePairID = &TradePairID{}
			}
			if err := m.TradePairID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickIndexTakerToMaker", wireType)
			}
			m.TickIndexTakerToMaker = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickIndexTakerToMaker |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			m.Fee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPoolReserves(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoolReserves
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
func (m *PoolReserves) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoolReserves
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
			return fmt.Errorf("proto: PoolReserves: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolReserves: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
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
				return ErrInvalidLengthPoolReserves
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPoolReserves
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Key == nil {
				m.Key = &PoolReservesKey{}
			}
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservesMakerDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
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
				return ErrInvalidLengthPoolReserves
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolReserves
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReservesMakerDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceTakerToMaker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
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
				return ErrInvalidLengthPoolReserves
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolReserves
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PriceTakerToMaker.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceOppositeTakerToMaker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
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
				return ErrInvalidLengthPoolReserves
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolReserves
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PriceOppositeTakerToMaker.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPoolReserves(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoolReserves
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
func skipPoolReserves(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoolReserves
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
					return 0, ErrIntOverflowPoolReserves
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
					return 0, ErrIntOverflowPoolReserves
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
				return 0, ErrInvalidLengthPoolReserves
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoolReserves
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoolReserves
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoolReserves        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoolReserves          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoolReserves = fmt.Errorf("proto: unexpected end of group")
)
