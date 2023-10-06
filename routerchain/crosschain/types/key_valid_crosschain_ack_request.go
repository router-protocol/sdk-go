package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ValidCrosschainAckRequestKeyPrefix is the prefix to retrieve all ValidCrosschainAckRequest
	ValidCrosschainAckRequestKeyPrefix = "ValidCrosschainAckRequest/value/"
)

// ValidCrosschainAckRequestKey returns the store key to retrieve a ValidCrosschainAckRequest from the index fields
func ValidCrosschainAckRequestKey(
	ackSrcChainId string,
	ackRequestIdentifier uint64,
) []byte {
	crosschainAckRequestKey := util.AppendBytes([]byte(ackSrcChainId), util.UInt64Bytes(ackRequestIdentifier))

	return crosschainAckRequestKey
}
