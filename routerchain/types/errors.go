package types

import (
	errorsmod "cosmossdk.io/errors"
)

const (
	// RootCodespace is the codespace for all errors defined in this package
	RootCodespace = "router"
)

// NOTE: We can't use 1 since that error code is reserved for internal errors.

var (
	// ErrInvalidValue returns an error resulting from an invalid value.
	ErrInvalidValue = errorsmod.Register(RootCodespace, 2, "invalid value")

	// ErrInvalidChainID returns an error resulting from an invalid chain ID.
	ErrInvalidChainID = errorsmod.Register(RootCodespace, 3, "invalid chain ID")

	// ErrVMExecution returns an error resulting from an error in EVM execution.
	ErrVMExecution = errorsmod.Register(RootCodespace, 4, "error while executing evm transaction")
)
