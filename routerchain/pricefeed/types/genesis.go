package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId: PortID,
		// PriceList: []Price{},
		SymbolRequestList:   []SymbolRequest{},
		PriceFeederInfoList: []PriceFeederInfo{},
		GasPriceList:        []GasPrice{},
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

	// Check for duplicated index in symbolRequest
	symbolRequestIndexMap := make(map[string]struct{})

	for _, elem := range gs.SymbolRequestList {
		index := string(SymbolRequestKey(elem.Symbol))
		if _, ok := symbolRequestIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for symbolRequest")
		}
		symbolRequestIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in priceFeederInfo
	priceFeederInfoIndexMap := make(map[string]struct{})

	for _, elem := range gs.PriceFeederInfoList {
		index := string(PriceFeederInfoKey(elem.Name))
		if _, ok := priceFeederInfoIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for priceFeederInfo")
		}
		priceFeederInfoIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in gasPrice
	gasPriceIndexMap := make(map[string]struct{})

	for _, elem := range gs.GasPriceList {
		index := string(GasPriceKey(elem.PriceFeeder, elem.ChainId))
		if _, ok := gasPriceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for gasPrice")
		}
		gasPriceIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
