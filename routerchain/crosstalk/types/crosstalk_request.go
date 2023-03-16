package types

import (
	"encoding/json"

	"github.com/tendermint/tendermint/crypto/tmhash"
)

func (c CrossTalkRequest) ValidateBasic() error {
	//TODO: Validate id?
	//TODO: Validate cosmos sender?

	return nil
}

// Hash implements cCrossTalkRequest.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the c ante-handler for signatures
func (c *CrossTalkRequest) ClaimHash() ([]byte, error) {
	crosstalkRequestClaimHash := NewCrossTalkRequestClaimHash(
		c.EventNonce,
		c.BlockHeight,
		c.SourceChainType,
		c.SourceChainId,
		c.SourceTxHash,
		c.SourceTimestamp,
		c.DestinationChainType,
		c.DestinationChainId,
		c.DestinationGasLimit,
		c.DestinationGasPrice,
		c.RequestSender,
		c.RequestNonce,
		c.IsAtomic,
		c.ExpiryTimestamp,
		c.DestContractAddresses,
		c.DestContractPayloads,
		c.AckType,
		c.AckGasLimit,
		c.AckGasPrice,
		c.RequestTxOrigin,
		c.IsReadCall,
		c.FeePayer,
		c.AsmAddress)

	out, err := json.Marshal(crosstalkRequestClaimHash)
	return tmhash.Sum([]byte(out)), err
}
