package types

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// AttestationVotesPowerThreshold threshold of votes power to succeed
var (
	AttestationVotesPowerThreshold = sdkmath.NewInt(66)
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ValsetList:              []Valset{},
		AttestationList:         []Attestation{},
		ValsetConfirmationList:  []ValsetConfirmation{},
		ValsetUpdatedClaimList:  []ValsetUpdatedClaim{},
		ObservedAttestationList: []Attestation{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in valset
	valsetIndexMap := make(map[string]struct{})

	for _, elem := range gs.ValsetList {
		index := string(ValsetKey(elem.Nonce))
		if _, ok := valsetIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for valset")
		}
		valsetIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in attestation
	// attestationIndexMap := make(map[string]struct{})

	// TODO @venky Unpack attestation list properly and init
	/*
		for _, elem := range gs.AttestationList {
			index := string(AttestationKey(elem.Index))
			if _, ok := attestationIndexMap[index]; ok {
				return fmt.Errorf("duplicated index for attestation")
			}
			attestationIndexMap[index] = struct{}{}
		}
		**/

	// Check for duplicated index in valsetConfirmation
	valsetConfirmationIndexMap := make(map[string]struct{})

	for _, elem := range gs.ValsetConfirmationList {
		index := string(ValsetConfirmationKey(elem.ValsetNonce, elem.DestChainType, sdk.AccAddress(elem.Orchestrator)))
		if _, ok := valsetConfirmationIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for valsetConfirmation")
		}
		valsetConfirmationIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in valsetUpdatedClaim
	// valsetUpdatedClaimIndexMap := make(map[string]struct{})

	// for _, elem := range gs.ValsetUpdatedClaimList {
	// 	index := string(ValsetUpdatedClaimKey(elem.Index))
	// 	if _, ok := valsetUpdatedClaimIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for valsetUpdatedClaim")
	// 	}
	// 	valsetUpdatedClaimIndexMap[index] = struct{}{}
	// }
	// Check for duplicated index in observedAttestation
	// observedAttestationIndexMap := make(map[string]struct{})

	// for _, elem := range gs.ObservedAttestationList {
	// 	index := string(ObservedAttestationKey(elem.Index))
	// 	if _, ok := observedAttestationIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for observedAttestation")
	// 	}
	// 	observedAttestationIndexMap[index] = struct{}{}
	// }
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
