package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

const (
	DefaultVoyagerContractAddress = ""
	DefaultFundDepositGasLimit    = uint64(5000000000000000)
	DefaultFundPaidGasLimit       = uint64(5000000000000000)
	DefaultVoyagerGasPrice        = uint64(50000000)
)

var (
	KeyVoyagerContractAddress = []byte("VoyagerContractAddress")
	KeyFundDepositGasLimit    = []byte("FundDepositGasLimit")
	KeyFundPaidGasLimit       = []byte("FundPaidGasLimit")
	KeyVoyagerGasPrice        = []byte("VoyagerGasPrice")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(voyagerContractAddress string,
	fundDepositGasLimit uint64,
	fundPaidGasLimit uint64,
	voyagerGasPrice uint64) Params {
	return Params{
		VoyagerContractAddress: voyagerContractAddress,
		FundDepositGasLimit:    fundDepositGasLimit,
		FundPaidGasLimit:       fundPaidGasLimit,
		VoyagerGasPrice:        voyagerGasPrice,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultVoyagerContractAddress,
		DefaultFundDepositGasLimit,
		DefaultFundPaidGasLimit,
		DefaultVoyagerGasPrice,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyVoyagerContractAddress, &p.VoyagerContractAddress, validateString("contract address")),
		paramtypes.NewParamSetPair(KeyFundDepositGasLimit, &p.FundDepositGasLimit, validateUint64("FundDepositGasLimit", true)),
		paramtypes.NewParamSetPair(KeyFundPaidGasLimit, &p.FundPaidGasLimit, validateUint64("FundPaidGasLimit", true)),
		paramtypes.NewParamSetPair(KeyVoyagerGasPrice, &p.VoyagerGasPrice, validateUint64("VoyagerGasPrice", true)),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateUint64(name string, positiveOnly bool) func(interface{}) error {
	return func(i interface{}) error {
		v, ok := i.(uint64)
		if !ok {
			return fmt.Errorf("invalid parameter type: %T", i)
		}
		if v <= 0 && positiveOnly {
			return fmt.Errorf("%s must be positive: %d", name, v)
		}
		return nil
	}
}

func validateString(name string) func(interface{}) error {
	return func(i interface{}) error {
		_, ok := i.(string)
		if !ok {
			return fmt.Errorf("%s must be string: %T", name, i)
		}
		return nil
	}
}
