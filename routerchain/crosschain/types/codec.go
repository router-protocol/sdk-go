package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	attestationtypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCrosschainRequest{}, "crosschain/CrosschainRequest", nil)
	cdc.RegisterConcrete(&MsgConfirmCrosschainRequest{}, "crosschain/ConfirmCrosschainRequest", nil)
	cdc.RegisterConcrete(&MsgCrosschainAckRequest{}, "crosschain/CrosschainAckRequest", nil)
	cdc.RegisterConcrete(&MsgConfirmCrosschainAckRequest{}, "crosschain/ConfirmCrosschainAckRequest", nil)
	cdc.RegisterConcrete(&MsgCrosschainAckReceipt{}, "crosschain/CrosschainAckReceipt", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) { //nolint:staticcheck
	legacy.RegisterAminoMsg(cdc, &MsgCrosschainRequest{}, "crosschain/CrosschainRequest")
	legacy.RegisterAminoMsg(cdc, &MsgConfirmCrosschainRequest{}, "crosschain/ConfirmCrosschainRequest")
	legacy.RegisterAminoMsg(cdc, &MsgCrosschainAckRequest{}, "crosschain/CrosschainAckRequest")
	legacy.RegisterAminoMsg(cdc, &MsgConfirmCrosschainAckRequest{}, "crosschain/ConfirmCrosschainAckRequest")
	legacy.RegisterAminoMsg(cdc, &MsgCrosschainAckReceipt{}, "crosschain/CrosschainAckReceipt")
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*attestationtypes.Claim)(nil),
		&MsgCrosschainRequest{}, &MsgCrosschainAckRequest{}, &MsgCrosschainAckReceipt{},
	)

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

// NOTE: This is required for the GetSignBytes function
func init() {
	RegisterLegacyAminoCodec(Amino)
	// RegisterCrypto(amino)
	Amino.Seal()
}
