// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerprotocol/routerchain/multichain/ibc_relayer_config.proto

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

type IbcRelayerConnectionType int32

const (
	GATEWAY_ENDPOINT IbcRelayerConnectionType = 0
	VOYAGER_ENDPOINT IbcRelayerConnectionType = 1
)

var IbcRelayerConnectionType_name = map[int32]string{
	0: "GATEWAY_ENDPOINT",
	1: "VOYAGER_ENDPOINT",
}

var IbcRelayerConnectionType_value = map[string]int32{
	"GATEWAY_ENDPOINT": 0,
	"VOYAGER_ENDPOINT": 1,
}

func (x IbcRelayerConnectionType) String() string {
	return proto.EnumName(IbcRelayerConnectionType_name, int32(x))
}

func (IbcRelayerConnectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_baacafeab7cd839d, []int{0}
}

type IbcRelayerConfig struct {
	ChainId        string                   `protobuf:"bytes,1,opt,name=chainId,proto3" json:"chainId,omitempty"`
	RelayerName    string                   `protobuf:"bytes,2,opt,name=relayerName,proto3" json:"relayerName,omitempty"`
	Channel        string                   `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	RelayerEnabled bool                     `protobuf:"varint,4,opt,name=relayerEnabled,proto3" json:"relayerEnabled,omitempty"`
	ConnectionType IbcRelayerConnectionType `protobuf:"varint,5,opt,name=connectionType,proto3,enum=routerprotocol.routerchain.multichain.IbcRelayerConnectionType" json:"connectionType,omitempty"`
}

func (m *IbcRelayerConfig) Reset()         { *m = IbcRelayerConfig{} }
func (m *IbcRelayerConfig) String() string { return proto.CompactTextString(m) }
func (*IbcRelayerConfig) ProtoMessage()    {}
func (*IbcRelayerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_baacafeab7cd839d, []int{0}
}
func (m *IbcRelayerConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IbcRelayerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IbcRelayerConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IbcRelayerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IbcRelayerConfig.Merge(m, src)
}
func (m *IbcRelayerConfig) XXX_Size() int {
	return m.Size()
}
func (m *IbcRelayerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_IbcRelayerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_IbcRelayerConfig proto.InternalMessageInfo

func (m *IbcRelayerConfig) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *IbcRelayerConfig) GetRelayerName() string {
	if m != nil {
		return m.RelayerName
	}
	return ""
}

func (m *IbcRelayerConfig) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *IbcRelayerConfig) GetRelayerEnabled() bool {
	if m != nil {
		return m.RelayerEnabled
	}
	return false
}

func (m *IbcRelayerConfig) GetConnectionType() IbcRelayerConnectionType {
	if m != nil {
		return m.ConnectionType
	}
	return GATEWAY_ENDPOINT
}

func init() {
	proto.RegisterEnum("routerprotocol.routerchain.multichain.IbcRelayerConnectionType", IbcRelayerConnectionType_name, IbcRelayerConnectionType_value)
	proto.RegisterType((*IbcRelayerConfig)(nil), "routerprotocol.routerchain.multichain.IbcRelayerConfig")
}

func init() {
	proto.RegisterFile("routerprotocol/routerchain/multichain/ibc_relayer_config.proto", fileDescriptor_baacafeab7cd839d)
}

var fileDescriptor_baacafeab7cd839d = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x2b, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0x2a, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x87, 0x70, 0x93, 0x33, 0x12,
	0x33, 0xf3, 0xf4, 0x73, 0x4b, 0x73, 0x4a, 0x32, 0x21, 0xcc, 0xcc, 0xa4, 0xe4, 0xf8, 0xa2, 0xd4,
	0x9c, 0xc4, 0xca, 0xd4, 0xa2, 0xf8, 0xe4, 0xfc, 0xbc, 0xb4, 0xcc, 0x74, 0x3d, 0xb0, 0x16, 0x21,
	0x55, 0x54, 0xfd, 0x7a, 0x48, 0xfa, 0xf5, 0x10, 0xfa, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1,
	0x8a, 0xf4, 0x41, 0x2c, 0x88, 0x66, 0xa5, 0x9f, 0x8c, 0x5c, 0x02, 0x9e, 0x49, 0xc9, 0x41, 0x10,
	0x83, 0x9d, 0xc1, 0xe6, 0x0a, 0x49, 0x70, 0xb1, 0x83, 0xf5, 0x78, 0xa6, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42, 0x0a, 0x5c, 0xdc, 0x50, 0x37, 0xf8, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x81, 0x65, 0x91, 0x85, 0xa0, 0x7a, 0xf3, 0xf2, 0x52, 0x73, 0x24, 0x98, 0xe1, 0x7a, 0x41,
	0x5c, 0x21, 0x35, 0x2e, 0x3e, 0xa8, 0x42, 0xd7, 0xbc, 0xc4, 0xa4, 0x9c, 0xd4, 0x14, 0x09, 0x16,
	0x05, 0x46, 0x0d, 0x8e, 0x20, 0x34, 0x51, 0xa1, 0x74, 0x2e, 0xbe, 0xe4, 0xfc, 0xbc, 0xbc, 0xd4,
	0xe4, 0x92, 0xcc, 0xfc, 0xbc, 0x90, 0xca, 0x82, 0x54, 0x09, 0x56, 0x05, 0x46, 0x0d, 0x3e, 0x23,
	0x7b, 0x3d, 0xa2, 0x3c, 0xaa, 0x87, 0xe2, 0x1d, 0x24, 0x63, 0x82, 0xd0, 0x8c, 0xd5, 0xaa, 0xe0,
	0x92, 0xc0, 0xa5, 0x56, 0x48, 0x8b, 0x4b, 0xc0, 0xdd, 0x31, 0xc4, 0x35, 0xdc, 0x31, 0x32, 0xde,
	0xd5, 0xcf, 0x25, 0xc0, 0xdf, 0xd3, 0x2f, 0x44, 0x80, 0x41, 0x4a, 0xa4, 0x6b, 0xae, 0x02, 0x86,
	0x38, 0x48, 0x6d, 0x98, 0x7f, 0xa4, 0xa3, 0xbb, 0x6b, 0x10, 0x42, 0x2d, 0x23, 0x44, 0x2d, 0xba,
	0xb8, 0x14, 0x4b, 0xc7, 0x62, 0x39, 0x06, 0xa7, 0x90, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c,
	0x96, 0x63, 0x88, 0xb2, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0x85, 0x26,
	0x04, 0x5d, 0xb4, 0x84, 0xa1, 0x0b, 0x49, 0x0e, 0x15, 0xc8, 0x69, 0xa3, 0xa4, 0xb2, 0x20, 0xb5,
	0x38, 0x89, 0x0d, 0xac, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x32, 0xfe, 0x20, 0x51,
	0x02, 0x00, 0x00,
}

func (m *IbcRelayerConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IbcRelayerConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IbcRelayerConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ConnectionType != 0 {
		i = encodeVarintIbcRelayerConfig(dAtA, i, uint64(m.ConnectionType))
		i--
		dAtA[i] = 0x28
	}
	if m.RelayerEnabled {
		i--
		if m.RelayerEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintIbcRelayerConfig(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RelayerName) > 0 {
		i -= len(m.RelayerName)
		copy(dAtA[i:], m.RelayerName)
		i = encodeVarintIbcRelayerConfig(dAtA, i, uint64(len(m.RelayerName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintIbcRelayerConfig(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIbcRelayerConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovIbcRelayerConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IbcRelayerConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovIbcRelayerConfig(uint64(l))
	}
	l = len(m.RelayerName)
	if l > 0 {
		n += 1 + l + sovIbcRelayerConfig(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovIbcRelayerConfig(uint64(l))
	}
	if m.RelayerEnabled {
		n += 2
	}
	if m.ConnectionType != 0 {
		n += 1 + sovIbcRelayerConfig(uint64(m.ConnectionType))
	}
	return n
}

func sovIbcRelayerConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIbcRelayerConfig(x uint64) (n int) {
	return sovIbcRelayerConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IbcRelayerConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbcRelayerConfig
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
			return fmt.Errorf("proto: IbcRelayerConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IbcRelayerConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcRelayerConfig
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
				return ErrInvalidLengthIbcRelayerConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbcRelayerConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayerName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcRelayerConfig
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
				return ErrInvalidLengthIbcRelayerConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbcRelayerConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RelayerName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcRelayerConfig
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
				return ErrInvalidLengthIbcRelayerConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbcRelayerConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayerEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcRelayerConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RelayerEnabled = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionType", wireType)
			}
			m.ConnectionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcRelayerConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConnectionType |= IbcRelayerConnectionType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIbcRelayerConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbcRelayerConfig
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
func skipIbcRelayerConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIbcRelayerConfig
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
					return 0, ErrIntOverflowIbcRelayerConfig
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
					return 0, ErrIntOverflowIbcRelayerConfig
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
				return 0, ErrInvalidLengthIbcRelayerConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIbcRelayerConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIbcRelayerConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIbcRelayerConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIbcRelayerConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIbcRelayerConfig = fmt.Errorf("proto: unexpected end of group")
)
