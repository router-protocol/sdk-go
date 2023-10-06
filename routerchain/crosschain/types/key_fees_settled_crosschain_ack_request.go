package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// FeesSettledCrosschainAckRequestKeyPrefix is the prefix to retrieve all FeesSettledCrosschainAckRequest
	FeesSettledCrosschainAckRequestKeyPrefix = "FeesSettledCrosschainAckRequest/value/"
)

// FeesSettledCrosschainAckRequestKey returns the store key to retrieve a FeesSettledCrosschainAckRequest from the index fields
func FeesSettledCrosschainAckRequestKey(
	ackSrcChainId string,
	ackRequestIdentifier uint64,
) []byte {
	crosschainAckRequestKey := util.AppendBytes([]byte(ackSrcChainId), util.UInt64Bytes(ackRequestIdentifier))

	return crosschainAckRequestKey
}
