package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/crosschain module sentinel errors
var (
	ErrSample         = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidVersion = sdkerrors.Register(ModuleName, 2, "invalid version")
)
