package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ExecutedCrosschainAckRequestKeyPrefix is the prefix to retrieve all ExecutedCrosschainAckRequest
	ExecutedCrosschainAckRequestKeyPrefix = "ExecutedCrosschainAckRequest/value/"
)

// ExecutedCrosschainAckRequestKey returns the store key to retrieve a ExecutedCrosschainAckRequest from the index fields
func ExecutedCrosschainAckRequestKey(
	ackSrcChainId string,
	ackRequestIdentifier uint64,
) []byte {
	crosschainAckRequestKey := util.AppendBytes([]byte(ackSrcChainId), util.UInt64Bytes(ackRequestIdentifier))

	return crosschainAckRequestKey
}
