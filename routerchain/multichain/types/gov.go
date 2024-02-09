package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeCreateChainConfig = "CreateChainConfig"
	ProposalTypeDeleteChainConfig = "DeleteChainConfig"
	ProposalTypeUpdateChainConfig = "UpdateChainConfig"

	ProposalTypeCreateContractConfig = "CreateContractConfig"
	ProposalTypeDeleteContractConfig = "DeleteContractConfig"
	ProposalTypeUpdateContractConfig = "UpdateContractConfig"

	ProposalTypeCreateIbcRelayerConfig = "CreateIbcRelayerConfig"
	ProposalTypeDeleteIbcRelayerConfig = "DeleteIbcRelayerConfig"
	ProposalTypeUpdateIbcRelayerConfig = "UpdateIbcRelayerConfig"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeCreateChainConfig)
	// govtypes.RegisterProposalTypeCodec(&MultichainCreateChainConfigProposal{}, "router/CreateChainConfigProposal")
	govtypes.RegisterProposalType(ProposalTypeDeleteChainConfig)
	// govtypes.RegisterProposalTypeCodec(&MultichainDeleteChainConfigProposal{}, "router/DeleteChainConfigProposal")
	govtypes.RegisterProposalType(ProposalTypeUpdateChainConfig)
	// govtypes.RegisterProposalTypeCodec(&MultichainUpdateChainConfigProposal{}, "router/UpdateChainConfigProposal")

	govtypes.RegisterProposalType(ProposalTypeCreateContractConfig)
	govtypes.RegisterProposalType(ProposalTypeDeleteContractConfig)
	govtypes.RegisterProposalType(ProposalTypeUpdateContractConfig)

	govtypes.RegisterProposalType(ProposalTypeCreateIbcRelayerConfig)
	govtypes.RegisterProposalType(ProposalTypeDeleteIbcRelayerConfig)
	govtypes.RegisterProposalType(ProposalTypeUpdateIbcRelayerConfig)
}

var (
	_ govtypes.Content = &MultichainCreateChainConfigProposal{}
	_ govtypes.Content = &MultichainDeleteChainConfigProposal{}
	_ govtypes.Content = &MultichainUpdateChainConfigProposal{}

	_ govtypes.Content = &MultichainCreateContractConfigProposal{}
	_ govtypes.Content = &MultichainDeleteContractConfigProposal{}
	_ govtypes.Content = &MultichainUpdateContractConfigProposal{}

	_ govtypes.Content = &MultichainCreateIbcRelayerConfigProposal{}
	_ govtypes.Content = &MultichainDeleteIbcRelayerConfigProposal{}
	_ govtypes.Content = &MultichainUpdateIbcRelayerConfigProposal{}
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

func (p *MultichainCreateContractConfigProposal) ProposalRoute() string { return RouterKey }

func (p *MultichainCreateContractConfigProposal) ProposalType() string {
	return ProposalTypeCreateContractConfig
}

func (p *MultichainCreateContractConfigProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func (p *MultichainUpdateContractConfigProposal) ProposalRoute() string { return RouterKey }

func (p *MultichainUpdateContractConfigProposal) ProposalType() string {
	return ProposalTypeUpdateContractConfig
}

func (p *MultichainUpdateContractConfigProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func (p *MultichainDeleteContractConfigProposal) ProposalRoute() string { return RouterKey }

func (p *MultichainDeleteContractConfigProposal) ProposalType() string {
	return ProposalTypeDeleteContractConfig
}

func (p *MultichainDeleteContractConfigProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func (p *MultichainCreateIbcRelayerConfigProposal) ProposalRoute() string { return RouterKey }

func (p *MultichainCreateIbcRelayerConfigProposal) ProposalType() string {
	return ProposalTypeCreateIbcRelayerConfig
}

func (p *MultichainCreateIbcRelayerConfigProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func (p *MultichainUpdateIbcRelayerConfigProposal) ProposalRoute() string { return RouterKey }

func (p *MultichainUpdateIbcRelayerConfigProposal) ProposalType() string {
	return ProposalTypeUpdateIbcRelayerConfig
}

func (p *MultichainUpdateIbcRelayerConfigProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func (p *MultichainDeleteIbcRelayerConfigProposal) ProposalRoute() string { return RouterKey }

func (p *MultichainDeleteIbcRelayerConfigProposal) ProposalType() string {
	return ProposalTypeDeleteIbcRelayerConfig
}

func (p *MultichainDeleteIbcRelayerConfigProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}
