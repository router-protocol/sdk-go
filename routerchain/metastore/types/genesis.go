package types

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MetaInfoList:        []MetaInfo{},
		MetadataRequestList: []MetadataRequest{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in metaInfo
	// metaInfoIndexMap := make(map[string]struct{})

	// for _, elem := range gs.MetaInfoList {
	// 	index := string(MetaInfoKey(elem.Index))
	// 	if _, ok := metaInfoIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for metaInfo")
	// 	}
	// 	metaInfoIndexMap[index] = struct{}{}
	// }
	// Check for duplicated index in metadataRequest
	// metadataRequestIndexMap := make(map[string]struct{})

	// for _, elem := range gs.MetadataRequestList {
	// 	index := string(MetadataRequestKey(elem.Index))
	// 	if _, ok := metadataRequestIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for metadataRequest")
	// 	}
	// 	metadataRequestIndexMap[index] = struct{}{}
	// }
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
