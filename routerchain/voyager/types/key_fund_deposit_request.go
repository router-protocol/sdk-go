package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// FundDepositRequestKeyPrefix is the prefix to retrieve all FundDepositRequest
	FundDepositRequestKeyPrefix = "FundDepositRequest/value/"
)

// FundDepositRequestKey returns the store key to retrieve a FundDepositRequest from the index fields
func FundDepositRequestKey(
	chainId string,
	depositId uint64,
	claimHash []byte,
) []byte {
	fundDepositRequestKey := util.AppendBytes([]byte(chainId), util.UInt64Bytes(depositId), claimHash)

	return fundDepositRequestKey
}
