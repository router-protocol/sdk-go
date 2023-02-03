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
	CROSSTALK_ACK_REQUEST_CREATED  CrossTalkAckRequestStatus = 0
	CROSSTALK_ACK_REQUEST_OBSERVED CrossTalkAckRequestStatus = 1
)

var CrossTalkAckRequestStatus_name = map[int32]string{
	0: "CROSSTALK_ACK_REQUEST_CREATED",
	1: "CROSSTALK_ACK_REQUEST_OBSERVED",
}

var CrossTalkAckRequestStatus_value = map[string]int32{
	"CROSSTALK_ACK_REQUEST_CREATED":  0,
	"CROSSTALK_ACK_REQUEST_OBSERVED": 1,
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
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0x2e, 0xca, 0x2f,
	0x2e, 0x2e, 0x49, 0xcc, 0xc9, 0xd6, 0x87, 0xb3, 0xe2, 0x13, 0x93, 0xb3, 0xe3, 0x8b, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0xe2, 0x8b, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x54, 0x8a, 0xf2, 0x4b, 0x4b, 0x52, 0x8b, 0xc0, 0x9c, 0xe4, 0xfc, 0x1c, 0x3d, 0x08,
	0x37, 0x39, 0x23, 0x31, 0x33, 0x4f, 0x0f, 0xae, 0x5d, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xac,
	0x46, 0x1f, 0xc4, 0x82, 0xe8, 0x95, 0x92, 0x4c, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3,
	0x92, 0x4a, 0xd3, 0xf4, 0x13, 0xf3, 0x2a, 0x21, 0x52, 0x5a, 0xeb, 0x19, 0xb9, 0x24, 0x9d, 0x41,
	0xda, 0x43, 0x12, 0x73, 0xb2, 0x1d, 0x93, 0xb3, 0x83, 0x20, 0x76, 0x07, 0x83, 0xad, 0x16, 0x72,
	0xe1, 0x92, 0x75, 0x0e, 0xf2, 0x0f, 0x0e, 0x0e, 0x71, 0xf4, 0xf1, 0x8e, 0x77, 0x74, 0xf6, 0x8e,
	0x0f, 0x72, 0x0d, 0x0c, 0x75, 0x0d, 0x0e, 0x89, 0x77, 0x0e, 0x72, 0x75, 0x0c, 0x71, 0x75, 0x11,
	0x60, 0x90, 0x52, 0xec, 0x9a, 0xab, 0x80, 0x5f, 0x91, 0x90, 0x1b, 0x97, 0x1c, 0x76, 0x05, 0xfe,
	0x4e, 0xc1, 0xae, 0x41, 0x61, 0xae, 0x2e, 0x02, 0x8c, 0x52, 0x4a, 0x5d, 0x73, 0x15, 0x08, 0xa8,
	0x92, 0x62, 0xe9, 0x58, 0x2c, 0xc7, 0xe0, 0x14, 0x7c, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72,
	0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7,
	0x72, 0x0c, 0x51, 0x96, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x90,
	0xe0, 0xd1, 0x85, 0x05, 0x17, 0x8c, 0x0f, 0x0e, 0x2f, 0xfd, 0x0a, 0x44, 0x80, 0xeb, 0x97, 0x54,
	0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x15, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x21,
	0x55, 0x7f, 0x94, 0x01, 0x00, 0x00,
}