package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	// ValsetConfirmationKeyPrefix is the prefix to retrieve all ValsetConfirmation
	ValsetConfirmationKeyPrefix = "ValsetConfirmation/value/"
)

// ValsetConfirmationKey returns the store key to retrieve a ValsetConfirmation from the index fields
func ValsetConfirmationKey(
	valsetNonce uint64,
	orchestrator sdk.AccAddress,
) []byte {
	return util.AppendBytes(util.UInt64Bytes(valsetNonce), orchestrator.Bytes())
}
