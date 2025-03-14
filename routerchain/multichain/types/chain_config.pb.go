// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerprotocol/routerchain/multichain/chain_config.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
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

type ContractType int32

const (
	GATEWAY ContractType = 0
	VOYAGER ContractType = 1
)

var ContractType_name = map[int32]string{
	0: "GATEWAY",
	1: "VOYAGER",
}

var ContractType_value = map[string]int32{
	"GATEWAY": 0,
	"VOYAGER": 1,
}

func (x ContractType) String() string {
	return proto.EnumName(ContractType_name, int32(x))
}

func (ContractType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f5dfeea5cb73ed75, []int{0}
}

type ContractConfig struct {
	ChainId                       string                      `protobuf:"bytes,1,opt,name=chainId,proto3" json:"chainId,omitempty"`
	ContractAddress               string                      `protobuf:"bytes,2,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	ContractHeight                uint64                      `protobuf:"varint,3,opt,name=contractHeight,proto3" json:"contractHeight,omitempty"`
	LastObservedEventNonce        uint64                      `protobuf:"varint,4,opt,name=lastObservedEventNonce,proto3" json:"lastObservedEventNonce,omitempty"`
	LastObservedEventBlockHeight  uint64                      `protobuf:"varint,5,opt,name=lastObservedEventBlockHeight,proto3" json:"lastObservedEventBlockHeight,omitempty"`
	ContractType                  ContractType                `protobuf:"varint,6,opt,name=contractType,proto3,enum=routerprotocol.routerchain.multichain.ContractType" json:"contractType,omitempty"`
	ClaimSlashingEnabled          bool                        `protobuf:"varint,7,opt,name=claimSlashingEnabled,proto3" json:"claimSlashingEnabled,omitempty"`
	ClaimSlashingWindow           uint64                      `protobuf:"varint,8,opt,name=claimSlashingWindow,proto3" json:"claimSlashingWindow,omitempty"`
	SlashFractionMissingClaim     cosmossdk_io_math.LegacyDec `protobuf:"bytes,9,opt,name=slash_fraction_missing_claim,json=slashFractionMissingClaim,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"slash_fraction_missing_claim"`
	SlashFractionConflictingClaim cosmossdk_io_math.LegacyDec `protobuf:"bytes,10,opt,name=slash_fraction_conflicting_claim,json=slashFractionConflictingClaim,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"slash_fraction_conflicting_claim"`
	ContractEnabled               bool                        `protobuf:"varint,11,opt,name=contract_enabled,json=contractEnabled,proto3" json:"contract_enabled,omitempty"`
}

func (m *ContractConfig) Reset()         { *m = ContractConfig{} }
func (m *ContractConfig) String() string { return proto.CompactTextString(m) }
func (*ContractConfig) ProtoMessage()    {}
func (*ContractConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5dfeea5cb73ed75, []int{0}
}
func (m *ContractConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractConfig.Merge(m, src)
}
func (m *ContractConfig) XXX_Size() int {
	return m.Size()
}
func (m *ContractConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ContractConfig proto.InternalMessageInfo

func (m *ContractConfig) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *ContractConfig) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ContractConfig) GetContractHeight() uint64 {
	if m != nil {
		return m.ContractHeight
	}
	return 0
}

func (m *ContractConfig) GetLastObservedEventNonce() uint64 {
	if m != nil {
		return m.LastObservedEventNonce
	}
	return 0
}

func (m *ContractConfig) GetLastObservedEventBlockHeight() uint64 {
	if m != nil {
		return m.LastObservedEventBlockHeight
	}
	return 0
}

func (m *ContractConfig) GetContractType() ContractType {
	if m != nil {
		return m.ContractType
	}
	return GATEWAY
}

func (m *ContractConfig) GetClaimSlashingEnabled() bool {
	if m != nil {
		return m.ClaimSlashingEnabled
	}
	return false
}

func (m *ContractConfig) GetClaimSlashingWindow() uint64 {
	if m != nil {
		return m.ClaimSlashingWindow
	}
	return 0
}

func (m *ContractConfig) GetContractEnabled() bool {
	if m != nil {
		return m.ContractEnabled
	}
	return false
}

type ChainConfig struct {
	ChainId                 string    `protobuf:"bytes,1,opt,name=chainId,proto3" json:"chainId,omitempty"`
	ChainName               string    `protobuf:"bytes,2,opt,name=chainName,proto3" json:"chainName,omitempty"`
	Symbol                  string    `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	NativeDecimals          int64     `protobuf:"varint,4,opt,name=native_decimals,json=nativeDecimals,proto3" json:"native_decimals,omitempty"`
	ChainType               ChainType `protobuf:"varint,5,opt,name=chainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"chainType,omitempty"`
	ConfirmationsRequired   uint64    `protobuf:"varint,6,opt,name=confirmationsRequired,proto3" json:"confirmationsRequired,omitempty"`
	LastObservedValsetNonce uint64    `protobuf:"varint,7,opt,name=lastObservedValsetNonce,proto3" json:"lastObservedValsetNonce,omitempty"`
	ChainEnabled            bool      `protobuf:"varint,8,opt,name=chain_enabled,json=chainEnabled,proto3" json:"chain_enabled,omitempty"`
	IsUnorderedNonces       bool      `protobuf:"varint,9,opt,name=is_unordered_nonces,json=isUnorderedNonces,proto3" json:"is_unordered_nonces,omitempty"`
}

func (m *ChainConfig) Reset()         { *m = ChainConfig{} }
func (m *ChainConfig) String() string { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()    {}
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5dfeea5cb73ed75, []int{1}
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

func (m *ChainConfig) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
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

func (m *ChainConfig) GetNativeDecimals() int64 {
	if m != nil {
		return m.NativeDecimals
	}
	return 0
}

func (m *ChainConfig) GetChainType() ChainType {
	if m != nil {
		return m.ChainType
	}
	return CHAIN_TYPE_NONE
}

func (m *ChainConfig) GetConfirmationsRequired() uint64 {
	if m != nil {
		return m.ConfirmationsRequired
	}
	return 0
}

func (m *ChainConfig) GetLastObservedValsetNonce() uint64 {
	if m != nil {
		return m.LastObservedValsetNonce
	}
	return 0
}

func (m *ChainConfig) GetChainEnabled() bool {
	if m != nil {
		return m.ChainEnabled
	}
	return false
}

func (m *ChainConfig) GetIsUnorderedNonces() bool {
	if m != nil {
		return m.IsUnorderedNonces
	}
	return false
}

func init() {
	proto.RegisterEnum("routerprotocol.routerchain.multichain.ContractType", ContractType_name, ContractType_value)
	proto.RegisterType((*ContractConfig)(nil), "routerprotocol.routerchain.multichain.ContractConfig")
	proto.RegisterType((*ChainConfig)(nil), "routerprotocol.routerchain.multichain.ChainConfig")
}

func init() {
	proto.RegisterFile("routerprotocol/routerchain/multichain/chain_config.proto", fileDescriptor_f5dfeea5cb73ed75)
}

var fileDescriptor_f5dfeea5cb73ed75 = []byte{
	// 718 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0xee, 0x4a, 0x81, 0x76, 0xa8, 0x05, 0x07, 0xc4, 0xa5, 0xe2, 0xd2, 0x40, 0xd4, 0x6a, 0xc2,
	0x2e, 0x3f, 0x86, 0x10, 0xef, 0xda, 0x52, 0xc1, 0x44, 0x21, 0x59, 0x11, 0x82, 0x37, 0x9b, 0xd9,
	0xd9, 0x61, 0x3b, 0x61, 0x77, 0x06, 0x77, 0xb6, 0xd5, 0xbe, 0x81, 0xf1, 0xca, 0x7b, 0xe3, 0x95,
	0xef, 0xe0, 0x33, 0x70, 0xc9, 0xa5, 0xf1, 0x82, 0x18, 0x78, 0x11, 0xb3, 0xb3, 0xbb, 0xf6, 0x47,
	0x50, 0xbc, 0x69, 0xe6, 0x7c, 0xdf, 0x39, 0xdf, 0xe9, 0x9e, 0xf9, 0xe6, 0x80, 0xf5, 0x80, 0xb7,
	0x42, 0x12, 0x1c, 0x07, 0x3c, 0xe4, 0x98, 0x7b, 0x46, 0x1c, 0xe2, 0x26, 0xa2, 0xcc, 0xf0, 0x5b,
	0x5e, 0x48, 0xe3, 0xa3, 0xfc, 0xb5, 0x30, 0x67, 0x87, 0xd4, 0xd5, 0x65, 0x32, 0xbc, 0xdf, 0x5f,
	0xa9, 0xf7, 0x54, 0xea, 0xdd, 0xca, 0xd2, 0xda, 0xff, 0x34, 0x08, 0x3b, 0xc7, 0x24, 0x96, 0x2f,
	0x69, 0x98, 0x0b, 0x9f, 0x0b, 0xc3, 0x46, 0x82, 0x18, 0xed, 0x65, 0x9b, 0x84, 0x68, 0xd9, 0xc0,
	0x9c, 0xb2, 0x84, 0x9f, 0x72, 0xb9, 0xcb, 0xe5, 0xd1, 0x88, 0x4e, 0x09, 0x3a, 0xe3, 0x72, 0xee,
	0x7a, 0xc4, 0x90, 0x91, 0xdd, 0x3a, 0x34, 0x10, 0xeb, 0xc4, 0xd4, 0xfc, 0xb7, 0x61, 0x50, 0xac,
	0x73, 0x16, 0x06, 0x08, 0x87, 0x75, 0xf9, 0x21, 0x50, 0x05, 0xa3, 0xb2, 0xef, 0x73, 0x47, 0x55,
	0xca, 0x4a, 0x25, 0x6f, 0xa6, 0x21, 0xac, 0x80, 0x71, 0x9c, 0xe4, 0x56, 0x1d, 0x27, 0x20, 0x42,
	0xa8, 0x37, 0x64, 0xc6, 0x20, 0x0c, 0x1f, 0x80, 0x62, 0x0a, 0x6d, 0x11, 0xea, 0x36, 0x43, 0x75,
	0xa8, 0xac, 0x54, 0xb2, 0xe6, 0x00, 0x0a, 0xd7, 0xc0, 0xb4, 0x87, 0x44, 0xb8, 0x63, 0x0b, 0x12,
	0xb4, 0x89, 0xd3, 0x68, 0x13, 0x16, 0x6e, 0x73, 0x86, 0x89, 0x9a, 0x95, 0xf9, 0x57, 0xb0, 0xb0,
	0x06, 0x66, 0xff, 0x60, 0x6a, 0x1e, 0xc7, 0x47, 0x49, 0xb7, 0x61, 0x59, 0xfd, 0xd7, 0x1c, 0xb8,
	0x0f, 0x0a, 0xe9, 0xbf, 0xd9, 0xed, 0x1c, 0x13, 0x75, 0xa4, 0xac, 0x54, 0x8a, 0x2b, 0xab, 0xfa,
	0xb5, 0x6e, 0x50, 0xaf, 0xf7, 0x94, 0x9a, 0x7d, 0x42, 0x70, 0x05, 0x4c, 0x61, 0x0f, 0x51, 0xff,
	0x95, 0x87, 0x44, 0x93, 0x32, 0xb7, 0xc1, 0x90, 0xed, 0x11, 0x47, 0x1d, 0x2d, 0x2b, 0x95, 0x9c,
	0x79, 0x29, 0x07, 0x97, 0xc0, 0x64, 0x1f, 0xbe, 0x4f, 0x99, 0xc3, 0xdf, 0xa9, 0x39, 0xf9, 0x1d,
	0x97, 0x51, 0xd0, 0x01, 0xb3, 0x22, 0x42, 0xac, 0xc3, 0xa8, 0x31, 0xe5, 0xcc, 0xf2, 0xa9, 0x10,
	0x94, 0xb9, 0x96, 0xcc, 0x56, 0xf3, 0x65, 0xa5, 0x52, 0xa8, 0x2d, 0x9c, 0x9c, 0xcd, 0x65, 0x7e,
	0x9c, 0xcd, 0xdd, 0x8d, 0x8d, 0x23, 0x9c, 0x23, 0x9d, 0x72, 0xc3, 0x47, 0x61, 0x53, 0x7f, 0x41,
	0x5c, 0x84, 0x3b, 0x1b, 0x04, 0x9b, 0x33, 0x52, 0xe8, 0x59, 0xa2, 0xf3, 0x32, 0x96, 0xa9, 0x47,
	0x2a, 0xd0, 0x03, 0xe5, 0x81, 0x2e, 0x91, 0xdd, 0x3d, 0x8a, 0xc3, 0x6e, 0x27, 0x70, 0xfd, 0x4e,
	0xf7, 0xfa, 0x3a, 0xd5, 0xbb, 0x52, 0x71, 0xb7, 0x47, 0x60, 0x22, 0x9d, 0xa4, 0x45, 0x92, 0xa9,
	0x8d, 0xc9, 0xa9, 0xfd, 0x76, 0x58, 0x32, 0xb0, 0xf9, 0xcf, 0x43, 0x60, 0xac, 0x1e, 0xdd, 0xc4,
	0x3f, 0x5d, 0x3b, 0x0b, 0xf2, 0xf2, 0xb8, 0x8d, 0x7c, 0x92, 0xf8, 0xb5, 0x0b, 0xc0, 0x69, 0x30,
	0x22, 0x3a, 0xbe, 0xcd, 0x3d, 0xe9, 0xd0, 0xbc, 0x99, 0x44, 0xf0, 0x21, 0x18, 0x67, 0x28, 0xa4,
	0x6d, 0x62, 0x39, 0x04, 0x53, 0x1f, 0x79, 0x42, 0x5a, 0x72, 0xc8, 0x2c, 0xc6, 0xf0, 0x46, 0x82,
	0xc2, 0xed, 0x44, 0x5e, 0x7a, 0x68, 0x58, 0x7a, 0x68, 0xe9, 0xba, 0x1e, 0x4a, 0xeb, 0xcc, 0xae,
	0x04, 0x7c, 0x02, 0x6e, 0xcb, 0x8d, 0x12, 0xf8, 0x28, 0x9a, 0x91, 0x30, 0xc9, 0xdb, 0x16, 0x0d,
	0x88, 0x23, 0xfd, 0x99, 0x35, 0x2f, 0x27, 0xe1, 0x3a, 0xb8, 0xd3, 0x6b, 0xf6, 0x3d, 0xe4, 0x09,
	0x92, 0xbc, 0xa4, 0x51, 0x59, 0x77, 0x15, 0x0d, 0x17, 0xc0, 0xcd, 0x78, 0xcd, 0xa4, 0x03, 0xcf,
	0xc9, 0x81, 0x17, 0x24, 0x98, 0xda, 0x53, 0x07, 0x93, 0x54, 0x58, 0x2d, 0xc6, 0x03, 0x87, 0x04,
	0xc4, 0xb1, 0x58, 0x54, 0x2a, 0xa4, 0xc7, 0x72, 0xe6, 0x2d, 0x2a, 0x5e, 0xa7, 0x8c, 0xd4, 0x14,
	0x8f, 0xb7, 0x40, 0xa1, 0xf7, 0x81, 0x44, 0xb7, 0xb3, 0x59, 0xdd, 0x6d, 0xec, 0x57, 0x0f, 0x26,
	0x32, 0xa5, 0xb1, 0x8f, 0x5f, 0xca, 0x69, 0x18, 0x31, 0x7b, 0x3b, 0x07, 0xd5, 0xcd, 0x86, 0x39,
	0xa1, 0xc4, 0x4c, 0x12, 0x96, 0xb2, 0x1f, 0xbe, 0x6a, 0x99, 0xda, 0xee, 0xc9, 0xb9, 0xa6, 0x9c,
	0x9e, 0x6b, 0xca, 0xcf, 0x73, 0x4d, 0xf9, 0x74, 0xa1, 0x65, 0x4e, 0x2f, 0xb4, 0xcc, 0xf7, 0x0b,
	0x2d, 0xf3, 0xe6, 0xa9, 0x4b, 0xc3, 0x66, 0xcb, 0xd6, 0x31, 0xf7, 0x93, 0xfd, 0xb9, 0x38, 0xb0,
	0x4f, 0x17, 0xe3, 0x2d, 0xfa, 0xbe, 0x77, 0xa5, 0x46, 0xcb, 0x54, 0xd8, 0x23, 0x32, 0x73, 0xf5,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xb1, 0x18, 0xe0, 0xe9, 0x05, 0x00, 0x00,
}

func (m *ContractConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ContractEnabled {
		i--
		if m.ContractEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	{
		size := m.SlashFractionConflictingClaim.Size()
		i -= size
		if _, err := m.SlashFractionConflictingClaim.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintChainConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.SlashFractionMissingClaim.Size()
		i -= size
		if _, err := m.SlashFractionMissingClaim.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintChainConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.ClaimSlashingWindow != 0 {
		i = encodeVarintChainConfig(dAtA, i, uint64(m.ClaimSlashingWindow))
		i--
		dAtA[i] = 0x40
	}
	if m.ClaimSlashingEnabled {
		i--
		if m.ClaimSlashingEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.ContractType != 0 {
		i = encodeVarintChainConfig(dAtA, i, uint64(m.ContractType))
		i--
		dAtA[i] = 0x30
	}
	if m.LastObservedEventBlockHeight != 0 {
		i = encodeVarintChainConfig(dAtA, i, uint64(m.LastObservedEventBlockHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.LastObservedEventNonce != 0 {
		i = encodeVarintChainConfig(dAtA, i, uint64(m.LastObservedEventNonce))
		i--
		dAtA[i] = 0x20
	}
	if m.ContractHeight != 0 {
		i = encodeVarintChainConfig(dAtA, i, uint64(m.ContractHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintChainConfig(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintChainConfig(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	if m.IsUnorderedNonces {
		i--
		if m.IsUnorderedNonces {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.ChainEnabled {
		i--
		if m.ChainEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.LastObservedValsetNonce != 0 {
		i = encodeVarintChainConfig(dAtA, i, uint64(m.LastObservedValsetNonce))
		i--
		dAtA[i] = 0x38
	}
	if m.ConfirmationsRequired != 0 {
		i = encodeVarintChainConfig(dAtA, i, uint64(m.ConfirmationsRequired))
		i--
		dAtA[i] = 0x30
	}
	if m.ChainType != 0 {
		i = encodeVarintChainConfig(dAtA, i, uint64(m.ChainType))
		i--
		dAtA[i] = 0x28
	}
	if m.NativeDecimals != 0 {
		i = encodeVarintChainConfig(dAtA, i, uint64(m.NativeDecimals))
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
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintChainConfig(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
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
func (m *ContractConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovChainConfig(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovChainConfig(uint64(l))
	}
	if m.ContractHeight != 0 {
		n += 1 + sovChainConfig(uint64(m.ContractHeight))
	}
	if m.LastObservedEventNonce != 0 {
		n += 1 + sovChainConfig(uint64(m.LastObservedEventNonce))
	}
	if m.LastObservedEventBlockHeight != 0 {
		n += 1 + sovChainConfig(uint64(m.LastObservedEventBlockHeight))
	}
	if m.ContractType != 0 {
		n += 1 + sovChainConfig(uint64(m.ContractType))
	}
	if m.ClaimSlashingEnabled {
		n += 2
	}
	if m.ClaimSlashingWindow != 0 {
		n += 1 + sovChainConfig(uint64(m.ClaimSlashingWindow))
	}
	l = m.SlashFractionMissingClaim.Size()
	n += 1 + l + sovChainConfig(uint64(l))
	l = m.SlashFractionConflictingClaim.Size()
	n += 1 + l + sovChainConfig(uint64(l))
	if m.ContractEnabled {
		n += 2
	}
	return n
}

func (m *ChainConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovChainConfig(uint64(l))
	}
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovChainConfig(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovChainConfig(uint64(l))
	}
	if m.NativeDecimals != 0 {
		n += 1 + sovChainConfig(uint64(m.NativeDecimals))
	}
	if m.ChainType != 0 {
		n += 1 + sovChainConfig(uint64(m.ChainType))
	}
	if m.ConfirmationsRequired != 0 {
		n += 1 + sovChainConfig(uint64(m.ConfirmationsRequired))
	}
	if m.LastObservedValsetNonce != 0 {
		n += 1 + sovChainConfig(uint64(m.LastObservedValsetNonce))
	}
	if m.ChainEnabled {
		n += 2
	}
	if m.IsUnorderedNonces {
		n += 2
	}
	return n
}

func sovChainConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChainConfig(x uint64) (n int) {
	return sovChainConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContractConfig) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ContractConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
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
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
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
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractHeight", wireType)
			}
			m.ContractHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ContractHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastObservedEventBlockHeight", wireType)
			}
			m.LastObservedEventBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastObservedEventBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractType", wireType)
			}
			m.ContractType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ContractType |= ContractType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimSlashingEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
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
			m.ClaimSlashingEnabled = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimSlashingWindow", wireType)
			}
			m.ClaimSlashingWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimSlashingWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionMissingClaim", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
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
				return ErrInvalidLengthChainConfig
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChainConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionMissingClaim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionConflictingClaim", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
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
				return ErrInvalidLengthChainConfig
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChainConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionConflictingClaim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
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
			m.ContractEnabled = bool(v != 0)
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
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
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
				return fmt.Errorf("proto: wrong wireType = %d for field NativeDecimals", wireType)
			}
			m.NativeDecimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NativeDecimals |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
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
		case 6:
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
		case 7:
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
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
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
			m.ChainEnabled = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsUnorderedNonces", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainConfig
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
			m.IsUnorderedNonces = bool(v != 0)
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
