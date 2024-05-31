// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerprotocol/routerchain/pricefeed/whitelist_ibc_channels.proto

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

type Channels struct {
	Channel []string `protobuf:"bytes,1,rep,name=channel,proto3" json:"channel,omitempty"`
}

func (m *Channels) Reset()         { *m = Channels{} }
func (m *Channels) String() string { return proto.CompactTextString(m) }
func (*Channels) ProtoMessage()    {}
func (*Channels) Descriptor() ([]byte, []int) {
	return fileDescriptor_72b732cf1832540e, []int{0}
}
func (m *Channels) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Channels) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Channels.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Channels) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Channels.Merge(m, src)
}
func (m *Channels) XXX_Size() int {
	return m.Size()
}
func (m *Channels) XXX_DiscardUnknown() {
	xxx_messageInfo_Channels.DiscardUnknown(m)
}

var xxx_messageInfo_Channels proto.InternalMessageInfo

func (m *Channels) GetChannel() []string {
	if m != nil {
		return m.Channel
	}
	return nil
}

type CreateWhitelistedIBCChannelProposal struct {
	Title       string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Channels    Channels `protobuf:"bytes,3,opt,name=channels,proto3" json:"channels"`
}

func (m *CreateWhitelistedIBCChannelProposal) Reset()         { *m = CreateWhitelistedIBCChannelProposal{} }
func (m *CreateWhitelistedIBCChannelProposal) String() string { return proto.CompactTextString(m) }
func (*CreateWhitelistedIBCChannelProposal) ProtoMessage()    {}
func (*CreateWhitelistedIBCChannelProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_72b732cf1832540e, []int{1}
}
func (m *CreateWhitelistedIBCChannelProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateWhitelistedIBCChannelProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateWhitelistedIBCChannelProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateWhitelistedIBCChannelProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWhitelistedIBCChannelProposal.Merge(m, src)
}
func (m *CreateWhitelistedIBCChannelProposal) XXX_Size() int {
	return m.Size()
}
func (m *CreateWhitelistedIBCChannelProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWhitelistedIBCChannelProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWhitelistedIBCChannelProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Channels)(nil), "routerprotocol.routerchain.pricefeed.Channels")
	proto.RegisterType((*CreateWhitelistedIBCChannelProposal)(nil), "routerprotocol.routerchain.pricefeed.CreateWhitelistedIBCChannelProposal")
}

func init() {
	proto.RegisterFile("routerprotocol/routerchain/pricefeed/whitelist_ibc_channels.proto", fileDescriptor_72b732cf1832540e)
}

var fileDescriptor_72b732cf1832540e = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x2c, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0x2a, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x87, 0x70, 0x93, 0x33, 0x12,
	0x33, 0xf3, 0xf4, 0x0b, 0x8a, 0x32, 0x93, 0x53, 0xd3, 0x52, 0x53, 0x53, 0xf4, 0xcb, 0x33, 0x32,
	0x4b, 0x52, 0x73, 0x32, 0x8b, 0x4b, 0xe2, 0x33, 0x93, 0x92, 0xe3, 0x93, 0x33, 0x12, 0xf3, 0xf2,
	0x52, 0x73, 0x8a, 0xf5, 0xc0, 0xba, 0x84, 0x54, 0x50, 0x8d, 0xd0, 0x43, 0x32, 0x42, 0x0f, 0x6e,
	0x84, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58, 0x8d, 0x3e, 0x88, 0x05, 0xd1, 0xab, 0xa4, 0xc2,
	0xc5, 0xe1, 0x0c, 0x35, 0x4d, 0x48, 0x82, 0x8b, 0x1d, 0x6a, 0xb2, 0x04, 0xa3, 0x02, 0xb3, 0x06,
	0x67, 0x10, 0x8c, 0xab, 0xb4, 0x95, 0x91, 0x4b, 0xd9, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0x35, 0x1c,
	0xe6, 0x90, 0xd4, 0x14, 0x4f, 0x27, 0x67, 0xa8, 0xd6, 0x80, 0xa2, 0xfc, 0x82, 0xfc, 0xe2, 0xc4,
	0x1c, 0x21, 0x11, 0x2e, 0xd6, 0x92, 0xcc, 0x92, 0x9c, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x08, 0x47, 0x48, 0x81, 0x8b, 0x3b, 0x25, 0xb5, 0x38, 0xb9, 0x28, 0xb3, 0xa0, 0x24, 0x33,
	0x3f, 0x4f, 0x82, 0x09, 0x2c, 0x87, 0x2c, 0x24, 0x14, 0xc0, 0xc5, 0x01, 0xf3, 0x93, 0x04, 0xb3,
	0x02, 0xa3, 0x06, 0xb7, 0x91, 0x9e, 0x1e, 0x31, 0x9e, 0xd2, 0x83, 0xb9, 0xdd, 0x89, 0xe5, 0xc4,
	0x3d, 0x79, 0x86, 0x20, 0xb8, 0x29, 0x56, 0x2c, 0x1d, 0x0b, 0xe4, 0x19, 0x9c, 0x82, 0x4f, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x32, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49,
	0x2f, 0x39, 0x3f, 0x17, 0x1a, 0xe4, 0xba, 0x68, 0x51, 0xa0, 0x0b, 0x89, 0x83, 0x0a, 0xa4, 0x58,
	0x28, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x2b, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xd2, 0xbc, 0x51, 0xe8, 0xba, 0x01, 0x00, 0x00,
}

func (m *Channels) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Channels) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Channels) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Channel) > 0 {
		for iNdEx := len(m.Channel) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Channel[iNdEx])
			copy(dAtA[i:], m.Channel[iNdEx])
			i = encodeVarintWhitelistIbcChannels(dAtA, i, uint64(len(m.Channel[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CreateWhitelistedIBCChannelProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateWhitelistedIBCChannelProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateWhitelistedIBCChannelProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Channels.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintWhitelistIbcChannels(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintWhitelistIbcChannels(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintWhitelistIbcChannels(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWhitelistIbcChannels(dAtA []byte, offset int, v uint64) int {
	offset -= sovWhitelistIbcChannels(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Channels) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Channel) > 0 {
		for _, s := range m.Channel {
			l = len(s)
			n += 1 + l + sovWhitelistIbcChannels(uint64(l))
		}
	}
	return n
}

func (m *CreateWhitelistedIBCChannelProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovWhitelistIbcChannels(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovWhitelistIbcChannels(uint64(l))
	}
	l = m.Channels.Size()
	n += 1 + l + sovWhitelistIbcChannels(uint64(l))
	return n
}

func sovWhitelistIbcChannels(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWhitelistIbcChannels(x uint64) (n int) {
	return sovWhitelistIbcChannels(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Channels) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhitelistIbcChannels
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
			return fmt.Errorf("proto: Channels: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Channels: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelistIbcChannels
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
				return ErrInvalidLengthWhitelistIbcChannels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhitelistIbcChannels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = append(m.Channel, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWhitelistIbcChannels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhitelistIbcChannels
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
func (m *CreateWhitelistedIBCChannelProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhitelistIbcChannels
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
			return fmt.Errorf("proto: CreateWhitelistedIBCChannelProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateWhitelistedIBCChannelProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelistIbcChannels
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
				return ErrInvalidLengthWhitelistIbcChannels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhitelistIbcChannels
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
					return ErrIntOverflowWhitelistIbcChannels
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
				return ErrInvalidLengthWhitelistIbcChannels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhitelistIbcChannels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelistIbcChannels
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
				return ErrInvalidLengthWhitelistIbcChannels
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhitelistIbcChannels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Channels.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWhitelistIbcChannels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhitelistIbcChannels
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
func skipWhitelistIbcChannels(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWhitelistIbcChannels
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
					return 0, ErrIntOverflowWhitelistIbcChannels
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
					return 0, ErrIntOverflowWhitelistIbcChannels
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
				return 0, ErrInvalidLengthWhitelistIbcChannels
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWhitelistIbcChannels
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWhitelistIbcChannels
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWhitelistIbcChannels        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWhitelistIbcChannels          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWhitelistIbcChannels = fmt.Errorf("proto: unexpected end of group")
)
