package types

import (
	"encoding/binary"

	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ValsetUpdatedClaimKeyPrefix is the prefix to retrieve all ValsetUpdatedClaim
	ValsetUpdatedClaimKeyPrefix = "ValsetUpdatedClaim/value/"
)

// ValsetUpdatedClaimKey returns the store key to retrieve a ValsetUpdatedClaim from the index fields
func ValsetUpdatedClaimKey(
	chainType multichainTypes.ChainType,
	chainId string,
	eventNonce uint64,
	claimHash []byte,
) []byte {
	return util.AppendBytes(util.UInt64Bytes(uint64(chainType)), []byte(chainId), util.UInt64Bytes(eventNonce), claimHash)
}
