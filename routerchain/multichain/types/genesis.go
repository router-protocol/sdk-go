package types

import (
	"fmt"
)

// DefaultIndex is the default capability global ChainId
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ChainConfigList: []ChainConfig{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ChainId in chainConfig
	chainConfigIndexMap := make(map[string]struct{})

	for _, elem := range gs.ChainConfigList {
		ChainId := string(ChainConfigKey(elem.ChainId))
		if _, ok := chainConfigIndexMap[ChainId]; ok {
			return fmt.Errorf("duplicated ChainId for chainConfig")
		}
		chainConfigIndexMap[ChainId] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
