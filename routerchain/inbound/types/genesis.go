package types

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		IncomingTxList: []IncomingTx{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in incomingTx
	// incomingTxIndexMap := make(map[string]struct{})
	// TODO @venky Unpack the attestation claim properly and calculate the claim hash
	/*
		for _, elem := range gs.IncomingTxList {
			index := string(IncomingTxKey(elem.Index))
			if _, ok := incomingTxIndexMap[index]; ok {
				return fmt.Errorf("duplicated index for incomingTx")
			}
			incomingTxIndexMap[index] = struct{}{}
		}
		**/
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
