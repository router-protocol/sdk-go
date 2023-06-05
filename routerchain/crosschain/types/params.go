package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

const (
	DefaultInboundGasPrice    = uint64(50000000)
	DefaultMinimumRelayerFees = uint64(25000000000000000)
)

var (
	KeyInboundGasPrice    = []byte("InboundGasPrice")
	KeyMinimumRelayerFees = []byte("MinimumRelayerFees")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	inboundGasPrice uint64,
	minimumRelayerFees uint64,
) Params {
	return Params{
		InboundGasPrice:    inboundGasPrice,
		MinimumRelayerFees: minimumRelayerFees,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultInboundGasPrice,
		DefaultMinimumRelayerFees,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyInboundGasPrice, &p.InboundGasPrice, validateUint64("InboundGasPrice", true)),
		paramtypes.NewParamSetPair(KeyMinimumRelayerFees, &p.MinimumRelayerFees, validateUint64("MinimumRelayerFees", true)),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
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

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
