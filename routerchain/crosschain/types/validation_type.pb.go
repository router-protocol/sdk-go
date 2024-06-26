// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerprotocol/routerchain/crosschain/validation_type.proto

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

type ValidationType int32

const (
	DEFAULT_VALIDATION      ValidationType = 0
	IBC_VALIDATION          ValidationType = 1
	ORCHESTRATOR_VALIDATION ValidationType = 2
)

var ValidationType_name = map[int32]string{
	0: "DEFAULT_VALIDATION",
	1: "IBC_VALIDATION",
	2: "ORCHESTRATOR_VALIDATION",
}

var ValidationType_value = map[string]int32{
	"DEFAULT_VALIDATION":      0,
	"IBC_VALIDATION":          1,
	"ORCHESTRATOR_VALIDATION": 2,
}

func (x ValidationType) String() string {
	return proto.EnumName(ValidationType_name, int32(x))
}

func (ValidationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49c617079252aa91, []int{0}
}

func init() {
	proto.RegisterEnum("routerprotocol.routerchain.crosschain.ValidationType", ValidationType_name, ValidationType_value)
}

func init() {
	proto.RegisterFile("routerprotocol/routerchain/crosschain/validation_type.proto", fileDescriptor_49c617079252aa91)
}

var fileDescriptor_49c617079252aa91 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x2e, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0x2a, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x87, 0x70, 0x93, 0x33, 0x12,
	0x33, 0xf3, 0xf4, 0x93, 0x8b, 0xf2, 0x8b, 0x8b, 0x21, 0xcc, 0xb2, 0xc4, 0x9c, 0xcc, 0x94, 0xc4,
	0x92, 0xcc, 0xfc, 0xbc, 0xf8, 0x92, 0xca, 0x82, 0x54, 0x3d, 0xb0, 0x7a, 0x21, 0x55, 0x54, 0xcd,
	0x7a, 0x48, 0x9a, 0xf5, 0x10, 0x9a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x8a, 0xf4, 0x41,
	0x2c, 0x88, 0x66, 0x29, 0xc9, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x7d, 0x30, 0x2f, 0xa9, 0x34,
	0x4d, 0x3f, 0x31, 0xaf, 0x12, 0x22, 0xa5, 0xb5, 0x82, 0x91, 0x8b, 0x2f, 0x0c, 0x6e, 0x63, 0x48,
	0x65, 0x41, 0xaa, 0x90, 0x1e, 0x97, 0x90, 0x8b, 0xab, 0x9b, 0x63, 0xa8, 0x4f, 0x48, 0x7c, 0x98,
	0xa3, 0x8f, 0xa7, 0x8b, 0x63, 0x88, 0xa7, 0xbf, 0x9f, 0x00, 0x83, 0x94, 0x58, 0xd7, 0x5c, 0x05,
	0x2c, 0x32, 0x42, 0x6a, 0x5c, 0x7c, 0x9e, 0x4e, 0xce, 0xc8, 0x6a, 0x19, 0xa5, 0x84, 0xba, 0xe6,
	0x2a, 0xa0, 0x89, 0x0a, 0x59, 0x70, 0x89, 0xfb, 0x07, 0x39, 0x7b, 0xb8, 0x06, 0x87, 0x04, 0x39,
	0x86, 0xf8, 0x07, 0x21, 0x6b, 0x60, 0x92, 0x92, 0xee, 0x9a, 0xab, 0x80, 0x4b, 0x5a, 0x8a, 0xa5,
	0x63, 0xb1, 0x1c, 0x83, 0x53, 0xc8, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78,
	0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44,
	0x59, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0x42, 0x43, 0x55, 0x17, 0x2d,
	0x94, 0x75, 0x21, 0x61, 0x5b, 0x81, 0x1c, 0xd0, 0xa0, 0xd0, 0x2d, 0x4e, 0x62, 0x03, 0xab, 0x34,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xea, 0x4f, 0x88, 0x9e, 0x01, 0x00, 0x00,
}
