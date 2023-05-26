// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerchain/pricefeed/price_feeder_info.proto

package types

import (
	fmt "fmt"
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

type PriceFeederInfo struct {
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Active  bool   `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
}

func (m *PriceFeederInfo) Reset()         { *m = PriceFeederInfo{} }
func (m *PriceFeederInfo) String() string { return proto.CompactTextString(m) }
func (*PriceFeederInfo) ProtoMessage()    {}
func (*PriceFeederInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_838d1869f3e2b656, []int{0}
}
func (m *PriceFeederInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriceFeederInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PriceFeederInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PriceFeederInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceFeederInfo.Merge(m, src)
}
func (m *PriceFeederInfo) XXX_Size() int {
	return m.Size()
}
func (m *PriceFeederInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceFeederInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PriceFeederInfo proto.InternalMessageInfo

func (m *PriceFeederInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PriceFeederInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PriceFeederInfo) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func init() {
	proto.RegisterType((*PriceFeederInfo)(nil), "routerprotocol.routerchain.pricefeed.PriceFeederInfo")
}

func init() {
	proto.RegisterFile("routerchain/pricefeed/price_feeder_info.proto", fileDescriptor_838d1869f3e2b656)
}

var fileDescriptor_838d1869f3e2b656 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2d, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0x4a, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x28, 0xca, 0x4c, 0x4e, 0x4d, 0x4b, 0x4d,
	0x4d, 0x81, 0xb0, 0xe2, 0x41, 0xcc, 0xd4, 0xa2, 0xf8, 0xcc, 0xbc, 0xb4, 0x7c, 0xbd, 0x82, 0xa2,
	0xfc, 0x92, 0x7c, 0x21, 0x15, 0x88, 0x72, 0x30, 0x27, 0x39, 0x3f, 0x47, 0x0f, 0x49, 0xb7, 0x1e,
	0x5c, 0xb7, 0x52, 0x38, 0x17, 0x7f, 0x00, 0x88, 0xe3, 0x06, 0xd6, 0xef, 0x99, 0x97, 0x96, 0x2f,
	0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66,
	0x0b, 0x49, 0x70, 0xb1, 0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x81, 0x85, 0x61,
	0x5c, 0x21, 0x31, 0x2e, 0xb6, 0xc4, 0xe4, 0x92, 0xcc, 0xb2, 0x54, 0x09, 0x66, 0x05, 0x46, 0x0d,
	0x8e, 0x20, 0x28, 0xcf, 0x29, 0xf8, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2,
	0x2c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x21, 0x8e, 0xd2, 0x85,
	0x39, 0x12, 0xc6, 0x87, 0xf8, 0xb1, 0x02, 0xc9, 0x97, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c,
	0x60, 0x85, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0x5c, 0xb8, 0x4e, 0x0b, 0x01, 0x00,
	0x00,
}

func (m *PriceFeederInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriceFeederInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriceFeederInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPriceFeederInfo(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPriceFeederInfo(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPriceFeederInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovPriceFeederInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PriceFeederInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPriceFeederInfo(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPriceFeederInfo(uint64(l))
	}
	if m.Active {
		n += 2
	}
	return n
}

func sovPriceFeederInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPriceFeederInfo(x uint64) (n int) {
	return sovPriceFeederInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PriceFeederInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPriceFeederInfo
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
			return fmt.Errorf("proto: PriceFeederInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriceFeederInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPriceFeederInfo
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
				return ErrInvalidLengthPriceFeederInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPriceFeederInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPriceFeederInfo
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
				return ErrInvalidLengthPriceFeederInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPriceFeederInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPriceFeederInfo
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
			m.Active = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPriceFeederInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPriceFeederInfo
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
func skipPriceFeederInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPriceFeederInfo
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
					return 0, ErrIntOverflowPriceFeederInfo
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
					return 0, ErrIntOverflowPriceFeederInfo
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
				return 0, ErrInvalidLengthPriceFeederInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPriceFeederInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPriceFeederInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPriceFeederInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPriceFeederInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPriceFeederInfo = fmt.Errorf("proto: unexpected end of group")
)
