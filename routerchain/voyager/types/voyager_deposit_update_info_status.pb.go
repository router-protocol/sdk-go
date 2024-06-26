// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: routerprotocol/routerchain/voyager/voyager_deposit_update_info_status.proto

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

type VoyagerDepositUpdateInfoStatus int32

const (
	VOYAGER_DEPOSIT_UPDATE_INFO_REQUEST_CREATED          VoyagerDepositUpdateInfoStatus = 0
	VOYAGER_DEPOSIT_UPDATE_INFO_REQUEST_VALIDATED        VoyagerDepositUpdateInfoStatus = 1
	VOYAGER_DEPOSIT_UPDATE_INFO_REQUEST_EXECUTED         VoyagerDepositUpdateInfoStatus = 2
	VOYAGER_DEPOSIT_UPDATE_INFO_REQUEST_EXECUTION_FAILED VoyagerDepositUpdateInfoStatus = 3
)

var VoyagerDepositUpdateInfoStatus_name = map[int32]string{
	0: "VOYAGER_DEPOSIT_UPDATE_INFO_REQUEST_CREATED",
	1: "VOYAGER_DEPOSIT_UPDATE_INFO_REQUEST_VALIDATED",
	2: "VOYAGER_DEPOSIT_UPDATE_INFO_REQUEST_EXECUTED",
	3: "VOYAGER_DEPOSIT_UPDATE_INFO_REQUEST_EXECUTION_FAILED",
}

var VoyagerDepositUpdateInfoStatus_value = map[string]int32{
	"VOYAGER_DEPOSIT_UPDATE_INFO_REQUEST_CREATED":          0,
	"VOYAGER_DEPOSIT_UPDATE_INFO_REQUEST_VALIDATED":        1,
	"VOYAGER_DEPOSIT_UPDATE_INFO_REQUEST_EXECUTED":         2,
	"VOYAGER_DEPOSIT_UPDATE_INFO_REQUEST_EXECUTION_FAILED": 3,
}

func (x VoyagerDepositUpdateInfoStatus) String() string {
	return proto.EnumName(VoyagerDepositUpdateInfoStatus_name, int32(x))
}

func (VoyagerDepositUpdateInfoStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_250605b53f330cd7, []int{0}
}

func init() {
	proto.RegisterEnum("routerprotocol.routerchain.voyager.VoyagerDepositUpdateInfoStatus", VoyagerDepositUpdateInfoStatus_name, VoyagerDepositUpdateInfoStatus_value)
}

func init() {
	proto.RegisterFile("routerprotocol/routerchain/voyager/voyager_deposit_update_info_status.proto", fileDescriptor_250605b53f330cd7)
}

var fileDescriptor_250605b53f330cd7 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x2e, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0x2a, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x87, 0x70, 0x93, 0x33, 0x12,
	0x33, 0xf3, 0xf4, 0xcb, 0xf2, 0x2b, 0x13, 0xd3, 0x53, 0x8b, 0x60, 0x74, 0x7c, 0x4a, 0x6a, 0x41,
	0x7e, 0x71, 0x66, 0x49, 0x7c, 0x69, 0x41, 0x4a, 0x62, 0x49, 0x6a, 0x7c, 0x66, 0x5e, 0x5a, 0x7e,
	0x7c, 0x71, 0x49, 0x62, 0x49, 0x69, 0xb1, 0x1e, 0x58, 0xbf, 0x90, 0x12, 0xaa, 0x61, 0x7a, 0x48,
	0x86, 0xe9, 0x41, 0x0d, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd0, 0x07, 0xb1, 0x20,
	0x3a, 0xa5, 0x24, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xbc, 0xa4, 0xd2, 0x34, 0xfd,
	0xc4, 0xbc, 0x4a, 0x88, 0x94, 0xd6, 0x31, 0x66, 0x2e, 0xb9, 0x30, 0x88, 0x66, 0x17, 0x88, 0x03,
	0x42, 0xc1, 0xf6, 0x7b, 0xe6, 0xa5, 0xe5, 0x07, 0x83, 0x6d, 0x17, 0x4a, 0xe0, 0xd2, 0x0e, 0xf3,
	0x8f, 0x74, 0x74, 0x77, 0x0d, 0x8a, 0x77, 0x71, 0x0d, 0xf0, 0x0f, 0xf6, 0x0c, 0x89, 0x0f, 0x0d,
	0x70, 0x71, 0x0c, 0x71, 0x8d, 0xf7, 0xf4, 0x73, 0xf3, 0x8f, 0x0f, 0x72, 0x0d, 0x0c, 0x75, 0x0d,
	0x0e, 0x89, 0x77, 0x0e, 0x72, 0x75, 0x0c, 0x71, 0x75, 0x11, 0x60, 0x90, 0xd2, 0xef, 0x9a, 0xab,
	0x40, 0x8a, 0x16, 0xa1, 0x14, 0x2e, 0x5d, 0x62, 0x94, 0x87, 0x39, 0xfa, 0x78, 0xba, 0x80, 0xed,
	0x60, 0x94, 0x32, 0xec, 0x9a, 0xab, 0x40, 0x9a, 0x26, 0xa1, 0x24, 0x2e, 0x1d, 0x62, 0x34, 0xb8,
	0x46, 0xb8, 0x3a, 0x87, 0x82, 0x2c, 0x61, 0x92, 0x32, 0xe8, 0x9a, 0xab, 0x40, 0x92, 0x1e, 0xa1,
	0x22, 0x2e, 0x13, 0xe2, 0xd5, 0x7b, 0xfa, 0xfb, 0xc5, 0xbb, 0x39, 0x7a, 0xfa, 0xb8, 0xba, 0x08,
	0x30, 0x4b, 0x59, 0x74, 0xcd, 0x55, 0x20, 0x4b, 0xaf, 0x14, 0x4b, 0xc7, 0x62, 0x39, 0x06, 0xa7,
	0xc0, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63,
	0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0x85, 0x26, 0x40, 0x5d, 0xb4, 0x04, 0xa9, 0x0b, 0x49, 0x91,
	0x15, 0xf0, 0x34, 0x59, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x56, 0x66, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xbd, 0x41, 0x3b, 0x73, 0xc6, 0x02, 0x00, 0x00,
}
