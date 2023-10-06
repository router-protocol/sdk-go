package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// BlockedCrosschainRequestKeyPrefix is the prefix to retrieve all BlockedCrosschainRequest
	BlockedCrosschainRequestKeyPrefix = "BlockedCrosschainRequest/value/"
)

// BlockedCrosschainRequestKey returns the store key to retrieve a BlockedCrosschainRequest from the index fields
func BlockedCrosschainRequestKey(
	srcChainId string,
	requestIdentifier uint64,
) []byte {
	crosschainRequestKey := util.AppendBytes([]byte(srcChainId), util.UInt64Bytes(requestIdentifier))

	return crosschainRequestKey
}
