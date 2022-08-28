package main

import (
	"context"
	"fmt"
	"os"
	"time"

	chainclient "github.com/router-protocol/sdk-go/client/chain"
	"github.com/router-protocol/sdk-go/client/common"
	log "github.com/xlab/suplog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	outboundTypes "github.com/router-protocol/sdk-go/routerchain/outbound/types"

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

	// prepare tx msg
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	grpcConn := chainClient.QueryClient()
	waitForService(ctx, grpcConn)
	outboundTypesGrpcQuerier := outboundTypes.NewQueryClient(grpcConn)

	outgoingBatchTxs, err := outboundTypesGrpcQuerier.OutgoingBatchTxAll(ctx, &outboundTypes.QueryAllOutgoingBatchTxRequest{})
	fmt.Println("OutgoingBatchTxs", outgoingBatchTxs, "Checkpoint", outgoingBatchTxs.OutgoingBatchTx[0].GetCheckpoint("routerchain"))

	signature := ""
	msg := outboundTypes.NewMsgOutgoingBatchConfirm(senderAddress.String(), multichainTypes.ChainType(CHAIN_TYPE), CHAIN_ID, ROUTER_BRIDGE_CONTRACT, BATCH_NONCE, ORCHESTRATOR_ETH_ADDRESS, signature)

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

// waitForService awaits an active ClientConn to a GRPC service.
func waitForService(ctx context.Context, clientconn *grpc.ClientConn) {
	for {
		select {
		case <-ctx.Done():
			log.Fatalln("GRPC service wait timed out")
		default:
			state := clientconn.GetState()

			if state != connectivity.Ready {
				log.WithField("state", state.String()).Warningln("state of GRPC connection not ready")
				time.Sleep(5 * time.Second)
				continue
			}

			return
		}
	}
}
