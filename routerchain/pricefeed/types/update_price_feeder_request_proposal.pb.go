// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerprotocol/routerchain/pricefeed/update_price_feeder_request_proposal.proto

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

// UpdatePriceFeederInfoProposal is a message that represents a proposal to update
// a pricefeeder info. It contains a title, description, and pricefeeder info
// requests to be updated. Once approved, the data will update the PriceFeederInfo
// in the state.
type UpdatePriceFeederInfoProposal struct {
	Title           string          `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description     string          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	PriceFeederInfo PriceFeederInfo `protobuf:"bytes,3,opt,name=price_feeder_info,json=priceFeederInfo,proto3" json:"price_feeder_info"`
}

func (m *UpdatePriceFeederInfoProposal) Reset()         { *m = UpdatePriceFeederInfoProposal{} }
func (m *UpdatePriceFeederInfoProposal) String() string { return proto.CompactTextString(m) }
func (*UpdatePriceFeederInfoProposal) ProtoMessage()    {}
func (*UpdatePriceFeederInfoProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b37b845c7ebafac, []int{0}
}
func (m *UpdatePriceFeederInfoProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdatePriceFeederInfoProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdatePriceFeederInfoProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdatePriceFeederInfoProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePriceFeederInfoProposal.Merge(m, src)
}
func (m *UpdatePriceFeederInfoProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdatePriceFeederInfoProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePriceFeederInfoProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePriceFeederInfoProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpdatePriceFeederInfoProposal)(nil), "routerprotocol.routerchain.pricefeed.UpdatePriceFeederInfoProposal")
}

func init() {
	proto.RegisterFile("routerprotocol/routerchain/pricefeed/update_price_feeder_request_proposal.proto", fileDescriptor_3b37b845c7ebafac)
}

var fileDescriptor_3b37b845c7ebafac = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x2f, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0x2a, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x87, 0x70, 0x93, 0x33, 0x12,
	0x33, 0xf3, 0xf4, 0x0b, 0x8a, 0x32, 0x93, 0x53, 0xd3, 0x52, 0x53, 0x53, 0xf4, 0x4b, 0x0b, 0x52,
	0x12, 0x4b, 0x52, 0xe3, 0xc1, 0x02, 0xf1, 0x20, 0x91, 0xd4, 0xa2, 0xf8, 0xa2, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x92, 0xf8, 0x82, 0xa2, 0xfc, 0x82, 0xfc, 0xe2, 0xc4, 0x1c, 0x3d, 0xb0, 0x19, 0x42,
	0x2a, 0xa8, 0x06, 0xea, 0x21, 0x19, 0xa8, 0x07, 0x37, 0x50, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f,
	0xac, 0x46, 0x1f, 0xc4, 0x82, 0xe8, 0x95, 0xb2, 0x21, 0xca, 0x31, 0x28, 0xae, 0xc8, 0xcc, 0x4b,
	0x83, 0xea, 0x56, 0x3a, 0xc6, 0xc8, 0x25, 0x1b, 0x0a, 0x76, 0x68, 0x00, 0x48, 0x85, 0x1b, 0x58,
	0x81, 0x67, 0x5e, 0x5a, 0x7e, 0x00, 0xd4, 0x85, 0x42, 0x22, 0x5c, 0xac, 0x25, 0x99, 0x25, 0x39,
	0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x02, 0x17, 0x77, 0x4a, 0x6a,
	0x71, 0x72, 0x51, 0x66, 0x41, 0x49, 0x66, 0x7e, 0x9e, 0x04, 0x13, 0x58, 0x0e, 0x59, 0x48, 0x28,
	0x9d, 0x4b, 0x10, 0xc3, 0x52, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x53, 0x3d, 0x62, 0xfc,
	0xab, 0x87, 0xe6, 0x22, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0xf8, 0x0b, 0x50, 0x85, 0xad,
	0x58, 0x3a, 0x16, 0xc8, 0x33, 0x38, 0x05, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3,
	0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c,
	0x43, 0x94, 0x65, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0x2e, 0x34, 0x70, 0x74,
	0xd1, 0x02, 0x4b, 0x17, 0x12, 0x5a, 0x15, 0x48, 0xe1, 0x55, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4,
	0x06, 0x56, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x04, 0xb0, 0xc3, 0xf1, 0x01, 0x00,
	0x00,
}

func (m *UpdatePriceFeederInfoProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdatePriceFeederInfoProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdatePriceFeederInfoProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PriceFeederInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintUpdatePriceFeederRequestProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintUpdatePriceFeederRequestProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintUpdatePriceFeederRequestProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUpdatePriceFeederRequestProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovUpdatePriceFeederRequestProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UpdatePriceFeederInfoProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovUpdatePriceFeederRequestProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovUpdatePriceFeederRequestProposal(uint64(l))
	}
	l = m.PriceFeederInfo.Size()
	n += 1 + l + sovUpdatePriceFeederRequestProposal(uint64(l))
	return n
}

func sovUpdatePriceFeederRequestProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUpdatePriceFeederRequestProposal(x uint64) (n int) {
	return sovUpdatePriceFeederRequestProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UpdatePriceFeederInfoProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpdatePriceFeederRequestProposal
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
			return fmt.Errorf("proto: UpdatePriceFeederInfoProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdatePriceFeederInfoProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdatePriceFeederRequestProposal
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
				return ErrInvalidLengthUpdatePriceFeederRequestProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdatePriceFeederRequestProposal
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
					return ErrIntOverflowUpdatePriceFeederRequestProposal
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
				return ErrInvalidLengthUpdatePriceFeederRequestProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdatePriceFeederRequestProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceFeederInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdatePriceFeederRequestProposal
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
				return ErrInvalidLengthUpdatePriceFeederRequestProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUpdatePriceFeederRequestProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PriceFeederInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpdatePriceFeederRequestProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpdatePriceFeederRequestProposal
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
func skipUpdatePriceFeederRequestProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUpdatePriceFeederRequestProposal
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
					return 0, ErrIntOverflowUpdatePriceFeederRequestProposal
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
					return 0, ErrIntOverflowUpdatePriceFeederRequestProposal
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
				return 0, ErrInvalidLengthUpdatePriceFeederRequestProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUpdatePriceFeederRequestProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUpdatePriceFeederRequestProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUpdatePriceFeederRequestProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUpdatePriceFeederRequestProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUpdatePriceFeederRequestProposal = fmt.Errorf("proto: unexpected end of group")
)
