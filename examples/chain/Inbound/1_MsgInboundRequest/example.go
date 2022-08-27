package main

import (
	"fmt"
	"os"
	"time"

	chainclient "github.com/router-protocol/sdk-go/client/chain"
	"github.com/router-protocol/sdk-go/client/common"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	inboundTypes "github.com/router-protocol/sdk-go/routerchain/inbound/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

const (
	// Tx Agrs
	CHAIN_TYPE             = 0
	CHAIN_ID               = "137"
	EVENT_NONCE            = 1
	BLOCK_HEIGHT           = 23423
	SOURCE_SENDER          = "0xdE23C5FfC7B045b48F0B85ADA2c518d213d9e24F"
	SOURCE_TX_HASH         = "0x93eafe329a93d3b83b7bc34852c1bcf0ed2094a1634e1b0296c14f4d156cf141"
	ROUTER_BRIDGE_CONTRACT = "router14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s00ztvk"
	PAYLOAD                = "router"
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
	msg := inboundTypes.NewMsgInboundRequest(senderAddress.String(), multichainTypes.ChainType(CHAIN_TYPE), CHAIN_ID, EVENT_NONCE, BLOCK_HEIGHT, SOURCE_SENDER, SOURCE_TX_HASH, ROUTER_BRIDGE_CONTRACT, []byte(PAYLOAD))

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		// common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("100000000000000router"),
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
