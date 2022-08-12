package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ValsetKeyPrefix is the prefix to retrieve all Valset
	ValsetKeyPrefix = "Valset/value/"
)

var (
	// LatestValsetNonce indexes the latest valset nonce
	// [0xba0fa05da166611b656bac7739a6e7d3]
	LatestValsetNonce = util.HashString("LatestValsetNonce")

	// LastUnBondingBlockHeight indexes the last validator unbonding block height
	// [0x06a6b30651341e80276e0d2e19449250]
	LastUnBondingBlockHeight = util.HashString("LastUnBondingBlockHeight")
)

// ValsetKey returns the store key to retrieve a Valset from the index fields
func ValsetKey(
	valsetNonce uint64,
) []byte {
	return util.AppendBytes(util.HashString(ValsetKeyPrefix), util.UInt64Bytes(valsetNonce))
}
