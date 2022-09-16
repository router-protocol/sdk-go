package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// BandOracleRequestKeyPrefix is the prefix to retrieve all BandOracleRequest
	BandOracleRequestKeyPrefix = "BandOracleRequest/value/"
)

// BandOracleRequestKey returns the store key to retrieve a BandOracleRequest from the index fields
func BandOracleRequestKey(
	requestId uint64,
) []byte {
	return util.AppendBytes(util.HashString(BandOracleRequestKeyPrefix), util.UInt64Bytes(requestId))
}
