package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCrosschainRequest{}, "crosschain/CrosschainRequest", nil)
	cdc.RegisterConcrete(&MsgConfirmCrosschainRequest{}, "crosschain/ConfirmCrosschainRequest", nil)
	cdc.RegisterConcrete(&MsgCrosschainAckRequest{}, "crosschain/CrosschainAckRequest", nil)
	cdc.RegisterConcrete(&MsgConfirmCrosschainAckRequest{}, "crosschain/ConfirmCrosschainAckRequest", nil)
	cdc.RegisterConcrete(&MsgCrosschainAckReceipt{}, "crosschain/CrosschainAckReceipt", nil)
	cdc.RegisterConcrete(&MsgCreateRelayerConfig{}, "crosschain/CreateRelayerConfig", nil)
	cdc.RegisterConcrete(&MsgUpdateRelayerConfig{}, "crosschain/UpdateRelayerConfig", nil)
	cdc.RegisterConcrete(&MsgDeleteRelayerConfig{}, "crosschain/DeleteRelayerConfig", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCrosschainRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgConfirmCrosschainRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCrosschainAckRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgConfirmCrosschainAckRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCrosschainAckReceipt{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRelayerConfig{},
		&MsgUpdateRelayerConfig{},
		&MsgDeleteRelayerConfig{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&CrosschainCreateIBCConfigProposal{},
		&CrosschainUpdateIBCConfigProposal{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
