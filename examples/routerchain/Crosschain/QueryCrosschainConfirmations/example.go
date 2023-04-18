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

const (
	// Tx Agrs
	ORCHESTRATOR_ETH_ADDRESS = "0x46e551558388e670ac58e6C84c0759Ca2dCBe6e3"
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

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	allCrosschainRequests, err := chainClient.GetAllCrosschainRequests(ctx, nil)
	// fmt.Println("allCrosschainRequests", allCrosschainRequests, "error", err)
	if err != nil {
		fmt.Println(err)
	}

	for _, crosschainRequest := range allCrosschainRequests.CrosschainRequest {
		if crosschainRequest.RequestIdentifier != 7 {
			continue
		}

		fmt.Println("crosschainRequest", crosschainRequest)
		claimhash, _ := crosschainRequest.ClaimHash()
		allCrosschainRequestConfirmations, err := chainClient.GetAllCrosschainRequestConfirmations(ctx, nil, crosschainRequest.SrcChainId, crosschainRequest.RequestIdentifier, claimhash)
		fmt.Println("allCrosschainRequestConfirmations", allCrosschainRequestConfirmations, "error", err)
		if err != nil {
			fmt.Println(err)
		}
	}

	time.Sleep(time.Second * 5)

	gasFee, err := chainClient.GetGasFee()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("gas fee:", gasFee)
}
