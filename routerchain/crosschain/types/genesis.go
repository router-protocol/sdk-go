package types

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// CrosschainRequestList: []CrosschainRequest{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in crosschainRequest
	// crosschainRequestIndexMap := make(map[string]struct{})

	// for _, elem := range gs.CrosschainRequestList {
	// 	index := string(CrosschainRequestKey(elem.Index))
	// 	if _, ok := crosschainRequestIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for crosschainRequest")
	// 	}
	// 	crosschainRequestIndexMap[index] = struct{}{}
	// }
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
