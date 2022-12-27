package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/crosstalk module sentinel errors
var (
	ErrSample  = sdkerrors.Register(ModuleName, 1, "sample error")
	ErrInvalid = sdkerrors.Register(ModuleName, 2, "invalid")
)
