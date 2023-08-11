package types

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		FundDepositRequestList:       []FundDepositRequest{},
		FundsPaidRequestList:         []FundsPaidRequest{},
		DepositUpdateInfoRequestList: []DepositUpdateInfoRequest{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// fmt.Println("validate")
	// Check for duplicated index in fundDepositRequest
	// fundDepositRequestIndexMap := make(map[string]struct{})

	// for _, elem := range gs.FundDepositRequestList {
	// 	index := string(FundDepositRequestKey(elem.SrcChainId))
	// 	if _, ok := fundDepositRequestIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for fundDepositRequest")
	// 	}
	// 	fundDepositRequestIndexMap[index] = struct{}{}
	// }
	// Check for duplicated index in fundsPaidRequest
	// fundsPaidRequestIndexMap := make(map[string]struct{})

	// for _, elem := range gs.FundsPaidRequestList {
	// 	index := string(FundsPaidRequestKey(elem.Index))
	// 	if _, ok := fundsPaidRequestIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for fundsPaidRequest")
	// 	}
	// 	fundsPaidRequestIndexMap[index] = struct{}{}
	// }
	// Check for duplicated index in depositUpdateInfoRequest
	// depositUpdateInfoRequestIndexMap := make(map[string]struct{})

	// for _, elem := range gs.DepositUpdateInfoRequestList {
	// 	index := DepositUpdateInfoRequestKey(elem.DepositId)
	// 	if _, ok := depositUpdateInfoRequestIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for depositUpdateInfoRequest")
	// 	}
	// 	depositUpdateInfoRequestIndexMap[index] = struct{}{}
	// }
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
