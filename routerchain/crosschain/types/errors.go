package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/crosschain module sentinel errors
var (
	ErrSample         = errorsmod.Register(ModuleName, 1100, "sample error")
	ErrInvalidVersion = errorsmod.Register(ModuleName, 2, "invalid version")
	ErrInvalidRequest = errorsmod.Register(ModuleName, 3, "invalid request")
	ErrNoIBCRelayer   = errorsmod.Register(ModuleName, 4, "no IBC relayer for that destination chain")
	ErrIBCAck         = errorsmod.Register(ModuleName, 5, "IBC acknowledgement failed")
	ErrIBCTransfer    = errorsmod.Register(ModuleName, 6, "IBC transfer failed")
	ErrNotSupported   = errorsmod.Register(ModuleName, 7, "crosschain request not supported")
)
