// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: multichain/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_232e36992b7f3d2c, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_232e36992b7f3d2c, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetChainConfigRequest struct {
	ChainType uint64 `protobuf:"varint,1,opt,name=chain_type,json=chainType,proto3" json:"chain_type,omitempty"`
	ChainId   string `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *QueryGetChainConfigRequest) Reset()         { *m = QueryGetChainConfigRequest{} }
func (m *QueryGetChainConfigRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetChainConfigRequest) ProtoMessage()    {}
func (*QueryGetChainConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_232e36992b7f3d2c, []int{2}
}
func (m *QueryGetChainConfigRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetChainConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetChainConfigRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetChainConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetChainConfigRequest.Merge(m, src)
}
func (m *QueryGetChainConfigRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetChainConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetChainConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetChainConfigRequest proto.InternalMessageInfo

func (m *QueryGetChainConfigRequest) GetChainType() uint64 {
	if m != nil {
		return m.ChainType
	}
	return 0
}

func (m *QueryGetChainConfigRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

type QueryGetChainConfigResponse struct {
	ChainConfig ChainConfig `protobuf:"bytes,1,opt,name=chainConfig,proto3" json:"chainConfig"`
}

func (m *QueryGetChainConfigResponse) Reset()         { *m = QueryGetChainConfigResponse{} }
func (m *QueryGetChainConfigResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetChainConfigResponse) ProtoMessage()    {}
func (*QueryGetChainConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_232e36992b7f3d2c, []int{3}
}
func (m *QueryGetChainConfigResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetChainConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetChainConfigResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetChainConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetChainConfigResponse.Merge(m, src)
}
func (m *QueryGetChainConfigResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetChainConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetChainConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetChainConfigResponse proto.InternalMessageInfo

func (m *QueryGetChainConfigResponse) GetChainConfig() ChainConfig {
	if m != nil {
		return m.ChainConfig
	}
	return ChainConfig{}
}

type QueryAllChainConfigRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllChainConfigRequest) Reset()         { *m = QueryAllChainConfigRequest{} }
func (m *QueryAllChainConfigRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllChainConfigRequest) ProtoMessage()    {}
func (*QueryAllChainConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_232e36992b7f3d2c, []int{4}
}
func (m *QueryAllChainConfigRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllChainConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllChainConfigRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllChainConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllChainConfigRequest.Merge(m, src)
}
func (m *QueryAllChainConfigRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllChainConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllChainConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllChainConfigRequest proto.InternalMessageInfo

func (m *QueryAllChainConfigRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllChainConfigResponse struct {
	ChainConfig []ChainConfig       `protobuf:"bytes,1,rep,name=chainConfig,proto3" json:"chainConfig"`
	Pagination  *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllChainConfigResponse) Reset()         { *m = QueryAllChainConfigResponse{} }
func (m *QueryAllChainConfigResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllChainConfigResponse) ProtoMessage()    {}
func (*QueryAllChainConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_232e36992b7f3d2c, []int{5}
}
func (m *QueryAllChainConfigResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllChainConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllChainConfigResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllChainConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllChainConfigResponse.Merge(m, src)
}
func (m *QueryAllChainConfigResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllChainConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllChainConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllChainConfigResponse proto.InternalMessageInfo

func (m *QueryAllChainConfigResponse) GetChainConfig() []ChainConfig {
	if m != nil {
		return m.ChainConfig
	}
	return nil
}

func (m *QueryAllChainConfigResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "routerprotocol.routerchain.multichain.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "routerprotocol.routerchain.multichain.QueryParamsResponse")
	proto.RegisterType((*QueryGetChainConfigRequest)(nil), "routerprotocol.routerchain.multichain.QueryGetChainConfigRequest")
	proto.RegisterType((*QueryGetChainConfigResponse)(nil), "routerprotocol.routerchain.multichain.QueryGetChainConfigResponse")
	proto.RegisterType((*QueryAllChainConfigRequest)(nil), "routerprotocol.routerchain.multichain.QueryAllChainConfigRequest")
	proto.RegisterType((*QueryAllChainConfigResponse)(nil), "routerprotocol.routerchain.multichain.QueryAllChainConfigResponse")
}

func init() { proto.RegisterFile("multichain/query.proto", fileDescriptor_232e36992b7f3d2c) }

var fileDescriptor_232e36992b7f3d2c = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0xb1, 0x8d, 0x76, 0x02, 0x1e, 0xc6, 0xa2, 0x75, 0x6b, 0xd7, 0xb2, 0xa0, 0x16,
	0xa1, 0x33, 0x24, 0x22, 0xc5, 0x82, 0x87, 0xa4, 0x60, 0x11, 0x11, 0xeb, 0x52, 0x3d, 0xf4, 0x22,
	0xb3, 0x9b, 0x71, 0xbb, 0xb0, 0xd9, 0xd9, 0xee, 0xce, 0x8a, 0x41, 0x7a, 0xf1, 0xe6, 0x4d, 0xf0,
	0xcb, 0x78, 0xf6, 0xd4, 0x63, 0x41, 0x10, 0x4f, 0x22, 0x89, 0x37, 0xbf, 0x84, 0xec, 0xcc, 0x84,
	0x9d, 0x9a, 0x55, 0x93, 0xda, 0x4b, 0xc8, 0x9b, 0xf7, 0xde, 0xff, 0xbd, 0xdf, 0xe4, 0x3f, 0x81,
	0x97, 0xfb, 0x79, 0x24, 0x42, 0x7f, 0x9f, 0x86, 0x31, 0x39, 0xc8, 0x59, 0x3a, 0xc0, 0x49, 0xca,
	0x05, 0x47, 0x37, 0x52, 0x9e, 0x0b, 0x96, 0xca, 0xc0, 0xe7, 0x11, 0x56, 0xa1, 0xac, 0xc3, 0x65,
	0x8b, 0xb5, 0x18, 0xf0, 0x80, 0xcb, 0x22, 0x52, 0x7c, 0x53, 0xcd, 0xd6, 0xb5, 0x80, 0xf3, 0x20,
	0x62, 0x84, 0x26, 0x21, 0xa1, 0x71, 0xcc, 0x05, 0x15, 0x21, 0x8f, 0x33, 0x9d, 0xbd, 0xed, 0xf3,
	0xac, 0xcf, 0x33, 0xe2, 0xd1, 0x8c, 0xa9, 0x99, 0xe4, 0x55, 0xcb, 0x63, 0x82, 0xb6, 0x48, 0x42,
	0x83, 0x30, 0x96, 0xc5, 0xba, 0xf6, 0x8a, 0xb1, 0x5e, 0x42, 0x53, 0xda, 0x1f, 0x8b, 0xac, 0x18,
	0x09, 0xf9, 0xf9, 0xc2, 0xe7, 0xf1, 0xcb, 0x30, 0xd0, 0xe9, 0xe5, 0x89, 0xb4, 0x18, 0x24, 0x4c,
	0x25, 0x9d, 0x45, 0x88, 0x9e, 0x16, 0x63, 0x77, 0xa4, 0xa0, 0xcb, 0x0e, 0x72, 0x96, 0x09, 0xc7,
	0x83, 0x97, 0x4e, 0x9c, 0x66, 0x09, 0x8f, 0x33, 0x86, 0x1e, 0xc1, 0x86, 0x1a, 0xbc, 0x04, 0x56,
	0xc1, 0x5a, 0xb3, 0xbd, 0x8e, 0xa7, 0xba, 0x19, 0xac, 0x64, 0xba, 0x73, 0x47, 0xdf, 0xae, 0xd7,
	0x5c, 0x2d, 0xe1, 0x3c, 0x87, 0x96, 0x9c, 0xb1, 0xcd, 0xc4, 0x56, 0x51, 0xb5, 0x25, 0x77, 0xd6,
	0x1b, 0xa0, 0x15, 0x08, 0xcb, 0x5d, 0xe5, 0xb8, 0x39, 0x77, 0x41, 0x9e, 0xec, 0x0e, 0x12, 0x86,
	0xae, 0xc2, 0x0b, 0x2a, 0x1d, 0xf6, 0x96, 0xea, 0xab, 0x60, 0x6d, 0xc1, 0x3d, 0x2f, 0xe3, 0x87,
	0x3d, 0x67, 0x00, 0x97, 0x2b, 0x75, 0x35, 0xc3, 0x1e, 0x6c, 0xfa, 0xe5, 0xb1, 0x06, 0x69, 0x4f,
	0x09, 0x62, 0x08, 0x6a, 0x1a, 0x53, 0xcc, 0xe9, 0x69, 0xa4, 0x4e, 0x14, 0x55, 0x20, 0x3d, 0x80,
	0xb0, 0xfc, 0x4d, 0xf5, 0xe0, 0x9b, 0x58, 0x19, 0x00, 0x17, 0x06, 0xc0, 0xca, 0x74, 0xda, 0x00,
	0x78, 0x87, 0x06, 0x4c, 0xf7, 0xba, 0x46, 0xa7, 0xf3, 0x09, 0x68, 0xc2, 0xdf, 0xc7, 0xfc, 0x89,
	0xf0, 0xdc, 0x99, 0x11, 0xa2, 0xed, 0x13, 0x0c, 0x75, 0xc9, 0x70, 0xeb, 0x9f, 0x0c, 0x6a, 0x31,
	0x13, 0xa2, 0xfd, 0x6e, 0x1e, 0xce, 0x4b, 0x08, 0xf4, 0x11, 0xc0, 0x86, 0x32, 0x08, 0xba, 0x37,
	0xe5, 0x92, 0x93, 0x8e, 0xb5, 0x36, 0x4f, 0xd3, 0xaa, 0xf6, 0x72, 0x36, 0xde, 0x7e, 0xfe, 0xf1,
	0xa1, 0xde, 0x42, 0x84, 0xa8, 0xa6, 0xf5, 0xb1, 0xc8, 0x38, 0x56, 0x6f, 0x67, 0xe2, 0xf9, 0xa1,
	0x9f, 0x00, 0x36, 0x8d, 0x0b, 0x43, 0x9d, 0x59, 0x96, 0xa8, 0xf4, 0xbd, 0xd5, 0xfd, 0x1f, 0x09,
	0xcd, 0xf3, 0x4c, 0xf2, 0x3c, 0x41, 0x8f, 0xa7, 0xe6, 0x31, 0xff, 0x35, 0xc8, 0x9b, 0xf2, 0xe1,
	0x1d, 0x8e, 0x83, 0xb0, 0x77, 0x88, 0xbe, 0x00, 0x78, 0xd1, 0x18, 0xd7, 0x89, 0xa2, 0xd9, 0x80,
	0x2b, 0x5f, 0xc5, 0x6c, 0xc0, 0xd5, 0x8e, 0x77, 0xee, 0x4b, 0xe0, 0x0d, 0x74, 0xf7, 0x54, 0xc0,
	0xdd, 0xdd, 0xa3, 0xa1, 0x0d, 0x8e, 0x87, 0x36, 0xf8, 0x3e, 0xb4, 0xc1, 0xfb, 0x91, 0x5d, 0x3b,
	0x1e, 0xd9, 0xb5, 0xaf, 0x23, 0xbb, 0xb6, 0xb7, 0x19, 0x84, 0x62, 0x3f, 0xf7, 0xb0, 0xcf, 0xfb,
	0x7f, 0x97, 0x7e, 0x6d, 0x8a, 0x17, 0x37, 0x97, 0x79, 0x0d, 0x59, 0x79, 0xe7, 0x57, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x10, 0xb6, 0x63, 0x18, 0x56, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a ChainConfig by chain_id.
	ChainConfig(ctx context.Context, in *QueryGetChainConfigRequest, opts ...grpc.CallOption) (*QueryGetChainConfigResponse, error)
	// Queries a list of ChainConfig items.
	ChainConfigAll(ctx context.Context, in *QueryAllChainConfigRequest, opts ...grpc.CallOption) (*QueryAllChainConfigResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/routerprotocol.routerchain.multichain.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ChainConfig(ctx context.Context, in *QueryGetChainConfigRequest, opts ...grpc.CallOption) (*QueryGetChainConfigResponse, error) {
	out := new(QueryGetChainConfigResponse)
	err := c.cc.Invoke(ctx, "/routerprotocol.routerchain.multichain.Query/ChainConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ChainConfigAll(ctx context.Context, in *QueryAllChainConfigRequest, opts ...grpc.CallOption) (*QueryAllChainConfigResponse, error) {
	out := new(QueryAllChainConfigResponse)
	err := c.cc.Invoke(ctx, "/routerprotocol.routerchain.multichain.Query/ChainConfigAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a ChainConfig by chain_id.
	ChainConfig(context.Context, *QueryGetChainConfigRequest) (*QueryGetChainConfigResponse, error)
	// Queries a list of ChainConfig items.
	ChainConfigAll(context.Context, *QueryAllChainConfigRequest) (*QueryAllChainConfigResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ChainConfig(ctx context.Context, req *QueryGetChainConfigRequest) (*QueryGetChainConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainConfig not implemented")
}
func (*UnimplementedQueryServer) ChainConfigAll(ctx context.Context, req *QueryAllChainConfigRequest) (*QueryAllChainConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainConfigAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routerprotocol.routerchain.multichain.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ChainConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetChainConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ChainConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routerprotocol.routerchain.multichain.Query/ChainConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ChainConfig(ctx, req.(*QueryGetChainConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ChainConfigAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllChainConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ChainConfigAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routerprotocol.routerchain.multichain.Query/ChainConfigAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ChainConfigAll(ctx, req.(*QueryAllChainConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routerprotocol.routerchain.multichain.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ChainConfig",
			Handler:    _Query_ChainConfig_Handler,
		},
		{
			MethodName: "ChainConfigAll",
			Handler:    _Query_ChainConfigAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "multichain/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetChainConfigRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetChainConfigRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetChainConfigRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if m.ChainType != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ChainType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetChainConfigResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetChainConfigResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetChainConfigResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ChainConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllChainConfigRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllChainConfigRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllChainConfigRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllChainConfigResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllChainConfigResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllChainConfigResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainConfig) > 0 {
		for iNdEx := len(m.ChainConfig) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainConfig[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetChainConfigRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainType != 0 {
		n += 1 + sovQuery(uint64(m.ChainType))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetChainConfigResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ChainConfig.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllChainConfigRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllChainConfigResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChainConfig) > 0 {
		for _, e := range m.ChainConfig {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetChainConfigRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetChainConfigRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetChainConfigRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainType", wireType)
			}
			m.ChainType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainType |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetChainConfigResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetChainConfigResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetChainConfigResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllChainConfigRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllChainConfigRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllChainConfigRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllChainConfigResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllChainConfigResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllChainConfigResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainConfig = append(m.ChainConfig, ChainConfig{})
			if err := m.ChainConfig[len(m.ChainConfig)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
