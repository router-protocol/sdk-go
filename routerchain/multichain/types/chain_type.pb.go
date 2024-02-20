// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerchain/multichain/chain_type.proto

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

type ChainType int32

const (
	CHAIN_TYPE_NONE     ChainType = 0
	CHAIN_TYPE_ROUTER   ChainType = 1
	CHAIN_TYPE_EVM      ChainType = 2
	CHAIN_TYPE_COSMOS   ChainType = 3
	CHAIN_TYPE_POLKADOT ChainType = 4
	CHAIN_TYPE_SOLANA   ChainType = 5
	CHAIN_TYPE_NEAR     ChainType = 6
	CHAIN_TYPE_TRON     ChainType = 7
	CHAIN_TYPE_STARKNET ChainType = 8
	CHAIN_TYPE_BITCOIN  ChainType = 9
)

var ChainType_name = map[int32]string{
	0: "CHAIN_TYPE_NONE",
	1: "CHAIN_TYPE_ROUTER",
	2: "CHAIN_TYPE_EVM",
	3: "CHAIN_TYPE_COSMOS",
	4: "CHAIN_TYPE_POLKADOT",
	5: "CHAIN_TYPE_SOLANA",
	6: "CHAIN_TYPE_NEAR",
	7: "CHAIN_TYPE_TRON",
	8: "CHAIN_TYPE_STARKNET",
	9: "CHAIN_TYPE_BITCOIN",
}

var ChainType_value = map[string]int32{
	"CHAIN_TYPE_NONE":     0,
	"CHAIN_TYPE_ROUTER":   1,
	"CHAIN_TYPE_EVM":      2,
	"CHAIN_TYPE_COSMOS":   3,
	"CHAIN_TYPE_POLKADOT": 4,
	"CHAIN_TYPE_SOLANA":   5,
	"CHAIN_TYPE_NEAR":     6,
	"CHAIN_TYPE_TRON":     7,
	"CHAIN_TYPE_STARKNET": 8,
	"CHAIN_TYPE_BITCOIN":  9,
}

func (x ChainType) String() string {
	return proto.EnumName(ChainType_name, int32(x))
}

func (ChainType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1772aa1ee5db7db1, []int{0}
}

func init() {
	proto.RegisterEnum("routerprotocol.routerchain.multichain.ChainType", ChainType_name, ChainType_value)
}

func init() {
	proto.RegisterFile("routerchain/multichain/chain_type.proto", fileDescriptor_1772aa1ee5db7db1)
}

var fileDescriptor_1772aa1ee5db7db1 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x3d, 0x4b, 0xc3, 0x40,
	0x18, 0x4e, 0x6c, 0xad, 0x36, 0x83, 0x9e, 0x57, 0x3f, 0xf0, 0x86, 0x90, 0xc5, 0x0f, 0x44, 0x13,
	0xc1, 0xcd, 0x2d, 0x8d, 0x01, 0x4b, 0xdb, 0xbb, 0x92, 0x9c, 0x82, 0x2e, 0xa5, 0x2d, 0x31, 0x2d,
	0xb4, 0xbd, 0x52, 0x13, 0xb0, 0xff, 0x40, 0x32, 0xf9, 0x07, 0x32, 0xf9, 0x67, 0x1c, 0xbb, 0x08,
	0x8e, 0xd2, 0xfe, 0x11, 0x49, 0x8e, 0x62, 0xb8, 0x74, 0x39, 0xee, 0x7d, 0xbe, 0x78, 0x79, 0x79,
	0x94, 0xb3, 0x29, 0x0b, 0x03, 0x6f, 0xda, 0xeb, 0x77, 0x06, 0x63, 0x63, 0x14, 0x0e, 0x83, 0x01,
	0xff, 0xa6, 0x6f, 0x3b, 0x98, 0x4d, 0x3c, 0x7d, 0x32, 0x65, 0x01, 0x83, 0x27, 0x5c, 0x98, 0x0e,
	0x3d, 0x36, 0xd4, 0x33, 0x3e, 0xfd, 0xdf, 0x87, 0xf6, 0x7d, 0xe6, 0xb3, 0x54, 0x64, 0x24, 0x3f,
	0x6e, 0x46, 0xc7, 0x3e, 0x63, 0xfe, 0xd0, 0x33, 0xd2, 0xa9, 0x1b, 0xbe, 0x18, 0x9d, 0xf1, 0x8c,
	0x53, 0x17, 0xdf, 0x05, 0xa5, 0x6c, 0x25, 0x56, 0x3a, 0x9b, 0x78, 0xf0, 0x5c, 0xd9, 0xb5, 0xee,
	0xcd, 0x1a, 0x6e, 0xd3, 0xa7, 0x96, 0xdd, 0xc6, 0x04, 0xdb, 0x40, 0x42, 0x95, 0x28, 0xd6, 0x44,
	0x18, 0x5e, 0x2a, 0x7b, 0x19, 0xc8, 0x21, 0x0f, 0xd4, 0x76, 0x80, 0x8c, 0x0e, 0xa2, 0x58, 0xcb,
	0x13, 0xf0, 0x54, 0xd9, 0xc9, 0x80, 0xf6, 0x63, 0x13, 0x6c, 0x20, 0x18, 0xc5, 0x9a, 0x80, 0x0a,
	0xa9, 0x16, 0x71, 0x9b, 0xc4, 0x05, 0x85, 0x5c, 0x2a, 0x27, 0xe0, 0xb5, 0x52, 0xc9, 0x80, 0x2d,
	0xd2, 0xa8, 0x9b, 0x77, 0x84, 0x82, 0x22, 0x3a, 0x8a, 0x62, 0x6d, 0x1d, 0x25, 0xe4, 0xbb, 0xa4,
	0x61, 0x62, 0x13, 0x6c, 0xe6, 0xf2, 0x39, 0x21, 0x5e, 0xc3, 0x36, 0x1d, 0x50, 0xca, 0x5f, 0xc3,
	0x36, 0x1d, 0x41, 0x49, 0x1d, 0x82, 0xc1, 0x56, 0x4e, 0x99, 0xc0, 0xc2, 0xce, 0x2e, 0x35, 0x9d,
	0x3a, 0xb6, 0x29, 0xd8, 0xce, 0xed, 0xbc, 0xa2, 0xa0, 0xae, 0xc0, 0x0c, 0x5c, 0xad, 0x51, 0x8b,
	0xd4, 0x30, 0x28, 0xa3, 0xc3, 0x28, 0xd6, 0xd6, 0x30, 0xa8, 0xf8, 0xfe, 0xa9, 0x4a, 0x55, 0xfa,
	0xb5, 0x50, 0xe5, 0xf9, 0x42, 0x95, 0x7f, 0x17, 0xaa, 0xfc, 0xb1, 0x54, 0xa5, 0xf9, 0x52, 0x95,
	0x7e, 0x96, 0xaa, 0xf4, 0x7c, 0xeb, 0x0f, 0x82, 0x7e, 0xd8, 0xd5, 0x7b, 0x6c, 0x64, 0xf0, 0x16,
	0x5d, 0xad, 0x5a, 0xb5, 0x9a, 0x79, 0x07, 0xdf, 0xb2, 0x85, 0x4c, 0xaa, 0xf8, 0xda, 0x2d, 0xa5,
	0xca, 0x9b, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x1b, 0x27, 0x6f, 0xb7, 0x02, 0x00, 0x00,
}
