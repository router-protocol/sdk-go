package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateChainConfig{}, "multichain/CreateChainConfig", nil)
	cdc.RegisterConcrete(&MsgUpdateChainConfig{}, "multichain/UpdateChainConfig", nil)
	cdc.RegisterConcrete(&MsgDeleteChainConfig{}, "multichain/DeleteChainConfig", nil)
	cdc.RegisterConcrete(&MultichainCreateChainConfigProposal{}, "multichain/MultichainCreateChainConfigProposal", nil)
	cdc.RegisterConcrete(&MultichainUpdateChainConfigProposal{}, "multichain/MultichainUpdateChainConfigProposal", nil)
	cdc.RegisterConcrete(&MultichainDeleteChainConfigProposal{}, "multichain/MultichainDeleteChainConfigProposal", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateChainConfig{},
		&MsgUpdateChainConfig{},
		&MsgDeleteChainConfig{},
	)

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&MultichainCreateChainConfigProposal{},
		&MultichainUpdateChainConfigProposal{},
		&MultichainDeleteChainConfigProposal{},
	)

	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
