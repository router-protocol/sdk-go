package types

import (
	fmt "fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CrossTalkRequestList:           []CrossTalkRequest{},
		CrosstalkRequestConfirmList:    []CrosstalkRequestConfirm{},
		CrossTalkAckRequestList:        []CrossTalkAckRequest{},
		CrosstalkAckRequestConfirmList: []CrosstalkAckRequestConfirm{},
		CrossTalkAckReceiptList:        []CrossTalkAckReceipt{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in crossTalkRequest
	crossTalkRequestIndexMap := make(map[string]struct{})

	for _, elem := range gs.CrossTalkRequestList {
		claimHash, _ := elem.ClaimHash()
		index := string(CrossTalkRequestKey(elem.SourceChainType, elem.SourceChainId, elem.EventNonce, claimHash))
		if _, ok := crossTalkRequestIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for crossTalkRequest")
		}
		crossTalkRequestIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in crossTalkAckReceipt
	// crossTalkAckReceiptIndexMap := make(map[string]struct{})

	// for _, elem := range gs.CrossTalkAckReceiptList {
	// 	index := string(CrossTalkAckReceiptKey(elem.Index))
	// 	if _, ok := crossTalkAckReceiptIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for crossTalkAckReceipt")
	// 	}
	// 	crossTalkAckReceiptIndexMap[index] = struct{}{}
	// }
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
