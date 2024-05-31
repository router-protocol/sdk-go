package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	// CrosschainRequestConfirmKeyPrefix is the prefix to retrieve all CrosschainRequestConfirm
	CrosschainRequestConfirmKeyPrefix = "CrosschainRequestConfirm/value/"
)

// CrosschainRequestConfirmKey returns the store key to retrieve a CrosschainRequestConfirm from the index fields
func CrosschainRequestConfirmKey(
	sourceChainId string,
	requestIdentifier uint64,
	claimHash []byte,
	orchestrator sdk.AccAddress,
) []byte {
	return util.AppendBytes([]byte(sourceChainId), util.UInt64Bytes(requestIdentifier), claimHash, orchestrator.Bytes())
}
