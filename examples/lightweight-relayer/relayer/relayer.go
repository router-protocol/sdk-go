package relayer

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	crosstalkTypes "github.com/router-protocol/sdk-go/routerchain/crosstalk/types"
	"github.com/router-protocol/sdk-go/routerchain/outbound/types"
	"github.com/router-protocol/sdk-go/routerchain/util"

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
		if outgoingBatchTx.Status == types.OUTGOING_TX_ACK_OBSERVED || outgoingBatchTx.Status == types.OUTGOING_TX_ACK_DELEGATED {
			continue
		}
		if outgoingBatchTx.SourceAddress != "router17p9rzwnnfxcjp32un9ug7yhhzgtkhvl9jfksztgw5uh69wac2pgsmpev85" {
			continue
		}
		signatures := relayer.collectSignatures(ctx, chainClient, outgoingBatchTx)
		fmt.Println("Sending tx", "outgoingBatchTx", outgoingBatchTx, "signatures", signatures)
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
	time.Sleep(15 * time.Second)
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

func (relayer *relayer) SubmitCrosstalkTxToGateway(ctx context.Context, chainClient routerclient.ChainClient) {

	valsetResponse, _ := chainClient.GetAllValsets(ctx)
	allCrosstalkRequests, _ := chainClient.GetAllCrossTalkRequest(ctx)

	for _, crosstalkRequest := range allCrosstalkRequests.CrossTalkRequest {
		if crosstalkRequest.Status == crosstalkTypes.CROSSTALK_REQUEST_CREATED {
			continue
		}

		if crosstalkRequest.EventNonce != 3 {
			continue
		}

		signatures := relayer.collectCrosstalkSignatures(ctx, chainClient, crosstalkRequest)
		fmt.Println("Sending tx", "crosstalkRequest", crosstalkRequest, "signatures", signatures)
		relayer.sendCrossTalkTx(signatures, crosstalkRequest, valsetResponse.Valset[0], relayer.relayerRouterAddress)
	}
}

func (relayer *relayer) collectCrosstalkSignatures(ctx context.Context, chainClient routerclient.ChainClient, crosstalkRequest crosstalkTypes.CrossTalkRequest) []string {
	claimHash, err := crosstalkRequest.ClaimHash()
	if err != nil {
		panic(err)
	}

	crosstalkConfirmations, err := chainClient.GetAllCrosstalkRequestConfirmations(ctx, uint64(crosstalkRequest.SourceChainType), crosstalkRequest.SourceChainId, crosstalkRequest.EventNonce, claimHash)
	if err != nil {
		panic(err)
	}
	signatures := make([]string, 0)
	for _, crosstalkConfirmation := range crosstalkConfirmations.CrosstalkRequestConfirm {
		signatures = append(signatures, crosstalkConfirmation.GetSignature())
	}

	return signatures
}

func (relayer *relayer) sendCrossTalkTx(signatures []string, crosstalkRequest crosstalkTypes.CrossTalkRequest, currentValset attestationTypes.Valset, relayerRouterAddress string) {
	// create auth and transaction package for deploying smart contract
	auth := getAccountAuth(relayer.ethClient, relayer.ethPrivatekey)

	sigs := make([]gatewayWrapper.UtilsSignature, len(signatures))
	for i := 0; i < len(signatures); i++ {
		v, r, s := sigToVRS(signatures[i])
		sigs[i].V = v
		sigs[i].R = r
		sigs[i].S = s
	}

	// Run through the elements of the crosstalk request and serialize them
	sourceParams := gatewayWrapper.UtilsSourceParams{
		Caller:    []byte(crosstalkRequest.RequestSender),
		ChainType: uint64(crosstalkRequest.SourceChainType),
		ChainId:   crosstalkRequest.SourceChainId,
	}

	contractCalls := gatewayWrapper.UtilsContractCalls{
		DestContractAddresses: crosstalkRequest.DestContractAddresses,
		Payloads:              crosstalkRequest.DestContractPayloads,
	}

	crosstalkRequestPayload := gatewayWrapper.UtilsCrossTalkPayload{
		RelayerRouterAddress: relayerRouterAddress,
		IsAtomic:             crosstalkRequest.IsAtomic,
		EventIdentifier:      crosstalkRequest.EventNonce,
		ExpTimestamp:         uint64(crosstalkRequest.ExpiryTimestamp),
		CrossTalkNonce:       crosstalkRequest.RequestNonce,
		SourceParams:         sourceParams,
		ContractCalls:        contractCalls,
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

	fmt.Println("currentValsetArgs", currentValsetArs)
	fmt.Println("sigs", sigs)
	fmt.Println("crosstalkRequestPayload", crosstalkRequestPayload)

	tx, err := relayer.gatewayContractClient.GatewayWrapper.RequestFromSource(auth, currentValsetArs, sigs, crosstalkRequestPayload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("crosstalkRequestPayload", crosstalkRequestPayload)

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	relayer.GetCheckpoint(crosstalkRequest, "")
	time.Sleep(15 * time.Second)
}

// GetCheckpoint gets the checkpoint signature from the given CrossTalkRequest
func (relayer *relayer) GetCheckpoint(msg crosstalkTypes.CrossTalkRequest, routerIDstring string) []byte {

	abi, err := abi.JSON(strings.NewReader(util.CrossTalkRequestCheckpointABIJSON))
	if err != nil {
		panic("Bad ABI constant!")
	}

	// the contract argument is not a arbitrary length array but a fixed length 32 byte
	// array, therefore we have to utf8 encode the string (the default in this case) and
	// then copy the variable length encoded data into a fixed length array. This function
	// will panic if gravityId is too long to fit in 32 bytes
	// routerID, err := util.StrToFixByteArray(routerIDstring)
	// if err != nil {
	// 	panic(err)
	// }

	// Create the methodName argument which salts the signature
	methodNameBytes := []uint8("requestFromSource")
	var crosstalkMethodName [32]uint8
	copy(crosstalkMethodName[:], methodNameBytes)

	eventIdentifier := &big.Int{}
	eventIdentifier.SetUint64(msg.EventNonce)

	crossTalkNonce := &big.Int{}
	crossTalkNonce.SetUint64(msg.RequestNonce)

	destChainType := &big.Int{}
	if msg.DestinationChainType == 0 {
		destChainType = big.NewInt(0)
	} else {
		destChainType.SetUint64(uint64(msg.DestinationChainType))
	}

	srcChainType := &big.Int{}
	srcChainType.SetUint64(uint64(msg.SourceChainType))

	caller := []byte(msg.RequestSender)

	expTimestamp := &big.Int{}
	expTimestamp.SetUint64(uint64(msg.ExpiryTimestamp))

	// the methodName needs to be the same as the 'name' above in the checkpointAbiJson
	// but other than that it's a constant that has no impact on the output. This is because
	// it gets encoded as a function name which we must then discard.
	abiEncodedCrosstalk, err := abi.Pack("checkpoint",
		crosstalkMethodName,
		eventIdentifier,
		crossTalkNonce,
		destChainType,
		msg.DestinationChainId,
		msg.SourceChainId,
		srcChainType,
		caller,
		msg.IsAtomic,
		expTimestamp,
		msg.DestContractAddresses,
		msg.DestContractPayloads,
	)

	// DATA
	fmt.Println("crosstalkMethodName", hex.EncodeToString(crosstalkMethodName[:]))
	fmt.Println("eventIdentifier", hex.EncodeToString(eventIdentifier.Bytes()))
	fmt.Println("crossTalkNonce", hex.EncodeToString(crossTalkNonce.Bytes()))
	fmt.Println("chainType", hex.EncodeToString(destChainType.Bytes()))
	fmt.Println("chainId", hex.EncodeToString([]byte(msg.DestinationChainId)))
	fmt.Println("SourceChainId", hex.EncodeToString([]byte(msg.SourceChainId)))
	fmt.Println("SourceChainType", hex.EncodeToString(srcChainType.Bytes()))
	fmt.Println("callerbytes32", hex.EncodeToString(caller[:]))
	fmt.Println("caller", hex.EncodeToString([]byte(msg.RequestSender)))
	fmt.Println("isAtomic", msg.IsAtomic)
	fmt.Println("expTimestamp", hex.EncodeToString(expTimestamp.Bytes()))
	fmt.Println("DestContractAddresses", msg.DestContractAddresses)
	fmt.Println("payloads", msg.DestContractPayloads)

	// this should never happen outside of test since any case that could crash on encoding
	// should be filtered above.
	if err != nil {
		panic(fmt.Sprintf("Error packing checkpoint! %s/n", err))
	}

	// we hash the resulting encoded bytes discarding the first 4 bytes these 4 bytes are the constant
	// method name 'checkpoint'. If you were to replace the checkpoint constant in this code you would
	// then need to adjust how many bytes you truncate off the front to get the output of abi.encode()
	fmt.Println("abiEncodedCrosstalk", hex.EncodeToString(abiEncodedCrosstalk))
	fmt.Println("messageHash", crypto.Keccak256Hash(abiEncodedCrosstalk[4:]))

	checkpoint := crypto.Keccak256Hash(abiEncodedCrosstalk[4:]).Bytes()
	mesageDigest := accounts.TextHash(checkpoint)
	fmt.Println("mesageDigest", hex.EncodeToString(mesageDigest))

	return crypto.Keccak256Hash(abiEncodedCrosstalk[4:]).Bytes()
}

func (relayer *relayer) SubmitCrosstalkAckTxToGateway(ctx context.Context, chainClient routerclient.ChainClient) {
	fmt.Println("SubmitCrosstalkAckTxToGateway")
	valsetResponse, _ := chainClient.GetAllValsets(ctx)
	allCrosstalkAckRequests, _ := chainClient.GetAllCrossTalkAckRequest(ctx)
	fmt.Println("allCrosstalkAckRequests", allCrosstalkAckRequests)
	for _, crosstalkAckRequest := range allCrosstalkAckRequests.CrossTalkAckRequest {
		if crosstalkAckRequest.CrosstalkNonce != 1 {
			continue
		}

		signatures := relayer.collectCrosstalkAckSignatures(ctx, chainClient, crosstalkAckRequest)
		fmt.Println("Sending tx", "crosstalkAckRequest", crosstalkAckRequest, "signatures", signatures)
		relayer.sendCrossTalkAckTx(signatures, crosstalkAckRequest, valsetResponse.Valset[0], relayer.relayerRouterAddress)
	}
}

func (relayer *relayer) collectCrosstalkAckSignatures(ctx context.Context, chainClient routerclient.ChainClient, crosstalkAckRequest crosstalkTypes.CrossTalkAckRequest) []string {
	claimHash, err := crosstalkAckRequest.ClaimHash()
	if err != nil {
		panic(err)
	}

	crosstalkAckConfirmations, err := chainClient.GetAllCrosstalkAckRequestConfirmations(ctx, uint64(crosstalkAckRequest.ChainType), crosstalkAckRequest.ChainId, crosstalkAckRequest.EventNonce, claimHash)
	if err != nil {
		panic(err)
	}
	signatures := make([]string, 0)
	for _, crosstalkAckConfirmation := range crosstalkAckConfirmations.CrosstalkAckRequestConfirm {
		signatures = append(signatures, crosstalkAckConfirmation.GetSignature())
	}

	return signatures
}

func (relayer *relayer) sendCrossTalkAckTx(signatures []string, crosstalkAckRequest crosstalkTypes.CrossTalkAckRequest, currentValset attestationTypes.Valset, relayerRouterAddress string) {
	// create auth and transaction package for deploying smart contract
	auth := getAccountAuth(relayer.ethClient, relayer.ethPrivatekey)

	sigs := make([]gatewayWrapper.UtilsSignature, len(signatures))
	for i := 0; i < len(signatures); i++ {
		v, r, s := sigToVRS(signatures[i])
		sigs[i].V = v
		sigs[i].R = r
		sigs[i].S = s
	}

	// Run through the elements of the crosstalk request and serialize them
	crosstalkAckPayload := gatewayWrapper.UtilsCrossTalkAckPayload{
		CrossTalkNonce:     crosstalkAckRequest.CrosstalkNonce,
		EventIdentifier:    crosstalkAckRequest.EventIdentifier,
		DestChainType:      uint64(crosstalkAckRequest.ChainType),
		DestChainId:        crosstalkAckRequest.ChainId,
		SrcContractAddress: crosstalkAckRequest.CrosstalkRequestSender,
		ExecFlags:          crosstalkAckRequest.ExecFlags,
		ExecData:           crosstalkAckRequest.ExecData,
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

	fmt.Println("currentValsetArgs", currentValsetArs)
	fmt.Println("sigs", sigs)
	fmt.Println("crosstalkAckPayload", crosstalkAckPayload)

	tx, err := relayer.gatewayContractClient.GatewayWrapper.CrossTalkAck(auth, currentValsetArs, sigs, crosstalkAckPayload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("crosstalkAckPayload", crosstalkAckPayload)

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	// relayer.GetCrosstalkCheckpoint(crosstalkAckPayload, "")
	time.Sleep(15 * time.Second)
}
