// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerchain/crosschain/relayer_type.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
	return fileDescriptor_ef51adb2db35d01e, []int{0}
}

func init() {
	proto.RegisterEnum("routerprotocol.routerchain.crosschain.RelayerType", RelayerType_name, RelayerType_value)
}

func init() {
	proto.RegisterFile("routerchain/crosschain/relayer_type.proto", fileDescriptor_ef51adb2db35d01e)
}

var fileDescriptor_ef51adb2db35d01e = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0x4a, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x2e, 0xca, 0x2f, 0x2e, 0x86, 0x30, 0x8b,
	0x52, 0x73, 0x12, 0x2b, 0x53, 0x8b, 0xe2, 0x4b, 0x2a, 0x0b, 0x52, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x54, 0x21, 0x4a, 0xc1, 0x9c, 0xe4, 0xfc, 0x1c, 0x3d, 0x24, 0x9d, 0x7a, 0x08, 0x9d,
	0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x45, 0xfa, 0x20, 0x16, 0x44, 0xb3, 0x94, 0x64, 0x7a,
	0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x98, 0x97, 0x54, 0x9a, 0xa6, 0x9f, 0x98, 0x57, 0x09, 0x91,
	0xd2, 0x6a, 0x64, 0xe4, 0xe2, 0x0e, 0x82, 0x58, 0x17, 0x52, 0x59, 0x90, 0x2a, 0xa4, 0xc4, 0xc5,
	0x13, 0xe4, 0xea, 0xe3, 0x18, 0xe9, 0x1a, 0x14, 0xef, 0xe7, 0xef, 0xe7, 0x2a, 0xc0, 0x20, 0x25,
	0xd0, 0x35, 0x57, 0x01, 0x45, 0x4c, 0x48, 0x8d, 0x8b, 0x2f, 0xc8, 0x3f, 0x34, 0xc4, 0x35, 0x28,
	0x1e, 0x2a, 0x2c, 0xc0, 0x28, 0x25, 0xd4, 0x35, 0x57, 0x01, 0x4d, 0x54, 0x48, 0x81, 0x8b, 0xdb,
	0xd3, 0xc9, 0x19, 0xae, 0x88, 0x49, 0x8a, 0xbf, 0x6b, 0xae, 0x02, 0xb2, 0x90, 0x14, 0x4b, 0xc7,
	0x62, 0x39, 0x06, 0xa7, 0x90, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xb2,
	0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0xf8, 0x58, 0x17, 0x16,
	0x02, 0x30, 0x3e, 0x24, 0xc4, 0x2a, 0x90, 0x83, 0x0f, 0x14, 0x6c, 0xc5, 0x49, 0x6c, 0x60, 0x95,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0xf5, 0x47, 0x05, 0x65, 0x01, 0x00, 0x00,
}
