package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// CrosschainAckRequestKeyPrefix is the prefix to retrieve all CrosschainAckRequest
	CrosschainAckRequestKeyPrefix = "CrosschainAckRequest/value/"
)

// CrosschainAckRequestKey returns the store key to retrieve a CrosschainAckRequest from the index fields
func CrosschainAckRequestKey(
	ackSrcChainId string,
	ackRequestIdentifier uint64,
	claimHash []byte,
) []byte {
	crosschainAckRequestKey := util.AppendBytes([]byte(ackSrcChainId), util.UInt64Bytes(ackRequestIdentifier), claimHash)

	return crosschainAckRequestKey
}
