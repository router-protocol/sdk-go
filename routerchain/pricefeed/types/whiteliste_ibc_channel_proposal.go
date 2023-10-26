package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	// UpdatePriceFeederInfo defines the type for a CreateWhitelistedIBCChannelProposal
	CreateWhitelistedIBCChannel = "CreateWhitelistedIBCChannel"
)

// Assert CreateWhitelistedIBCChannelProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &CreateWhitelistedIBCChannelProposal{}

func init() {
	govtypes.RegisterProposalType(CreateWhitelistedIBCChannel)
}

func NewCreateWhitelistedIBCChannelProposal(
	title, description string, channels Channels,
) *CreateWhitelistedIBCChannelProposal {
	return &CreateWhitelistedIBCChannelProposal{title, description, channels}
}

// GetTitle returns the title of a update price feeder info proposal.
func (p *CreateWhitelistedIBCChannelProposal) GetTitle() string { return p.Title }

// GetDescription returns the description of a update price feeder info proposal.
func (p *CreateWhitelistedIBCChannelProposal) GetDescription() string { return p.Description }

// ProposalRoute returns the routing key of a update price feeder info proposal.
func (*CreateWhitelistedIBCChannelProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a update price feeder info proposal.
func (*CreateWhitelistedIBCChannelProposal) ProposalType() string { return UpdatePriceFeederInfo }

// ValidateBasic validates the update price feeder info proposal.
func (p *CreateWhitelistedIBCChannelProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}
