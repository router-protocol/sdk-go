package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/crosstalk module sentinel errors
var (
	ErrSample                      = sdkerrors.Register(ModuleName, 1, "sample error")
	ErrInvalid                     = sdkerrors.Register(ModuleName, 2, "invalid")
	ErrCrosstalkNotFound           = sdkerrors.Register(ModuleName, 3, "Crosstalk request not found")
	ErrCrosstalkAckNotFound        = sdkerrors.Register(ModuleName, 4, "Crosstalk Ack request not found")
	ErrCrosstalkFeePayerAlreadySet = sdkerrors.Register(ModuleName, 5, "Fee payer already set for crosstalk request")
	ErrCrosstalkFeePayerNotSet     = sdkerrors.Register(ModuleName, 6, "Fee payer not set for crosstalk request")
)
