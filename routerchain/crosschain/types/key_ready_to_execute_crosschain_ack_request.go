package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ReadyToExecuteCrosschainAckRequestKeyPrefix is the prefix to retrieve all ReadyToExecuteCrosschainAckRequest
	ReadyToExecuteCrosschainAckRequestKeyPrefix = "ReadyToExecuteCrosschainAckRequest/value/"
)

// ReadyToExecuteCrosschainAckRequestKey returns the store key to retrieve a ReadyToExecuteCrosschainAckRequest from the index fields
func ReadyToExecuteCrosschainAckRequestKey(
	ackWorkflow WorkflowType,
	relayerType RelayerType,
	ackSrcChainId string,
	ackRequestIdentifier uint64,

) []byte {
	crosschainAckRequestKey := util.AppendBytes([]byte(ackWorkflow.String()), []byte(relayerType.String()), []byte(ackSrcChainId), util.UInt64Bytes(ackRequestIdentifier))

	return crosschainAckRequestKey
}
