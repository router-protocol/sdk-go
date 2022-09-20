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
	"github.com/router-protocol/sdk-go/routerchain/attestation/types"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	inboundTypes "github.com/router-protocol/sdk-go/routerchain/inbound/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	outboundTypes "github.com/router-protocol/sdk-go/routerchain/outbound/types"
)

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

func sendInboundRequest(ctx context.Context, chainClient chainclient.ChainClient, sendToRouterEvent *gatewayWrapper.GatewayRequestToRouterEvent) {
	// prepare tx msg
	msg := inboundTypes.NewMsgInboundRequest(chainClient.FromAddress().String(),
		multichainTypes.ChainType(sendToRouterEvent.SrcChainType),
		sendToRouterEvent.SrcChainId, sendToRouterEvent.EventNonce,
		sendToRouterEvent.Raw.BlockNumber,
		sendToRouterEvent.ApplicationContract.Hex(),
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

func sendValsetUpdateRequest(ctx context.Context, chainClient chainclient.ChainClient, valsetUpdatedEvent *gatewayWrapper.GatewayValsetUpdatedEvent) {
	// prepare tx msg
	members := make([]types.BridgeValidator, len(valsetUpdatedEvent.Validators))
	for i, val := range valsetUpdatedEvent.Validators {
		members[i] = types.BridgeValidator{
			EthereumAddress: val.Hex(),
			Power:           valsetUpdatedEvent.Powers[i],
		}
	}

	msg := attestationTypes.NewMsgValsetUpdatedClaim(chainClient.FromAddress().String(),
		multichainTypes.ChainType(valsetUpdatedEvent.SrcChainType),
		valsetUpdatedEvent.SrcChainId,
		valsetUpdatedEvent.EventNonce,
		valsetUpdatedEvent.NewValsetNonce,
		valsetUpdatedEvent.Raw.BlockNumber,
		valsetUpdatedEvent.Raw.TxHash.String(),
		members)

	fmt.Println("Sending Valset Updated claim Request", msg)
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

func sendOutboundAckRequest(ctx context.Context, chainClient chainclient.ChainClient, outboundAckEvent *gatewayWrapper.GatewayEventOutboundAck) {
	// prepare tx msg
	contractAckResponses := make([]*outboundTypes.ContractAckResponse, len(outboundAckEvent.ContractAckResponses))
	for i, _ := range outboundAckEvent.ContractAckResponses {
		// TODO: @harsh Build Contract Ack Responses
		contractAckResponses[i] = &outboundTypes.ContractAckResponse{}
	}

	msg := outboundTypes.NewMsgOutboundAckRequest(chainClient.FromAddress().String(),
		multichainTypes.ChainType(outboundAckEvent.ChainType),
		outboundAckEvent.ChainId,
		outboundAckEvent.OutboundTxNonce,
		outboundAckEvent.OutboundTxRequestedBy,
		outboundAckEvent.RelayerRouterAddress,
		outboundAckEvent.Raw.TxHash.Hex(),
		// TODO: @harsh Fetch Fee Consumed and remove hardcoded values
		uint64(245234),
		contractAckResponses,
		outboundAckEvent.EventNonce,
		outboundAckEvent.Raw.BlockNumber,
	)

	fmt.Println("Sending Outbound Ack Request", msg)
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

	for _, outgoingBatchTx := range outgoingBatchTxs.OutgoingBatchTx {

		outgoingBatchConfirm, err := orchestrator.routerChainClient.GetOutgoingBatchTxConfirm(ctx, uint64(outgoingBatchTx.DestinationChainType), outgoingBatchTx.DestinationChainId, outgoingBatchTx.SourceAddress, outgoingBatchTx.Nonce, orchestrator.routerChainClient.FromAddress().String())
		if outgoingBatchConfirm.OutgoingBatchConfirm.Orchestrator == orchestrator.routerChainClient.FromAddress().String() {
			return
		}

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
}
