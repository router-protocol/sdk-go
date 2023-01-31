package main

import (
	"context"
	"fmt"
	"os"
	"time"

	chainclient "github.com/router-protocol/sdk-go/client/routerchain"
	"github.com/router-protocol/sdk-go/client/routerchain/common"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

func main() {
	network := common.LoadNetwork("devnet", "k8s")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Network", network)
	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.routerd",
		"routerd",
		"file",
		"orchestrator",
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

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		// common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("100000000000000route"),
	)

	if err != nil {
		fmt.Println(err)
	}

	// prepare query request
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)

	chainType := uint64(0)
	chainID := "80001"
	eventNonce := uint64(32)
	crosstalkResponse, err := chainClient.GetCrossTalkRequest(ctx, chainType, chainID, eventNonce)
	if err != nil {
		panic(err)
	}

	fmt.Println("crosstalkResponse", crosstalkResponse)
}
