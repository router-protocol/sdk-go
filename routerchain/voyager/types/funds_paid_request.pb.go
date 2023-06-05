// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerchain/voyager/funds_paid_request.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/router-protocol/sdk-go/routerchain/multichain/types"
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

type FundsPaidRequest struct {
	SrcChainId          string          `protobuf:"bytes,1,opt,name=srcChainId,proto3" json:"srcChainId,omitempty"`
	SrcChainType        types.ChainType `protobuf:"varint,2,opt,name=srcChainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"srcChainType,omitempty"`
	SrcTxHash           string          `protobuf:"bytes,3,opt,name=srcTxHash,proto3" json:"srcTxHash,omitempty"`
	SrcTimestamp        uint64          `protobuf:"varint,4,opt,name=srcTimestamp,proto3" json:"srcTimestamp,omitempty"`
	Contract            string          `protobuf:"bytes,5,opt,name=contract,proto3" json:"contract,omitempty"`
	EventNonce          uint64          `protobuf:"varint,6,opt,name=eventNonce,proto3" json:"eventNonce,omitempty"`
	BlockHeight         uint64          `protobuf:"varint,7,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	MessageHash         []byte          `protobuf:"bytes,8,opt,name=messageHash,proto3" json:"messageHash,omitempty"`
	Forwarder           string          `protobuf:"bytes,9,opt,name=forwarder,proto3" json:"forwarder,omitempty"`
	ForwarderRouterAddr []byte          `protobuf:"bytes,10,opt,name=forwarderRouterAddr,proto3" json:"forwarderRouterAddr,omitempty"`
	Status              string          `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *FundsPaidRequest) Reset()         { *m = FundsPaidRequest{} }
func (m *FundsPaidRequest) String() string { return proto.CompactTextString(m) }
func (*FundsPaidRequest) ProtoMessage()    {}
func (*FundsPaidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d80dcc5f7c208491, []int{0}
}
func (m *FundsPaidRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FundsPaidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FundsPaidRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FundsPaidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundsPaidRequest.Merge(m, src)
}
func (m *FundsPaidRequest) XXX_Size() int {
	return m.Size()
}
func (m *FundsPaidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FundsPaidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FundsPaidRequest proto.InternalMessageInfo

func (m *FundsPaidRequest) GetSrcChainId() string {
	if m != nil {
		return m.SrcChainId
	}
	return ""
}

func (m *FundsPaidRequest) GetSrcChainType() types.ChainType {
	if m != nil {
		return m.SrcChainType
	}
	return types.CHAIN_TYPE_NONE
}

func (m *FundsPaidRequest) GetSrcTxHash() string {
	if m != nil {
		return m.SrcTxHash
	}
	return ""
}

func (m *FundsPaidRequest) GetSrcTimestamp() uint64 {
	if m != nil {
		return m.SrcTimestamp
	}
	return 0
}

func (m *FundsPaidRequest) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *FundsPaidRequest) GetEventNonce() uint64 {
	if m != nil {
		return m.EventNonce
	}
	return 0
}

func (m *FundsPaidRequest) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *FundsPaidRequest) GetMessageHash() []byte {
	if m != nil {
		return m.MessageHash
	}
	return nil
}

func (m *FundsPaidRequest) GetForwarder() string {
	if m != nil {
		return m.Forwarder
	}
	return ""
}

func (m *FundsPaidRequest) GetForwarderRouterAddr() []byte {
	if m != nil {
		return m.ForwarderRouterAddr
	}
	return nil
}

func (m *FundsPaidRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type FundsPaidRequestClaimHash struct {
	SrcChainId          string          `protobuf:"bytes,1,opt,name=srcChainId,proto3" json:"srcChainId,omitempty"`
	SrcChainType        types.ChainType `protobuf:"varint,2,opt,name=srcChainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"srcChainType,omitempty"`
	SrcTxHash           string          `protobuf:"bytes,3,opt,name=srcTxHash,proto3" json:"srcTxHash,omitempty"`
	SrcTimestamp        uint64          `protobuf:"varint,4,opt,name=srcTimestamp,proto3" json:"srcTimestamp,omitempty"`
	Contract            string          `protobuf:"bytes,5,opt,name=contract,proto3" json:"contract,omitempty"`
	EventNonce          uint64          `protobuf:"varint,6,opt,name=eventNonce,proto3" json:"eventNonce,omitempty"`
	BlockHeight         uint64          `protobuf:"varint,7,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	MessageHash         []byte          `protobuf:"bytes,8,opt,name=messageHash,proto3" json:"messageHash,omitempty"`
	Forwarder           string          `protobuf:"bytes,9,opt,name=forwarder,proto3" json:"forwarder,omitempty"`
	ForwarderRouterAddr []byte          `protobuf:"bytes,10,opt,name=forwarderRouterAddr,proto3" json:"forwarderRouterAddr,omitempty"`
}

func (m *FundsPaidRequestClaimHash) Reset()         { *m = FundsPaidRequestClaimHash{} }
func (m *FundsPaidRequestClaimHash) String() string { return proto.CompactTextString(m) }
func (*FundsPaidRequestClaimHash) ProtoMessage()    {}
func (*FundsPaidRequestClaimHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_d80dcc5f7c208491, []int{1}
}
func (m *FundsPaidRequestClaimHash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FundsPaidRequestClaimHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FundsPaidRequestClaimHash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FundsPaidRequestClaimHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundsPaidRequestClaimHash.Merge(m, src)
}
func (m *FundsPaidRequestClaimHash) XXX_Size() int {
	return m.Size()
}
func (m *FundsPaidRequestClaimHash) XXX_DiscardUnknown() {
	xxx_messageInfo_FundsPaidRequestClaimHash.DiscardUnknown(m)
}

var xxx_messageInfo_FundsPaidRequestClaimHash proto.InternalMessageInfo

func (m *FundsPaidRequestClaimHash) GetSrcChainId() string {
	if m != nil {
		return m.SrcChainId
	}
	return ""
}

func (m *FundsPaidRequestClaimHash) GetSrcChainType() types.ChainType {
	if m != nil {
		return m.SrcChainType
	}
	return types.CHAIN_TYPE_NONE
}

func (m *FundsPaidRequestClaimHash) GetSrcTxHash() string {
	if m != nil {
		return m.SrcTxHash
	}
	return ""
}

func (m *FundsPaidRequestClaimHash) GetSrcTimestamp() uint64 {
	if m != nil {
		return m.SrcTimestamp
	}
	return 0
}

func (m *FundsPaidRequestClaimHash) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *FundsPaidRequestClaimHash) GetEventNonce() uint64 {
	if m != nil {
		return m.EventNonce
	}
	return 0
}

func (m *FundsPaidRequestClaimHash) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *FundsPaidRequestClaimHash) GetMessageHash() []byte {
	if m != nil {
		return m.MessageHash
	}
	return nil
}

func (m *FundsPaidRequestClaimHash) GetForwarder() string {
	if m != nil {
		return m.Forwarder
	}
	return ""
}

func (m *FundsPaidRequestClaimHash) GetForwarderRouterAddr() []byte {
	if m != nil {
		return m.ForwarderRouterAddr
	}
	return nil
}

func init() {
	proto.RegisterType((*FundsPaidRequest)(nil), "routerprotocol.routerchain.voyager.FundsPaidRequest")
	proto.RegisterType((*FundsPaidRequestClaimHash)(nil), "routerprotocol.routerchain.voyager.FundsPaidRequestClaimHash")
}

func init() {
	proto.RegisterFile("routerchain/voyager/funds_paid_request.proto", fileDescriptor_d80dcc5f7c208491)
}

var fileDescriptor_d80dcc5f7c208491 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0xad, 0xe9, 0x52, 0xb6, 0xde, 0x15, 0x42, 0x06, 0x21, 0x53, 0x90, 0x15, 0xf5, 0xd4, 0x03,
	0x24, 0xbb, 0x70, 0xe0, 0x0c, 0x2b, 0xa1, 0xe5, 0x82, 0x20, 0xea, 0x89, 0x4b, 0xe5, 0x38, 0xb3,
	0x89, 0x45, 0x12, 0x07, 0xdb, 0x29, 0xdb, 0xbf, 0xe0, 0x1f, 0xf8, 0x03, 0xbe, 0x82, 0xe3, 0x1e,
	0x39, 0xa2, 0xf6, 0x47, 0x50, 0xec, 0x36, 0x9b, 0x22, 0xc4, 0x99, 0xc3, 0x5e, 0x22, 0xbf, 0xe7,
	0x99, 0x37, 0x6f, 0x66, 0xe4, 0xe0, 0xa7, 0x5a, 0x35, 0x16, 0xb4, 0xc8, 0xb9, 0xac, 0xa2, 0xa5,
	0x5a, 0xf1, 0x0c, 0x74, 0x74, 0xd1, 0x54, 0xa9, 0x59, 0xd4, 0x5c, 0xa6, 0x0b, 0x0d, 0x9f, 0x1b,
	0x30, 0x36, 0xac, 0xb5, 0xb2, 0x8a, 0x4c, 0x7d, 0xb4, 0x03, 0x42, 0x15, 0x61, 0x2f, 0x39, 0xdc,
	0x26, 0x4f, 0x1e, 0x97, 0x4d, 0x61, 0xa5, 0x17, 0x74, 0xdf, 0x85, 0x5d, 0xd5, 0xe0, 0x05, 0x26,
	0x4c, 0x28, 0x53, 0x2a, 0x13, 0x25, 0xdc, 0x40, 0xb4, 0x3c, 0x4d, 0xc0, 0xf2, 0xd3, 0x48, 0x28,
	0x59, 0x6d, 0xef, 0x1f, 0x64, 0x2a, 0x53, 0xee, 0x18, 0xb5, 0x27, 0xcf, 0x4e, 0xbf, 0x0f, 0xf1,
	0xbd, 0x37, 0xad, 0xa7, 0xf7, 0x5c, 0xa6, 0xb1, 0x77, 0x44, 0x18, 0xc6, 0x46, 0x8b, 0xb3, 0xb6,
	0xc2, 0xdb, 0x94, 0xa2, 0x00, 0xcd, 0xc6, 0x71, 0x8f, 0x21, 0x73, 0x7c, 0xbc, 0x43, 0xf3, 0x55,
	0x0d, 0xf4, 0x56, 0x80, 0x66, 0x77, 0x9f, 0x9f, 0x84, 0xff, 0x68, 0xe1, 0xda, 0x79, 0xd8, 0xe5,
	0xc5, 0x7b, 0x2a, 0xe4, 0x09, 0x1e, 0x1b, 0x2d, 0xe6, 0x97, 0xe7, 0xdc, 0xe4, 0x74, 0xe8, 0x8a,
	0x5e, 0x13, 0x64, 0xea, 0x6a, 0xce, 0x65, 0x09, 0xc6, 0xf2, 0xb2, 0xa6, 0x07, 0x01, 0x9a, 0x1d,
	0xc4, 0x7b, 0x1c, 0x99, 0xe0, 0x43, 0xa1, 0x2a, 0xab, 0xb9, 0xb0, 0xf4, 0xb6, 0x13, 0xe8, 0x70,
	0xdb, 0x13, 0x2c, 0xa1, 0xb2, 0xef, 0x54, 0x25, 0x80, 0x8e, 0x5c, 0x76, 0x8f, 0x21, 0x01, 0x3e,
	0x4a, 0x0a, 0x25, 0x3e, 0x9d, 0x83, 0xcc, 0x72, 0x4b, 0xef, 0xb8, 0x80, 0x3e, 0xd5, 0x46, 0x94,
	0x60, 0x0c, 0xcf, 0xc0, 0x39, 0x3c, 0x0c, 0xd0, 0xec, 0x38, 0xee, 0x53, 0x6d, 0x07, 0x17, 0x4a,
	0x7f, 0xe1, 0x3a, 0x05, 0x4d, 0xc7, 0xbe, 0x83, 0x8e, 0x20, 0x27, 0xf8, 0x7e, 0x07, 0x62, 0x37,
	0x9a, 0x57, 0x69, 0xaa, 0x29, 0x76, 0x3a, 0x7f, 0xbb, 0x22, 0x0f, 0xf1, 0xc8, 0x58, 0x6e, 0x1b,
	0x43, 0x8f, 0x9c, 0xd8, 0x16, 0x4d, 0xbf, 0x0d, 0xf1, 0xa3, 0x3f, 0x97, 0x76, 0x56, 0x70, 0x59,
	0x3a, 0x17, 0x37, 0xdb, 0xfb, 0x2f, 0xb6, 0xf7, 0xfa, 0xc3, 0x8f, 0x35, 0x43, 0x57, 0x6b, 0x86,
	0x7e, 0xad, 0x19, 0xfa, 0xba, 0x61, 0x83, 0xab, 0x0d, 0x1b, 0xfc, 0xdc, 0xb0, 0xc1, 0xc7, 0x97,
	0x99, 0xb4, 0x79, 0x93, 0x84, 0x42, 0x95, 0x91, 0x1f, 0xf3, 0xb3, 0xdd, 0xd8, 0x77, 0xd8, 0x3f,
	0xf2, 0xcb, 0xee, 0xbf, 0xd1, 0x3e, 0x74, 0x93, 0x8c, 0x5c, 0xd8, 0x8b, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x1d, 0xbf, 0x03, 0x8d, 0x5b, 0x04, 0x00, 0x00,
}

func (m *FundsPaidRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FundsPaidRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FundsPaidRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.ForwarderRouterAddr) > 0 {
		i -= len(m.ForwarderRouterAddr)
		copy(dAtA[i:], m.ForwarderRouterAddr)
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(len(m.ForwarderRouterAddr)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Forwarder) > 0 {
		i -= len(m.Forwarder)
		copy(dAtA[i:], m.Forwarder)
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(len(m.Forwarder)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.MessageHash) > 0 {
		i -= len(m.MessageHash)
		copy(dAtA[i:], m.MessageHash)
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(len(m.MessageHash)))
		i--
		dAtA[i] = 0x42
	}
	if m.BlockHeight != 0 {
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x38
	}
	if m.EventNonce != 0 {
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(m.EventNonce))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x2a
	}
	if m.SrcTimestamp != 0 {
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(m.SrcTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.SrcTxHash) > 0 {
		i -= len(m.SrcTxHash)
		copy(dAtA[i:], m.SrcTxHash)
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(len(m.SrcTxHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SrcChainType != 0 {
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(m.SrcChainType))
		i--
		dAtA[i] = 0x10
	}
	if len(m.SrcChainId) > 0 {
		i -= len(m.SrcChainId)
		copy(dAtA[i:], m.SrcChainId)
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(len(m.SrcChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FundsPaidRequestClaimHash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FundsPaidRequestClaimHash) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FundsPaidRequestClaimHash) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ForwarderRouterAddr) > 0 {
		i -= len(m.ForwarderRouterAddr)
		copy(dAtA[i:], m.ForwarderRouterAddr)
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(len(m.ForwarderRouterAddr)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Forwarder) > 0 {
		i -= len(m.Forwarder)
		copy(dAtA[i:], m.Forwarder)
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(len(m.Forwarder)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.MessageHash) > 0 {
		i -= len(m.MessageHash)
		copy(dAtA[i:], m.MessageHash)
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(len(m.MessageHash)))
		i--
		dAtA[i] = 0x42
	}
	if m.BlockHeight != 0 {
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x38
	}
	if m.EventNonce != 0 {
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(m.EventNonce))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x2a
	}
	if m.SrcTimestamp != 0 {
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(m.SrcTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.SrcTxHash) > 0 {
		i -= len(m.SrcTxHash)
		copy(dAtA[i:], m.SrcTxHash)
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(len(m.SrcTxHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SrcChainType != 0 {
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(m.SrcChainType))
		i--
		dAtA[i] = 0x10
	}
	if len(m.SrcChainId) > 0 {
		i -= len(m.SrcChainId)
		copy(dAtA[i:], m.SrcChainId)
		i = encodeVarintFundsPaidRequest(dAtA, i, uint64(len(m.SrcChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFundsPaidRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovFundsPaidRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FundsPaidRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SrcChainId)
	if l > 0 {
		n += 1 + l + sovFundsPaidRequest(uint64(l))
	}
	if m.SrcChainType != 0 {
		n += 1 + sovFundsPaidRequest(uint64(m.SrcChainType))
	}
	l = len(m.SrcTxHash)
	if l > 0 {
		n += 1 + l + sovFundsPaidRequest(uint64(l))
	}
	if m.SrcTimestamp != 0 {
		n += 1 + sovFundsPaidRequest(uint64(m.SrcTimestamp))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovFundsPaidRequest(uint64(l))
	}
	if m.EventNonce != 0 {
		n += 1 + sovFundsPaidRequest(uint64(m.EventNonce))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovFundsPaidRequest(uint64(m.BlockHeight))
	}
	l = len(m.MessageHash)
	if l > 0 {
		n += 1 + l + sovFundsPaidRequest(uint64(l))
	}
	l = len(m.Forwarder)
	if l > 0 {
		n += 1 + l + sovFundsPaidRequest(uint64(l))
	}
	l = len(m.ForwarderRouterAddr)
	if l > 0 {
		n += 1 + l + sovFundsPaidRequest(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovFundsPaidRequest(uint64(l))
	}
	return n
}

func (m *FundsPaidRequestClaimHash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SrcChainId)
	if l > 0 {
		n += 1 + l + sovFundsPaidRequest(uint64(l))
	}
	if m.SrcChainType != 0 {
		n += 1 + sovFundsPaidRequest(uint64(m.SrcChainType))
	}
	l = len(m.SrcTxHash)
	if l > 0 {
		n += 1 + l + sovFundsPaidRequest(uint64(l))
	}
	if m.SrcTimestamp != 0 {
		n += 1 + sovFundsPaidRequest(uint64(m.SrcTimestamp))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovFundsPaidRequest(uint64(l))
	}
	if m.EventNonce != 0 {
		n += 1 + sovFundsPaidRequest(uint64(m.EventNonce))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovFundsPaidRequest(uint64(m.BlockHeight))
	}
	l = len(m.MessageHash)
	if l > 0 {
		n += 1 + l + sovFundsPaidRequest(uint64(l))
	}
	l = len(m.Forwarder)
	if l > 0 {
		n += 1 + l + sovFundsPaidRequest(uint64(l))
	}
	l = len(m.ForwarderRouterAddr)
	if l > 0 {
		n += 1 + l + sovFundsPaidRequest(uint64(l))
	}
	return n
}

func sovFundsPaidRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFundsPaidRequest(x uint64) (n int) {
	return sovFundsPaidRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FundsPaidRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFundsPaidRequest
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
			return fmt.Errorf("proto: FundsPaidRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FundsPaidRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
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
				return ErrInvalidLengthFundsPaidRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFundsPaidRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcChainType", wireType)
			}
			m.SrcChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SrcChainType |= types.ChainType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
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
				return ErrInvalidLengthFundsPaidRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFundsPaidRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcTimestamp", wireType)
			}
			m.SrcTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SrcTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
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
				return ErrInvalidLengthFundsPaidRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFundsPaidRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventNonce", wireType)
			}
			m.EventNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
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
				return ErrInvalidLengthFundsPaidRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFundsPaidRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageHash = append(m.MessageHash[:0], dAtA[iNdEx:postIndex]...)
			if m.MessageHash == nil {
				m.MessageHash = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Forwarder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
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
				return ErrInvalidLengthFundsPaidRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFundsPaidRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Forwarder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForwarderRouterAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
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
				return ErrInvalidLengthFundsPaidRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFundsPaidRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ForwarderRouterAddr = append(m.ForwarderRouterAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.ForwarderRouterAddr == nil {
				m.ForwarderRouterAddr = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
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
				return ErrInvalidLengthFundsPaidRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFundsPaidRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFundsPaidRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFundsPaidRequest
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
func (m *FundsPaidRequestClaimHash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFundsPaidRequest
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
			return fmt.Errorf("proto: FundsPaidRequestClaimHash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FundsPaidRequestClaimHash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
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
				return ErrInvalidLengthFundsPaidRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFundsPaidRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcChainType", wireType)
			}
			m.SrcChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SrcChainType |= types.ChainType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
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
				return ErrInvalidLengthFundsPaidRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFundsPaidRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcTimestamp", wireType)
			}
			m.SrcTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SrcTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
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
				return ErrInvalidLengthFundsPaidRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFundsPaidRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventNonce", wireType)
			}
			m.EventNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
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
				return ErrInvalidLengthFundsPaidRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFundsPaidRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageHash = append(m.MessageHash[:0], dAtA[iNdEx:postIndex]...)
			if m.MessageHash == nil {
				m.MessageHash = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Forwarder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
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
				return ErrInvalidLengthFundsPaidRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFundsPaidRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Forwarder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForwarderRouterAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFundsPaidRequest
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
				return ErrInvalidLengthFundsPaidRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFundsPaidRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ForwarderRouterAddr = append(m.ForwarderRouterAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.ForwarderRouterAddr == nil {
				m.ForwarderRouterAddr = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFundsPaidRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFundsPaidRequest
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
func skipFundsPaidRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFundsPaidRequest
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
					return 0, ErrIntOverflowFundsPaidRequest
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
					return 0, ErrIntOverflowFundsPaidRequest
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
				return 0, ErrInvalidLengthFundsPaidRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFundsPaidRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFundsPaidRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFundsPaidRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFundsPaidRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFundsPaidRequest = fmt.Errorf("proto: unexpected end of group")
)
