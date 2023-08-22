package types

import (
	proto "github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

func (c DepositUpdateInfoRequest) ValidateBasic() error {
	//TODO: Validate id?
	//TODO: Validate cosmos sender?

	return nil
}

// Hash implements IncomingTx.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *DepositUpdateInfoRequest) ClaimHash() ([]byte, error) {
	depositInfoUpdatedRequestClaimHash := NewDepositInfoUpdatedRequestClaimHash(
		msg.SrcChainId,
		msg.SrcChainType,
		msg.SrcTxHash,
		msg.SrcTimestamp,
		msg.DepositId,
		msg.Contract,
		msg.EventNonce,
		msg.BlockHeight,
		msg.FeeAmount,
		msg.Initiatewithdrawal,
		msg.SrcToken,
		msg.Depositor)

	out, err := proto.Marshal(depositInfoUpdatedRequestClaimHash)
	return tmhash.Sum([]byte(out)), err
}
