package main

import (
	"context"
	"time"

	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/router-protocol/sdk-go/client/evm/gateway"
	chainclient "github.com/router-protocol/sdk-go/client/routerchain"
	"github.com/router-protocol/sdk-go/examples/lightweight-relayer/relayer"
)

const (
	// Tx Agrs
	NETWORK_NAME              = "local"
	ORCHESTRATOR_KEYRING_FROM = "genesis"
	PASSPHRASE                = "12345678"

	RELAYER_PRIVATE_KEY    = "843EE93DA70C08B88C726C43329B8DA92CC26AEFE2A0F3F33832A0540E66EA53"
	RELAYER_ETH_ADDRESS    = "0x46e551558388e670ac58e6C84c0759Ca2dCBe6e3"
	RELAYER_ROUTER_ADDRESS = "0x46e551558388e670ac58e6C84c0759Ca2dCBe6e3"
	ETH_RPC                = "https://ropsten.infura.io/v3/74c8b06d2481408c86fe936c11657def"
)

var (
	GATEWAY_CONTRACT_ADDRESS = ethcmn.HexToAddress("0xEE39624b883fe60b1B6b77686ecB0Be10c990341")
)

func main() {
	// eth client
	ethClient, err := ethclient.Dial(ETH_RPC)
	if err != nil {
		panic(err)
	}

	// INITIALIZE RELAYER
	gatewayContractClient := gateway.NewGatewayContractClient(ETH_RPC, GATEWAY_CONTRACT_ADDRESS)
	routerchainClient := chainclient.InitialiseChainClient(NETWORK_NAME, ORCHESTRATOR_KEYRING_FROM, PASSPHRASE, "")
	relayer := relayer.NewRelayer(ethClient, gatewayContractClient, routerchainClient, RELAYER_PRIVATE_KEY, RELAYER_ETH_ADDRESS, RELAYER_ROUTER_ADDRESS)

	// TRIGGER ACTIONS
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	relayer.SubmitBatchTxToGateway(ctx, routerchainClient)

}
