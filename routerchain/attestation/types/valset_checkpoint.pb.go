// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerprotocol/routerchain/attestation/valset_checkpoint.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/router-protocol/sdk-go/routerchain/multichain/types"
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

type ValsetCheckpoint struct {
	DestChainType types.ChainType `protobuf:"varint,1,opt,name=dest_chain_type,json=destChainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"dest_chain_type,omitempty"`
	Signature     string          `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ValsetCheckpoint) Reset()         { *m = ValsetCheckpoint{} }
func (m *ValsetCheckpoint) String() string { return proto.CompactTextString(m) }
func (*ValsetCheckpoint) ProtoMessage()    {}
func (*ValsetCheckpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d15c94797ad7b44, []int{0}
}
func (m *ValsetCheckpoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValsetCheckpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValsetCheckpoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValsetCheckpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValsetCheckpoint.Merge(m, src)
}
func (m *ValsetCheckpoint) XXX_Size() int {
	return m.Size()
}
func (m *ValsetCheckpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_ValsetCheckpoint.DiscardUnknown(m)
}

var xxx_messageInfo_ValsetCheckpoint proto.InternalMessageInfo

func (m *ValsetCheckpoint) GetDestChainType() types.ChainType {
	if m != nil {
		return m.DestChainType
	}
	return types.CHAIN_TYPE_NONE
}

func (m *ValsetCheckpoint) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*ValsetCheckpoint)(nil), "routerprotocol.routerchain.attestation.ValsetCheckpoint")
}

func init() {
	proto.RegisterFile("routerprotocol/routerchain/attestation/valset_checkpoint.proto", fileDescriptor_0d15c94797ad7b44)
}

var fileDescriptor_0d15c94797ad7b44 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x2b, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0x2a, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x87, 0x70, 0x93, 0x33, 0x12,
	0x33, 0xf3, 0xf4, 0x13, 0x4b, 0x4a, 0x52, 0x8b, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0xf4, 0xcb,
	0x12, 0x73, 0x8a, 0x53, 0x4b, 0xe2, 0x93, 0x33, 0x52, 0x93, 0xb3, 0x0b, 0xf2, 0x33, 0xf3, 0x4a,
	0xf4, 0xc0, 0x5a, 0x84, 0xd4, 0x50, 0xf5, 0xeb, 0x21, 0xe9, 0xd7, 0x43, 0xd2, 0x2f, 0x65, 0x86,
	0xc7, 0x9e, 0xdc, 0xd2, 0x9c, 0x92, 0x4c, 0x08, 0x13, 0x4c, 0xc6, 0x97, 0x54, 0x16, 0xa4, 0x42,
	0xcc, 0x57, 0xea, 0x62, 0xe4, 0x12, 0x08, 0x03, 0xdb, 0xed, 0x0c, 0xb7, 0x5a, 0x28, 0x82, 0x8b,
	0x3f, 0x25, 0xb5, 0x18, 0xe4, 0x1a, 0x98, 0x6a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x03,
	0x3d, 0x3c, 0xce, 0x41, 0x58, 0xa3, 0xe7, 0x0c, 0x22, 0x43, 0x2a, 0x0b, 0x52, 0x83, 0x78, 0x41,
	0x06, 0xc1, 0xb9, 0x42, 0x32, 0x5c, 0x9c, 0xc5, 0x99, 0xe9, 0x79, 0x89, 0x25, 0xa5, 0x45, 0xa9,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x08, 0x01, 0xa7, 0xd0, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18,
	0x6e, 0x3c, 0x96, 0x63, 0x88, 0xb2, 0x4e, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0x85, 0x7a, 0x4d, 0x17, 0xcd, 0xab, 0xba, 0x10, 0x0f, 0x56, 0xa0, 0x84, 0x2a, 0xc8, 0xe9, 0xc5,
	0x49, 0x6c, 0x60, 0xa5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x2a, 0xfa, 0xc5, 0x8c,
	0x01, 0x00, 0x00,
}

func (m *ValsetCheckpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValsetCheckpoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValsetCheckpoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintValsetCheckpoint(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if m.DestChainType != 0 {
		i = encodeVarintValsetCheckpoint(dAtA, i, uint64(m.DestChainType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintValsetCheckpoint(dAtA []byte, offset int, v uint64) int {
	offset -= sovValsetCheckpoint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValsetCheckpoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DestChainType != 0 {
		n += 1 + sovValsetCheckpoint(uint64(m.DestChainType))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovValsetCheckpoint(uint64(l))
	}
	return n
}

func sovValsetCheckpoint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValsetCheckpoint(x uint64) (n int) {
	return sovValsetCheckpoint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValsetCheckpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValsetCheckpoint
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
			return fmt.Errorf("proto: ValsetCheckpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValsetCheckpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestChainType", wireType)
			}
			m.DestChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValsetCheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DestChainType |= types.ChainType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValsetCheckpoint
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
				return ErrInvalidLengthValsetCheckpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValsetCheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValsetCheckpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValsetCheckpoint
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
func skipValsetCheckpoint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValsetCheckpoint
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
					return 0, ErrIntOverflowValsetCheckpoint
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
					return 0, ErrIntOverflowValsetCheckpoint
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
				return 0, ErrInvalidLengthValsetCheckpoint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValsetCheckpoint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValsetCheckpoint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValsetCheckpoint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValsetCheckpoint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValsetCheckpoint = fmt.Errorf("proto: unexpected end of group")
)
