package orchestrator

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/accounts"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	gatewayWrapper "github.com/router-protocol/router-gateway-contracts/evm/wrappers"
	"github.com/router-protocol/sdk-go/client/evm/gateway"
	chainclient "github.com/router-protocol/sdk-go/client/routerchain"
	routerclient "github.com/router-protocol/sdk-go/client/routerchain"
	inboundTypes "github.com/router-protocol/sdk-go/routerchain/inbound/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	outboundTypes "github.com/router-protocol/sdk-go/routerchain/outbound/types"
)

type Orchestrator interface {
	fetchAndProcessGatewayEvents()
	ConfirmOutgoingBatches(ctx context.Context)
}

type orchestrator struct {
	clientContext         client.Context
	ethClient             *ethclient.Client
	gatewayContractClient *gateway.GatewayContractClient
	routerChainClient     routerclient.ChainClient

	ethPrivatekey string
	ethAddress    string
}

func NewOrchestrator(ethClient *ethclient.Client, gatewayContractClient *gateway.GatewayContractClient, routerChainClient routerclient.ChainClient, orchestratorEthPrivKey string, orchestratorEthAddress string) *orchestrator {

	return &orchestrator{
		ethClient:             ethClient,
		gatewayContractClient: gatewayContractClient,
		routerChainClient:     routerChainClient,
		ethPrivatekey:         orchestratorEthPrivKey,
		ethAddress:            orchestratorEthAddress,
	}
}

func (orchestrator *orchestrator) FetchAndProcessGatewayEvents(startBlock uint64, endBlock uint64) {

	// fmt.Println("LatestBlockNumber", currentBlock)
	sendToRouterEvents := orchestrator.gatewayContractClient.QueryGatewaySendToRouterEvents(startBlock, endBlock)
	fmt.Println("Found SendToRouter Events", sendToRouterEvents)
	for _, sendToRouterEvent := range sendToRouterEvents {
		ctx, _ := context.WithTimeout(context.Background(), time.Minute)
		sendInboundRequest(ctx, orchestrator.routerChainClient, sendToRouterEvent)
	}

}

func sendInboundRequest(ctx context.Context, chainClient chainclient.ChainClient, sendToRouterEvent *gatewayWrapper.GatewaySendToRouterEvent) {
	// prepare tx msg
	msg := inboundTypes.NewMsgInboundRequest(chainClient.FromAddress().String(),
		multichainTypes.ChainType(sendToRouterEvent.ChainType.Int64()),
		sendToRouterEvent.ChainId, sendToRouterEvent.EventNonce.Uint64(),
		sendToRouterEvent.Raw.BlockNumber,
		sendToRouterEvent.Sender.Hex(),
		sendToRouterEvent.Raw.TxHash.Hex(),
		sendToRouterEvent.RouterBridgeContract, sendToRouterEvent.Payload)

	fmt.Println("Sending Inbound Request", msg)
	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	err := chainClient.QueueBroadcastMsg(msg)

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

func (orchestrator *orchestrator) ConfirmOutgoingBatches() {

	// prepare tx msg
	privateKey, err := crypto.HexToECDSA(orchestrator.ethPrivatekey)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	outgoingBatchTxs, err := orchestrator.routerChainClient.GetAllOutgoingBatchTx(ctx)
	outgoingBatchTx := outgoingBatchTxs.OutgoingBatchTx[1]
	checkpoint := outgoingBatchTx.GetCheckpoint("routerchain")
	protectedHash := accounts.TextHash(checkpoint)

	signature, err := crypto.Sign(protectedHash, privateKey)
	if err != nil {
		err = errors.New("failed to sign validator address")
		fmt.Println("Error", err)
	}

	fmt.Println("sending outgoing batch confirm")
	msg := outboundTypes.NewMsgOutgoingBatchConfirm(orchestrator.routerChainClient.FromAddress().String(), outgoingBatchTx.GetDestinationChainType(),
		outgoingBatchTx.GetDestinationChainId(), outgoingBatchTx.GetSourceAddress(), outgoingBatchTx.GetNonce(), orchestrator.ethAddress, ethcmn.Bytes2Hex(signature))

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	err = orchestrator.routerChainClient.QueueBroadcastMsg(msg)

	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 5)
}
