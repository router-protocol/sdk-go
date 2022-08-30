package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	chainclient "github.com/router-protocol/sdk-go/client/routerchain"
	"github.com/router-protocol/sdk-go/client/routerchain/common"
	"github.com/router-protocol/sdk-go/routerchain/outbound/types"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	gatewayWrapper "github.com/router-protocol/router-gateway-contracts/evm/wrappers"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
)

const (
	// Tx Agrs
	RELAYER_PRIVATE_KEY = "843EE93DA70C08B88C726C43329B8DA92CC26AEFE2A0F3F33832A0540E66EA53"
	ETH_RPC             = "https://ropsten.infura.io/v3/74c8b06d2481408c86fe936c11657def"
)

var (
	GATEWAY_CONTRACT_ADDRESS = ethcmn.HexToAddress("0xEE39624b883fe60b1B6b77686ecB0Be10c990341")
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

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	SubmitBatchTxToGateway(ctx, chainClient)

}

func SubmitBatchTxToGateway(ctx context.Context, chainClient chainclient.ChainClient) {

	valsetResponse, _ := chainClient.GetAllValsets(ctx)
	outgoingBatchTxs, _ := chainClient.GetAllOutgoingBatchTx(ctx)

	for _, outgoingBatchTx := range outgoingBatchTxs.OutgoingBatchTx {
		signatures := CollectSignatures(ctx, chainClient, outgoingBatchTx)
		SendTx(signatures, outgoingBatchTx, valsetResponse.Valset[0])
	}
}

func CollectSignatures(ctx context.Context, chainClient chainclient.ChainClient, outgoingBatchTx types.OutgoingBatchTx) []string {
	outgoingBatchConfirms, _ := chainClient.GetAllOutgoingBatchTxConfirms(ctx, uint64(outgoingBatchTx.GetDestinationChainType()), outgoingBatchTx.DestinationChainId, outgoingBatchTx.SourceAddress, outgoingBatchTx.Nonce)
	signatures := make([]string, 0)
	for _, outgoingBatchConfirm := range outgoingBatchConfirms.OutgoingBatchConfirm {
		signatures = append(signatures, outgoingBatchConfirm.GetSignature())
	}

	return signatures
}

func SendTx(signatures []string, outgoingBatchTx types.OutgoingBatchTx, currentValset attestationTypes.Valset) {
	// address of etherum env
	client, err := ethclient.Dial(ETH_RPC)
	if err != nil {
		panic(err)
	}

	// create auth and transaction package for deploying smart contract
	auth := getAccountAuth(client, RELAYER_PRIVATE_KEY)

	instance, err := gatewayWrapper.NewGateway(GATEWAY_CONTRACT_ADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}

	sigs := make([]gatewayWrapper.Signature, len(signatures))
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

	chainType := &big.Int{}
	chainType.SetUint64(uint64(outgoingBatchTx.DestinationChainType))

	batchNonce := &big.Int{}
	batchNonce.SetUint64(outgoingBatchTx.Nonce)

	routerRequestPayload := gatewayWrapper.RouterRequestPayload{
		Sender:        outgoingBatchTx.SourceAddress,
		ChainId:       outgoingBatchTx.DestinationChainId,
		ChainType:     chainType,
		RelayerFee:    outgoingBatchTx.RelayerFee.Amount.BigInt(),
		OutgoingTxFee: outgoingBatchTx.OutgoingTxFee.Amount.BigInt(),
		IsAtomic:      outgoingBatchTx.IsAtomic,
		Handlers:      handlers,
		Payloads:      payloads,
		Nonce:         batchNonce,
	}

	currentValsetArs := gatewayWrapper.ValsetArgs{
		Validators:  make([]ethcmn.Address, len(currentValset.Members)),
		Powers:      make([]*big.Int, len(currentValset.Members)),
		ValsetNonce: &big.Int{},
	}
	for i, valsetMember := range currentValset.Members {
		currentValsetArs.Validators[i] = ethcmn.HexToAddress(valsetMember.EthereumAddress)
		power := &big.Int{}
		power.SetUint64(valsetMember.Power)
		currentValsetArs.Powers[i] = power
	}

	valsetNonce := &big.Int{}
	// TODO: change valset nonce
	valsetNonce.SetUint64(0)
	// valsetNonce.SetUint64(currentValset.Nonce)
	currentValsetArs.ValsetNonce = valsetNonce

	// fmt.Println("currentValsetArgs", currentValsetArs)
	// fmt.Println("sigs", sigs)
	// fmt.Println("routerRequestPayload", routerRequestPayload)

	tx, err := instance.RequestFromRouter(auth, currentValsetArs, sigs, routerRequestPayload)
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
