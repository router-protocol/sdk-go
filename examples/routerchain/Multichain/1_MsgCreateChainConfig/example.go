package main

import (
	"fmt"
	"os"
	"time"

	chainclient "github.com/router-protocol/sdk-go/client/routerchain"
	"github.com/router-protocol/sdk-go/client/routerchain/common"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

const (
	// Tx Agrs
	CHAIN_ID                   = "137"
	CHAIN_NAME                 = "Ethereum"
	SYMBOL                     = "ETH"
	CHAIN_TYPE                 = 0
	CONFIRMATIONS_REQUIRED     = 12
	GATEWAY_CONTRACT_ADDRESS   = "0xef1c3fd2a191b557813df429fa095ac0cae0f159"
	GATEWAY_CONTRACT_HEIGHT    = 15307415
	ROUTER_CONTRACT_ADDRESS    = "0x0fb1097e5b4bce1b5ca83f53e29752d60dfa10aa"
	LAST_OBSERVED_EVENT_NONCE  = 0
	LAST_OBSERVED_VALSET_NONCE = 0
)

func main() {
	network := common.LoadNetwork("local", "k8s")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Network", network)
	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.routerd",
		"routerd",
		"file",
		"genesis",
		"12345678",
		"", // keyring will be used if pk not provided
		false,
	)

	if err != nil {
		panic(err)
	}

	// initialize grpc client
	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)

	if err != nil {
		fmt.Println(err)
	}

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmRPC)

	// prepare tx msg
	msg := multichainTypes.NewMsgCreateChainConfig(senderAddress.String(), CHAIN_ID, CHAIN_NAME, SYMBOL, CHAIN_TYPE, CONFIRMATIONS_REQUIRED, GATEWAY_CONTRACT_ADDRESS, GATEWAY_CONTRACT_HEIGHT, ROUTER_CONTRACT_ADDRESS,
		LAST_OBSERVED_EVENT_NONCE, LAST_OBSERVED_VALSET_NONCE)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		// common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("100000000000000route"),
	)

	if err != nil {
		fmt.Println(err)
	}

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	err = chainClient.QueueBroadcastMsg(msg)

	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 5)

	gasFee, err := chainClient.GetGasFee()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("gas fee:", gasFee)
}
