package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

// DefaultIndex is the default capability global symbol
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:                  PortID,
		BandTokenPriceStateList: []BandTokenPriceState{},
		CallDataRecordList:      []CallDataRecord{},
		BandOracleRequestList:   []BandOracleRequest{},
		GasPriceStateList:       []GasPriceState{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated symbol in bandTokenPriceState
	bandTokenPriceStateIndexMap := make(map[string]struct{})

	for _, elem := range gs.BandTokenPriceStateList {
		symbol := string(BandTokenPriceStateKey(elem.Symbol))
		if _, ok := bandTokenPriceStateIndexMap[symbol]; ok {
			return fmt.Errorf("duplicated symbol for bandTokenPriceState")
		}
		bandTokenPriceStateIndexMap[symbol] = struct{}{}
	}
	// Check for duplicated index in callDataRecord
	callDataRecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.CallDataRecordList {
		index := string(CallDataRecordKey(elem.ClientId))
		if _, ok := callDataRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for callDataRecord")
		}
		callDataRecordIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in bandOracleRequest
	bandOracleRequestIndexMap := make(map[string]struct{})

	for _, elem := range gs.BandOracleRequestList {
		index := string(BandOracleRequestKey(elem.RequestId))
		if _, ok := bandOracleRequestIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for bandOracleRequest")
		}
		bandOracleRequestIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in gasPriceState
	gasPriceStateIndexMap := make(map[string]struct{})

	for _, elem := range gs.GasPriceStateList {
		index := string(GasPriceStateKey(elem.ChainType, elem.ChainId))
		if _, ok := gasPriceStateIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for gasPriceState")
		}
		gasPriceStateIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
