// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metastore/meta_tx_status.proto

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

type MetaTxStatus int32

const (
	META_TX_CREATED   MetaTxStatus = 0
	META_TX_VALIDATED MetaTxStatus = 1
)

var MetaTxStatus_name = map[int32]string{
	0: "META_TX_CREATED",
	1: "META_TX_VALIDATED",
}

var MetaTxStatus_value = map[string]int32{
	"META_TX_CREATED":   0,
	"META_TX_VALIDATED": 1,
}

func (x MetaTxStatus) String() string {
	return proto.EnumName(MetaTxStatus_name, int32(x))
}

func (MetaTxStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_47e742fb9b8ae2cf, []int{0}
}

func init() {
	proto.RegisterEnum("routerprotocol.routerchain.metastore.MetaTxStatus", MetaTxStatus_name, MetaTxStatus_value)
}

func init() { proto.RegisterFile("metastore/meta_tx_status.proto", fileDescriptor_47e742fb9b8ae2cf) }

var fileDescriptor_47e742fb9b8ae2cf = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0x4d, 0x2d, 0x49,
	0x2c, 0x2e, 0xc9, 0x2f, 0x4a, 0xd5, 0x07, 0xb1, 0xe2, 0x4b, 0x2a, 0xe2, 0x8b, 0x4b, 0x12, 0x4b,
	0x4a, 0x8b, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x54, 0x8a, 0xf2, 0x4b, 0x4b, 0x52, 0x8b,
	0xc0, 0x9c, 0xe4, 0xfc, 0x1c, 0x3d, 0x08, 0x37, 0x39, 0x23, 0x31, 0x33, 0x4f, 0x0f, 0xae, 0x55,
	0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xac, 0x46, 0x1f, 0xc4, 0x82, 0xe8, 0x95, 0x92, 0x4c, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x13, 0xf3, 0x2a, 0x21, 0x52,
	0x5a, 0x39, 0x5c, 0x3c, 0xbe, 0xa9, 0x25, 0x89, 0x21, 0x15, 0xc1, 0x60, 0xcb, 0x84, 0x34, 0xb8,
	0xf8, 0x7d, 0x5d, 0x43, 0x1c, 0xe3, 0x43, 0x22, 0xe2, 0x9d, 0x83, 0x5c, 0x1d, 0x43, 0x5c, 0x5d,
	0x04, 0x18, 0xa4, 0x84, 0xbb, 0xe6, 0x2a, 0xa0, 0x0b, 0x0b, 0xe9, 0x70, 0x09, 0xc2, 0x84, 0xc2,
	0x1c, 0x7d, 0x3c, 0x5d, 0xc0, 0x6a, 0x19, 0xa5, 0x44, 0xbb, 0xe6, 0x2a, 0x60, 0x4a, 0x48, 0xb1,
	0x74, 0x2c, 0x96, 0x63, 0x70, 0x0a, 0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86,
	0x28, 0xcb, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x88, 0xd7, 0x74,
	0x61, 0x5e, 0x85, 0xf1, 0xc1, 0x7e, 0xd5, 0xaf, 0xd0, 0x47, 0x04, 0x54, 0x49, 0x65, 0x41, 0x6a,
	0x71, 0x12, 0x1b, 0x58, 0xa1, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x87, 0xcb, 0xeb, 0xd1, 0x42,
	0x01, 0x00, 0x00,
}
