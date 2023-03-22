// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metastore/tx.proto

package types

import (
	context "context"
	fmt "fmt"
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

type MsgCreateMetadataRequest struct {
	Orchestrator string          `protobuf:"bytes,1,opt,name=orchestrator,proto3" json:"orchestrator,omitempty"`
	ChainType    types.ChainType `protobuf:"varint,2,opt,name=chainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"chainType,omitempty"`
	ChainId      string          `protobuf:"bytes,3,opt,name=chainId,proto3" json:"chainId,omitempty"`
	EventNonce   uint64          `protobuf:"varint,4,opt,name=eventNonce,proto3" json:"eventNonce,omitempty"`
	BlockHeight  uint64          `protobuf:"varint,5,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	DaapAddress  []byte          `protobuf:"bytes,6,opt,name=daapAddress,proto3" json:"daapAddress,omitempty"`
	FeePayer     string          `protobuf:"bytes,7,opt,name=feePayer,proto3" json:"feePayer,omitempty"`
}

func (m *MsgCreateMetadataRequest) Reset()         { *m = MsgCreateMetadataRequest{} }
func (m *MsgCreateMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*MsgCreateMetadataRequest) ProtoMessage()    {}
func (*MsgCreateMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f44ee1d497dfd26a, []int{0}
}
func (m *MsgCreateMetadataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateMetadataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateMetadataRequest.Merge(m, src)
}
func (m *MsgCreateMetadataRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateMetadataRequest proto.InternalMessageInfo

func (m *MsgCreateMetadataRequest) GetOrchestrator() string {
	if m != nil {
		return m.Orchestrator
	}
	return ""
}

func (m *MsgCreateMetadataRequest) GetChainType() types.ChainType {
	if m != nil {
		return m.ChainType
	}
	return types.CHAIN_TYPE_EVM
}

func (m *MsgCreateMetadataRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *MsgCreateMetadataRequest) GetEventNonce() uint64 {
	if m != nil {
		return m.EventNonce
	}
	return 0
}

func (m *MsgCreateMetadataRequest) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *MsgCreateMetadataRequest) GetDaapAddress() []byte {
	if m != nil {
		return m.DaapAddress
	}
	return nil
}

func (m *MsgCreateMetadataRequest) GetFeePayer() string {
	if m != nil {
		return m.FeePayer
	}
	return ""
}

type MsgCreateMetadataRequestResponse struct {
}

func (m *MsgCreateMetadataRequestResponse) Reset()         { *m = MsgCreateMetadataRequestResponse{} }
func (m *MsgCreateMetadataRequestResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateMetadataRequestResponse) ProtoMessage()    {}
func (*MsgCreateMetadataRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f44ee1d497dfd26a, []int{1}
}
func (m *MsgCreateMetadataRequestResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateMetadataRequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateMetadataRequestResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateMetadataRequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateMetadataRequestResponse.Merge(m, src)
}
func (m *MsgCreateMetadataRequestResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateMetadataRequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateMetadataRequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateMetadataRequestResponse proto.InternalMessageInfo

type MsgApproveFeepayerRequest struct {
	Feepayer      string          `protobuf:"bytes,1,opt,name=feepayer,proto3" json:"feepayer,omitempty"`
	ChainType     types.ChainType `protobuf:"varint,2,opt,name=chainType,proto3,enum=routerprotocol.routerchain.multichain.ChainType" json:"chainType,omitempty"`
	ChainId       string          `protobuf:"bytes,3,opt,name=chainId,proto3" json:"chainId,omitempty"`
	DaapAddresses [][]byte        `protobuf:"bytes,4,rep,name=daapAddresses,proto3" json:"daapAddresses,omitempty"`
}

func (m *MsgApproveFeepayerRequest) Reset()         { *m = MsgApproveFeepayerRequest{} }
func (m *MsgApproveFeepayerRequest) String() string { return proto.CompactTextString(m) }
func (*MsgApproveFeepayerRequest) ProtoMessage()    {}
func (*MsgApproveFeepayerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f44ee1d497dfd26a, []int{2}
}
func (m *MsgApproveFeepayerRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgApproveFeepayerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgApproveFeepayerRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgApproveFeepayerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgApproveFeepayerRequest.Merge(m, src)
}
func (m *MsgApproveFeepayerRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgApproveFeepayerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgApproveFeepayerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgApproveFeepayerRequest proto.InternalMessageInfo

func (m *MsgApproveFeepayerRequest) GetFeepayer() string {
	if m != nil {
		return m.Feepayer
	}
	return ""
}

func (m *MsgApproveFeepayerRequest) GetChainType() types.ChainType {
	if m != nil {
		return m.ChainType
	}
	return types.CHAIN_TYPE_EVM
}

func (m *MsgApproveFeepayerRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *MsgApproveFeepayerRequest) GetDaapAddresses() [][]byte {
	if m != nil {
		return m.DaapAddresses
	}
	return nil
}

type MsgApproveFeepayerRequestResponse struct {
}

func (m *MsgApproveFeepayerRequestResponse) Reset()         { *m = MsgApproveFeepayerRequestResponse{} }
func (m *MsgApproveFeepayerRequestResponse) String() string { return proto.CompactTextString(m) }
func (*MsgApproveFeepayerRequestResponse) ProtoMessage()    {}
func (*MsgApproveFeepayerRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f44ee1d497dfd26a, []int{3}
}
func (m *MsgApproveFeepayerRequestResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgApproveFeepayerRequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgApproveFeepayerRequestResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgApproveFeepayerRequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgApproveFeepayerRequestResponse.Merge(m, src)
}
func (m *MsgApproveFeepayerRequestResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgApproveFeepayerRequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgApproveFeepayerRequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgApproveFeepayerRequestResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateMetadataRequest)(nil), "routerprotocol.routerchain.metastore.MsgCreateMetadataRequest")
	proto.RegisterType((*MsgCreateMetadataRequestResponse)(nil), "routerprotocol.routerchain.metastore.MsgCreateMetadataRequestResponse")
	proto.RegisterType((*MsgApproveFeepayerRequest)(nil), "routerprotocol.routerchain.metastore.MsgApproveFeepayerRequest")
	proto.RegisterType((*MsgApproveFeepayerRequestResponse)(nil), "routerprotocol.routerchain.metastore.MsgApproveFeepayerRequestResponse")
}

func init() { proto.RegisterFile("metastore/tx.proto", fileDescriptor_f44ee1d497dfd26a) }

var fileDescriptor_f44ee1d497dfd26a = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xb1, 0x6e, 0x13, 0x41,
	0x10, 0xf5, 0xda, 0x26, 0x21, 0x83, 0xa1, 0x58, 0x09, 0xb4, 0x18, 0xe9, 0x74, 0x1c, 0x29, 0xae,
	0x61, 0x8d, 0x42, 0x45, 0x03, 0x0a, 0x91, 0x02, 0x14, 0x8e, 0xd0, 0x41, 0x45, 0x83, 0xd6, 0x77,
	0xc3, 0xd9, 0xc2, 0xf6, 0x2e, 0xbb, 0xe3, 0x28, 0xfe, 0x0b, 0xfe, 0x00, 0x89, 0x8a, 0x3f, 0x81,
	0x82, 0x22, 0x25, 0x25, 0xb2, 0x7f, 0x04, 0x65, 0x8f, 0x35, 0x17, 0x29, 0x46, 0x91, 0x2c, 0xa5,
	0x39, 0xe9, 0x3d, 0xcd, 0xbc, 0x79, 0xf3, 0x6e, 0xee, 0x80, 0x4f, 0x90, 0x94, 0x23, 0x6d, 0xb1,
	0x47, 0x27, 0xd2, 0x58, 0x4d, 0x9a, 0xef, 0x5a, 0x3d, 0x23, 0xb4, 0x1e, 0xe4, 0x7a, 0x2c, 0x2b,
	0x98, 0x0f, 0xd5, 0x68, 0x2a, 0x57, 0xe5, 0xdd, 0x7b, 0x93, 0xd9, 0x98, 0x46, 0x9e, 0xed, 0xf9,
	0xe7, 0x7b, 0x9a, 0x1b, 0xac, 0x24, 0x92, 0x6f, 0x4d, 0x10, 0x7d, 0x57, 0x1e, 0x58, 0x54, 0x84,
	0x7d, 0x24, 0x55, 0x28, 0x52, 0x19, 0x7e, 0x9a, 0xa1, 0x23, 0x9e, 0x40, 0x47, 0xdb, 0x7c, 0x88,
	0x8e, 0xac, 0x22, 0x6d, 0x05, 0x8b, 0x59, 0xba, 0x93, 0x9d, 0xe3, 0xf8, 0x11, 0xec, 0x78, 0xd1,
	0xb7, 0x73, 0x83, 0xa2, 0x19, 0xb3, 0xf4, 0xd6, 0xde, 0x23, 0xf9, 0x3f, 0x5f, 0x2b, 0x33, 0xf2,
	0x20, 0xf4, 0x65, 0xff, 0x24, 0xb8, 0x80, 0x6d, 0x0f, 0x5e, 0x15, 0xa2, 0xe5, 0xc7, 0x05, 0xc8,
	0x23, 0x00, 0x3c, 0xc6, 0x29, 0x1d, 0xe9, 0x69, 0x8e, 0xa2, 0x1d, 0xb3, 0xb4, 0x9d, 0xd5, 0x18,
	0x1e, 0xc3, 0x8d, 0xc1, 0x58, 0xe7, 0x1f, 0x5f, 0xe2, 0xa8, 0x1c, 0x92, 0xb8, 0xe6, 0x0b, 0xea,
	0xd4, 0x59, 0x45, 0xa1, 0x94, 0xd9, 0x2f, 0x0a, 0x8b, 0xce, 0x89, 0xad, 0x98, 0xa5, 0x9d, 0xac,
	0x4e, 0xf1, 0x2e, 0x5c, 0xff, 0x80, 0xf8, 0x5a, 0xcd, 0xd1, 0x8a, 0x6d, 0x3f, 0x7e, 0x85, 0x93,
	0x04, 0xe2, 0x75, 0x49, 0x65, 0xe8, 0x8c, 0x9e, 0x3a, 0x4c, 0xbe, 0x33, 0xb8, 0xdb, 0x77, 0xe5,
	0xbe, 0x31, 0x56, 0x1f, 0xe3, 0x21, 0xa2, 0x39, 0x6b, 0x0d, 0x79, 0x56, 0xea, 0x9e, 0xfa, 0x9b,
	0xe5, 0x0a, 0x5f, 0x61, 0x8e, 0xbb, 0x70, 0xb3, 0xb6, 0x32, 0x3a, 0xd1, 0x8e, 0x5b, 0x69, 0x27,
	0x3b, 0x4f, 0x26, 0x0f, 0xe0, 0xfe, 0xda, 0x45, 0xc2, 0xba, 0x7b, 0x3f, 0x9b, 0xd0, 0xea, 0xbb,
	0x92, 0x7f, 0x61, 0x70, 0xfb, 0xe2, 0x13, 0x7a, 0x2a, 0x2f, 0x73, 0xa3, 0x72, 0x5d, 0xb0, 0xdd,
	0xc3, 0xcd, 0xfa, 0x83, 0x53, 0xfe, 0x95, 0xc1, 0x9d, 0x35, 0x6f, 0xe5, 0xd9, 0xa5, 0x47, 0x5c,
	0x2c, 0xd0, 0x7d, 0xb1, 0xa1, 0x40, 0x30, 0xf9, 0xfc, 0xcd, 0x8f, 0x45, 0xc4, 0x4e, 0x17, 0x11,
	0xfb, 0xbd, 0x88, 0xd8, 0xe7, 0x65, 0xd4, 0x38, 0x5d, 0x46, 0x8d, 0x5f, 0xcb, 0xa8, 0xf1, 0xee,
	0x49, 0x39, 0xa2, 0xe1, 0x6c, 0x20, 0x73, 0x3d, 0xe9, 0x55, 0xea, 0x0f, 0xc3, 0xb4, 0x80, 0xab,
	0x0f, 0xfc, 0xa4, 0x57, 0xfb, 0x4f, 0xcc, 0x0d, 0xba, 0xc1, 0x96, 0x2f, 0x7c, 0xfc, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xbc, 0x2c, 0x69, 0x26, 0x41, 0x04, 0x00, 0x00,
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
	CreateMetadataRequest(ctx context.Context, in *MsgCreateMetadataRequest, opts ...grpc.CallOption) (*MsgCreateMetadataRequestResponse, error)
	ApproveFeepayerRequest(ctx context.Context, in *MsgApproveFeepayerRequest, opts ...grpc.CallOption) (*MsgApproveFeepayerRequestResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateMetadataRequest(ctx context.Context, in *MsgCreateMetadataRequest, opts ...grpc.CallOption) (*MsgCreateMetadataRequestResponse, error) {
	out := new(MsgCreateMetadataRequestResponse)
	err := c.cc.Invoke(ctx, "/routerprotocol.routerchain.metastore.Msg/CreateMetadataRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ApproveFeepayerRequest(ctx context.Context, in *MsgApproveFeepayerRequest, opts ...grpc.CallOption) (*MsgApproveFeepayerRequestResponse, error) {
	out := new(MsgApproveFeepayerRequestResponse)
	err := c.cc.Invoke(ctx, "/routerprotocol.routerchain.metastore.Msg/ApproveFeepayerRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateMetadataRequest(context.Context, *MsgCreateMetadataRequest) (*MsgCreateMetadataRequestResponse, error)
	ApproveFeepayerRequest(context.Context, *MsgApproveFeepayerRequest) (*MsgApproveFeepayerRequestResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateMetadataRequest(ctx context.Context, req *MsgCreateMetadataRequest) (*MsgCreateMetadataRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMetadataRequest not implemented")
}
func (*UnimplementedMsgServer) ApproveFeepayerRequest(ctx context.Context, req *MsgApproveFeepayerRequest) (*MsgApproveFeepayerRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveFeepayerRequest not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateMetadataRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateMetadataRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routerprotocol.routerchain.metastore.Msg/CreateMetadataRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateMetadataRequest(ctx, req.(*MsgCreateMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ApproveFeepayerRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgApproveFeepayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ApproveFeepayerRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routerprotocol.routerchain.metastore.Msg/ApproveFeepayerRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ApproveFeepayerRequest(ctx, req.(*MsgApproveFeepayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routerprotocol.routerchain.metastore.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMetadataRequest",
			Handler:    _Msg_CreateMetadataRequest_Handler,
		},
		{
			MethodName: "ApproveFeepayerRequest",
			Handler:    _Msg_ApproveFeepayerRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metastore/tx.proto",
}

func (m *MsgCreateMetadataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateMetadataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateMetadataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeePayer) > 0 {
		i -= len(m.FeePayer)
		copy(dAtA[i:], m.FeePayer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FeePayer)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.DaapAddress) > 0 {
		i -= len(m.DaapAddress)
		copy(dAtA[i:], m.DaapAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DaapAddress)))
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

func (m *MsgCreateMetadataRequestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateMetadataRequestResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateMetadataRequestResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgApproveFeepayerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgApproveFeepayerRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgApproveFeepayerRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DaapAddresses) > 0 {
		for iNdEx := len(m.DaapAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DaapAddresses[iNdEx])
			copy(dAtA[i:], m.DaapAddresses[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.DaapAddresses[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
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
	if len(m.Feepayer) > 0 {
		i -= len(m.Feepayer)
		copy(dAtA[i:], m.Feepayer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Feepayer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgApproveFeepayerRequestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgApproveFeepayerRequestResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgApproveFeepayerRequestResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreateMetadataRequest) Size() (n int) {
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
	l = len(m.DaapAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.FeePayer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateMetadataRequestResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgApproveFeepayerRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Feepayer)
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
	if len(m.DaapAddresses) > 0 {
		for _, b := range m.DaapAddresses {
			l = len(b)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgApproveFeepayerRequestResponse) Size() (n int) {
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
func (m *MsgCreateMetadataRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateMetadataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateMetadataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field DaapAddress", wireType)
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
			m.DaapAddress = append(m.DaapAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.DaapAddress == nil {
				m.DaapAddress = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeePayer", wireType)
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
			m.FeePayer = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCreateMetadataRequestResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateMetadataRequestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateMetadataRequestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgApproveFeepayerRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgApproveFeepayerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgApproveFeepayerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feepayer", wireType)
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
			m.Feepayer = string(dAtA[iNdEx:postIndex])
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaapAddresses", wireType)
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
			m.DaapAddresses = append(m.DaapAddresses, make([]byte, postIndex-iNdEx))
			copy(m.DaapAddresses[len(m.DaapAddresses)-1], dAtA[iNdEx:postIndex])
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
func (m *MsgApproveFeepayerRequestResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgApproveFeepayerRequestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgApproveFeepayerRequestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
