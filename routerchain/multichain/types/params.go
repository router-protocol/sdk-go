package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

const DefaultRouterAdmin = "router1xw62qplveq9un9tccxx6qld8qnz5qv3kj8cnja"

var KeyRouterAdmin = []byte("RouterAdmin")

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	routerAdmin string,
) Params {
	return Params{
		RouterAdmin: routerAdmin,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultRouterAdmin,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyRouterAdmin, &p.RouterAdmin, validateAddress("RouterAdmin")),
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

func validateAddress(name string) func(interface{}) error {
	return func(i interface{}) error {
		v, ok := i.(string)
		if !ok {
			return fmt.Errorf("invalid parameter type: %T", i)
		}
		_, err := sdk.AccAddressFromBech32(v)
		if err != nil {
			return fmt.Errorf("%s is not a valid address: %s", name, v)
		}

		return nil
	}
}
