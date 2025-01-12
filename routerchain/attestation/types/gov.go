package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalResetAttestationStates = "ResetAttestationStates"
)

func init() {
	govtypes.RegisterProposalType(ProposalResetAttestationStates)
}

var (
	_ govtypes.Content = &ResetAttestationStatesProposal{}
)

func (p *ResetAttestationStatesProposal) ProposalRoute() string { return RouterKey }

func (p *ResetAttestationStatesProposal) ProposalType() string {
	return ProposalResetAttestationStates
}

func (p *ResetAttestationStatesProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}
