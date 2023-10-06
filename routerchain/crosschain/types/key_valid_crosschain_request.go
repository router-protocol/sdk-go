package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ValidCrosschainRequestKeyPrefix is the prefix to retrieve all ValidCrosschainRequest
	ValidCrosschainRequestKeyPrefix = "ValidCrosschainRequest/value/"
)

// ValidCrosschainRequestKey returns the store key to retrieve a ValidCrosschainRequest from the index fields
func ValidCrosschainRequestKey(
	srcChainId string,
	requestIdentifier uint64,
) []byte {
	crosschainRequestKey := util.AppendBytes([]byte(srcChainId), util.UInt64Bytes(requestIdentifier))

	return crosschainRequestKey
}
