package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/types/query"
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

	var nextKey []byte
	for {
		crosstalkAckResponse, err := chainClient.GetAllCrossTalkAckRequest(ctx, &query.PageRequest{
			Key:   nextKey,
			Limit: 50,
		})

		if err != nil {
			panic(err)
		}
		fmt.Println("crosstalkack length: ", len(crosstalkAckResponse.CrossTalkAckRequest))
		nextKey = crosstalkAckResponse.Pagination.NextKey
		if uint64(len(nextKey)) == 0 {
			break
		}
	}

	// for _, crosstalkRequest := range allCrosstalkRequests.CrossTalkRequest {
	// 	fmt.Println("crosstalk Request", crosstalkRequest)
	// 	claimHash, err := crosstalkRequest.ClaimHash()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	crosstalkConfirmations, err := chainClient.GetAllCrosstalkRequestConfirmations(ctx, &query.PageRequest{}, uint64(crosstalkRequest.SourceChainType), crosstalkRequest.SourceChainId, crosstalkRequest.EventNonce, claimHash)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("crosstalkConfirmations", crosstalkConfirmations)
	// }
}
