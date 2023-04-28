// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/token_price.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type TokenPrice struct {
	Symbol     string                                 `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	TokenPrice github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=tokenPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tokenPrice"`
	Timestamp  uint64                                 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *TokenPrice) Reset()         { *m = TokenPrice{} }
func (m *TokenPrice) String() string { return proto.CompactTextString(m) }
func (*TokenPrice) ProtoMessage()    {}
func (*TokenPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_77cf10cdbd0dd626, []int{0}
}
func (m *TokenPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenPrice.Merge(m, src)
}
func (m *TokenPrice) XXX_Size() int {
	return m.Size()
}
func (m *TokenPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenPrice.DiscardUnknown(m)
}

var xxx_messageInfo_TokenPrice proto.InternalMessageInfo

func (m *TokenPrice) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *TokenPrice) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*TokenPrice)(nil), "routerprotocol.routerchain.oracle.TokenPrice")
}

func init() { proto.RegisterFile("oracle/token_price.proto", fileDescriptor_77cf10cdbd0dd626) }

var fileDescriptor_77cf10cdbd0dd626 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0x2b, 0x31,
	0x14, 0x86, 0x27, 0xf7, 0x4a, 0xa1, 0x59, 0x0e, 0x22, 0x43, 0x91, 0xb4, 0xba, 0x90, 0x6e, 0x9a,
	0x50, 0x04, 0x1f, 0xa0, 0xb8, 0x96, 0x32, 0xb8, 0x72, 0x23, 0x49, 0x0c, 0xd3, 0xd0, 0xc9, 0x9c,
	0x21, 0x49, 0xc5, 0xbe, 0x85, 0xf8, 0x54, 0x5d, 0x76, 0x29, 0x2e, 0x8a, 0xcc, 0xbc, 0x88, 0x64,
	0xd2, 0xd1, 0xae, 0x92, 0x3f, 0xe7, 0x3f, 0xf9, 0xe0, 0xc3, 0x19, 0x58, 0x2e, 0x4b, 0xc5, 0x3c,
	0xac, 0x55, 0xf5, 0x5c, 0x5b, 0x2d, 0x15, 0xad, 0x2d, 0x78, 0x48, 0xaf, 0x2c, 0x6c, 0xbc, 0xb2,
	0x5d, 0x90, 0x50, 0xd2, 0x18, 0xe5, 0x8a, 0xeb, 0x8a, 0xc6, 0xa5, 0xd1, 0x79, 0x01, 0x05, 0x74,
	0x05, 0x16, 0x6e, 0x71, 0x71, 0x44, 0x24, 0x38, 0x03, 0x8e, 0x09, 0xee, 0x14, 0x7b, 0x9d, 0x0b,
	0xe5, 0xf9, 0x9c, 0x49, 0xd0, 0x55, 0x9c, 0x5f, 0x7f, 0x20, 0x8c, 0x1f, 0x03, 0x6e, 0x19, 0x68,
	0xe9, 0x05, 0x1e, 0xb8, 0xad, 0x11, 0x50, 0x66, 0x68, 0x82, 0xa6, 0xc3, 0xfc, 0x98, 0xd2, 0x07,
	0x8c, 0xfd, 0x6f, 0x2b, 0xfb, 0x17, 0x66, 0x0b, 0xba, 0x3b, 0x8c, 0x93, 0xaf, 0xc3, 0xf8, 0xa6,
	0xd0, 0x7e, 0xb5, 0x11, 0x54, 0x82, 0x61, 0x47, 0x5a, 0x3c, 0x66, 0xee, 0x65, 0xcd, 0xfc, 0xb6,
	0x56, 0x8e, 0xde, 0x2b, 0x99, 0x9f, 0xfc, 0x90, 0x5e, 0xe2, 0xa1, 0xd7, 0x46, 0x39, 0xcf, 0x4d,
	0x9d, 0xfd, 0x9f, 0xa0, 0xe9, 0x59, 0xfe, 0xf7, 0xb0, 0x58, 0xee, 0x1a, 0x82, 0xf6, 0x0d, 0x41,
	0xdf, 0x0d, 0x41, 0xef, 0x2d, 0x49, 0xf6, 0x2d, 0x49, 0x3e, 0x5b, 0x92, 0x3c, 0xdd, 0x9d, 0xb0,
	0xa2, 0x83, 0x59, 0xef, 0xa4, 0xcf, 0x9d, 0x14, 0xf6, 0xc6, 0x7a, 0x97, 0x81, 0x2f, 0x06, 0x5d,
	0xeb, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x22, 0x76, 0x44, 0x71, 0x62, 0x01, 0x00, 0x00,
}

func (m *TokenPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		i = encodeVarintTokenPrice(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.TokenPrice.Size()
		i -= size
		if _, err := m.TokenPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTokenPrice(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTokenPrice(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTokenPrice(dAtA []byte, offset int, v uint64) int {
	offset -= sovTokenPrice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTokenPrice(uint64(l))
	}
	l = m.TokenPrice.Size()
	n += 1 + l + sovTokenPrice(uint64(l))
	if m.Timestamp != 0 {
		n += 1 + sovTokenPrice(uint64(m.Timestamp))
	}
	return n
}

func sovTokenPrice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTokenPrice(x uint64) (n int) {
	return sovTokenPrice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenPrice
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
			return fmt.Errorf("proto: TokenPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenPrice
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
				return ErrInvalidLengthTokenPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenPrice
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
				return ErrInvalidLengthTokenPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTokenPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenPrice
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
func skipTokenPrice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokenPrice
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
					return 0, ErrIntOverflowTokenPrice
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
					return 0, ErrIntOverflowTokenPrice
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
				return 0, ErrInvalidLengthTokenPrice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTokenPrice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTokenPrice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTokenPrice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokenPrice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTokenPrice = fmt.Errorf("proto: unexpected end of group")
)
