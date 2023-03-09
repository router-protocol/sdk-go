package types

import (
	"encoding/json"

	"github.com/tendermint/tendermint/crypto/tmhash"
)

func (c IncomingTx) ValidateBasic() error {
	//TODO: Validate id?
	//TODO: Validate cosmos sender?

	return nil
}

// Hash implements IncomingTx.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *IncomingTx) ClaimHash() ([]byte, error) {
	inboundRequestClaimHash := NewInboundRequestClaimHash(
		msg.ChainType,
		msg.ChainId,
		msg.EventNonce,
		msg.BlockHeight,
		msg.SourceSender,
		msg.SourceTxHash,
		msg.SourceTimestamp,
		msg.RouterBridgeContract,
		msg.GasLimit,
		msg.RouteAmount,
		msg.RouteRecipient,
		msg.Payload)

	out, err := json.Marshal(inboundRequestClaimHash)
	return tmhash.Sum([]byte(out)), err
}
