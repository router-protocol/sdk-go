package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	chainclient "github.com/router-protocol/sdk-go/client/routerchain"
	"github.com/router-protocol/sdk-go/client/routerchain/common"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	outboundTypes "github.com/router-protocol/sdk-go/routerchain/outbound/types"

	sdktypes "github.com/cosmos/cosmos-sdk/types"

	gatewayWrapper "github.com/router-protocol/router-gateway-contracts/evm/build/bindings/go/GatewayUpgradeable"
)

const (
	// Tx Agrs
	ORCHESTRATOR_ETH_ADDRESS = "0x46e551558388e670ac58e6C84c0759Ca2dCBe6e3"
	ORCHESTRATOR_PRIVATE_KEY = "843EE93DA70C08B88C726C43329B8DA92CC26AEFE2A0F3F33832A0540E66EA53"
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
	privateKey, err := crypto.HexToECDSA(ORCHESTRATOR_PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	outgoingBatchTxs, err := chainClient.GetAllOutgoingBatchTx(ctx)
	checkpoint := outgoingBatchTxs.OutgoingBatchTx[0].GetCheckpoint("routerchain")
	protectedHash := accounts.TextHash(checkpoint)

	outgoingBatchTx := outgoingBatchTxs.OutgoingBatchTx[0]

	signature, err := crypto.Sign(protectedHash, privateKey)
	if err != nil {
		err = errors.New("failed to sign validator address")
		fmt.Println("Error", err)
	}

	msg := outboundTypes.NewMsgOutgoingBatchConfirm(senderAddress.String(), outgoingBatchTx.GetDestinationChainType(),
		outgoingBatchTx.GetDestinationChainId(), outgoingBatchTx.GetSourceAddress(), outgoingBatchTx.GetNonce(), ORCHESTRATOR_ETH_ADDRESS, ethcmn.Bytes2Hex(signature))

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	err = chainClient.QueueBroadcastMsg(msg)

	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 5)

	gatewayAbi := gatewayWrapper.GatewayABI
	fmt.Println(gatewayAbi)

	gasFee, err := chainClient.GetGasFee()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("gas fee:", gasFee)
}
