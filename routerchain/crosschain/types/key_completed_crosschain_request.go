package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// CompletedCrosschainRequestKeyPrefix is the prefix to retrieve all CompletedCrosschainRequest
	CompletedCrosschainRequestKeyPrefix = "CompletedCrosschainRequest/value/"
)

// CompletedCrosschainRequestKey returns the store key to retrieve a CompletedCrosschainRequest from the index fields
func CompletedCrosschainRequestKey(
	srcChainId string,
	requestIdentifier uint64,
) []byte {
	crosschainRequestKey := util.AppendBytes([]byte(srcChainId), util.UInt64Bytes(requestIdentifier))

	return crosschainRequestKey
}
