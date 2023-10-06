package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ExpiredCrosschainRequestKeyPrefix is the prefix to retrieve all ExpiredCrosschainRequest
	ExpiredCrosschainRequestKeyPrefix = "ExpiredCrosschainRequest/value/"
)

// ExpiredCrosschainRequestKey returns the store key to retrieve a ExpiredCrosschainRequest from the index fields
func ExpiredCrosschainRequestKey(
	srcChainId string,
	requestIdentifier uint64,
) []byte {
	crosschainRequestKey := util.AppendBytes([]byte(srcChainId), util.UInt64Bytes(requestIdentifier))

	return crosschainRequestKey
}
