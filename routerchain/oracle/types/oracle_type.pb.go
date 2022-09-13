// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/oracle_type.proto

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

type OracleType int32

const (
	ORACLE_TYPE_BAND      OracleType = 0
	ORACLE_TYPE_CHAINLINK OracleType = 1
)

var OracleType_name = map[int32]string{
	0: "ORACLE_TYPE_BAND",
	1: "ORACLE_TYPE_CHAINLINK",
}

var OracleType_value = map[string]int32{
	"ORACLE_TYPE_BAND":      0,
	"ORACLE_TYPE_CHAINLINK": 1,
}

func (x OracleType) String() string {
	return proto.EnumName(OracleType_name, int32(x))
}

func (OracleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9a86ee8a17a7f6bf, []int{0}
}

func init() {
	proto.RegisterEnum("routerprotocol.routerchain.oracle.OracleType", OracleType_name, OracleType_value)
}

func init() { proto.RegisterFile("oracle/oracle_type.proto", fileDescriptor_9a86ee8a17a7f6bf) }

var fileDescriptor_9a86ee8a17a7f6bf = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0x2f, 0x4a, 0x4c,
	0xce, 0x49, 0xd5, 0x87, 0x50, 0xf1, 0x25, 0x95, 0x05, 0xa9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x8a, 0x45, 0xf9, 0xa5, 0x25, 0xa9, 0x45, 0x60, 0x4e, 0x72, 0x7e, 0x8e, 0x1e, 0x84, 0x9b,
	0x9c, 0x91, 0x98, 0x99, 0xa7, 0x07, 0x51, 0x2d, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xa0,
	0x0f, 0x62, 0x41, 0x34, 0x4a, 0x49, 0xa6, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0x79, 0x49,
	0xa5, 0x69, 0xfa, 0x89, 0x79, 0x95, 0x10, 0x29, 0xad, 0x12, 0x2e, 0x2e, 0x7f, 0xb0, 0xd6, 0x90,
	0xca, 0x82, 0x54, 0x21, 0x2d, 0x2e, 0x01, 0xff, 0x20, 0x47, 0x67, 0x1f, 0xd7, 0xf8, 0x90, 0xc8,
	0x00, 0xd7, 0x78, 0x27, 0x47, 0x3f, 0x17, 0x01, 0x06, 0x29, 0x91, 0xae, 0xb9, 0x0a, 0x18, 0xe2,
	0x42, 0x26, 0x5c, 0xa2, 0xc8, 0x62, 0xce, 0x1e, 0x8e, 0x9e, 0x7e, 0x3e, 0x9e, 0x7e, 0xde, 0x02,
	0x8c, 0x52, 0x92, 0x5d, 0x73, 0x15, 0xb0, 0x4b, 0x4a, 0xb1, 0x74, 0x2c, 0x96, 0x63, 0x70, 0x0a,
	0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96,
	0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xb3, 0xf4, 0xcc, 0x92, 0x8c,
	0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x88, 0xff, 0x74, 0x61, 0xfe, 0x85, 0xf1, 0xc1, 0x1e,
	0xd6, 0xaf, 0x80, 0x06, 0x90, 0x3e, 0x28, 0x80, 0x8a, 0x93, 0xd8, 0xc0, 0xaa, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xc8, 0x4a, 0x02, 0xf1, 0x3e, 0x01, 0x00, 0x00,
}
