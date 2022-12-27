package types

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CrossTalkRequestList:        []CrossTalkRequest{},
		CrosstalkRequestConfirmList: []CrosstalkRequestConfirm{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in crossTalkRequest
	// crossTalkRequestIndexMap := make(map[string]struct{})

	// for _, elem := range gs.CrossTalkRequestList {
	// 	index := string(CrossTalkRequestKey(elem.Index))
	// 	if _, ok := crossTalkRequestIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for crossTalkRequest")
	// 	}
	// 	crossTalkRequestIndexMap[index] = struct{}{}
	// }
	// Check for duplicated index in crosstalkRequestConfirm
	// crosstalkRequestConfirmIndexMap := make(map[string]struct{})

	// for _, elem := range gs.CrosstalkRequestConfirmList {
	// 	index := string(CrosstalkRequestConfirmKey(elem.Index))
	// 	if _, ok := crosstalkRequestConfirmIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for crosstalkRequestConfirm")
	// 	}
	// 	crosstalkRequestConfirmIndexMap[index] = struct{}{}
	// }
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
