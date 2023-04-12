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
	chainId string,
	eventNonce uint64,
	claimHash []byte,
) []byte {
	crosschainAckRequestKey := util.AppendBytes([]byte(chainId), util.UInt64Bytes(eventNonce), claimHash)

	return crosschainAckRequestKey
}
