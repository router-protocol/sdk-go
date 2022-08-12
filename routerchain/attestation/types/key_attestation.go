package types

import (
	"encoding/binary"
	"strings"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// AttestationKeyPrefix is the prefix to retrieve all Attestation
	AttestationKeyPrefix = "Attestation/value/"
)

var (
	// PastEthSignatureCheckpointKey indexes eth signature checkpoints that have existed
	// [0x1cbe0be407a979331b98e599eeedd09f]
	PastEthSignatureCheckpointKey = util.HashString("PastEthSignatureCheckpointKey")
)

// AttestationKey returns the store key to retrieve a Attestation from the index fields
func AttestationKey(
	chainId uint64,
	eventNonce uint64,
	claimHash []byte,
) []byte {

	return util.AppendBytes(util.UInt64Bytes(chainId), util.UInt64Bytes(eventNonce), claimHash)
}

// GetPastEthSignatureCheckpointKey returns the following key format
// prefix    checkpoint
// [0x0][ checkpoint bytes ]
func GetPastEthSignatureCheckpointKey(checkpoint []byte) []byte {
	return util.AppendBytes(PastEthSignatureCheckpointKey, []byte(convertByteArrToString(checkpoint)))
}

// This function is broken and it should not be used in other places except in GetPastEthSignatureCheckpointKey
func convertByteArrToString(value []byte) string {
	var ret strings.Builder
	for i := 0; i < len(value); i++ {
		ret.WriteString(string(value[i]))
	}
	return ret.String()
}
