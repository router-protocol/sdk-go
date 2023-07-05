package types

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// CrosschainRequestList: []CrosschainRequest{},
		CrosschainRequestConfirmList:    []CrosschainRequestConfirm{},
		CrosschainAckRequestList:        []CrosschainAckRequest{},
		CrosschainAckRequestConfirmList: []CrosschainAckRequestConfirm{},
		CrosschainAckReceiptList:        []CrosschainAckReceipt{},
		// BlockedCrosschainAckRequestList: []BlockedCrosschainAckRequest{},
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
	// Check for duplicated index in crosschainRequestConfirm
	// crosschainRequestConfirmIndexMap := make(map[string]struct{})

	// for _, elem := range gs.CrosschainRequestConfirmList {
	// 	index := string(CrosschainRequestConfirmKey(elem.Index))
	// 	if _, ok := crosschainRequestConfirmIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for crosschainRequestConfirm")
	// 	}
	// 	crosschainRequestConfirmIndexMap[index] = struct{}{}
	// }
	// Check for duplicated index in crosschainAckRequest
	// crosschainAckRequestIndexMap := make(map[string]struct{})

	// for _, elem := range gs.CrosschainAckRequestList {
	// 	index := string(CrosschainAckRequestKey(elem.Index))
	// 	if _, ok := crosschainAckRequestIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for crosschainAckRequest")
	// 	}
	// 	crosschainAckRequestIndexMap[index] = struct{}{}
	// }
	// Check for duplicated index in crosschainAckRequestConfirm
	// crosschainAckRequestConfirmIndexMap := make(map[string]struct{})

	// for _, elem := range gs.CrosschainAckRequestConfirmList {
	// 	index := string(CrosschainAckRequestConfirmKey(elem.Index))
	// 	if _, ok := crosschainAckRequestConfirmIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for crosschainAckRequestConfirm")
	// 	}
	// 	crosschainAckRequestConfirmIndexMap[index] = struct{}{}
	// }
	// Check for duplicated index in crosschainAckReceipt
	// crosschainAckReceiptIndexMap := make(map[string]struct{})

	// for _, elem := range gs.CrosschainAckReceiptList {
	// 	index := string(CrosschainAckReceiptKey(elem.Index))
	// 	if _, ok := crosschainAckReceiptIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for crosschainAckReceipt")
	// 	}
	// 	crosschainAckReceiptIndexMap[index] = struct{}{}
	// }
	// Check for duplicated index in blockedCrosschainRequest
	// blockedCrosschainRequestIndexMap := make(map[string]struct{})

	// for _, elem := range gs.BlockedCrosschainRequestList {
	// 	index := string(BlockedCrosschainRequestKey(elem.Index))
	// 	if _, ok := blockedCrosschainRequestIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for blockedCrosschainRequest")
	// 	}
	// 	blockedCrosschainRequestIndexMap[index] = struct{}{}
	// }
	// Check for duplicated index in blockedCrosschainAckRequest
	// blockedCrosschainAckRequestIndexMap := make(map[string]struct{})

	// for _, elem := range gs.BlockedCrosschainAckRequestList {
	// 	index := string(BlockedCrosschainAckRequestKey(elem.Index))
	// 	if _, ok := blockedCrosschainAckRequestIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for blockedCrosschainAckRequest")
	// 	}
	// 	blockedCrosschainAckRequestIndexMap[index] = struct{}{}
	// }
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
