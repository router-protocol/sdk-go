package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/router-protocol/sdk-go/routerchain/outbound/types"

	ethcmn "github.com/ethereum/go-ethereum/common"

	"github.com/router-protocol/sdk-go/client/evm/gateway"
	routerclient "github.com/router-protocol/sdk-go/client/routerchain"

	gatewayWrapper "github.com/router-protocol/router-gateway-contracts/evm/build/bindings/go/GatewayUpgradeable"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
)

type Relayer interface {
	fetchAndProcessGatewayEvents()
}

type relayer struct {
	clientContext         client.Context
	ethClient             *ethclient.Client
	gatewayContractClient *gateway.GatewayContractClient
	routerChainClient     routerclient.ChainClient

	ethPrivatekey        string
	ethAddress           string
	relayerRouterAddress string
}

func NewRelayer(ethClient *ethclient.Client, gatewayContractClient *gateway.GatewayContractClient, routerChainClient routerclient.ChainClient, orchestratorEthPrivKey string, orchestratorEthAddress string, relayerRouterAddress string) *relayer {

	return &relayer{
		ethClient:             ethClient,
		gatewayContractClient: gatewayContractClient,
		routerChainClient:     routerChainClient,
		ethPrivatekey:         orchestratorEthPrivKey,
		ethAddress:            orchestratorEthAddress,
		relayerRouterAddress:  relayerRouterAddress,
	}
}

func (relayer *relayer) SubmitBatchTxToGateway(ctx context.Context, chainClient routerclient.ChainClient) {

	valsetResponse, _ := chainClient.GetAllValsets(ctx)
	outgoingBatchTxs, _ := chainClient.GetAllOutgoingBatchTx(ctx)

	for _, outgoingBatchTx := range outgoingBatchTxs.OutgoingBatchTx {
		signatures := relayer.collectSignatures(ctx, chainClient, outgoingBatchTx)
		relayer.sendTx(signatures, outgoingBatchTx, valsetResponse.Valset[0], relayer.relayerRouterAddress)
	}
}

func (relayer *relayer) collectSignatures(ctx context.Context, chainClient routerclient.ChainClient, outgoingBatchTx types.OutgoingBatchTx) []string {
	outgoingBatchConfirms, _ := chainClient.GetAllOutgoingBatchTxConfirms(ctx, uint64(outgoingBatchTx.GetDestinationChainType()), outgoingBatchTx.DestinationChainId, outgoingBatchTx.SourceAddress, outgoingBatchTx.Nonce)
	signatures := make([]string, 0)
	for _, outgoingBatchConfirm := range outgoingBatchConfirms.OutgoingBatchConfirm {
		signatures = append(signatures, outgoingBatchConfirm.GetSignature())
	}

	return signatures
}

func (relayer *relayer) sendTx(signatures []string, outgoingBatchTx types.OutgoingBatchTx, currentValset attestationTypes.Valset, relayerRouterAddress string) {
	// create auth and transaction package for deploying smart contract
	auth := getAccountAuth(relayer.ethClient, relayer.ethPrivatekey)

	sigs := make([]gatewayWrapper.UtilsSignature, len(signatures))
	for i := 0; i < len(signatures); i++ {
		v, r, s := sigToVRS(signatures[i])
		sigs[i].V = v
		sigs[i].R = r
		sigs[i].S = s
	}

	// Run through the elements of the batch and serialize them
	payloads := make([][]byte, len(outgoingBatchTx.ContractCalls))
	handlers := make([][]byte, len(outgoingBatchTx.ContractCalls))

	for j, contractCal := range outgoingBatchTx.ContractCalls {
		handlers[j] = contractCal.DestinationContractAddress
		payloads[j] = contractCal.Payload
	}

	routerRequestPayload := gatewayWrapper.UtilsRouterRequestPayload{
		RouterBridgeAddress:  outgoingBatchTx.SourceAddress,
		RelayerRouterAddress: relayerRouterAddress,
		RelayerFee:           outgoingBatchTx.RelayerFee.Amount.BigInt(),
		OutgoingTxFee:        outgoingBatchTx.OutgoingTxFee.Amount.BigInt(),
		IsAtomic:             outgoingBatchTx.IsAtomic,
		ExpTimestamp:         uint64(outgoingBatchTx.ExpiryTimestamp),
		Handlers:             handlers,
		Payloads:             payloads,
		OutboundTxNonce:      outgoingBatchTx.Nonce,
	}

	currentValsetArs := gatewayWrapper.UtilsValsetArgs{
		Validators:  make([]ethcmn.Address, len(currentValset.Members)),
		Powers:      make([]uint64, len(currentValset.Members)),
		ValsetNonce: currentValset.Nonce,
	}
	for i, valsetMember := range currentValset.Members {
		currentValsetArs.Validators[i] = ethcmn.HexToAddress(valsetMember.EthereumAddress)
		// power := &big.Int{}
		// power.SetUint64(valsetMember.Power)
		currentValsetArs.Powers[i] = valsetMember.Power
	}

	currentValsetArs.ValsetNonce = currentValset.Nonce

	// fmt.Println("currentValsetArgs", currentValsetArs)
	// fmt.Println("sigs", sigs)
	// fmt.Println("routerRequestPayload", routerRequestPayload)

	tx, err := relayer.gatewayContractClient.GatewayWrapper.RequestFromRouter(auth, currentValsetArs, sigs, routerRequestPayload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func getAccountAuth(client *ethclient.Client, privateKeyStr string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000000000)

	return auth
}

func sigToVRS(sigHex string) (v uint8, r, s ethcmn.Hash) {
	signatureBytes := ethcmn.FromHex(sigHex)
	vParam := signatureBytes[64]
	if vParam == byte(0) {
		vParam = byte(27)
	} else if vParam == byte(1) {
		vParam = byte(28)
	}

	v = vParam
	r = ethcmn.BytesToHash(signatureBytes[0:32])
	s = ethcmn.BytesToHash(signatureBytes[32:64])

	return
}
