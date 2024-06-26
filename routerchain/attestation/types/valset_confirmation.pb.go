// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerprotocol/routerchain/attestation/valset_confirmation.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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
	ValsetNonce  uint64 `protobuf:"varint,1,opt,name=valsetNonce,proto3" json:"valsetNonce,omitempty"`
	EthAddress   string `protobuf:"bytes,2,opt,name=ethAddress,proto3" json:"ethAddress,omitempty"`
	Signature    string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Orchestrator string `protobuf:"bytes,4,opt,name=orchestrator,proto3" json:"orchestrator,omitempty"`
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
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x63, 0xa8, 0x90, 0x6a, 0x98, 0x3c, 0x65, 0x40, 0x56, 0xd4, 0x01, 0x75, 0x69, 0x32,
	0x30, 0xb2, 0xf0, 0x67, 0x67, 0xa8, 0x04, 0x03, 0x0b, 0x72, 0xdd, 0xa3, 0xb1, 0xd4, 0xfa, 0xaa,
	0xf3, 0x05, 0xc1, 0x5b, 0x30, 0xf2, 0x48, 0x8c, 0x19, 0x19, 0x51, 0xf2, 0x22, 0x48, 0x0e, 0x12,
	0x4e, 0x47, 0x7f, 0xfe, 0xdd, 0x37, 0x7c, 0xf2, 0x9a, 0xb0, 0x61, 0xa0, 0x3d, 0x21, 0xa3, 0xc5,
	0x6d, 0x35, 0x3c, 0x6d, 0x6d, 0x9c, 0xaf, 0x0c, 0x33, 0x04, 0x36, 0xec, 0xd0, 0x57, 0xaf, 0x66,
	0x1b, 0x80, 0x9f, 0x2d, 0xfa, 0x17, 0x47, 0xbb, 0xc8, 0xca, 0x78, 0xa4, 0x2e, 0xc6, 0x86, 0x32,
	0x31, 0x94, 0x89, 0x61, 0xf6, 0x29, 0xa4, 0x7a, 0x8c, 0x96, 0xbb, 0x44, 0xa2, 0x0a, 0x79, 0x3a,
	0xb8, 0xef, 0xd1, 0x5b, 0xc8, 0x45, 0x21, 0xe6, 0x93, 0x65, 0x8a, 0x94, 0x96, 0x12, 0xb8, 0xbe,
	0x59, 0xaf, 0x09, 0x42, 0xc8, 0x8f, 0x0a, 0x31, 0x9f, 0x2e, 0x13, 0xa2, 0xce, 0xe5, 0x34, 0xb8,
	0x8d, 0x37, 0xdc, 0x10, 0xe4, 0xc7, 0xf1, 0xfb, 0x1f, 0xa8, 0x99, 0x3c, 0x43, 0xb2, 0x35, 0x04,
	0x26, 0xc3, 0x48, 0xf9, 0x24, 0x0e, 0x46, 0xec, 0xf6, 0xe1, 0xab, 0xd3, 0xa2, 0xed, 0xb4, 0xf8,
	0xe9, 0xb4, 0xf8, 0xe8, 0x75, 0xd6, 0xf6, 0x3a, 0xfb, 0xee, 0x75, 0xf6, 0x74, 0xb5, 0x71, 0x5c,
	0x37, 0xab, 0xd2, 0xe2, 0xee, 0x2f, 0xcd, 0xe2, 0x20, 0xd5, 0x62, 0x68, 0xf5, 0x36, 0xaa, 0xc5,
	0xef, 0x7b, 0x08, 0xab, 0x93, 0x38, 0xbd, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x77, 0x4a,
	0xac, 0x64, 0x01, 0x00, 0x00,
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
		dAtA[i] = 0x22
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintValsetConfirmation(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EthAddress) > 0 {
		i -= len(m.EthAddress)
		copy(dAtA[i:], m.EthAddress)
		i = encodeVarintValsetConfirmation(dAtA, i, uint64(len(m.EthAddress)))
		i--
		dAtA[i] = 0x12
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
		case 3:
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
		case 4:
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
