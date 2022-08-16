package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type (
	ExtensionOptionsWeb3TxI interface{}
)

// RegisterInterfaces registers the tendermint concrete client-related
// implementations and interfaces.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface("routerprotocol.routerchain.types.EthAccount", (*authtypes.AccountI)(nil))

	registry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&EthAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&EthAccount{},
	)

	registry.RegisterInterface("routerprotocol.routerchain.types.ExtensionOptionsWeb3Tx", (*ExtensionOptionsWeb3TxI)(nil))
	registry.RegisterImplementations(
		(*ExtensionOptionsWeb3TxI)(nil),
		&ExtensionOptionsWeb3Tx{},
	)
}
