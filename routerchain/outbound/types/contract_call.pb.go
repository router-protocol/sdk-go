// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: outbound/contract_call.proto

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

type ContractCall struct {
	DestinationContractAddress []byte `protobuf:"bytes,1,opt,name=destinationContractAddress,proto3" json:"destinationContractAddress,omitempty"`
	Payload                    []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *ContractCall) Reset()         { *m = ContractCall{} }
func (m *ContractCall) String() string { return proto.CompactTextString(m) }
func (*ContractCall) ProtoMessage()    {}
func (*ContractCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a829a097a466a14, []int{0}
}
func (m *ContractCall) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractCall.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractCall.Merge(m, src)
}
func (m *ContractCall) XXX_Size() int {
	return m.Size()
}
func (m *ContractCall) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractCall.DiscardUnknown(m)
}

var xxx_messageInfo_ContractCall proto.InternalMessageInfo

func (m *ContractCall) GetDestinationContractAddress() []byte {
	if m != nil {
		return m.DestinationContractAddress
	}
	return nil
}

func (m *ContractCall) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*ContractCall)(nil), "routerprotocol.routerchain.outbound.ContractCall")
}

func init() { proto.RegisterFile("outbound/contract_call.proto", fileDescriptor_5a829a097a466a14) }

var fileDescriptor_5a829a097a466a14 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0x2f, 0x2d, 0x49,
	0xca, 0x2f, 0xcd, 0x4b, 0xd1, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0x89, 0x4f, 0x4e,
	0xcc, 0xc9, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2e, 0xca, 0x2f, 0x2d, 0x49, 0x2d,
	0x02, 0x73, 0x92, 0xf3, 0x73, 0xf4, 0x20, 0xdc, 0xe4, 0x8c, 0xc4, 0xcc, 0x3c, 0x3d, 0x98, 0x46,
	0xa5, 0x0c, 0x2e, 0x1e, 0x67, 0xa8, 0x5e, 0xe7, 0xc4, 0x9c, 0x1c, 0x21, 0x3b, 0x2e, 0xa9, 0x94,
	0xd4, 0xe2, 0x92, 0xcc, 0xbc, 0xc4, 0x92, 0xcc, 0xfc, 0x3c, 0x98, 0x94, 0x63, 0x4a, 0x4a, 0x51,
	0x6a, 0x71, 0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x1e, 0x15, 0x42, 0x12, 0x5c, 0xec,
	0x05, 0x89, 0x95, 0x39, 0xf9, 0x89, 0x29, 0x12, 0x4c, 0x60, 0xc5, 0x30, 0xae, 0x53, 0xd0, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3,
	0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x59, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0x1c, 0xa9, 0x0b, 0x73, 0x34, 0x8c, 0x0f, 0x76, 0xb5, 0x7e,
	0x85, 0x3e, 0xdc, 0xc3, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x75, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x04, 0x9c, 0x28, 0xf6, 0x09, 0x01, 0x00, 0x00,
}

func (m *ContractCall) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractCall) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractCall) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintContractCall(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DestinationContractAddress) > 0 {
		i -= len(m.DestinationContractAddress)
		copy(dAtA[i:], m.DestinationContractAddress)
		i = encodeVarintContractCall(dAtA, i, uint64(len(m.DestinationContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintContractCall(dAtA []byte, offset int, v uint64) int {
	offset -= sovContractCall(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ContractCall) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DestinationContractAddress)
	if l > 0 {
		n += 1 + l + sovContractCall(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovContractCall(uint64(l))
	}
	return n
}

func sovContractCall(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContractCall(x uint64) (n int) {
	return sovContractCall(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContractCall) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContractCall
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
			return fmt.Errorf("proto: ContractCall: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractCall: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationContractAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthContractCall
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthContractCall
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationContractAddress = append(m.DestinationContractAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.DestinationContractAddress == nil {
				m.DestinationContractAddress = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthContractCall
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthContractCall
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContractCall(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContractCall
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
func skipContractCall(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContractCall
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
					return 0, ErrIntOverflowContractCall
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
					return 0, ErrIntOverflowContractCall
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
				return 0, ErrInvalidLengthContractCall
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContractCall
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContractCall
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContractCall        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContractCall          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContractCall = fmt.Errorf("proto: unexpected end of group")
)