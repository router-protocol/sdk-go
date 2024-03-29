package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// CrosschainRequestKeyPrefix is the prefix to retrieve all CrosschainRequest
	CrosschainRequestKeyPrefix = "CrosschainRequest/value/"
)

// CrosschainRequestKey returns the store key to retrieve a CrosschainRequest from the index fields
func CrosschainRequestKey(
	srcChainId string,
	requestIdentifier uint64,
	claimHash []byte,
) []byte {
	crosschainRequestKey := util.AppendBytes([]byte(srcChainId), util.UInt64Bytes(requestIdentifier), claimHash)

	return crosschainRequestKey
}
