package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ObservedAttestationKeyPrefix is the prefix to retrieve all ObservedAttestation
	ObservedAttestationKeyPrefix = "ObservedAttestation/value/"
)

// ObservedAttestationKey returns the store key to retrieve a ObservedAttestation from the index fields
func ObservedAttestationKey(
	chainId string,
	contract string,
	eventNonce uint64,
	claimHash []byte,
) []byte {
	return util.AppendBytes([]byte(chainId), []byte(contract), util.UInt64Bytes(eventNonce), claimHash)
}
