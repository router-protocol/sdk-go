package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

const (
	DefaultInboundGasPrice       = uint64(50000000)
	DefaultMinimumRelayerFees    = uint64(25000000000000000)
	DefaultCleanupInterval       = int64(100)
	DefaultBlockedRetryInterval  = int64(355)
	DefaultBlockedExpiryInterval = int64(3550)
	DefaultRouterAdmin           = "router1xw62qplveq9un9tccxx6qld8qnz5qv3kj8cnja"
)

var (
	KeyInboundGasPrice       = []byte("InboundGasPrice")
	KeyMinimumRelayerFees    = []byte("MinimumRelayerFees")
	KeyCleanupInterval       = []byte("CleanupInterval")
	KeyBlockedRetryInterval  = []byte("BlockedRetryInterval")
	KeyBlockedExpiryInterval = []byte("BlockedExpiryInterval")
	KeyRouterAdmin           = []byte("RouterAdmin")
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
	cleanupInterval int64,
	blockedRetryInterval int64,
	blockedExpiryInterval int64,
	routerAdmin string,
) Params {
	return Params{
		InboundGasPrice:       inboundGasPrice,
		MinimumRelayerFees:    minimumRelayerFees,
		CleanupInterval:       cleanupInterval,
		BlockedRetryInterval:  blockedRetryInterval,
		BlockedExpiryInterval: blockedExpiryInterval,
		RouterAdmin:           routerAdmin,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultInboundGasPrice,
		DefaultMinimumRelayerFees,
		DefaultCleanupInterval,
		DefaultBlockedRetryInterval,
		DefaultBlockedExpiryInterval,
		DefaultRouterAdmin,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyInboundGasPrice, &p.InboundGasPrice, validateUint64("InboundGasPrice", true)),
		paramtypes.NewParamSetPair(KeyMinimumRelayerFees, &p.MinimumRelayerFees, validateUint64("MinimumRelayerFees", true)),
		paramtypes.NewParamSetPair(KeyCleanupInterval, &p.CleanupInterval, validateInt64("CleanupInterval", true)),
		paramtypes.NewParamSetPair(KeyBlockedRetryInterval, &p.BlockedRetryInterval, validateInt64("BlockedRetryInterval", true)),
		paramtypes.NewParamSetPair(KeyBlockedExpiryInterval, &p.BlockedExpiryInterval, validateInt64("BlockedExpiryInterval", true)),
		paramtypes.NewParamSetPair(KeyRouterAdmin, &p.RouterAdmin, validateAddress("RouterAdmin")),
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

func validateInt64(name string, positiveOnly bool) func(interface{}) error {
	return func(i interface{}) error {
		v, ok := i.(int64)
		if !ok {
			return fmt.Errorf("invalid parameter type: %T", i)
		}
		if v <= 0 && positiveOnly {
			return fmt.Errorf("%s must be positive: %d", name, v)
		}
		return nil
	}
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

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
