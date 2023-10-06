package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/crosschain module sentinel errors
var (
	ErrSample         = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidVersion = sdkerrors.Register(ModuleName, 2, "invalid version")
	ErrInvalidRequest = sdkerrors.Register(ModuleName, 3, "invalid request")
	ErrNoIBCRelayer   = sdkerrors.Register(ModuleName, 4, "no IBC relayer for that destination chain")
	ErrIBCAck         = sdkerrors.Register(ModuleName, 5, "IBC acknowledgement failed")
	ErrIBCTransfer    = sdkerrors.Register(ModuleName, 6, "IBC transfer failed")
)
