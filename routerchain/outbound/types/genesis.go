package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		OutgoingBatchTxList:      []OutgoingBatchTx{},
		OutgoingBatchConfirmList: []OutgoingBatchConfirm{},
		OutboundAckList:          []OutboundAck{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in outgoingBatchTx
	outgoingBatchTxIndexMap := make(map[string]struct{})

	for _, elem := range gs.OutgoingBatchTxList {
		index := string(OutgoingBatchTxKey(elem.GetDestinationChainType(), elem.GetDestinationChainId(), elem.GetSourceAddress(), elem.GetNonce()))
		if _, ok := outgoingBatchTxIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for outgoingBatchTx")
		}
		outgoingBatchTxIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in outgoingBatchConfirm
	outgoingBatchConfirmIndexMap := make(map[string]struct{})

	for _, elem := range gs.OutgoingBatchConfirmList {
		orchaddr, err := sdk.AccAddressFromBech32(elem.Orchestrator)
		if err != nil {
			return fmt.Errorf("acc address invalid")
		}

		index := string(OutgoingBatchConfirmKey(elem.DestinationChainType, elem.DestinationChainId, elem.OutgoingBatchSender, elem.OutgoingBatchNonce, orchaddr))
		if _, ok := outgoingBatchConfirmIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for outgoingBatchConfirm")
		}
		outgoingBatchConfirmIndexMap[index] = struct{}{}
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
