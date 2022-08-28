package main

import (
	"context"
	"fmt"
	"os"
	"time"

	chainclient "github.com/router-protocol/sdk-go/client/chain"
	"github.com/router-protocol/sdk-go/client/common"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	outboundTypes "github.com/router-protocol/sdk-go/routerchain/outbound/types"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

const (
	// Tx Agrs
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

	// prepare tx msg
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	outgoingBatchTxs, err := chainClient.GetAllOutgoingBatchTx(ctx)
	fmt.Println("OutgoingBatchTxs", outgoingBatchTxs, "Checkpoint", outgoingBatchTxs.OutgoingBatchTx[0].GetCheckpoint("routerchain"))

	outgoingBatchTx := outgoingBatchTxs.OutgoingBatchTx[0]
	signature := ""
	msg := outboundTypes.NewMsgOutgoingBatchConfirm(senderAddress.String(), outgoingBatchTx.GetDestinationChainType(),
		outgoingBatchTx.GetDestinationChainId(), outgoingBatchTx.GetSourceAddress(), outgoingBatchTx.GetNonce(), ORCHESTRATOR_ETH_ADDRESS, signature)

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
