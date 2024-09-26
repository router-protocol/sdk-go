// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerprotocol/routerchain/attestation/valset_confirmation.proto

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

type ValsetConfirmation struct {
	ValsetNonce   uint64          `protobuf:"varint,1,opt,name=valsetNonce,proto3" json:"valsetNonce,omitempty"`
	DestChainType types.ChainType `protobuf:"varint,2,opt,name=dest_chain_type,json=destChainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"dest_chain_type,omitempty"`
	EthAddress    string          `protobuf:"bytes,3,opt,name=ethAddress,proto3" json:"ethAddress,omitempty"`
	Signature     string          `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	Orchestrator  string          `protobuf:"bytes,5,opt,name=orchestrator,proto3" json:"orchestrator,omitempty"`
}

func (m *ValsetConfirmation) Reset()         { *m = ValsetConfirmation{} }
func (m *ValsetConfirmation) String() string { return proto.CompactTextString(m) }
func (*ValsetConfirmation) ProtoMessage()    {}
func (*ValsetConfirmation) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d6ec7171fa43206, []int{0}
}
func (m *ValsetConfirmation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValsetConfirmation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValsetConfirmation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValsetConfirmation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValsetConfirmation.Merge(m, src)
}
func (m *ValsetConfirmation) XXX_Size() int {
	return m.Size()
}
func (m *ValsetConfirmation) XXX_DiscardUnknown() {
	xxx_messageInfo_ValsetConfirmation.DiscardUnknown(m)
}

var xxx_messageInfo_ValsetConfirmation proto.InternalMessageInfo

func (m *ValsetConfirmation) GetValsetNonce() uint64 {
	if m != nil {
		return m.ValsetNonce
	}
	return 0
}

func (m *ValsetConfirmation) GetDestChainType() types.ChainType {
	if m != nil {
		return m.DestChainType
	}
	return types.CHAIN_TYPE_NONE
}

func (m *ValsetConfirmation) GetEthAddress() string {
	if m != nil {
		return m.EthAddress
	}
	return ""
}

func (m *ValsetConfirmation) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *ValsetConfirmation) GetOrchestrator() string {
	if m != nil {
		return m.Orchestrator
	}
	return ""
}

func init() {
	proto.RegisterType((*ValsetConfirmation)(nil), "routerprotocol.routerchain.attestation.ValsetConfirmation")
}

func init() {
	proto.RegisterFile("routerprotocol/routerchain/attestation/valset_confirmation.proto", fileDescriptor_6d6ec7171fa43206)
}

var fileDescriptor_6d6ec7171fa43206 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x6b, 0x28, 0x48, 0x35, 0xff, 0x24, 0x4f, 0x11, 0x42, 0x56, 0xd4, 0x01, 0x75, 0xa9,
	0x8b, 0x40, 0x62, 0x61, 0x01, 0xba, 0x33, 0x44, 0x80, 0x10, 0x4b, 0xe4, 0x3a, 0xa6, 0xb1, 0x94,
	0xc4, 0x91, 0x7d, 0x41, 0xf4, 0x2d, 0x78, 0x2c, 0xc6, 0x8e, 0x8c, 0x28, 0x79, 0x03, 0x9e, 0x00,
	0xc5, 0x41, 0x24, 0x61, 0xe8, 0x72, 0xca, 0x7d, 0x39, 0xff, 0xee, 0xf3, 0x67, 0x7c, 0x6d, 0x74,
	0x01, 0xd2, 0xe4, 0x46, 0x83, 0x16, 0x3a, 0x99, 0x35, 0xad, 0x88, 0xb9, 0xca, 0x66, 0x1c, 0x40,
	0x5a, 0xe0, 0xa0, 0x74, 0x36, 0x7b, 0xe5, 0x89, 0x95, 0x10, 0x0a, 0x9d, 0xbd, 0x28, 0x93, 0x3a,
	0x8d, 0xb9, 0x43, 0xe4, 0xb4, 0x4f, 0x60, 0x1d, 0x02, 0xeb, 0x10, 0x8e, 0x2f, 0x37, 0x6c, 0x4a,
	0x8b, 0x04, 0x54, 0xf3, 0xe9, 0x6a, 0x08, 0xab, 0x5c, 0x36, 0xfc, 0xf1, 0x37, 0xc2, 0xe4, 0xd1,
	0x6d, 0x9f, 0x77, 0x96, 0x13, 0x1f, 0xef, 0x35, 0x9e, 0xee, 0x74, 0x26, 0xa4, 0x87, 0x7c, 0x34,
	0x19, 0x06, 0x5d, 0x89, 0x3c, 0xe1, 0xa3, 0x48, 0x5a, 0x08, 0x5b, 0xa2, 0xb7, 0xe5, 0xa3, 0xc9,
	0xe1, 0xf9, 0x19, 0xdb, 0x60, 0xb9, 0xb5, 0xc2, 0xe6, 0x75, 0xbd, 0x5f, 0xe5, 0x32, 0x38, 0xa8,
	0x41, 0x7f, 0x2d, 0xa1, 0x18, 0x4b, 0x88, 0x6f, 0xa2, 0xc8, 0x48, 0x6b, 0xbd, 0x6d, 0x1f, 0x4d,
	0x46, 0x41, 0x47, 0x21, 0x27, 0x78, 0x64, 0xd5, 0x32, 0xe3, 0x50, 0x18, 0xe9, 0x0d, 0xdd, 0xef,
	0x56, 0x20, 0x63, 0xbc, 0xaf, 0x8d, 0x88, 0xa5, 0x05, 0xc3, 0x41, 0x1b, 0x6f, 0xc7, 0x0d, 0xf4,
	0xb4, 0xdb, 0x87, 0x8f, 0x92, 0xa2, 0x75, 0x49, 0xd1, 0x57, 0x49, 0xd1, 0x7b, 0x45, 0x07, 0xeb,
	0x8a, 0x0e, 0x3e, 0x2b, 0x3a, 0x78, 0xbe, 0x5a, 0x2a, 0x88, 0x8b, 0x05, 0x13, 0x3a, 0xfd, 0x8d,
	0x70, 0xfa, 0x2f, 0xd2, 0x69, 0x13, 0xe4, 0x5b, 0xef, 0xfd, 0xea, 0xeb, 0xdb, 0xc5, 0xae, 0x1b,
	0xbd, 0xf8, 0x09, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x3a, 0x21, 0x5b, 0xf6, 0x01, 0x00, 0x00,
}

func (m *ValsetConfirmation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValsetConfirmation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValsetConfirmation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Orchestrator) > 0 {
		i -= len(m.Orchestrator)
		copy(dAtA[i:], m.Orchestrator)
		i = encodeVarintValsetConfirmation(dAtA, i, uint64(len(m.Orchestrator)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintValsetConfirmation(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.EthAddress) > 0 {
		i -= len(m.EthAddress)
		copy(dAtA[i:], m.EthAddress)
		i = encodeVarintValsetConfirmation(dAtA, i, uint64(len(m.EthAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.DestChainType != 0 {
		i = encodeVarintValsetConfirmation(dAtA, i, uint64(m.DestChainType))
		i--
		dAtA[i] = 0x10
	}
	if m.ValsetNonce != 0 {
		i = encodeVarintValsetConfirmation(dAtA, i, uint64(m.ValsetNonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintValsetConfirmation(dAtA []byte, offset int, v uint64) int {
	offset -= sovValsetConfirmation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValsetConfirmation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValsetNonce != 0 {
		n += 1 + sovValsetConfirmation(uint64(m.ValsetNonce))
	}
	if m.DestChainType != 0 {
		n += 1 + sovValsetConfirmation(uint64(m.DestChainType))
	}
	l = len(m.EthAddress)
	if l > 0 {
		n += 1 + l + sovValsetConfirmation(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovValsetConfirmation(uint64(l))
	}
	l = len(m.Orchestrator)
	if l > 0 {
		n += 1 + l + sovValsetConfirmation(uint64(l))
	}
	return n
}

func sovValsetConfirmation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValsetConfirmation(x uint64) (n int) {
	return sovValsetConfirmation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValsetConfirmation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValsetConfirmation
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
			return fmt.Errorf("proto: ValsetConfirmation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValsetConfirmation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValsetNonce", wireType)
			}
			m.ValsetNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValsetConfirmation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValsetNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestChainType", wireType)
			}
			m.DestChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValsetConfirmation
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValsetConfirmation
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
				return ErrInvalidLengthValsetConfirmation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValsetConfirmation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValsetConfirmation
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
				return ErrInvalidLengthValsetConfirmation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValsetConfirmation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Orchestrator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValsetConfirmation
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
				return ErrInvalidLengthValsetConfirmation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValsetConfirmation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Orchestrator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValsetConfirmation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValsetConfirmation
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
func skipValsetConfirmation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValsetConfirmation
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
					return 0, ErrIntOverflowValsetConfirmation
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
					return 0, ErrIntOverflowValsetConfirmation
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
				return 0, ErrInvalidLengthValsetConfirmation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValsetConfirmation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValsetConfirmation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValsetConfirmation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValsetConfirmation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValsetConfirmation = fmt.Errorf("proto: unexpected end of group")
)
