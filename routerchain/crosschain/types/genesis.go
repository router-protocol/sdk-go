package types

import (
	fmt "fmt"

	types "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// CrosschainRequestList: []CrosschainRequest{},
		PortId:                          PortID,
		CrosschainRequestConfirmList:    []CrosschainRequestConfirm{},
		CrosschainAckRequestList:        []CrosschainAckRequest{},
		CrosschainAckRequestConfirmList: []CrosschainAckRequestConfirm{},
		CrosschainAckReceiptList:        []CrosschainAckReceipt{},
		// BlockedCrosschainAckRequestList: []BlockedCrosschainAckRequest{},
		RelayerConfigList: []RelayerConfig{},
		// ExpiredCrosschainRequestList: []ExpiredCrosschainRequest{},
		// ExpiredCrosschainAckRequestList: []ExpiredCrosschainAckRequest{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in crosschainRequest
	crosschainRequestIndexMap := make(map[string]struct{})

	for _, elem := range gs.CrosschainRequestList {
		hashValue, err := elem.ClaimHash()

		if err != nil {
			return fmt.Errorf("cannot calculate hash value")
		}
		index := string(CrosschainRequestKey(elem.SrcChainId, elem.RequestIdentifier, hashValue))
		if _, ok := crosschainRequestIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for crosschainRequest")
		}
		crosschainRequestIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in crosschainRequestConfirm
	crosschainRequestConfirmIndexMap := make(map[string]struct{})

	for _, elem := range gs.CrosschainRequestConfirmList {
		index := string(CrosschainRequestConfirmKey(elem.SourceChainId, elem.RequestIdentifier, elem.ClaimHash, types.AccAddress(elem.Orchestrator)))
		if _, ok := crosschainRequestConfirmIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for crosschainRequestConfirm")
		}
		crosschainRequestConfirmIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in crosschainAckRequest
	crosschainAckRequestIndexMap := make(map[string]struct{})

	for _, elem := range gs.CrosschainAckRequestList {
		hashValue, err := elem.ClaimHash()

		if err != nil {
			return fmt.Errorf("cannot calculate hash value")
		}
		index := string(CrosschainAckRequestKey(elem.AckSrcChainId, elem.AckRequestIdentifier, hashValue))
		if _, ok := crosschainAckRequestIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for crosschainAckRequest")
		}
		crosschainAckRequestIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in crosschainAckRequestConfirm
	crosschainAckRequestConfirmIndexMap := make(map[string]struct{})

	for _, elem := range gs.CrosschainAckRequestConfirmList {
		index := string(CrosschainAckRequestConfirmKey(elem.AckSrcChainId, elem.AckRequestIdentifier, elem.ClaimHash, types.AccAddress(elem.Orchestrator)))
		if _, ok := crosschainAckRequestConfirmIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for crosschainAckRequestConfirm")
		}
		crosschainAckRequestConfirmIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in crosschainAckReceipt
	crosschainAckReceiptIndexMap := make(map[string]struct{})

	for _, elem := range gs.CrosschainAckReceiptList {
		hashValue, err := elem.ClaimHash()

		if err != nil {
			return fmt.Errorf("cannot calculate hash value")
		}
		index := string(CrosschainAckReceiptKey(elem.AckReceiptSrcChainId, elem.AckReceiptIdentifier, hashValue))
		if _, ok := crosschainAckReceiptIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for crosschainAckReceipt")
		}
		crosschainAckReceiptIndexMap[index] = struct{}{}
	}
	// // Check for duplicated index in blockedCrosschainRequest
	// blockedCrosschainRequestIndexMap := make(map[string]struct{})

	// for _, elem := range gs.BlockedCrosschainRequestList {
	// 	index := string(BlockedCrosschainRequestKey(elem.SrcChainId, elem.RequestIdentifier, ))
	// 	if _, ok := blockedCrosschainRequestIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for blockedCrosschainRequest")
	// 	}
	// 	blockedCrosschainRequestIndexMap[index] = struct{}{}
	// }
	// // Check for duplicated index in blockedCrosschainAckRequest
	// blockedCrosschainAckRequestIndexMap := make(map[string]struct{})

	// for _, elem := range gs.BlockedCrosschainAckRequestList {
	// 	index := string(BlockedCrosschainAckRequestKey(elem.Index))
	// 	if _, ok := blockedCrosschainAckRequestIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for blockedCrosschainAckRequest")
	// 	}
	// 	blockedCrosschainAckRequestIndexMap[index] = struct{}{}
	// }
	// Check for duplicated index in expiredCrosschainRequest
	// expiredCrosschainRequestIndexMap := make(map[string]struct{})

	// for _, elem := range gs.ExpiredCrosschainRequestList {
	// 	index := string(ExpiredCrosschainRequestKey(elem.Index))
	// 	if _, ok := expiredCrosschainRequestIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for expiredCrosschainRequest")
	// 	}
	// 	expiredCrosschainRequestIndexMap[index] = struct{}{}
	// }
	// Check for duplicated index in expiredCrosschainAckRequest
	// expiredCrosschainAckRequestIndexMap := make(map[string]struct{})

	// for _, elem := range gs.ExpiredCrosschainAckRequestList {
	// 	index := string(ExpiredCrosschainAckRequestKey(elem.Index))
	// 	if _, ok := expiredCrosschainAckRequestIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for expiredCrosschainAckRequest")
	// 	}
	// 	expiredCrosschainAckRequestIndexMap[index] = struct{}{}
	// }
	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in relayerConfig
	relayerConfigIndexMap := make(map[string]struct{})

	for _, elem := range gs.RelayerConfigList {
		index := string(RelayerConfigKey(elem.ChainId))
		if _, ok := relayerConfigIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for relayerConfig")
		}
		relayerConfigIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
