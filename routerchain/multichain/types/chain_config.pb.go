// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: multichain/chain_config.proto

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

type ChainConfig struct {
	ChainId                 uint64    `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ChainName               string    `protobuf:"bytes,2,opt,name=chainName,proto3" json:"chainName,omitempty"`
	Symbol                  string    `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	ChainType               ChainType `protobuf:"varint,4,opt,name=chainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"chainType,omitempty"`
	ConfirmationsRequired   uint64    `protobuf:"varint,5,opt,name=confirmationsRequired,proto3" json:"confirmationsRequired,omitempty"`
	GatewayContractAddress  string    `protobuf:"bytes,6,opt,name=gatewayContractAddress,proto3" json:"gatewayContractAddress,omitempty"`
	GatewayContractHeight   uint64    `protobuf:"varint,7,opt,name=gatewayContractHeight,proto3" json:"gatewayContractHeight,omitempty"`
	RouterContractAddress   string    `protobuf:"bytes,8,opt,name=routerContractAddress,proto3" json:"routerContractAddress,omitempty"`
	LastObservedEventNonce  uint64    `protobuf:"varint,9,opt,name=lastObservedEventNonce,proto3" json:"lastObservedEventNonce,omitempty"`
	LastObservedValsetNonce uint64    `protobuf:"varint,10,opt,name=lastObservedValsetNonce,proto3" json:"lastObservedValsetNonce,omitempty"`
	Creator                 string    `protobuf:"bytes,11,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *ChainConfig) Reset()         { *m = ChainConfig{} }
func (m *ChainConfig) String() string { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()    {}
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_706820093b076a45, []int{0}
}
func (m *ChainConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainConfig.Merge(m, src)
}
func (m *ChainConfig) XXX_Size() int {
	return m.Size()
}
func (m *ChainConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ChainConfig proto.InternalMessageInfo

func (m *ChainConfig) GetChainId() uint64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *ChainConfig) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *ChainConfig) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ChainConfig) GetChainType() ChainType {
	if m != nil {
		return m.ChainType
	}
	return CHAIN_TYPE_EVM
}

func (m *ChainConfig) GetConfirmationsRequired() uint64 {
	if m != nil {
		return m.ConfirmationsRequired
	}
	return 0
}

func (m *ChainConfig) GetGatewayContractAddress() string {
	if m != nil {
		return m.GatewayContractAddress
	}
	return ""
}

func (m *ChainConfig) GetGatewayContractHeight() uint64 {
	if m != nil {
		return m.GatewayContractHeight
	}
	return 0
}

func (m *ChainConfig) GetRouterContractAddress() string {
	if m != nil {
		return m.RouterContractAddress
	}
	return ""
}

func (m *ChainConfig) GetLastObservedEventNonce() uint64 {
	if m != nil {
		return m.LastObservedEventNonce
	}
	return 0
}

func (m *ChainConfig) GetLastObservedValsetNonce() uint64 {
	if m != nil {
		return m.LastObservedValsetNonce
	}
	return 0
}

func (m *ChainConfig) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*ChainConfig)(nil), "routerprotocol.routerchain.multichain.ChainConfig")
}

func init() { proto.RegisterFile("multichain/chain_config.proto", fileDescriptor_706820093b076a45) }

var fileDescriptor_706820093b076a45 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x1b, 0xad, 0xfd, 0x99, 0x82, 0x8b, 0x01, 0xeb, 0xf8, 0x17, 0x8a, 0x20, 0x74, 0x63,
	0x2a, 0x2a, 0x22, 0xee, 0x34, 0x08, 0xba, 0xa9, 0x10, 0x8a, 0x0b, 0x37, 0x32, 0x49, 0xc6, 0x74,
	0x20, 0xc9, 0xd4, 0x99, 0x49, 0x35, 0x6f, 0xe1, 0xc2, 0x87, 0x72, 0xd9, 0xa5, 0x4b, 0x69, 0x5f,
	0x44, 0x72, 0xd3, 0xd0, 0x52, 0x5b, 0x70, 0x13, 0x72, 0xe6, 0xbb, 0xe7, 0xde, 0x33, 0xc9, 0x45,
	0x07, 0x51, 0x12, 0x6a, 0xee, 0xf5, 0x29, 0x8f, 0x3b, 0xf0, 0x7c, 0xf6, 0x44, 0xfc, 0xc2, 0x03,
	0x6b, 0x20, 0x85, 0x16, 0xf8, 0x48, 0x8a, 0x44, 0x33, 0x09, 0xc2, 0x13, 0xa1, 0x95, 0x4b, 0x28,
	0xb4, 0x66, 0xce, 0xdd, 0xbd, 0x3f, 0x5d, 0x74, 0x3a, 0x60, 0x79, 0x8f, 0xc3, 0xcf, 0x32, 0x6a,
	0xd8, 0xd9, 0xa1, 0x0d, 0x9d, 0xf1, 0x0e, 0xaa, 0xe5, 0x35, 0xdc, 0x27, 0x46, 0xcb, 0x68, 0x97,
	0x9d, 0x2a, 0xe8, 0x7b, 0x1f, 0xef, 0xa3, 0x3a, 0xbc, 0x76, 0x69, 0xc4, 0xc8, 0x5a, 0xcb, 0x68,
	0xd7, 0x9d, 0xd9, 0x01, 0x6e, 0xa2, 0x8a, 0x4a, 0x23, 0x57, 0x84, 0x64, 0x1d, 0xd0, 0x54, 0xe1,
	0xee, 0xd4, 0xd5, 0x4b, 0x07, 0x8c, 0x94, 0x5b, 0x46, 0x7b, 0xf3, 0xf4, 0xc4, 0xfa, 0x57, 0x70,
	0xcb, 0x2e, 0x7c, 0xce, 0xac, 0x05, 0x3e, 0x47, 0x5b, 0xf0, 0x11, 0x64, 0x44, 0x35, 0x17, 0xb1,
	0x72, 0xd8, 0x6b, 0xc2, 0x25, 0xf3, 0xc9, 0x06, 0xa4, 0x5d, 0x0e, 0xf1, 0x05, 0x6a, 0x06, 0x54,
	0xb3, 0x37, 0x9a, 0xda, 0x22, 0xd6, 0x92, 0x7a, 0xfa, 0xda, 0xf7, 0x25, 0x53, 0x8a, 0x54, 0x20,
	0xed, 0x0a, 0x9a, 0x4d, 0x5b, 0x20, 0x77, 0x8c, 0x07, 0x7d, 0x4d, 0xaa, 0xf9, 0xb4, 0xa5, 0x30,
	0x73, 0xe5, 0x57, 0x5a, 0x1c, 0x56, 0x83, 0x61, 0xcb, 0x61, 0x96, 0x31, 0xa4, 0x4a, 0x3f, 0xb8,
	0x8a, 0xc9, 0x21, 0xf3, 0x6f, 0x87, 0x2c, 0xd6, 0x5d, 0x11, 0x7b, 0x8c, 0xd4, 0x61, 0xd8, 0x0a,
	0x8a, 0x2f, 0xd1, 0xf6, 0x3c, 0x79, 0xa4, 0xa1, 0x62, 0x53, 0x23, 0x02, 0xe3, 0x2a, 0x8c, 0x09,
	0xaa, 0x7a, 0x92, 0x51, 0x2d, 0x24, 0x69, 0x40, 0xb2, 0x42, 0xde, 0xf4, 0xbe, 0xc6, 0xa6, 0x31,
	0x1a, 0x9b, 0xc6, 0xcf, 0xd8, 0x34, 0x3e, 0x26, 0x66, 0x69, 0x34, 0x31, 0x4b, 0xdf, 0x13, 0xb3,
	0xf4, 0x74, 0x15, 0x70, 0xdd, 0x4f, 0x5c, 0xcb, 0x13, 0x51, 0x27, 0xbf, 0xc7, 0x71, 0xf1, 0x1f,
	0x0b, 0x9d, 0xaf, 0xda, 0x7b, 0x67, 0x6e, 0xef, 0xb2, 0x8d, 0x53, 0x6e, 0x05, 0x2a, 0xcf, 0x7e,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x8a, 0x91, 0xdd, 0xd8, 0x02, 0x00, 0x00,
}

func (m *ChainConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintChainConfig(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x5a
	}
	if m.LastObservedValsetNonce != 0 {
		i = encodeVarintChainConfig(dAtA, i, uint64(m.LastObservedValsetNonce))
		i--
		dAtA[i] = 0x50
	}
	if m.LastObservedEventNonce != 0 {
		i = encodeVarintChainConfig(dAtA, i, uint64(m.LastObservedEventNonce))
		i--
		dAtA[i] = 0x48
	}
	if len(m.RouterContractAddress) > 0 {
		i -= len(m.RouterContractAddress)
		copy(dAtA[i:], m.RouterContractAddress)
		i = encodeVarintChainConfig(dAtA, i, uint64(len(m.RouterContractAddress)))
		i--
		dAtA[i] = 0x42
	}
	if m.GatewayContractHeight != 0 {
		i = encodeVarintChainConfig(dAtA, i, uint64(m.GatewayContractHeight))
		i--
		dAtA[i] = 0x38
	}
	if len(m.GatewayContractAddress) > 0 {
		i -= len(m.GatewayContractAddress)
		copy(dAtA[i:], m.GatewayContractAddress)
		i = encodeVarintChainConfig(dAtA, i, uint64(len(m.GatewayContractAddress)))
		i--
		dAtA[i] = 0x32
	}
	if m.ConfirmationsRequired != 0 {
		i = encodeVarintChainConfig(dAtA, i, uint64(m.ConfirmationsRequired))
		i--
		dAtA[i] = 0x28
	}
	if m.ChainType != 0 {
		i = encodeVarintChainConfig(dAtA, i, uint64(m.ChainType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintChainConfig(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintChainConfig(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0x12
	}
	if m.ChainId != 0 {
		i = encodeVarintChainConfig(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintChainConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovChainConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainId != 0 {
		n += 1 + sovChainConfig(uint64(m.ChainId))
	}
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovChainConfig(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovChainConfig(uint64(l))
	}
	if m.ChainType != 0 {
		n += 1 + sovChainConfig(uint64(m.ChainType))
	}
	if m.ConfirmationsRequired != 0 {
		n += 1 + sovChainConfig(uint64(m.ConfirmationsRequired))
	}
	l = len(m.GatewayContractAddress)
	if l > 0 {
		n += 1 + l + sovChainConfig(uint64(l))
	}
	if m.GatewayContractHeight != 0 {
		n += 1 + sovChainConfig(uint64(m.GatewayContractHeight))
	}
	l = len(m.RouterContractAddress)
	if l > 0 {
		n += 1 + l + sovChainConfig(uint64(l))
	}
	if m.LastObservedEventNonce != 0 {
		n += 1 + sovChainConfig(uint64(m.LastObservedEventNonce))
	}
	if m.LastObservedValsetNonce != 0 {
		n += 1 + sovChainConfig(uint64(m.LastObservedValsetNonce))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovChainConfig(uint64(l))
	}
	return n
}

func sovChainConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChainConfig(x uint64) (n int) {
	return sovChainConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChainConfig
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
			return fmt.Errorf("proto: ChainConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
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
				return ErrInvalidLengthChainConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
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
				return ErrInvalidLengthChainConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainType", wireType)
			}
			m.ChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainType |= ChainType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmationsRequired", wireType)
			}
			m.ConfirmationsRequired = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConfirmationsRequired |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GatewayContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
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
				return ErrInvalidLengthChainConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GatewayContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GatewayContractHeight", wireType)
			}
			m.GatewayContractHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GatewayContractHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouterContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
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
				return ErrInvalidLengthChainConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RouterContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastObservedEventNonce", wireType)
			}
			m.LastObservedEventNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastObservedEventNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastObservedValsetNonce", wireType)
			}
			m.LastObservedValsetNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastObservedValsetNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
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
				return ErrInvalidLengthChainConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChainConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChainConfig
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
func skipChainConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChainConfig
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
					return 0, ErrIntOverflowChainConfig
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
					return 0, ErrIntOverflowChainConfig
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
				return 0, ErrInvalidLengthChainConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChainConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChainConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChainConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChainConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChainConfig = fmt.Errorf("proto: unexpected end of group")
)
