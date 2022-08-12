package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/attestation module sentinel errors
var (
	ErrSample                   = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInternal                 = sdkerrors.Register(ModuleName, 1101, "internal")
	ErrDuplicate                = sdkerrors.Register(ModuleName, 1102, "duplicate")
	ErrInvalid                  = sdkerrors.Register(ModuleName, 1103, "invalid")
	ErrTimeout                  = sdkerrors.Register(ModuleName, 1104, "timeout")
	ErrUnknown                  = sdkerrors.Register(ModuleName, 1105, "unknown")
	ErrEmpty                    = sdkerrors.Register(ModuleName, 1106, "empty")
	ErrOutdated                 = sdkerrors.Register(ModuleName, 1107, "outdated")
	ErrUnsupported              = sdkerrors.Register(ModuleName, 1108, "unsupported")
	ErrNonContiguousEventNonce  = sdkerrors.Register(ModuleName, 1109, "non contiguous event nonce, expected: %v received: %v")
	ErrResetDelegateKeys        = sdkerrors.Register(ModuleName, 1110, "can not set orchestrator addresses more than once")
	ErrMismatched               = sdkerrors.Register(ModuleName, 1111, "mismatched")
	ErrNoValidators             = sdkerrors.Register(ModuleName, 1112, "no bonded validators in active set")
	ErrInvalidValAddress        = sdkerrors.Register(ModuleName, 1113, "invalid validator address in current valset %v")
	ErrInvalidEthAddress        = sdkerrors.Register(ModuleName, 1114, "discovered invalid eth address stored for validator %v")
	ErrInvalidValset            = sdkerrors.Register(ModuleName, 1115, "generated invalid valset")
	ErrDuplicateEthereumKey     = sdkerrors.Register(ModuleName, 1116, "duplicate ethereum key")
	ErrDuplicateOrchestratorKey = sdkerrors.Register(ModuleName, 1117, "duplicate orchestrator key")
)
