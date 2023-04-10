// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerchain/crosschain/contract_call_metadata.proto

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

type ContractCallMetadata struct {
}

func (m *ContractCallMetadata) Reset()         { *m = ContractCallMetadata{} }
func (m *ContractCallMetadata) String() string { return proto.CompactTextString(m) }
func (*ContractCallMetadata) ProtoMessage()    {}
func (*ContractCallMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e8fc4f6c4aa7d12, []int{0}
}
func (m *ContractCallMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractCallMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractCallMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractCallMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractCallMetadata.Merge(m, src)
}
func (m *ContractCallMetadata) XXX_Size() int {
	return m.Size()
}
func (m *ContractCallMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractCallMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ContractCallMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ContractCallMetadata)(nil), "routerprotocol.routerchain.crosschain.ContractCallMetadata")
}

func init() {
	proto.RegisterFile("routerchain/crosschain/contract_call_metadata.proto", fileDescriptor_4e8fc4f6c4aa7d12)
}

var fileDescriptor_4e8fc4f6c4aa7d12 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2e, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0x4a, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x2e, 0xca, 0x2f, 0x2e, 0x86, 0x32, 0xf3,
	0xf3, 0x4a, 0x8a, 0x12, 0x93, 0x4b, 0xe2, 0x93, 0x13, 0x73, 0x72, 0xe2, 0x73, 0x53, 0x4b, 0x12,
	0x53, 0x12, 0x4b, 0x12, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x54, 0x21, 0x9a, 0xc0, 0x9c,
	0xe4, 0xfc, 0x1c, 0x3d, 0x24, 0x33, 0xf4, 0x10, 0x66, 0x28, 0x89, 0x71, 0x89, 0x38, 0x43, 0x8d,
	0x71, 0x4e, 0xcc, 0xc9, 0xf1, 0x85, 0x1a, 0xe2, 0x14, 0x72, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47,
	0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d,
	0xc7, 0x72, 0x0c, 0x51, 0x56, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x10, 0x43, 0x75, 0x61, 0x96, 0xc0, 0xf8, 0x10, 0xe7, 0x55, 0x20, 0xbb, 0xb5, 0xa4, 0xb2, 0x20,
	0xb5, 0x38, 0x89, 0x0d, 0xac, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x20, 0x7d, 0x2f,
	0xd2, 0x00, 0x00, 0x00,
}

func (m *ContractCallMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractCallMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractCallMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintContractCallMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovContractCallMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ContractCallMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovContractCallMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContractCallMetadata(x uint64) (n int) {
	return sovContractCallMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContractCallMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContractCallMetadata
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
			return fmt.Errorf("proto: ContractCallMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractCallMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipContractCallMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContractCallMetadata
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
func skipContractCallMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContractCallMetadata
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
					return 0, ErrIntOverflowContractCallMetadata
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
					return 0, ErrIntOverflowContractCallMetadata
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
				return 0, ErrInvalidLengthContractCallMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContractCallMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContractCallMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContractCallMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContractCallMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContractCallMetadata = fmt.Errorf("proto: unexpected end of group")
)
