package types

func (c CrosschainAckRequest) ValidateBasic() error {
	//TODO: Validate id?
	//TODO: Validate cosmos sender?

	return nil
}

// Hash implements IncomingTx.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *CrosschainAckRequest) ClaimHash() ([]byte, error) {
	// TODO : remove hardcoded value
	return []byte{10}, nil
}
