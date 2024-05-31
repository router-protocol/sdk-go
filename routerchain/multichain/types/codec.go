package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MultichainCreateChainConfigProposal{}, "multichain/MultichainCreateChainConfigProposal", nil)
	cdc.RegisterConcrete(&MultichainUpdateChainConfigProposal{}, "multichain/MultichainUpdateChainConfigProposal", nil)
	cdc.RegisterConcrete(&MultichainDeleteChainConfigProposal{}, "multichain/MultichainDeleteChainConfigProposal", nil)
	cdc.RegisterConcrete(&MultichainCreateContractConfigProposal{}, "multichain/MultichainCreateContractConfigProposal", nil)
	cdc.RegisterConcrete(&MultichainUpdateContractConfigProposal{}, "multichain/MultichainUpdateContractConfigProposal", nil)
	cdc.RegisterConcrete(&MultichainDeleteContractConfigProposal{}, "multichain/MultichainDeleteContractConfigProposal", nil)
	cdc.RegisterConcrete(&MultichainCreateIbcRelayerConfigProposal{}, "multichain/CreateIbcRelayerConfigProposal", nil)
	cdc.RegisterConcrete(&MultichainUpdateIbcRelayerConfigProposal{}, "multichain/UpdateIbcRelayerConfigProposal", nil)
	cdc.RegisterConcrete(&MultichainDeleteIbcRelayerConfigProposal{}, "multichain/DeleteIbcRelayerConfigProposal", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) { //nolint:staticcheck
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&MultichainCreateChainConfigProposal{},
		&MultichainUpdateChainConfigProposal{},
		&MultichainDeleteChainConfigProposal{},
		&MultichainCreateContractConfigProposal{},
		&MultichainUpdateContractConfigProposal{},
		&MultichainDeleteContractConfigProposal{},
		&MultichainCreateIbcRelayerConfigProposal{},
		&MultichainUpdateIbcRelayerConfigProposal{},
		&MultichainDeleteIbcRelayerConfigProposal{},
	)

	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

// NOTE: This is required for the GetSignBytes function
func init() {
	RegisterLegacyAminoCodec(Amino)
	// RegisterCrypto(amino)
	Amino.Seal()
}
