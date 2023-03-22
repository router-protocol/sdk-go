package main

import (
	"context"
	"fmt"
	"os"
	"time"

	chainclient "github.com/router-protocol/sdk-go/client/routerchain"
	"github.com/router-protocol/sdk-go/client/routerchain/common"

	metastoreTypes "github.com/router-protocol/sdk-go/routerchain/metastore/types"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

func main() {
	network := common.LoadNetwork("devnet-alpha", "k8s")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Network", network)
	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.routerd",
		"routerd",
		"file",
		"gaurav",
		"12345678",
		"", // keyring will be used if pk not provided
		false,
	)

	if err != nil {
		panic(err)
	}

	// initialize grpc client
	clientCtx, err := chainclient.NewClientContext(network.ChainId, senderAddress.String(), cosmosKeyring)
	if err != nil {
		fmt.Println(err)
	}
	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmRPC)

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)

	routerChainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		// common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("100000000000000route"),
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	allMetaInfo, err := routerChainClient.GetAllMetaInfo(ctx)
	if err != nil {
		fmt.Println("Error fetching all meta info")
		panic(err)
	}

	fmt.Println("allMetaInfo", allMetaInfo)
	for _, metaInfo := range allMetaInfo.MetaInfo {

		if metaInfo.FeePayerApproved {
			continue
		}
		// prepare tx msg
		msg := metastoreTypes.NewMsgApproveFeepayerRequest(senderAddress.String(), metaInfo.ChainType, metaInfo.ChainId, metaInfo.DappAddress)

		//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
		err = routerChainClient.QueueBroadcastMsg(msg)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Second * 5)
		gasFee, err := routerChainClient.GetGasFee()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("gas fee:", gasFee)
	}
}
