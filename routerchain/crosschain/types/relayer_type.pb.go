// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerprotocol/routerchain/crosschain/relayer_type.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	math "math"
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

type RelayerType int32

const (
	RELAYER_NONE   RelayerType = 0
	ROUTER_RELAYER RelayerType = 1
	IBC_RELAYER    RelayerType = 2
)

var RelayerType_name = map[int32]string{
	0: "RELAYER_NONE",
	1: "ROUTER_RELAYER",
	2: "IBC_RELAYER",
}

var RelayerType_value = map[string]int32{
	"RELAYER_NONE":   0,
	"ROUTER_RELAYER": 1,
	"IBC_RELAYER":    2,
}

func (x RelayerType) String() string {
	return proto.EnumName(RelayerType_name, int32(x))
}

func (RelayerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d6cfbc0b2d0242eb, []int{0}
}

func init() {
	proto.RegisterEnum("routerprotocol.routerchain.crosschain.RelayerType", RelayerType_name, RelayerType_value)
}

func init() {
	proto.RegisterFile("routerprotocol/routerchain/crosschain/relayer_type.proto", fileDescriptor_d6cfbc0b2d0242eb)
}

var fileDescriptor_d6cfbc0b2d0242eb = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x28, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0x2a, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x87, 0x70, 0x93, 0x33, 0x12,
	0x33, 0xf3, 0xf4, 0x93, 0x8b, 0xf2, 0x8b, 0x8b, 0x21, 0xcc, 0xa2, 0xd4, 0x9c, 0xc4, 0xca, 0xd4,
	0xa2, 0xf8, 0x92, 0xca, 0x82, 0x54, 0x3d, 0xb0, 0x62, 0x21, 0x55, 0x54, 0x9d, 0x7a, 0x48, 0x3a,
	0xf5, 0x10, 0x3a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x8a, 0xf4, 0x41, 0x2c, 0x88, 0x66,
	0x29, 0xc9, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x7d, 0x30, 0x2f, 0xa9, 0x34, 0x4d, 0x3f, 0x31,
	0xaf, 0x12, 0x22, 0xa5, 0xd5, 0xc8, 0xc8, 0xc5, 0x1d, 0x04, 0xb1, 0x2e, 0xa4, 0xb2, 0x20, 0x55,
	0x48, 0x89, 0x8b, 0x27, 0xc8, 0xd5, 0xc7, 0x31, 0xd2, 0x35, 0x28, 0xde, 0xcf, 0xdf, 0xcf, 0x55,
	0x80, 0x41, 0x4a, 0xa0, 0x6b, 0xae, 0x02, 0x8a, 0x98, 0x90, 0x1a, 0x17, 0x5f, 0x90, 0x7f, 0x68,
	0x88, 0x6b, 0x50, 0x3c, 0x54, 0x58, 0x80, 0x51, 0x4a, 0xa8, 0x6b, 0xae, 0x02, 0x9a, 0xa8, 0x90,
	0x02, 0x17, 0xb7, 0xa7, 0x93, 0x33, 0x5c, 0x11, 0x93, 0x14, 0x7f, 0xd7, 0x5c, 0x05, 0x64, 0x21,
	0x29, 0x96, 0x8e, 0xc5, 0x72, 0x0c, 0x4e, 0x21, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7,
	0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c,
	0xc7, 0x10, 0x65, 0x95, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0x0b, 0x0d, 0x2b,
	0x5d, 0xb4, 0xb0, 0xd3, 0x85, 0x84, 0x58, 0x05, 0x72, 0xf0, 0x81, 0x82, 0xad, 0x38, 0x89, 0x0d,
	0xac, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x43, 0x7c, 0x44, 0x74, 0x01, 0x00, 0x00,
}
