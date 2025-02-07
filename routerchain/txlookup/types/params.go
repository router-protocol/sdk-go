package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

const (
	DefaultGasPrice                 = uint64(50000000)
	DefaultProcessingFeeForOneAdhoc = uint64(10000000000000000)
	DefaultRequestFeeForOneAdhoc    = uint64(1000000000000000)
)

var (
	KeyGasPrice      = []byte("GasPrice")
	KeyProcessingFee = []byte("ProcessingFee")
	KeyRequestFee    = []byte("RequestFee")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	gasPrice uint64,
	processingFee uint64,
	requestFee uint64,
) Params {
	return Params{
		GasPrice:      gasPrice,
		ProcessingFee: processingFee,
		RequestFee:    requestFee,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultGasPrice,
		DefaultProcessingFeeForOneAdhoc,
		DefaultRequestFeeForOneAdhoc,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyGasPrice, &p.GasPrice, validateUint64("gas price", true)),
		paramtypes.NewParamSetPair(KeyProcessingFee, &p.ProcessingFee, validateUint64("minimum processing fee", true)),
		paramtypes.NewParamSetPair(KeyRequestFee, &p.RequestFee, validateUint64("minimum request fee", true)),
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
