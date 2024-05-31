// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerprotocol/routerchain/pricefeed/update_symbol_request_proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// UpdateSymbolRequestProposal is a message that represents a proposal to update
// a symbol request. It contains a title, description, and a list of symbol
// requests to be updated. Once approved, the data will update the SymbolRequest
// in the state.
type UpdateSymbolRequestProposal struct {
	Title          string          `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description    string          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	SymbolRequests []SymbolRequest `protobuf:"bytes,3,rep,name=symbol_requests,json=symbolRequests,proto3" json:"symbol_requests"`
}

func (m *UpdateSymbolRequestProposal) Reset()         { *m = UpdateSymbolRequestProposal{} }
func (m *UpdateSymbolRequestProposal) String() string { return proto.CompactTextString(m) }
func (*UpdateSymbolRequestProposal) ProtoMessage()    {}
func (*UpdateSymbolRequestProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_de8d08612a7f0923, []int{0}
}
func (m *UpdateSymbolRequestProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateSymbolRequestProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateSymbolRequestProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateSymbolRequestProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSymbolRequestProposal.Merge(m, src)
}
func (m *UpdateSymbolRequestProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateSymbolRequestProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSymbolRequestProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSymbolRequestProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpdateSymbolRequestProposal)(nil), "routerprotocol.routerchain.pricefeed.UpdateSymbolRequestProposal")
}

func init() {
	proto.RegisterFile("routerprotocol/routerchain/pricefeed/update_symbol_request_proposal.proto", fileDescriptor_de8d08612a7f0923)
}

var fileDescriptor_de8d08612a7f0923 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x2c, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0x2a, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x87, 0x70, 0x93, 0x33, 0x12,
	0x33, 0xf3, 0xf4, 0x0b, 0x8a, 0x32, 0x93, 0x53, 0xd3, 0x52, 0x53, 0x53, 0xf4, 0x4b, 0x0b, 0x52,
	0x12, 0x4b, 0x52, 0xe3, 0x8b, 0x2b, 0x73, 0x93, 0xf2, 0x73, 0xe2, 0x8b, 0x52, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0xe2, 0x0b, 0x8a, 0xf2, 0x0b, 0xf2, 0x8b, 0x13, 0x73, 0xf4, 0xc0, 0xba, 0x85, 0x54,
	0x50, 0x8d, 0xd2, 0x43, 0x32, 0x4a, 0x0f, 0x6e, 0x94, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58,
	0x8d, 0x3e, 0x88, 0x05, 0xd1, 0x2b, 0x65, 0x49, 0x94, 0x33, 0x50, 0xed, 0x87, 0x68, 0x55, 0xda,
	0xcf, 0xc8, 0x25, 0x1d, 0x0a, 0x76, 0x5f, 0x30, 0x58, 0x3a, 0x08, 0x22, 0x1b, 0x00, 0x75, 0x9c,
	0x90, 0x08, 0x17, 0x6b, 0x49, 0x66, 0x49, 0x4e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x84, 0x23, 0xa4, 0xc0, 0xc5, 0x9d, 0x92, 0x5a, 0x9c, 0x5c, 0x94, 0x59, 0x50, 0x92, 0x99, 0x9f,
	0x27, 0xc1, 0x04, 0x96, 0x43, 0x16, 0x12, 0x4a, 0xe2, 0xe2, 0x47, 0xb5, 0xaf, 0x58, 0x82, 0x59,
	0x81, 0x59, 0x83, 0xdb, 0xc8, 0x58, 0x8f, 0x18, 0x8f, 0xea, 0xa1, 0xb8, 0xc6, 0x89, 0xe5, 0xc4,
	0x3d, 0x79, 0x86, 0x20, 0xbe, 0x62, 0x64, 0xc1, 0x62, 0x2b, 0x96, 0x8e, 0x05, 0xf2, 0x0c, 0x4e,
	0xc1, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x99, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0x0b, 0x0d, 0x12, 0x5d, 0xb4, 0x20, 0xd2, 0x85, 0x84, 0x51,
	0x05, 0x52, 0x28, 0x95, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x15, 0x1a, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xe6, 0x76, 0x8d, 0xab, 0xe1, 0x01, 0x00, 0x00,
}

func (m *UpdateSymbolRequestProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateSymbolRequestProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateSymbolRequestProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SymbolRequests) > 0 {
		for iNdEx := len(m.SymbolRequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SymbolRequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUpdateSymbolRequestProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintUpdateSymbolRequestProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintUpdateSymbolRequestProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUpdateSymbolRequestProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovUpdateSymbolRequestProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UpdateSymbolRequestProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovUpdateSymbolRequestProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovUpdateSymbolRequestProposal(uint64(l))
	}
	if len(m.SymbolRequests) > 0 {
		for _, e := range m.SymbolRequests {
			l = e.Size()
			n += 1 + l + sovUpdateSymbolRequestProposal(uint64(l))
		}
	}
	return n
}

func sovUpdateSymbolRequestProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUpdateSymbolRequestProposal(x uint64) (n int) {
	return sovUpdateSymbolRequestProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UpdateSymbolRequestProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpdateSymbolRequestProposal
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
			return fmt.Errorf("proto: UpdateSymbolRequestProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateSymbolRequestProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdateSymbolRequestProposal
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
				return ErrInvalidLengthUpdateSymbolRequestProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdateSymbolRequestProposal
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
					return ErrIntOverflowUpdateSymbolRequestProposal
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
				return ErrInvalidLengthUpdateSymbolRequestProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdateSymbolRequestProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SymbolRequests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdateSymbolRequestProposal
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
				return ErrInvalidLengthUpdateSymbolRequestProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUpdateSymbolRequestProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SymbolRequests = append(m.SymbolRequests, SymbolRequest{})
			if err := m.SymbolRequests[len(m.SymbolRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpdateSymbolRequestProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpdateSymbolRequestProposal
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
func skipUpdateSymbolRequestProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUpdateSymbolRequestProposal
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
					return 0, ErrIntOverflowUpdateSymbolRequestProposal
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
					return 0, ErrIntOverflowUpdateSymbolRequestProposal
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
				return 0, ErrInvalidLengthUpdateSymbolRequestProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUpdateSymbolRequestProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUpdateSymbolRequestProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUpdateSymbolRequestProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUpdateSymbolRequestProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUpdateSymbolRequestProposal = fmt.Errorf("proto: unexpected end of group")
)
