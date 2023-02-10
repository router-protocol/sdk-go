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
	INCOMING_TX_CREATED          IncomingTxStatus = 0
	INCOMING_TX_VALIDATED        IncomingTxStatus = 1
	INCOMING_TX_READY_TO_EXECUTE IncomingTxStatus = 2
	INCOMING_TX_EXECUTED         IncomingTxStatus = 3
	INCOMING_TX_EXECUTION_FAILED IncomingTxStatus = 4
)

var IncomingTxStatus_name = map[int32]string{
	0: "INCOMING_TX_CREATED",
	1: "INCOMING_TX_VALIDATED",
	2: "INCOMING_TX_READY_TO_EXECUTE",
	3: "INCOMING_TX_EXECUTED",
	4: "INCOMING_TX_EXECUTION_FAILED",
}

var IncomingTxStatus_value = map[string]int32{
	"INCOMING_TX_CREATED":          0,
	"INCOMING_TX_VALIDATED":        1,
	"INCOMING_TX_READY_TO_EXECUTE": 2,
	"INCOMING_TX_EXECUTED":         3,
	"INCOMING_TX_EXECUTION_FAILED": 4,
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
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0xcc, 0x4b, 0xca,
	0x2f, 0xcd, 0x4b, 0xd1, 0xcf, 0xcc, 0x4b, 0xce, 0xcf, 0xcd, 0xcc, 0x4b, 0x8f, 0x2f, 0xa9, 0x88,
	0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2a, 0xca,
	0x2f, 0x2d, 0x49, 0x2d, 0x02, 0x73, 0x92, 0xf3, 0x73, 0xf4, 0x20, 0xdc, 0xe4, 0x8c, 0xc4, 0xcc,
	0x3c, 0x3d, 0xa8, 0x66, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0x0a, 0x7d, 0x10, 0x0b, 0xa2,
	0x53, 0x4a, 0x32, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55, 0x1f, 0xcc, 0x4b, 0x2a, 0x4d, 0xd3, 0x4f,
	0xcc, 0xab, 0x84, 0x48, 0x69, 0xed, 0x61, 0xe2, 0x12, 0xf0, 0x84, 0xda, 0x18, 0x52, 0x11, 0x0c,
	0xb6, 0x4f, 0xc8, 0x80, 0x4b, 0xd8, 0xd3, 0xcf, 0xd9, 0xdf, 0xd7, 0xd3, 0xcf, 0x3d, 0x3e, 0x24,
	0x22, 0xde, 0x39, 0xc8, 0xd5, 0x31, 0xc4, 0xd5, 0x45, 0x80, 0x41, 0x4a, 0xbc, 0x6b, 0xae, 0x02,
	0x36, 0x29, 0x21, 0x13, 0x2e, 0x51, 0x64, 0xe1, 0x30, 0x47, 0x1f, 0x4f, 0x17, 0xb0, 0x1e, 0x46,
	0x29, 0xc9, 0xae, 0xb9, 0x0a, 0xd8, 0x25, 0x85, 0x9c, 0xb8, 0x64, 0x90, 0x25, 0x82, 0x5c, 0x1d,
	0x5d, 0x22, 0xe3, 0x43, 0xfc, 0xe3, 0x5d, 0x23, 0x5c, 0x9d, 0x43, 0x43, 0x5c, 0x05, 0x98, 0xa4,
	0x14, 0xba, 0xe6, 0x2a, 0xe0, 0x55, 0x23, 0x64, 0xc4, 0x25, 0x82, 0x2c, 0x0f, 0x15, 0x76, 0x11,
	0x60, 0x96, 0x92, 0xe8, 0x9a, 0xab, 0x80, 0x55, 0x0e, 0xdd, 0x5e, 0x88, 0xb8, 0xa7, 0xbf, 0x5f,
	0xbc, 0x9b, 0xa3, 0xa7, 0x8f, 0xab, 0x8b, 0x00, 0x0b, 0xa6, 0xbd, 0xe8, 0x6a, 0xa4, 0x58, 0x3a,
	0x16, 0xcb, 0x31, 0x38, 0x05, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47,
	0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94,
	0x79, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x24, 0xa6, 0x74, 0x61,
	0x31, 0x07, 0xe3, 0x83, 0xa3, 0x4e, 0xbf, 0x42, 0x1f, 0x16, 0xf3, 0x25, 0x95, 0x05, 0xa9, 0xc5,
	0x49, 0x6c, 0x60, 0x65, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x87, 0xd9, 0xdc, 0x11,
	0x02, 0x00, 0x00,
}
