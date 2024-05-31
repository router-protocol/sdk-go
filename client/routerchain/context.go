package routerchain

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/pkg/errors"

	keyscodec "github.com/router-protocol/sdk-go/routerchain/crypto/codec"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	crosschainTypes "github.com/router-protocol/sdk-go/routerchain/crosschain/types"
	metastoreTypes "github.com/router-protocol/sdk-go/routerchain/metastore/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	pricefeedTypes "github.com/router-protocol/sdk-go/routerchain/pricefeed/types"
	voyagerTypes "github.com/router-protocol/sdk-go/routerchain/voyager/types"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	paramproposaltypes "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	ibcapplicationtypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibccoretypes "github.com/cosmos/ibc-go/v7/modules/core/types"

	enccodec "github.com/evmos/ethermint/encoding/codec"
)

func NewCodec() (*codec.ProtoCodec, types.InterfaceRegistry) {

	interfaceRegistry := types.NewInterfaceRegistry()
	keyscodec.RegisterInterfaces(interfaceRegistry)
	enccodec.RegisterInterfaces(interfaceRegistry)

	std.RegisterInterfaces(interfaceRegistry)
	evmtypes.RegisterInterfaces(interfaceRegistry)
	attestationTypes.RegisterInterfaces(interfaceRegistry)
	metastoreTypes.RegisterInterfaces(interfaceRegistry)
	multichainTypes.RegisterInterfaces(interfaceRegistry)
	pricefeedTypes.RegisterInterfaces(interfaceRegistry)
	crosschainTypes.RegisterInterfaces(interfaceRegistry)
	voyagerTypes.RegisterInterfaces(interfaceRegistry)

	// more cosmos types
	authtypes.RegisterInterfaces(interfaceRegistry)
	authztypes.RegisterInterfaces(interfaceRegistry)
	vestingtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	distributiontypes.RegisterInterfaces(interfaceRegistry)
	evidencetypes.RegisterInterfaces(interfaceRegistry)
	// gov.RegisterInterfaces(interfaceRegistry)
	paramproposaltypes.RegisterInterfaces(interfaceRegistry)
	ibcapplicationtypes.RegisterInterfaces(interfaceRegistry)
	ibccoretypes.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	wasmtypes.RegisterInterfaces(interfaceRegistry)
	icatypes.RegisterInterfaces(interfaceRegistry)

	marshaler := codec.NewProtoCodec(interfaceRegistry)
	return marshaler, interfaceRegistry
}

// NewTxConfig initializes new Cosmos TxConfig with certain signModes enabled.
func NewTxConfig(signModes []signingtypes.SignMode) client.TxConfig {
	marshaler, _ := NewCodec()
	return tx.NewTxConfig(marshaler, signModes)
}

// NewClientContext creates a new Cosmos Client context, where chainID
// corresponds to Cosmos chain ID, fromSpec is either name of the key, or bech32-address
// of the Cosmos account. Keyring is required to contain the specified key.
func NewClientContext(
	chainId, fromSpec string, kb keyring.Keyring,
) (client.Context, error) {
	clientCtx := client.Context{}
	marshaler, interfaceRegistry := NewCodec()
	encodingConfig := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig: NewTxConfig([]signingtypes.SignMode{
			signingtypes.SignMode_SIGN_MODE_DIRECT,
		}),
	}

	var keyInfo *keyring.Record

	if kb != nil {
		addr, err := cosmostypes.AccAddressFromBech32(fromSpec)
		if err == nil {
			keyInfo, err = kb.KeyByAddress(addr)
			if err != nil {
				err = errors.Wrapf(err, "failed to load key info by address %s", addr.String())
				return clientCtx, err
			}
		} else {
			// failed to parse Bech32, is it a name?
			keyInfo, err = kb.Key(fromSpec)
			if err != nil {
				err = errors.Wrapf(err, "no key in keyring for name: %s", fromSpec)
				return clientCtx, err
			}
		}
	}

	clientCtx = newContext(
		chainId,
		encodingConfig,
		kb,
		keyInfo,
	)

	return clientCtx, nil
}

type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
}

func newContext(
	chainId string,
	encodingConfig EncodingConfig,
	kb keyring.Keyring,
	keyInfo *keyring.Record,

) client.Context {
	clientCtx := client.Context{
		ChainID:           chainId,
		Codec:             encodingConfig.Marshaler,
		InterfaceRegistry: encodingConfig.InterfaceRegistry,
		Output:            os.Stderr,
		OutputFormat:      "json",
		BroadcastMode:     "block",
		UseLedger:         false,
		Simulate:          true,
		GenerateOnly:      false,
		Offline:           false,
		SkipConfirm:       true,
		TxConfig:          encodingConfig.TxConfig,
		AccountRetriever:  authtypes.AccountRetriever{},
	}

	if keyInfo != nil {
		clientCtx = clientCtx.WithKeyring(kb)
		addr, _ := keyInfo.GetAddress()
		clientCtx = clientCtx.WithFromAddress(addr)
		clientCtx = clientCtx.WithFromName(keyInfo.Name)
		clientCtx = clientCtx.WithFrom(keyInfo.Name)
	}

	return clientCtx
}
