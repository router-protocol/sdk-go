package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// DepositUpdateInfoRequestKeyPrefix is the prefix to retrieve all DepositUpdateInfoRequest
	DepositUpdateInfoRequestKeyPrefix = "DepositUpdateInfoRequest/value/"
)

// DepositUpdateInfoRequestKey returns the store key to retrieve a DepositUpdateInfoRequest from the index fields
func DepositUpdateInfoRequestKey(
	chainId string,
	eventNonce uint64,
	claimHash []byte,
) []byte {
	fundDepositUpdateInfoRequestKey := util.AppendBytes([]byte(chainId), util.UInt64Bytes(eventNonce), claimHash)

	return fundDepositUpdateInfoRequestKey
}
