// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: inbound/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgInboundRequest struct {
	Orchestrator         string                                 `protobuf:"bytes,1,opt,name=orchestrator,proto3" json:"orchestrator,omitempty"`
	ChainType            types.ChainType                        `protobuf:"varint,2,opt,name=chainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"chainType,omitempty"`
	ChainId              string                                 `protobuf:"bytes,3,opt,name=chainId,proto3" json:"chainId,omitempty"`
	EventNonce           uint64                                 `protobuf:"varint,4,opt,name=eventNonce,proto3" json:"eventNonce,omitempty"`
	BlockHeight          uint64                                 `protobuf:"varint,5,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	SourceTxHash         string                                 `protobuf:"bytes,6,opt,name=sourceTxHash,proto3" json:"sourceTxHash,omitempty"`
	SourceTimestamp      uint64                                 `protobuf:"varint,7,opt,name=sourceTimestamp,proto3" json:"sourceTimestamp,omitempty"`
	SourceSender         []byte                                 `protobuf:"bytes,8,opt,name=sourceSender,proto3" json:"sourceSender,omitempty"`
	RouterBridgeContract string                                 `protobuf:"bytes,9,opt,name=routerBridgeContract,proto3" json:"routerBridgeContract,omitempty"`
	Payload              []byte                                 `protobuf:"bytes,10,opt,name=payload,proto3" json:"payload,omitempty"`
	GasLimit             uint64                                 `protobuf:"varint,11,opt,name=gasLimit,proto3" json:"gasLimit,omitempty"`
	RouteAmount          github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,12,opt,name=routeAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"routeAmount"`
	RouteRecipient       string                                 `protobuf:"bytes,13,opt,name=routeRecipient,proto3" json:"routeRecipient,omitempty"`
	AsmAddress           []byte                                 `protobuf:"bytes,14,opt,name=asmAddress,proto3" json:"asmAddress,omitempty"`
}

func (m *MsgInboundRequest) Reset()         { *m = MsgInboundRequest{} }
func (m *MsgInboundRequest) String() string { return proto.CompactTextString(m) }
func (*MsgInboundRequest) ProtoMessage()    {}
func (*MsgInboundRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e828bdc3262cd782, []int{0}
}
func (m *MsgInboundRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInboundRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInboundRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInboundRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInboundRequest.Merge(m, src)
}
func (m *MsgInboundRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgInboundRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInboundRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInboundRequest proto.InternalMessageInfo

func (m *MsgInboundRequest) GetOrchestrator() string {
	if m != nil {
		return m.Orchestrator
	}
	return ""
}

func (m *MsgInboundRequest) GetChainType() types.ChainType {
	if m != nil {
		return m.ChainType
	}
	return types.CHAIN_TYPE_EVM
}

func (m *MsgInboundRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *MsgInboundRequest) GetEventNonce() uint64 {
	if m != nil {
		return m.EventNonce
	}
	return 0
}

func (m *MsgInboundRequest) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *MsgInboundRequest) GetSourceTxHash() string {
	if m != nil {
		return m.SourceTxHash
	}
	return ""
}

func (m *MsgInboundRequest) GetSourceTimestamp() uint64 {
	if m != nil {
		return m.SourceTimestamp
	}
	return 0
}

func (m *MsgInboundRequest) GetSourceSender() []byte {
	if m != nil {
		return m.SourceSender
	}
	return nil
}

func (m *MsgInboundRequest) GetRouterBridgeContract() string {
	if m != nil {
		return m.RouterBridgeContract
	}
	return ""
}

func (m *MsgInboundRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *MsgInboundRequest) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *MsgInboundRequest) GetRouteRecipient() string {
	if m != nil {
		return m.RouteRecipient
	}
	return ""
}

func (m *MsgInboundRequest) GetAsmAddress() []byte {
	if m != nil {
		return m.AsmAddress
	}
	return nil
}

type MsgInboundRequestResponse struct {
}

func (m *MsgInboundRequestResponse) Reset()         { *m = MsgInboundRequestResponse{} }
func (m *MsgInboundRequestResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInboundRequestResponse) ProtoMessage()    {}
func (*MsgInboundRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e828bdc3262cd782, []int{1}
}
func (m *MsgInboundRequestResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInboundRequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInboundRequestResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInboundRequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInboundRequestResponse.Merge(m, src)
}
func (m *MsgInboundRequestResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInboundRequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInboundRequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInboundRequestResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgInboundRequest)(nil), "routerprotocol.routerchain.inbound.MsgInboundRequest")
	proto.RegisterType((*MsgInboundRequestResponse)(nil), "routerprotocol.routerchain.inbound.MsgInboundRequestResponse")
}

func init() { proto.RegisterFile("inbound/tx.proto", fileDescriptor_e828bdc3262cd782) }

var fileDescriptor_e828bdc3262cd782 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0x58, 0xd9, 0x56, 0xb7, 0x14, 0xb0, 0x76, 0x30, 0x9d, 0x94, 0x55, 0x3d, 0x4c, 0xbd,
	0x2c, 0x61, 0x45, 0x88, 0x13, 0x87, 0x75, 0x97, 0x55, 0x62, 0x13, 0x84, 0x9d, 0xb8, 0x20, 0xc7,
	0x79, 0x4a, 0xad, 0x35, 0x76, 0xb0, 0x9d, 0xa9, 0xfd, 0x03, 0x1c, 0x11, 0x3f, 0x6b, 0xc7, 0x1d,
	0x81, 0xc3, 0x84, 0xda, 0x3f, 0x82, 0xe2, 0x34, 0x6b, 0xd6, 0x21, 0x90, 0x76, 0x69, 0xfd, 0x7d,
	0x7e, 0xfe, 0xde, 0xf7, 0x5e, 0xde, 0x43, 0xcf, 0xb8, 0x08, 0x65, 0x26, 0x22, 0xdf, 0x4c, 0xbd,
	0x54, 0x49, 0x23, 0x71, 0x4f, 0xc9, 0xcc, 0x80, 0xb2, 0x80, 0xc9, 0x89, 0x57, 0x40, 0x36, 0xa6,
	0x5c, 0x78, 0xcb, 0xe0, 0xce, 0x6e, 0x92, 0x4d, 0x0c, 0xb7, 0x9c, 0x6f, 0x7f, 0x3f, 0x9b, 0x59,
	0x0a, 0x85, 0x40, 0xc7, 0x65, 0x52, 0x27, 0x52, 0xfb, 0x21, 0xd5, 0xe0, 0x5f, 0x1e, 0x86, 0x60,
	0xe8, 0xa1, 0xcf, 0x24, 0x17, 0xcb, 0xfb, 0x9d, 0x58, 0xc6, 0xd2, 0x1e, 0xfd, 0xfc, 0x54, 0xb0,
	0xbd, 0x9f, 0x75, 0xf4, 0xfc, 0x54, 0xc7, 0xa3, 0x22, 0x43, 0x00, 0x5f, 0x32, 0xd0, 0x06, 0xf7,
	0x50, 0x4b, 0x2a, 0x36, 0x06, 0x6d, 0x14, 0x35, 0x52, 0x11, 0xa7, 0xeb, 0xf4, 0x1b, 0xc1, 0x1d,
	0x0e, 0x9f, 0xa1, 0x86, 0xf5, 0x70, 0x3e, 0x4b, 0x81, 0x3c, 0xea, 0x3a, 0xfd, 0xf6, 0xe0, 0xa5,
	0xf7, 0x8f, 0x22, 0x56, 0xde, 0xbd, 0xe3, 0xf2, 0x5d, 0xb0, 0x92, 0xc0, 0x04, 0x6d, 0x59, 0x30,
	0x8a, 0xc8, 0x86, 0x4d, 0x57, 0x42, 0xec, 0x22, 0x04, 0x97, 0x20, 0xcc, 0x99, 0x14, 0x0c, 0x48,
	0xbd, 0xeb, 0xf4, 0xeb, 0x41, 0x85, 0xc1, 0x5d, 0xd4, 0x0c, 0x27, 0x92, 0x5d, 0x9c, 0x00, 0x8f,
	0xc7, 0x86, 0x3c, 0xb6, 0x01, 0x55, 0x2a, 0xaf, 0x47, 0xcb, 0x4c, 0x31, 0x38, 0x9f, 0x9e, 0x50,
	0x3d, 0x26, 0x9b, 0x45, 0x3d, 0x55, 0x0e, 0xf7, 0xd1, 0xd3, 0x25, 0xe6, 0x09, 0x68, 0x43, 0x93,
	0x94, 0x6c, 0x59, 0xa5, 0x75, 0x7a, 0xa5, 0xf6, 0x11, 0x44, 0x04, 0x8a, 0x6c, 0x77, 0x9d, 0x7e,
	0x2b, 0xb8, 0xc3, 0xe1, 0x01, 0xda, 0x29, 0x8a, 0x1f, 0x2a, 0x1e, 0xc5, 0x70, 0x2c, 0x85, 0x51,
	0x94, 0x19, 0xd2, 0xb0, 0x99, 0xff, 0x7a, 0x97, 0x77, 0x20, 0xa5, 0xb3, 0x89, 0xa4, 0x11, 0x41,
	0x56, 0xb2, 0x84, 0xb8, 0x83, 0xb6, 0x63, 0xaa, 0xdf, 0xf1, 0x84, 0x1b, 0xd2, 0xb4, 0xa6, 0x6e,
	0x31, 0x7e, 0x8f, 0x9a, 0x56, 0xed, 0x28, 0x91, 0x99, 0x30, 0xa4, 0x95, 0x27, 0x18, 0x7a, 0x57,
	0x37, 0x7b, 0xb5, 0x5f, 0x37, 0x7b, 0xfb, 0x31, 0x37, 0xe3, 0x2c, 0xf4, 0x98, 0x4c, 0xfc, 0xe5,
	0x7c, 0x14, 0x7f, 0x07, 0x3a, 0xba, 0xf0, 0xf3, 0xe9, 0xd1, 0xde, 0x48, 0x98, 0xa0, 0x2a, 0x81,
	0xf7, 0x51, 0xdb, 0xc2, 0x00, 0x18, 0x4f, 0x39, 0x08, 0x43, 0x9e, 0x58, 0xd7, 0x6b, 0x6c, 0xfe,
	0x5d, 0xa8, 0x4e, 0x8e, 0xa2, 0x48, 0x81, 0xd6, 0xa4, 0x6d, 0x2d, 0x57, 0x98, 0xde, 0x2e, 0x7a,
	0x71, 0x6f, 0xb4, 0x02, 0xd0, 0xa9, 0x14, 0x1a, 0x06, 0xdf, 0x1c, 0xb4, 0x71, 0xaa, 0x63, 0xfc,
	0xd5, 0x41, 0xed, 0xb5, 0xe9, 0x7b, 0xed, 0xfd, 0x7f, 0x17, 0xbc, 0x7b, 0xca, 0x9d, 0xb7, 0x0f,
	0x7a, 0x56, 0x1a, 0x1a, 0x7e, 0xb8, 0x9a, 0xbb, 0xce, 0xf5, 0xdc, 0x75, 0x7e, 0xcf, 0x5d, 0xe7,
	0xfb, 0xc2, 0xad, 0x5d, 0x2f, 0xdc, 0xda, 0x8f, 0x85, 0x5b, 0xfb, 0xf4, 0xa6, 0xd2, 0xc4, 0x42,
	0xf3, 0xa0, 0xcc, 0x51, 0xe2, 0x62, 0x27, 0xa7, 0xfe, 0xed, 0x5a, 0xe7, 0x9d, 0x0d, 0x37, 0x6d,
	0xd8, 0xab, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x70, 0xc2, 0x2a, 0x81, 0xee, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	InboundRequest(ctx context.Context, in *MsgInboundRequest, opts ...grpc.CallOption) (*MsgInboundRequestResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) InboundRequest(ctx context.Context, in *MsgInboundRequest, opts ...grpc.CallOption) (*MsgInboundRequestResponse, error) {
	out := new(MsgInboundRequestResponse)
	err := c.cc.Invoke(ctx, "/routerprotocol.routerchain.inbound.Msg/InboundRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	InboundRequest(context.Context, *MsgInboundRequest) (*MsgInboundRequestResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) InboundRequest(ctx context.Context, req *MsgInboundRequest) (*MsgInboundRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InboundRequest not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_InboundRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInboundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InboundRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routerprotocol.routerchain.inbound.Msg/InboundRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InboundRequest(ctx, req.(*MsgInboundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routerprotocol.routerchain.inbound.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InboundRequest",
			Handler:    _Msg_InboundRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inbound/tx.proto",
}

func (m *MsgInboundRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInboundRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInboundRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AsmAddress) > 0 {
		i -= len(m.AsmAddress)
		copy(dAtA[i:], m.AsmAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AsmAddress)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.RouteRecipient) > 0 {
		i -= len(m.RouteRecipient)
		copy(dAtA[i:], m.RouteRecipient)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RouteRecipient)))
		i--
		dAtA[i] = 0x6a
	}
	{
		size := m.RouteAmount.Size()
		i -= size
		if _, err := m.RouteAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if m.GasLimit != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x58
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.RouterBridgeContract) > 0 {
		i -= len(m.RouterBridgeContract)
		copy(dAtA[i:], m.RouterBridgeContract)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RouterBridgeContract)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.SourceSender) > 0 {
		i -= len(m.SourceSender)
		copy(dAtA[i:], m.SourceSender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceSender)))
		i--
		dAtA[i] = 0x42
	}
	if m.SourceTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.SourceTimestamp))
		i--
		dAtA[i] = 0x38
	}
	if len(m.SourceTxHash) > 0 {
		i -= len(m.SourceTxHash)
		copy(dAtA[i:], m.SourceTxHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceTxHash)))
		i--
		dAtA[i] = 0x32
	}
	if m.BlockHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.EventNonce != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.EventNonce))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ChainType != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ChainType))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Orchestrator) > 0 {
		i -= len(m.Orchestrator)
		copy(dAtA[i:], m.Orchestrator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Orchestrator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInboundRequestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInboundRequestResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInboundRequestResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgInboundRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Orchestrator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ChainType != 0 {
		n += 1 + sovTx(uint64(m.ChainType))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.EventNonce != 0 {
		n += 1 + sovTx(uint64(m.EventNonce))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTx(uint64(m.BlockHeight))
	}
	l = len(m.SourceTxHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.SourceTimestamp != 0 {
		n += 1 + sovTx(uint64(m.SourceTimestamp))
	}
	l = len(m.SourceSender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RouterBridgeContract)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GasLimit != 0 {
		n += 1 + sovTx(uint64(m.GasLimit))
	}
	l = m.RouteAmount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.RouteRecipient)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.AsmAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgInboundRequestResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInboundRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInboundRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInboundRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Orchestrator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Orchestrator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainType", wireType)
			}
			m.ChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainType |= types.ChainType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventNonce", wireType)
			}
			m.EventNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceTimestamp", wireType)
			}
			m.SourceTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SourceTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceSender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceSender = append(m.SourceSender[:0], dAtA[iNdEx:postIndex]...)
			if m.SourceSender == nil {
				m.SourceSender = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouterBridgeContract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RouterBridgeContract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouteAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RouteAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouteRecipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RouteRecipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AsmAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AsmAddress = append(m.AsmAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.AsmAddress == nil {
				m.AsmAddress = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgInboundRequestResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInboundRequestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInboundRequestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
