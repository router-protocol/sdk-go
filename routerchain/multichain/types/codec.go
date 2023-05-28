package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

	//sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	grpc "google.golang.org/grpc"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MultichainCreateChainConfigProposal{}, "multichain/MultichainCreateChainConfigProposal", nil)
	cdc.RegisterConcrete(&MultichainUpdateChainConfigProposal{}, "multichain/MultichainUpdateChainConfigProposal", nil)
	cdc.RegisterConcrete(&MultichainDeleteChainConfigProposal{}, "multichain/MultichainDeleteChainConfigProposal", nil)
	cdc.RegisterConcrete(&MultichainCreateContractConfigProposal{}, "multichain/MultichainCreateContractConfigProposal", nil)
	cdc.RegisterConcrete(&MultichainUpdateContractConfigProposal{}, "multichain/MultichainUpdateContractConfigProposal", nil)
	cdc.RegisterConcrete(&MultichainDeleteContractConfigProposal{}, "multichain/MultichainDeleteContractConfigProposal", nil)
	// this line is used by starport scaffolding # 2
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
	)

	// this line is used by starport scaffolding # 3
	msgservice.RegisterMsgServiceDesc(registry, &serviceDesc)
}

var serviceDesc = grpc.ServiceDesc{
	ServiceName: "routerprotocol.routerchain.multichain.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "multichain/tx.proto",
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
