package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// FundsPaidRequestKeyPrefix is the prefix to retrieve all FundsPaidRequest
	FundsPaidRequestKeyPrefix = "FundsPaidRequest/value/"
)

// FundsPaidRequestKey returns the store key to retrieve a FundsPaidRequest from the index fields
func FundsPaidRequestKey(
	chainId string,
	eventNonce uint64,
	claimHash []byte,
) []byte {
	fundPaidRequestKey := util.AppendBytes([]byte(chainId), util.UInt64Bytes(eventNonce), claimHash)
	return fundPaidRequestKey
}
