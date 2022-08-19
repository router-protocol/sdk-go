// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: inbound/incoming_tx_status.proto

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

type IncomingTxStatus int32

const (
	INCOMING_TX_CREATED   IncomingTxStatus = 0
	INCOMING_TX_OBSERVED  IncomingTxStatus = 1
	INCOMING_TX_DELIVERED IncomingTxStatus = 2
)

var IncomingTxStatus_name = map[int32]string{
	0: "INCOMING_TX_CREATED",
	1: "INCOMING_TX_OBSERVED",
	2: "INCOMING_TX_DELIVERED",
}

var IncomingTxStatus_value = map[string]int32{
	"INCOMING_TX_CREATED":   0,
	"INCOMING_TX_OBSERVED":  1,
	"INCOMING_TX_DELIVERED": 2,
}

func (x IncomingTxStatus) String() string {
	return proto.EnumName(IncomingTxStatus_name, int32(x))
}

func (IncomingTxStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e77c756f19aa3322, []int{0}
}

func init() {
	proto.RegisterEnum("routerprotocol.routerchain.inbound.IncomingTxStatus", IncomingTxStatus_name, IncomingTxStatus_value)
}

func init() { proto.RegisterFile("inbound/incoming_tx_status.proto", fileDescriptor_e77c756f19aa3322) }

var fileDescriptor_e77c756f19aa3322 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0xcc, 0x4b, 0xca,
	0x2f, 0xcd, 0x4b, 0xd1, 0xcf, 0xcc, 0x4b, 0xce, 0xcf, 0xcd, 0xcc, 0x4b, 0x8f, 0x2f, 0xa9, 0x88,
	0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2a, 0xca,
	0x2f, 0x2d, 0x49, 0x2d, 0x02, 0x73, 0x92, 0xf3, 0x73, 0xf4, 0x20, 0xdc, 0xe4, 0x8c, 0xc4, 0xcc,
	0x3c, 0x3d, 0xa8, 0x66, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0x0a, 0x7d, 0x10, 0x0b, 0xa2,
	0x53, 0x4a, 0x32, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55, 0x1f, 0xcc, 0x4b, 0x2a, 0x4d, 0xd3, 0x4f,
	0xcc, 0xab, 0x84, 0x48, 0x69, 0x6d, 0x61, 0xe4, 0x12, 0xf0, 0x84, 0xda, 0x18, 0x52, 0x11, 0x0c,
	0xb6, 0x4f, 0xc8, 0x80, 0x4b, 0xd8, 0xd3, 0xcf, 0xd9, 0xdf, 0xd7, 0xd3, 0xcf, 0x3d, 0x3e, 0x24,
	0x22, 0xde, 0x39, 0xc8, 0xd5, 0x31, 0xc4, 0xd5, 0x45, 0x80, 0x41, 0x4a, 0xbc, 0x6b, 0xae, 0x02,
	0x36, 0x29, 0x21, 0x23, 0x2e, 0x11, 0x64, 0x61, 0x7f, 0xa7, 0x60, 0xd7, 0xa0, 0x30, 0x57, 0x17,
	0x01, 0x46, 0x29, 0x89, 0xae, 0xb9, 0x0a, 0x58, 0xe5, 0x84, 0x4c, 0xb8, 0x44, 0x91, 0xc5, 0x5d,
	0x5c, 0x7d, 0x3c, 0xc3, 0x5c, 0x83, 0x5c, 0x5d, 0x04, 0x98, 0xa4, 0x24, 0xbb, 0xe6, 0x2a, 0x60,
	0x97, 0x94, 0x62, 0xe9, 0x58, 0x2c, 0xc7, 0xe0, 0x14, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47,
	0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d,
	0xc7, 0x72, 0x0c, 0x51, 0xe6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x90, 0x10, 0xd2, 0x85, 0x85, 0x18, 0x8c, 0x0f, 0x0e, 0x32, 0xfd, 0x0a, 0x7d, 0x58, 0x88, 0x97,
	0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x95, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x42,
	0x7e, 0x6d, 0x26, 0x89, 0x01, 0x00, 0x00,
}
