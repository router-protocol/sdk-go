package types

import (
	"encoding/json"

	"github.com/tendermint/tendermint/crypto/tmhash"
)

func (c CrossTalkAckRequest) ValidateBasic() error {
	//TODO: Validate id?
	//TODO: Validate cosmos sender?

	return nil
}

// Hash implements cCrossTalkAckRequest.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the c ante-handler for signatures
func (c *CrossTalkAckRequest) ClaimHash() ([]byte, error) {
	crosstalkAckRequestClaimHash := NewCrossTalkAckRequestClaimHash(
		c.EventNonce,
		c.BlockHeight,
		c.RelayerRouterAddress,
		c.SourceChainType,
		c.SourceChainId,
		c.ChainType,
		c.ChainId,
		c.DestinationTxHash,
		c.EventIdentifier,
		c.CrosstalkRequestSender,
		c.CrosstalkNonce,
		c.ContractAckResponses,
		c.ExeCode,
		c.ExecStatus,
		c.ExecFlags,
		c.ExecData)

	out, err := json.Marshal(crosstalkAckRequestClaimHash)
	return tmhash.Sum([]byte(out)), err
}
