package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// FeesSettledCrosschainRequestKeyPrefix is the prefix to retrieve all FeesSettledCrosschainRequest
	FeesSettledCrosschainRequestKeyPrefix = "FeesSettledCrosschainRequest/value/"
)

// FeesSettledCrosschainRequestKey returns the store key to retrieve a FeesSettledCrosschainRequest from the index fields
func FeesSettledCrosschainRequestKey(
	srcChainId string,
	requestIdentifier uint64,
) []byte {
	crosschainRequestKey := util.AppendBytes([]byte(srcChainId), util.UInt64Bytes(requestIdentifier))

	return crosschainRequestKey
}
