package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeCreateIBCConfig = "CreateIBCConfig"
	ProposalTypeUpdateIBCConfig = "UpdateIBCConfig"
	// ProposalTypeDeleteIBCConfig = "DeleteIBCConfig"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeCreateIBCConfig)
	govtypes.RegisterProposalType(ProposalTypeUpdateIBCConfig)
	// govtypes.RegisterProposalType(ProposalTypeDeleteIBCConfig)
}

var (
	_ govtypes.Content = &CrosschainCreateIBCConfigProposal{}
	_ govtypes.Content = &CrosschainUpdateIBCConfigProposal{}
	// _ govtypes.Content = &CrosschainDeleteIBCConfigProposal{}
)

func (p *CrosschainCreateIBCConfigProposal) ProposalRoute() string { return RouterKey }

func (p *CrosschainCreateIBCConfigProposal) ProposalType() string {
	return ProposalTypeCreateIBCConfig
}

func (p *CrosschainCreateIBCConfigProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func (p *CrosschainUpdateIBCConfigProposal) ProposalRoute() string { return RouterKey }

func (p *CrosschainUpdateIBCConfigProposal) ProposalType() string {
	return ProposalTypeUpdateIBCConfig
}

func (p *CrosschainUpdateIBCConfigProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

// func (p *CrosschainDeleteIBCConfigProposal) ProposalRoute() string { return RouterKey }

// func (p *CrosschainDeleteIBCConfigProposal) ProposalType() string {
// 	return ProposalTypeCreateIBCConfig
// }

// func (p *CrosschainDeleteIBCConfigProposal) ValidateBasic() error {
// 	err := govtypes.ValidateAbstract(p)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
