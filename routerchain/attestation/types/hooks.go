package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AttestationHooks interface {
	AttestationObservedHook(ctx sdk.Context, claim Claim)
}

var _ AttestationHooks = MultiAttestationHooks{}

type MultiAttestationHooks []AttestationHooks

func NewMultiAttestationHooks(hooks ...AttestationHooks) MultiAttestationHooks {
	return hooks
}

func (h MultiAttestationHooks) AttestationObservedHook(ctx sdk.Context, claim Claim) {
	for i := range h {
		h[i].AttestationObservedHook(ctx, claim)
	}
}
