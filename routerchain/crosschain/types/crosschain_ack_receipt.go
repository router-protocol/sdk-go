package types

import (
	"github.com/cometbft/cometbft/crypto/tmhash"
	proto "github.com/cosmos/gogoproto/proto"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

func (c CrosschainAckReceipt) ValidateBasic() error {
	// TODO: Validate id?
	// TODO: Validate cosmos sender?

	return nil
}

// Hash implements IncomingTx.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *CrosschainAckReceipt) ClaimHash() ([]byte, error) {
	crosschainAckReceiptClaimHash := NewCrosschainAckReceiptClaimHash(
		msg.AckReceiptSrcChainId,
		msg.Contract,
		msg.AckReceiptIdentifier,
		msg.AckReceiptBlockHeight,
		msg.AckReceiptTxHash,
		msg.RelayerRouterAddress,
		msg.RequestIdentifier,
		msg.AckSrcChainId,
		msg.AckRequestIdentifier,
		msg.FeeConsumed,
		msg.AckExecData,
		msg.AckExecStatus,
	)

	out, err := proto.Marshal(crosschainAckReceiptClaimHash)
	return tmhash.Sum([]byte(out)), err
}

func (msg *CrosschainAckReceipt) GetChainId() string {
	return msg.AckReceiptSrcChainId
}

func (msg *CrosschainAckReceipt) ValidationType() ValidationType {
	switch msg.AckReceiptSrcChainType {
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return IBC_VALIDATION

	case multichainTypes.CHAIN_TYPE_ROUTER:
		return DEFAULT_VALIDATION

	default:
		return ORCHESTRATOR_VALIDATION
	}
}
