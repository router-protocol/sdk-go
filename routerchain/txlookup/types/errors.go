package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/pricefeed module sentinel errors
var (
	ErrInvalidAmount       = errorsmod.Register(ModuleName, 1, "invalid amount")
	ErrInsufficientFunds   = errorsmod.Register(ModuleName, 2, "insufficient funds")
	ErrUnregisteredChain   = errorsmod.Register(ModuleName, 3, "unregistered chain id")
	ErrInvalidGasLimit     = errorsmod.Register(ModuleName, 4, "invalid gas limit")
	ErrAlreadyExisted      = errorsmod.Register(ModuleName, 5, "already existed")
	ErrInvalidInput        = errorsmod.Register(ModuleName, 6, "invalid input")
	ErrInvalidNumberOfArgs = errorsmod.Register(ModuleName, 7, "invalid number of args")
)
