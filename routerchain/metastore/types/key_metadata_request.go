package types

import (
	"encoding/binary"

	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// MetadataRequestKeyPrefix is the prefix to retrieve all MetadataRequest
	MetadataRequestKeyPrefix = "MetadataRequest/value/"
)

// MetadataRequestKey returns the store key to retrieve a MetadataRequest from the index fields
func MetadataRequestKey(
	chainType multichainTypes.ChainType,
	chainId string,
	eventNonce uint64,
	claimHash []byte,
) []byte {
	return util.AppendBytes(util.UInt64Bytes(uint64(chainType)), []byte(chainId), util.UInt64Bytes(eventNonce), claimHash)
}
