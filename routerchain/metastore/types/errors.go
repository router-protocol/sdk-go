package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/metastore module sentinel errors
var (
	ErrMetaInfoNotFound        = errorsmod.Register(ModuleName, 1100, "meta info not found")
	ErrFeePayerAlreadyApproved = errorsmod.Register(ModuleName, 1101, "fee payer already approved")
	ErrFeePayerMismatch        = errorsmod.Register(ModuleName, 1102, "fee payer mismatch")
	ErrFeePayerAlreadyRevoked  = errorsmod.Register(ModuleName, 1103, "fee payer already revoked")
	ErrInvalidAmount           = errorsmod.Register(ModuleName, 1104, "invalid amount")
	ErrBurnAmountTooLarge      = errorsmod.Register(ModuleName, 1105, "burn amount should be less than the account balance")
	ErrUnknown                 = errorsmod.Register(ModuleName, 1106, "unknown")
)
