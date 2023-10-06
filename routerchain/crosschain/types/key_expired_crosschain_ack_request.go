package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ExpiredCrosschainAckRequestKeyPrefix is the prefix to retrieve all ExpiredCrosschainAckRequest
	ExpiredCrosschainAckRequestKeyPrefix = "ExpiredCrosschainAckRequest/value/"
)

// ExpiredCrosschainAckRequestKey returns the store key to retrieve a ExpiredCrosschainAckRequest from the index fields
func ExpiredCrosschainAckRequestKey(
	ackSrcChainId string,
	ackRequestIdentifier uint64,
) []byte {
	crosschainAckRequestKey := util.AppendBytes([]byte(ackSrcChainId), util.UInt64Bytes(ackRequestIdentifier))

	return crosschainAckRequestKey
}
