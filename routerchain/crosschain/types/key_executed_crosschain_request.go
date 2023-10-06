package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ExecutedCrosschainRequestKeyPrefix is the prefix to retrieve all ExecutedCrosschainRequest
	ExecutedCrosschainRequestKeyPrefix = "ExecutedCrosschainRequest/value/"
)

// ExecutedCrosschainRequestKey returns the store key to retrieve a ExecutedCrosschainRequest from the index fields
func ExecutedCrosschainRequestKey(
	srcChainId string,
	requestIdentifier uint64,
) []byte {
	crosschainRequestKey := util.AppendBytes([]byte(srcChainId), util.UInt64Bytes(requestIdentifier))

	return crosschainRequestKey

}
