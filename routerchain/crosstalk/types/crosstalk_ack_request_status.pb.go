// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crosstalk/crosstalk_ack_request_status.proto

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

type CrossTalkAckRequestStatus int32

const (
	CROSSTALK_ACK_REQUEST_CREATED          CrossTalkAckRequestStatus = 0
	CROSSTALK_ACK_REQUEST_VALIDATED        CrossTalkAckRequestStatus = 1
	CROSSTALK_ACK_REQUEST_READY_TO_EXECUTE CrossTalkAckRequestStatus = 2
	CROSSTALK_ACK_REQUEST_COMPLETED        CrossTalkAckRequestStatus = 3
)

var CrossTalkAckRequestStatus_name = map[int32]string{
	0: "CROSSTALK_ACK_REQUEST_CREATED",
	1: "CROSSTALK_ACK_REQUEST_VALIDATED",
	2: "CROSSTALK_ACK_REQUEST_READY_TO_EXECUTE",
	3: "CROSSTALK_ACK_REQUEST_COMPLETED",
}

var CrossTalkAckRequestStatus_value = map[string]int32{
	"CROSSTALK_ACK_REQUEST_CREATED":          0,
	"CROSSTALK_ACK_REQUEST_VALIDATED":        1,
	"CROSSTALK_ACK_REQUEST_READY_TO_EXECUTE": 2,
	"CROSSTALK_ACK_REQUEST_COMPLETED":        3,
}

func (x CrossTalkAckRequestStatus) String() string {
	return proto.EnumName(CrossTalkAckRequestStatus_name, int32(x))
}

func (CrossTalkAckRequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d2c7aed92efd961, []int{0}
}

func init() {
	proto.RegisterEnum("routerprotocol.routerchain.crosstalk.CrossTalkAckRequestStatus", CrossTalkAckRequestStatus_name, CrossTalkAckRequestStatus_value)
}

func init() {
	proto.RegisterFile("crosstalk/crosstalk_ack_request_status.proto", fileDescriptor_0d2c7aed92efd961)
}

var fileDescriptor_0d2c7aed92efd961 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0x2e, 0xca, 0x2f,
	0x2e, 0x2e, 0x49, 0xcc, 0xc9, 0xd6, 0x87, 0xb3, 0xe2, 0x13, 0x93, 0xb3, 0xe3, 0x8b, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0xe2, 0x8b, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x54, 0x8a, 0xf2, 0x4b, 0x4b, 0x52, 0x8b, 0xc0, 0x9c, 0xe4, 0xfc, 0x1c, 0x3d, 0x08,
	0x37, 0x39, 0x23, 0x31, 0x33, 0x4f, 0x0f, 0xae, 0x5d, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xac,
	0x46, 0x1f, 0xc4, 0x82, 0xe8, 0x95, 0x92, 0x4c, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3,
	0x92, 0x4a, 0xd3, 0xf4, 0x13, 0xf3, 0x2a, 0x21, 0x52, 0x5a, 0x97, 0x99, 0xb8, 0x24, 0x9d, 0x41,
	0xda, 0x43, 0x12, 0x73, 0xb2, 0x1d, 0x93, 0xb3, 0x83, 0x20, 0x76, 0x07, 0x83, 0xad, 0x16, 0x72,
	0xe1, 0x92, 0x75, 0x0e, 0xf2, 0x0f, 0x0e, 0x0e, 0x71, 0xf4, 0xf1, 0x8e, 0x77, 0x74, 0xf6, 0x8e,
	0x0f, 0x72, 0x0d, 0x0c, 0x75, 0x0d, 0x0e, 0x89, 0x77, 0x0e, 0x72, 0x75, 0x0c, 0x71, 0x75, 0x11,
	0x60, 0x90, 0x52, 0xec, 0x9a, 0xab, 0x80, 0x5f, 0x91, 0x90, 0x07, 0x97, 0x3c, 0x76, 0x05, 0x61,
	0x8e, 0x3e, 0x9e, 0x2e, 0x60, 0x73, 0x18, 0xa5, 0x94, 0xbb, 0xe6, 0x2a, 0x10, 0x52, 0x26, 0x14,
	0xc6, 0xa5, 0x86, 0x5d, 0x49, 0x90, 0xab, 0xa3, 0x4b, 0x64, 0x7c, 0x88, 0x7f, 0xbc, 0x6b, 0x84,
	0xab, 0x73, 0x68, 0x88, 0xab, 0x00, 0x93, 0x94, 0x56, 0xd7, 0x5c, 0x05, 0x22, 0x55, 0xe3, 0x76,
	0xa1, 0xb3, 0xbf, 0x6f, 0x80, 0x8f, 0x2b, 0xc8, 0x85, 0xcc, 0xf8, 0x5c, 0x08, 0x57, 0x26, 0xc5,
	0xd2, 0xb1, 0x58, 0x8e, 0xc1, 0x29, 0xf8, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18,
	0xa2, 0x2c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x21, 0x51, 0xa8,
	0x0b, 0x8b, 0x52, 0x18, 0x1f, 0x1c, 0xa7, 0xfa, 0x15, 0x88, 0x44, 0xa1, 0x5f, 0x52, 0x59, 0x90,
	0x5a, 0x9c, 0xc4, 0x06, 0x56, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x18, 0x24, 0xf6, 0x13,
	0x38, 0x02, 0x00, 0x00,
}
