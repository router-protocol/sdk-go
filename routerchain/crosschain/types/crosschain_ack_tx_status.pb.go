// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerchain/crosschain/crosschain_ack_tx_status.proto

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

type CrosschainAckTxStatus int32

const (
	CROSSCHAIN_ACK_TX_CREATED          CrosschainAckTxStatus = 0
	CROSSCHAIN_ACK_TX_VALIDATED        CrosschainAckTxStatus = 1
	CROSSCHAIN_ACK_TX_READY_TO_EXECUTE CrosschainAckTxStatus = 3
	CROSSCHAIN_ACK_TX_EXECUTED         CrosschainAckTxStatus = 4
	CROSSCHAIN_ACK_TX_EXECUTION_FAILED CrosschainAckTxStatus = 5
	CROSSCHAIN_ACK_TX_FEES_SETTLED     CrosschainAckTxStatus = 6
	CROSSCHAIN_ACK_TX_COMPLETED        CrosschainAckTxStatus = 7
)

var CrosschainAckTxStatus_name = map[int32]string{
	0: "CROSSCHAIN_ACK_TX_CREATED",
	1: "CROSSCHAIN_ACK_TX_VALIDATED",
	3: "CROSSCHAIN_ACK_TX_READY_TO_EXECUTE",
	4: "CROSSCHAIN_ACK_TX_EXECUTED",
	5: "CROSSCHAIN_ACK_TX_EXECUTION_FAILED",
	6: "CROSSCHAIN_ACK_TX_FEES_SETTLED",
	7: "CROSSCHAIN_ACK_TX_COMPLETED",
}

var CrosschainAckTxStatus_value = map[string]int32{
	"CROSSCHAIN_ACK_TX_CREATED":          0,
	"CROSSCHAIN_ACK_TX_VALIDATED":        1,
	"CROSSCHAIN_ACK_TX_READY_TO_EXECUTE": 3,
	"CROSSCHAIN_ACK_TX_EXECUTED":         4,
	"CROSSCHAIN_ACK_TX_EXECUTION_FAILED": 5,
	"CROSSCHAIN_ACK_TX_FEES_SETTLED":     6,
	"CROSSCHAIN_ACK_TX_COMPLETED":        7,
}

func (x CrosschainAckTxStatus) String() string {
	return proto.EnumName(CrosschainAckTxStatus_name, int32(x))
}

func (CrosschainAckTxStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bddb92c34d61e985, []int{0}
}

func init() {
	proto.RegisterEnum("routerprotocol.routerchain.crosschain.CrosschainAckTxStatus", CrosschainAckTxStatus_name, CrosschainAckTxStatus_value)
}

func init() {
	proto.RegisterFile("routerchain/crosschain/crosschain_ack_tx_status.proto", fileDescriptor_bddb92c34d61e985)
}

var fileDescriptor_bddb92c34d61e985 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2d, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0x4a, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x2e, 0xca, 0x2f, 0x2e, 0x46, 0x67, 0xc6,
	0x27, 0x26, 0x67, 0xc7, 0x97, 0x54, 0xc4, 0x17, 0x97, 0x24, 0x96, 0x94, 0x16, 0xeb, 0x15, 0x14,
	0xe5, 0x97, 0xe4, 0x0b, 0xa9, 0x42, 0xb4, 0x81, 0x39, 0xc9, 0xf9, 0x39, 0x7a, 0x48, 0xa6, 0xe8,
	0x21, 0xb4, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x15, 0xe9, 0x83, 0x58, 0x10, 0xcd, 0x52,
	0x92, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0x5e, 0x52, 0x69, 0x9a, 0x7e, 0x62, 0x5e,
	0x25, 0x44, 0x4a, 0xab, 0x9d, 0x85, 0x4b, 0xd4, 0x19, 0xae, 0xdf, 0x31, 0x39, 0x3b, 0xa4, 0x22,
	0x18, 0x6c, 0xaf, 0x90, 0x0d, 0x97, 0xa4, 0x73, 0x90, 0x7f, 0x70, 0xb0, 0xb3, 0x87, 0xa3, 0xa7,
	0x5f, 0xbc, 0xa3, 0xb3, 0x77, 0x7c, 0x48, 0x44, 0xbc, 0x73, 0x90, 0xab, 0x63, 0x88, 0xab, 0x8b,
	0x00, 0x83, 0x94, 0x6c, 0xd7, 0x5c, 0x05, 0xdc, 0x0a, 0x84, 0x1c, 0xb8, 0xa4, 0x31, 0x25, 0xc3,
	0x1c, 0x7d, 0x3c, 0x5d, 0xc0, 0xfa, 0x19, 0xa5, 0xe4, 0xbb, 0xe6, 0x2a, 0xe0, 0x53, 0x22, 0xe4,
	0xc7, 0xa5, 0x84, 0x29, 0x1d, 0xe4, 0xea, 0xe8, 0x12, 0x19, 0x1f, 0xe2, 0x1f, 0xef, 0x1a, 0xe1,
	0xea, 0x1c, 0x1a, 0xe2, 0x2a, 0xc0, 0x2c, 0xa5, 0xd6, 0x35, 0x57, 0x81, 0x08, 0x95, 0x42, 0x76,
	0x5c, 0x52, 0x98, 0xaa, 0xa0, 0x92, 0x2e, 0x02, 0x2c, 0x52, 0x72, 0x5d, 0x73, 0x15, 0xf0, 0xa8,
	0xc0, 0xee, 0x1e, 0x88, 0xac, 0xa7, 0xbf, 0x5f, 0xbc, 0x9b, 0xa3, 0xa7, 0x8f, 0xab, 0x8b, 0x00,
	0x2b, 0x2e, 0xf7, 0xa0, 0xab, 0x14, 0x72, 0xe3, 0x92, 0xc3, 0x54, 0xe5, 0xe6, 0xea, 0x1a, 0x1c,
	0x1f, 0xec, 0x1a, 0x12, 0x02, 0x32, 0x8b, 0x4d, 0x4a, 0xa9, 0x6b, 0xae, 0x02, 0x01, 0x55, 0xd8,
	0x43, 0xda, 0xd9, 0xdf, 0x37, 0xc0, 0xc7, 0x15, 0xe4, 0x31, 0x76, 0x5c, 0x21, 0x0d, 0x57, 0x22,
	0xc5, 0xd2, 0xb1, 0x58, 0x8e, 0xc1, 0x29, 0xe4, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5,
	0x18, 0xa2, 0xac, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x21, 0xe9,
	0x4e, 0x17, 0x96, 0x0e, 0x61, 0x7c, 0x48, 0x1a, 0xae, 0x40, 0x4e, 0xd0, 0x25, 0x95, 0x05, 0xa9,
	0xc5, 0x49, 0x6c, 0x60, 0x95, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0xb9, 0xf1, 0xbb,
	0xf7, 0x02, 0x00, 0x00,
}