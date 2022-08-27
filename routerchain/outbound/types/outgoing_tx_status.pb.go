// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: outbound/outgoing_tx_status.proto

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

type OutgoingTxStatus int32

const (
	OUTGOING_TX_CREATED       OutgoingTxStatus = 0
	OUTGOING_TX_CONFIRMED     OutgoingTxStatus = 1
	OUTGOING_TX_ACK_RECEIVED  OutgoingTxStatus = 2
	OUTGOING_TX_ACK_OBSERVED  OutgoingTxStatus = 3
	OUTGOING_TX_ACK_DELEGATED OutgoingTxStatus = 4
)

var OutgoingTxStatus_name = map[int32]string{
	0: "OUTGOING_TX_CREATED",
	1: "OUTGOING_TX_CONFIRMED",
	2: "OUTGOING_TX_ACK_RECEIVED",
	3: "OUTGOING_TX_ACK_OBSERVED",
	4: "OUTGOING_TX_ACK_DELEGATED",
}

var OutgoingTxStatus_value = map[string]int32{
	"OUTGOING_TX_CREATED":       0,
	"OUTGOING_TX_CONFIRMED":     1,
	"OUTGOING_TX_ACK_RECEIVED":  2,
	"OUTGOING_TX_ACK_OBSERVED":  3,
	"OUTGOING_TX_ACK_DELEGATED": 4,
}

func (x OutgoingTxStatus) String() string {
	return proto.EnumName(OutgoingTxStatus_name, int32(x))
}

func (OutgoingTxStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7e38f3412884ec83, []int{0}
}

func init() {
	proto.RegisterEnum("routerprotocol.routerchain.outbound.OutgoingTxStatus", OutgoingTxStatus_name, OutgoingTxStatus_value)
}

func init() { proto.RegisterFile("outbound/outgoing_tx_status.proto", fileDescriptor_7e38f3412884ec83) }

var fileDescriptor_7e38f3412884ec83 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcc, 0x2f, 0x2d, 0x49,
	0xca, 0x2f, 0xcd, 0x4b, 0xd1, 0xcf, 0x2f, 0x2d, 0x49, 0xcf, 0xcf, 0xcc, 0x4b, 0x8f, 0x2f, 0xa9,
	0x88, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2e,
	0xca, 0x2f, 0x2d, 0x49, 0x2d, 0x02, 0x73, 0x92, 0xf3, 0x73, 0xf4, 0x20, 0xdc, 0xe4, 0x8c, 0xc4,
	0xcc, 0x3c, 0x3d, 0x98, 0x6e, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0x12, 0x7d, 0x10, 0x0b,
	0xa2, 0x55, 0x4a, 0x32, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55, 0x1f, 0xcc, 0x4b, 0x2a, 0x4d, 0xd3,
	0x4f, 0xcc, 0xab, 0x84, 0x48, 0x69, 0x6d, 0x63, 0xe2, 0x12, 0xf0, 0x87, 0x5a, 0x19, 0x52, 0x11,
	0x0c, 0xb6, 0x50, 0xc8, 0x80, 0x4b, 0xd8, 0x3f, 0x34, 0xc4, 0xdd, 0xdf, 0xd3, 0xcf, 0x3d, 0x3e,
	0x24, 0x22, 0xde, 0x39, 0xc8, 0xd5, 0x31, 0xc4, 0xd5, 0x45, 0x80, 0x41, 0x4a, 0xbc, 0x6b, 0xae,
	0x02, 0x36, 0x29, 0x21, 0x13, 0x2e, 0x51, 0x14, 0x61, 0x7f, 0x3f, 0x37, 0xcf, 0x20, 0x5f, 0x57,
	0x17, 0x01, 0x46, 0x29, 0xc9, 0xae, 0xb9, 0x0a, 0xd8, 0x25, 0x85, 0xac, 0xb8, 0x24, 0x90, 0x25,
	0x1c, 0x9d, 0xbd, 0xe3, 0x83, 0x5c, 0x9d, 0x5d, 0x3d, 0xc3, 0x5c, 0x5d, 0x04, 0x98, 0xa4, 0x64,
	0xba, 0xe6, 0x2a, 0xe0, 0x94, 0xc7, 0xa6, 0xd7, 0xdf, 0x29, 0xd8, 0x35, 0x08, 0xa4, 0x97, 0x19,
	0xbb, 0x5e, 0x98, 0xbc, 0x90, 0x0d, 0x97, 0x24, 0xba, 0x9c, 0x8b, 0xab, 0x8f, 0xab, 0x3b, 0xd8,
	0x97, 0x2c, 0x52, 0xb2, 0x5d, 0x73, 0x15, 0x70, 0x2b, 0x90, 0x62, 0xe9, 0x58, 0x2c, 0xc7, 0xe0,
	0x14, 0x74, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x16, 0xe9, 0x99, 0x25,
	0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x90, 0x48, 0xd2, 0x85, 0x45, 0x1a, 0x8c, 0x0f,
	0x8e, 0x35, 0xfd, 0x0a, 0x7d, 0x78, 0xac, 0x97, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xd5,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x93, 0x1e, 0x45, 0x49, 0x0e, 0x02, 0x00, 0x00,
}
