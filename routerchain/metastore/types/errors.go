package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/metastore module sentinel errors
var (
	ErrMetaInfoNotFound        = sdkerrors.Register(ModuleName, 1100, "meta info not found")
	ErrFeePayerAlreadyApproved = sdkerrors.Register(ModuleName, 1101, "fee payer already approved")
	ErrFeePayerMismatch        = sdkerrors.Register(ModuleName, 1102, "fee payer mismatch")
)
