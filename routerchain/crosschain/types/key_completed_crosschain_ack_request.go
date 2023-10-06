package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// CompletedCrosschainAckRequestKeyPrefix is the prefix to retrieve all CompletedCrosschainAckRequest
	CompletedCrosschainAckRequestKeyPrefix = "CompletedCrosschainAckRequest/value/"
)

// CompletedCrosschainAckRequestKey returns the store key to retrieve a CompletedCrosschainAckRequest from the index fields
func CompletedCrosschainAckRequestKey(
	ackSrcChainId string,
	ackRequestIdentifier uint64,
) []byte {
	crosschainAckRequestKey := util.AppendBytes([]byte(ackSrcChainId), util.UInt64Bytes(ackRequestIdentifier))

	return crosschainAckRequestKey
}
