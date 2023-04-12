package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ValsetUpdatedClaimKeyPrefix is the prefix to retrieve all ValsetUpdatedClaim
	ValsetUpdatedClaimKeyPrefix = "ValsetUpdatedClaim/value/"
)

// ValsetUpdatedClaimKey returns the store key to retrieve a ValsetUpdatedClaim from the index fields
func ValsetUpdatedClaimKey(
	chainId string,
	eventNonce uint64,
	claimHash []byte,
) []byte {
	return util.AppendBytes([]byte(chainId), util.UInt64Bytes(eventNonce), claimHash)
}
