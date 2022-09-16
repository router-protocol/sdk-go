package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// CallDataRecordKeyPrefix is the prefix to retrieve all CallDataRecord
	CallDataRecordKeyPrefix = "CallDataRecord/value/"
)

// CallDataRecordKey returns the store key to retrieve a CallDataRecord from the index fields
func CallDataRecordKey(
	clientId uint64,
) []byte {
	return util.AppendBytes(util.HashString(CallDataRecordKeyPrefix), util.UInt64Bytes(clientId))
}
