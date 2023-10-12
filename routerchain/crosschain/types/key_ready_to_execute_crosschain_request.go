package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ReadyToExecuteCrosschainRequestKeyPrefix is the prefix to retrieve all ReadyToExecuteCrosschainRequest
	ReadyToExecuteCrosschainRequestKeyPrefix = "ReadyToExecuteCrosschainRequest/value/"
)

// ReadyToExecuteCrosschainRequestKey returns the store key to retrieve a ReadyToExecuteCrosschainRequest from the index fields
func ReadyToExecuteCrosschainRequestKey(
	workflowType WorkflowType,
	relayerType RelayerType,
	srcChainId string,
	requestIdentifier uint64,

) []byte {
	crosschainRequestKey := util.AppendBytes([]byte(workflowType.String()), []byte(relayerType.String()), []byte(srcChainId), util.UInt64Bytes(requestIdentifier))
	return crosschainRequestKey
}
