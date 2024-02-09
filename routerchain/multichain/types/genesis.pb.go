// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerchain/multichain/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// GenesisState defines the multichain module's genesis state.
type GenesisState struct {
	Params               Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	ChainConfigList      []ChainConfig      `protobuf:"bytes,2,rep,name=chainConfigList,proto3" json:"chainConfigList"`
	ContractConfigList   []ContractConfig   `protobuf:"bytes,3,rep,name=contractConfigList,proto3" json:"contractConfigList"`
	IbcRelayerConfigList []IbcRelayerConfig `protobuf:"bytes,4,rep,name=ibcRelayerConfigList,proto3" json:"ibcRelayerConfigList"`
	// this line is used by starport scaffolding # genesis/proto/state
	NonceObservedStatusList []NonceObservedStatus `protobuf:"bytes,5,rep,name=nonceObservedStatusList,proto3" json:"nonceObservedStatusList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e71e214b7249721a, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetChainConfigList() []ChainConfig {
	if m != nil {
		return m.ChainConfigList
	}
	return nil
}

func (m *GenesisState) GetContractConfigList() []ContractConfig {
	if m != nil {
		return m.ContractConfigList
	}
	return nil
}

func (m *GenesisState) GetIbcRelayerConfigList() []IbcRelayerConfig {
	if m != nil {
		return m.IbcRelayerConfigList
	}
	return nil
}

func (m *GenesisState) GetNonceObservedStatusList() []NonceObservedStatus {
	if m != nil {
		return m.NonceObservedStatusList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "routerprotocol.routerchain.multichain.GenesisState")
}

func init() {
	proto.RegisterFile("routerchain/multichain/genesis.proto", fileDescriptor_e71e214b7249721a)
}

var fileDescriptor_e71e214b7249721a = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x93, 0xdb, 0xde, 0x2e, 0xa6, 0x17, 0x2e, 0x0c, 0x05, 0x4b, 0x17, 0xb1, 0xf8, 0x07,
	0xea, 0xa2, 0x29, 0x44, 0x44, 0xe8, 0xb2, 0x5d, 0x88, 0x28, 0x2a, 0xad, 0x2b, 0x37, 0x61, 0x32,
	0x8e, 0xe9, 0x60, 0x9b, 0xa9, 0x93, 0x89, 0x58, 0x9f, 0xc2, 0xbd, 0x2f, 0xd4, 0x65, 0x97, 0xae,
	0x44, 0xda, 0x17, 0x11, 0xcf, 0x4c, 0x69, 0x5a, 0x1b, 0xc8, 0x26, 0xcc, 0x21, 0xdf, 0xf7, 0xfd,
	0x0e, 0xe7, 0x1c, 0x74, 0x20, 0x45, 0xa2, 0x98, 0xa4, 0x03, 0xc2, 0xa3, 0xd6, 0x28, 0x19, 0x2a,
	0xae, 0x9f, 0x21, 0x8b, 0x58, 0xcc, 0x63, 0x77, 0x2c, 0x85, 0x12, 0xf8, 0x50, 0xab, 0xa0, 0xa0,
	0x62, 0xe8, 0xa6, 0x4c, 0xee, 0xca, 0x54, 0xab, 0x84, 0x22, 0x14, 0x20, 0x6a, 0xfd, 0xbc, 0xb4,
	0xb9, 0xb6, 0x9f, 0x81, 0x18, 0x13, 0x49, 0x46, 0x86, 0x50, 0x3b, 0xca, 0x10, 0xc1, 0xd7, 0xa7,
	0x22, 0x7a, 0xe0, 0xa1, 0x91, 0xb6, 0x32, 0xa4, 0x3c, 0xa0, 0xbe, 0x64, 0x43, 0x32, 0x61, 0x72,
	0xdd, 0xe0, 0x65, 0x18, 0x22, 0x11, 0x51, 0xe6, 0x8b, 0x20, 0x66, 0xf2, 0x99, 0xdd, 0xfb, 0xb1,
	0x22, 0x2a, 0x31, 0xfd, 0xec, 0xbd, 0x17, 0xd1, 0xbf, 0x33, 0x3d, 0x83, 0xbe, 0x22, 0x8a, 0xe1,
	0x0b, 0x54, 0xd2, 0x0d, 0x57, 0xed, 0xba, 0xdd, 0x28, 0x7b, 0x4d, 0x37, 0xd7, 0x4c, 0xdc, 0x1b,
	0x30, 0x75, 0x8a, 0xd3, 0xcf, 0x5d, 0xab, 0x67, 0x22, 0x70, 0x80, 0xfe, 0xc3, 0xdf, 0x2e, 0xb4,
	0x79, 0xc9, 0x63, 0x55, 0xfd, 0x53, 0x2f, 0x34, 0xca, 0x9e, 0x97, 0x33, 0xb5, 0xbb, 0x72, 0x9b,
	0xe8, 0xcd, 0x40, 0xfc, 0x88, 0x30, 0x15, 0x91, 0x92, 0x84, 0xaa, 0x14, 0xa6, 0x00, 0x98, 0x93,
	0xbc, 0x98, 0xb5, 0x00, 0x43, 0xda, 0x12, 0x8b, 0x9f, 0x50, 0x85, 0x07, 0xb4, 0xa7, 0xa7, 0x9f,
	0xc2, 0x15, 0x01, 0x77, 0x9a, 0x13, 0x77, 0xbe, 0x11, 0x61, 0x80, 0x5b, 0xa3, 0xf1, 0x2b, 0xda,
	0x81, 0x05, 0x5e, 0x9b, 0xfd, 0xf5, 0x61, 0x7d, 0x40, 0xfd, 0x0b, 0xd4, 0x76, 0x4e, 0xea, 0xd5,
	0xef, 0x14, 0x03, 0xce, 0x02, 0x74, 0x6e, 0xa7, 0x73, 0xc7, 0x9e, 0xcd, 0x1d, 0xfb, 0x6b, 0xee,
	0xd8, 0x6f, 0x0b, 0xc7, 0x9a, 0x2d, 0x1c, 0xeb, 0x63, 0xe1, 0x58, 0x77, 0xed, 0x90, 0xab, 0x41,
	0x12, 0xb8, 0x54, 0x8c, 0xcc, 0x9d, 0x36, 0x97, 0xfc, 0x65, 0xad, 0x8f, 0xef, 0x25, 0x7d, 0x89,
	0x6a, 0x32, 0x66, 0x71, 0x50, 0x02, 0xe5, 0xf1, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x7d,
	0xc4, 0xca, 0x94, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NonceObservedStatusList) > 0 {
		for iNdEx := len(m.NonceObservedStatusList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NonceObservedStatusList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.IbcRelayerConfigList) > 0 {
		for iNdEx := len(m.IbcRelayerConfigList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IbcRelayerConfigList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ContractConfigList) > 0 {
		for iNdEx := len(m.ContractConfigList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractConfigList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ChainConfigList) > 0 {
		for iNdEx := len(m.ChainConfigList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainConfigList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ChainConfigList) > 0 {
		for _, e := range m.ChainConfigList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ContractConfigList) > 0 {
		for _, e := range m.ContractConfigList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.IbcRelayerConfigList) > 0 {
		for _, e := range m.IbcRelayerConfigList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NonceObservedStatusList) > 0 {
		for _, e := range m.NonceObservedStatusList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainConfigList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainConfigList = append(m.ChainConfigList, ChainConfig{})
			if err := m.ChainConfigList[len(m.ChainConfigList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractConfigList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractConfigList = append(m.ContractConfigList, ContractConfig{})
			if err := m.ContractConfigList[len(m.ContractConfigList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcRelayerConfigList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcRelayerConfigList = append(m.IbcRelayerConfigList, IbcRelayerConfig{})
			if err := m.IbcRelayerConfigList[len(m.IbcRelayerConfigList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NonceObservedStatusList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NonceObservedStatusList = append(m.NonceObservedStatusList, NonceObservedStatus{})
			if err := m.NonceObservedStatusList[len(m.NonceObservedStatusList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
