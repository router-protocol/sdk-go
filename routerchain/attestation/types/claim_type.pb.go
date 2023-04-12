// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: attestation/claim_type.proto

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

type ClaimType int32

const (
	CLAIM_TYPE_UNSPECIFIED            ClaimType = 0
	CLAIM_TYPE_VALSET_UPDATED         ClaimType = 1
	CLAIM_TYPE_SET_METADATA_REQUEST   ClaimType = 2
	CLAIM_TYPE_CROSSCHAIN_REQUEST     ClaimType = 3
	CLAIM_TYPE_CROSSCHAIN_ACK_REQUEST ClaimType = 4
)

var ClaimType_name = map[int32]string{
	0: "CLAIM_TYPE_UNSPECIFIED",
	1: "CLAIM_TYPE_VALSET_UPDATED",
	2: "CLAIM_TYPE_SET_METADATA_REQUEST",
	3: "CLAIM_TYPE_CROSSCHAIN_REQUEST",
	4: "CLAIM_TYPE_CROSSCHAIN_ACK_REQUEST",
}

var ClaimType_value = map[string]int32{
	"CLAIM_TYPE_UNSPECIFIED":            0,
	"CLAIM_TYPE_VALSET_UPDATED":         1,
	"CLAIM_TYPE_SET_METADATA_REQUEST":   2,
	"CLAIM_TYPE_CROSSCHAIN_REQUEST":     3,
	"CLAIM_TYPE_CROSSCHAIN_ACK_REQUEST": 4,
}

func (x ClaimType) String() string {
	return proto.EnumName(ClaimType_name, int32(x))
}

func (ClaimType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3f7dc0a8906b6b4d, []int{0}
}

func init() {
	proto.RegisterEnum("routerprotocol.routerchain.attestation.ClaimType", ClaimType_name, ClaimType_value)
}

func init() { proto.RegisterFile("attestation/claim_type.proto", fileDescriptor_3f7dc0a8906b6b4d) }

var fileDescriptor_3f7dc0a8906b6b4d = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x2c, 0x29, 0x49,
	0x2d, 0x2e, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x4f, 0xce, 0x49, 0xcc, 0xcc, 0x8d, 0x2f, 0xa9,
	0x2c, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2b, 0xca, 0x2f, 0x2d, 0x49, 0x2d,
	0x02, 0x73, 0x92, 0xf3, 0x73, 0xf4, 0x20, 0xdc, 0xe4, 0x8c, 0xc4, 0xcc, 0x3c, 0x3d, 0x24, 0x8d,
	0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x55, 0xfa, 0x20, 0x16, 0x44, 0xb7, 0x94, 0x64, 0x7a,
	0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x98, 0x97, 0x54, 0x9a, 0xa6, 0x9f, 0x98, 0x57, 0x09, 0x91,
	0xd2, 0xda, 0xcd, 0xc8, 0xc5, 0xe9, 0x0c, 0xb2, 0x2d, 0xa4, 0xb2, 0x20, 0x55, 0x48, 0x8a, 0x4b,
	0xcc, 0xd9, 0xc7, 0xd1, 0xd3, 0x37, 0x3e, 0x24, 0x32, 0xc0, 0x35, 0x3e, 0xd4, 0x2f, 0x38, 0xc0,
	0xd5, 0xd9, 0xd3, 0xcd, 0xd3, 0xd5, 0x45, 0x80, 0x41, 0x48, 0x96, 0x4b, 0x12, 0x49, 0x2e, 0xcc,
	0xd1, 0x27, 0xd8, 0x35, 0x24, 0x3e, 0x34, 0xc0, 0xc5, 0x31, 0xc4, 0xd5, 0x45, 0x80, 0x51, 0x48,
	0x99, 0x4b, 0x1e, 0x49, 0x1a, 0x24, 0xe7, 0xeb, 0x1a, 0xe2, 0xe8, 0xe2, 0x18, 0xe2, 0x18, 0x1f,
	0xe4, 0x1a, 0x18, 0xea, 0x1a, 0x1c, 0x22, 0xc0, 0x24, 0xa4, 0xc8, 0x25, 0x8b, 0xa4, 0xc8, 0x39,
	0xc8, 0x3f, 0x38, 0xd8, 0xd9, 0xc3, 0xd1, 0xd3, 0x0f, 0xae, 0x84, 0x59, 0x48, 0x95, 0x4b, 0x11,
	0xbb, 0x12, 0x47, 0x67, 0x6f, 0xb8, 0x32, 0x16, 0x29, 0x96, 0x8e, 0xc5, 0x72, 0x0c, 0x4e, 0xa1,
	0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72,
	0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x9d, 0x9e, 0x59, 0x92, 0x51,
	0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x09, 0x2c, 0x5d, 0x58, 0xe0, 0xc1, 0xf8, 0xe0, 0xd0,
	0xd3, 0xaf, 0xd0, 0x47, 0x0e, 0x78, 0x50, 0x90, 0x17, 0x27, 0xb1, 0x81, 0x95, 0x1a, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x52, 0x7c, 0x7f, 0x05, 0x94, 0x01, 0x00, 0x00,
}
