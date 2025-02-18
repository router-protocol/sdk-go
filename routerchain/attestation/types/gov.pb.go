// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerprotocol/routerchain/attestation/gov.proto

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

type ResetAttestationStatesProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ChainId     string `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Contract    string `protobuf:"bytes,4,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (m *ResetAttestationStatesProposal) Reset()         { *m = ResetAttestationStatesProposal{} }
func (m *ResetAttestationStatesProposal) String() string { return proto.CompactTextString(m) }
func (*ResetAttestationStatesProposal) ProtoMessage()    {}
func (*ResetAttestationStatesProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e686becc7c9f249, []int{0}
}
func (m *ResetAttestationStatesProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResetAttestationStatesProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResetAttestationStatesProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResetAttestationStatesProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetAttestationStatesProposal.Merge(m, src)
}
func (m *ResetAttestationStatesProposal) XXX_Size() int {
	return m.Size()
}
func (m *ResetAttestationStatesProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetAttestationStatesProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ResetAttestationStatesProposal proto.InternalMessageInfo

func (m *ResetAttestationStatesProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ResetAttestationStatesProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ResetAttestationStatesProposal) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *ResetAttestationStatesProposal) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func init() {
	proto.RegisterType((*ResetAttestationStatesProposal)(nil), "routerprotocol.routerchain.attestation.ResetAttestationStatesProposal")
}

func init() {
	proto.RegisterFile("routerprotocol/routerchain/attestation/gov.proto", fileDescriptor_3e686becc7c9f249)
}

var fileDescriptor_3e686becc7c9f249 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x28, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0x2a, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x87, 0x70, 0x93, 0x33, 0x12,
	0x33, 0xf3, 0xf4, 0x13, 0x4b, 0x4a, 0x52, 0x8b, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0xf4, 0xd3,
	0xf3, 0xcb, 0xf4, 0xc0, 0x8a, 0x84, 0xd4, 0x50, 0x75, 0xe8, 0x21, 0xe9, 0xd0, 0x43, 0xd2, 0xa1,
	0xd4, 0xcf, 0xc8, 0x25, 0x17, 0x94, 0x5a, 0x9c, 0x5a, 0xe2, 0x88, 0x10, 0x0c, 0x2e, 0x49, 0x2c,
	0x49, 0x2d, 0x0e, 0x28, 0xca, 0x2f, 0xc8, 0x2f, 0x4e, 0xcc, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0xc9,
	0x2c, 0xc9, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0x14, 0xb8, 0xb8,
	0x53, 0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x40, 0x5a, 0x24, 0x98, 0xc0, 0x72, 0xc8, 0x42, 0x42,
	0x92, 0x5c, 0x1c, 0x60, 0xfb, 0xe2, 0x33, 0x53, 0x24, 0x98, 0xc1, 0xd2, 0xec, 0x60, 0xbe, 0x67,
	0x8a, 0x90, 0x14, 0x17, 0x47, 0x72, 0x7e, 0x5e, 0x49, 0x51, 0x62, 0x72, 0x89, 0x04, 0x0b, 0x58,
	0x0a, 0xce, 0x77, 0x0a, 0x3d, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xeb,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x68, 0x08, 0xe8, 0xa2, 0x85, 0x88,
	0x2e, 0x24, 0x48, 0x2a, 0x50, 0x02, 0xa5, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xac, 0xd4,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x72, 0x5c, 0x88, 0x4b, 0x01, 0x00, 0x00,
}

func (m *ResetAttestationStatesProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResetAttestationStatesProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResetAttestationStatesProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintGov(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ResetAttestationStatesProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResetAttestationStatesProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: ResetAttestationStatesProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResetAttestationStatesProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
