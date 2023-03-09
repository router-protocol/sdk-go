package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeCreateChainConfig = "CreateChainConfig"
	ProposalTypeDeleteChainConfig = "DeleteChainConfig"
	ProposalTypeUpdateChainConfig = "UpdateChainConfig"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeCreateChainConfig)
	// govtypes.RegisterProposalTypeCodec(&MultichainCreateChainConfigProposal{}, "router/CreateChainConfigProposal")
	govtypes.RegisterProposalType(ProposalTypeDeleteChainConfig)
	// govtypes.RegisterProposalTypeCodec(&MultichainDeleteChainConfigProposal{}, "router/DeleteChainConfigProposal")
	govtypes.RegisterProposalType(ProposalTypeUpdateChainConfig)
	// govtypes.RegisterProposalTypeCodec(&MultichainUpdateChainConfigProposal{}, "router/UpdateChainConfigProposal")
}

var (
	_ govtypes.Content = &MultichainCreateChainConfigProposal{}
	_ govtypes.Content = &MultichainDeleteChainConfigProposal{}
	_ govtypes.Content = &MultichainUpdateChainConfigProposal{}
)

func (p *MultichainCreateChainConfigProposal) ProposalRoute() string { return RouterKey }

func (p *MultichainCreateChainConfigProposal) ProposalType() string {
	return ProposalTypeCreateChainConfig
}

func (p *MultichainCreateChainConfigProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func (p *MultichainUpdateChainConfigProposal) ProposalRoute() string { return RouterKey }

func (p *MultichainUpdateChainConfigProposal) ProposalType() string {
	return ProposalTypeCreateChainConfig
}

func (p *MultichainUpdateChainConfigProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func (p *MultichainDeleteChainConfigProposal) ProposalRoute() string { return RouterKey }

func (p *MultichainDeleteChainConfigProposal) ProposalType() string {
	return ProposalTypeCreateChainConfig
}

func (p *MultichainDeleteChainConfigProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}
