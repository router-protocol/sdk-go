package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// BlockedCrosschainAckRequestKeyPrefix is the prefix to retrieve all BlockedCrosschainAckRequest
	BlockedCrosschainAckRequestKeyPrefix = "BlockedCrosschainAckRequest/value/"
)

// BlockedCrosschainAckRequestKey returns the store key to retrieve a BlockedCrosschainAckRequest from the index fields
func BlockedCrosschainAckRequestKey(
	ackSrcChainId string,
	ackRequestIdentifier uint64,
	claimHash []byte,
) []byte {
	crosschainAckRequestKey := util.AppendBytes([]byte(ackSrcChainId), util.UInt64Bytes(ackRequestIdentifier), claimHash)

	return crosschainAckRequestKey
}
