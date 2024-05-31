package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/attestation module sentinel errors
var (
	ErrSample                   = errorsmod.Register(ModuleName, 1100, "sample error")
	ErrInternal                 = errorsmod.Register(ModuleName, 1101, "internal")
	ErrDuplicate                = errorsmod.Register(ModuleName, 1102, "duplicate")
	ErrInvalid                  = errorsmod.Register(ModuleName, 1103, "invalid")
	ErrTimeout                  = errorsmod.Register(ModuleName, 1104, "timeout")
	ErrUnknown                  = errorsmod.Register(ModuleName, 1105, "unknown")
	ErrEmpty                    = errorsmod.Register(ModuleName, 1106, "empty")
	ErrOutdated                 = errorsmod.Register(ModuleName, 1107, "outdated")
	ErrUnsupported              = errorsmod.Register(ModuleName, 1108, "unsupported")
	ErrNonContiguousEventNonce  = errorsmod.Register(ModuleName, 1109, "non contiguous event nonce, expected: %v received: %v")
	ErrResetDelegateKeys        = errorsmod.Register(ModuleName, 1110, "can not set orchestrator addresses more than once")
	ErrMismatched               = errorsmod.Register(ModuleName, 1111, "mismatched")
	ErrNoValidators             = errorsmod.Register(ModuleName, 1112, "no bonded validators in active set")
	ErrInvalidValAddress        = errorsmod.Register(ModuleName, 1113, "invalid validator address in current valset %v")
	ErrInvalidEthAddress        = errorsmod.Register(ModuleName, 1114, "discovered invalid eth address stored for validator %v")
	ErrInvalidValset            = errorsmod.Register(ModuleName, 1115, "generated invalid valset")
	ErrDuplicateEthereumKey     = errorsmod.Register(ModuleName, 1116, "duplicate ethereum key")
	ErrDuplicateOrchestratorKey = errorsmod.Register(ModuleName, 1117, "duplicate orchestrator key")
)
