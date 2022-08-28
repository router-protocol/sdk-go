package main

import (
	"context"
	"fmt"
	"os"
	"time"

	chainclient "github.com/router-protocol/sdk-go/client/routerchain"
	"github.com/router-protocol/sdk-go/client/routerchain/common"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

const (
	// Tx Agrs
	CHAIN_TYPE               = 0
	CHAIN_ID                 = "137"
	BATCH_NONCE              = 1
	ROUTER_BRIDGE_CONTRACT   = "router14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s00ztvk"
	ORCHESTRATOR_ETH_ADDRESS = "0xdE23C5FfC7B045b48F0B85ADA2c518d213d9e24F"
	IS_ATOMIC                = false
)

var (
	RELAYER_FEE     = sdktypes.NewCoin("router", sdktypes.NewIntFromUint64(uint64(100000000)))
	OUTGOING_TX_FEE = sdktypes.NewCoin("router", sdktypes.NewIntFromUint64(uint64(900000000)))
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

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		// common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("100000000000000router"),
	)

	if err != nil {
		fmt.Println(err)
	}

	// prepare query request
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	outgoingBatchTxs, err := chainClient.GetAllOutgoingBatchTx(ctx)
	fmt.Println("OutgoingBatchTxs", outgoingBatchTxs, "Checkpoint", outgoingBatchTxs.OutgoingBatchTx[0].GetCheckpoint("routerchain"))
}
