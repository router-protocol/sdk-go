package types

import (
	proto "github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

func (c OutboundAck) ValidateBasic() error {
	//TODO: Validate id?
	//TODO: Validate cosmos sender?

	return nil
}

// Hash implements OutboundAck.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *OutboundAck) ClaimHash() ([]byte, error) {
	outboundAckClaimHash := NewOutboundAckClaimHash(
		msg.ChainType,
		msg.ChainId,
		msg.EventNonce,
		msg.BlockHeight,
		msg.GetOutboundTxNonce(),
		msg.GetOutboundTxRequestedBy(),
		msg.GetRelayerRouterAddress(),
		msg.GetDestinationTxHash(),
		msg.GetContractAckResponses(),
		msg.GetExeCode(),
		msg.GetExecStatus(),
		msg.GetExecFlags(),
		msg.GetExecData())

	out, err := proto.Marshal(outboundAckClaimHash)
	return tmhash.Sum([]byte(out)), err
}
