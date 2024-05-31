package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	// CrosschainAckRequestConfirmKeyPrefix is the prefix to retrieve all CrosschainAckRequestConfirm
	CrosschainAckRequestConfirmKeyPrefix = "CrosschainAckRequestConfirm/value/"
)

// CrosschainAckRequestConfirmKey returns the store key to retrieve a CrosschainAckRequestConfirm from the index fields
func CrosschainAckRequestConfirmKey(
	ackSrcChainId string,
	ackrequestIdentifier uint64,
	claimHash []byte,
	orchestrator sdk.AccAddress,
) []byte {
	return util.AppendBytes([]byte(ackSrcChainId), util.UInt64Bytes(ackrequestIdentifier), claimHash, orchestrator.Bytes())
}
