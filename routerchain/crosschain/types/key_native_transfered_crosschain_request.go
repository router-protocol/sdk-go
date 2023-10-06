package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// NativeTransferedCrosschainRequestKeyPrefix is the prefix to retrieve all NativeTransferedCrosschainRequest
	NativeTransferedCrosschainRequestKeyPrefix = "NativeTransferedCrosschainRequest/value/"
)

// NativeTransferedCrosschainRequestKey returns the store key to retrieve a NativeTransferedCrosschainRequest from the index fields
func NativeTransferedCrosschainRequestKey(
	srcChainId string,
	requestIdentifier uint64,
) []byte {
	crosschainRequestKey := util.AppendBytes([]byte(srcChainId), util.UInt64Bytes(requestIdentifier))

	return crosschainRequestKey
}
