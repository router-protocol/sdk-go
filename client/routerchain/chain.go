package routerchain

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/pkg/errors"
	"github.com/router-protocol/sdk-go/client/routerchain/common"
	log "github.com/xlab/suplog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	"github.com/CosmWasm/wasmd/x/wasm/ioutils"
	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"

	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	crosschainTypes "github.com/router-protocol/sdk-go/routerchain/crosschain/types"
	metastoreTypes "github.com/router-protocol/sdk-go/routerchain/metastore/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	pricefeedTypes "github.com/router-protocol/sdk-go/routerchain/pricefeed/types"
	txlookupTypes "github.com/router-protocol/sdk-go/routerchain/txlookup/types"
	voyagerTypes "github.com/router-protocol/sdk-go/routerchain/voyager/types"
)

const (
	msgCommitBatchSizeLimit          = 1024
	msgCommitBatchTimeLimit          = 500 * time.Millisecond
	defaultBroadcastStatusPoll       = 100 * time.Millisecond
	defaultBroadcastTimeout          = 40 * time.Second
	defaultTimeoutHeight             = 20
	defaultTimeoutHeightSyncInterval = 10 * time.Second
	defaultSessionRenewalOffset      = 120
	defaultBlockTime                 = 3 * time.Second
	defaultChainCookieName           = ".chain_cookie"
)

var (
	ErrTimedOut       = errors.New("tx timed out")
	ErrQueueClosed    = errors.New("queue is closed")
	ErrEnqueueTimeout = errors.New("enqueue timeout")
	ErrReadOnly       = errors.New("client is in read-only mode")
)

type ChainClient interface {
	CanSignTransactions() bool
	FromAddress() sdk.AccAddress
	QueryClient() *grpc.ClientConn
	ClientContext() client.Context

	SimulateMsg(clientCtx client.Context, msgs ...sdk.Msg) (*txtypes.SimulateResponse, error)
	AsyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error)
	SyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error)
	QueueBroadcastMsg(msgs ...sdk.Msg) error

	// bank
	GetBankBalances(ctx context.Context, address string) (*banktypes.QueryAllBalancesResponse, error)
	GetBankBalance(ctx context.Context, address string, denom string) (*banktypes.QueryBalanceResponse, error)
	GetAccount(ctx context.Context, address string) (*authtypes.QueryAccountResponse, error)

	// staking
	GetValidator(ctx context.Context, validatorAddr string) (*stakingtypes.QueryValidatorResponse, error)
	GetAllValidators(ctx context.Context, status string, pagination *query.PageRequest) (*stakingtypes.QueryValidatorsResponse, error)

	// MultiChain
	GetAllChainConfig(ctx context.Context, pagination *query.PageRequest) (*multichainTypes.QueryAllChainConfigResponse, error)
	GetChainConfig(ctx context.Context, chainId string) (*multichainTypes.QueryGetChainConfigResponse, error)
	GetContractConfig(ctx context.Context, chainId string, contract string) (*multichainTypes.QueryGetContractConfigResponse, error)
	GetAllContractConfig(ctx context.Context, pagination *query.PageRequest) (*multichainTypes.QueryAllContractConfigResponse, error)
	GetAllContractConfigByChainId(ctx context.Context, chainId string) (*multichainTypes.QueryAllContractConfigByChainIdResponse, error)
	GetNonceObservedStatus(ctx sdk.Context, chainId string, contractAddress string, nonce uint64) (*multichainTypes.QueryGetNonceObservedStatusResponse, error)

	// Attestation
	GetOrchestratorValidator(ctx context.Context, orchestratorAddr sdk.AccAddress) (*attestationTypes.QueryFetchOrchestratorValidatorResponse, error)
	GetValsetByNonce(c context.Context, valsetNonce uint64) (*attestationTypes.QueryGetValsetResponse, error)
	GetLastEventByValidator(ctx context.Context, chainId string, contract string, validator sdk.ValAddress) (*attestationTypes.QueryLastEventNonceResponse, error)
	GetLatestValset(ctx context.Context) (*attestationTypes.QueryLatestValsetResponse, error)
	GetLatestValsetNonce(ctx context.Context) (*attestationTypes.QueryLatestValsetNonceResponse, error)
	GetAllAttestations(ctx context.Context, pagination *query.PageRequest) (*attestationTypes.QueryAllAttestationResponse, error)
	GetAllObservedAttestations(ctx context.Context, pagination *query.PageRequest) (*attestationTypes.QueryAllObservedAttestationResponse, error)
	GetAllOrchestrators(ctx context.Context) (*attestationTypes.QueryListOrchestratorsResponse, error)
	GetAllValsets(ctx context.Context, pagination *query.PageRequest) (*attestationTypes.QueryAllValsetResponse, error)
	GetAllValsetConfirms(ctx context.Context, valsetNonce uint64, destChainType multichainTypes.ChainType, pagination *query.PageRequest) (*attestationTypes.QueryAllValsetConfirmationResponse, error)
	GetValsetConfirm(ctx context.Context, valsetNonce uint64, destChainType multichainTypes.ChainType, orchestrator string) (*attestationTypes.QueryGetValsetConfirmationResponse, error)

	// PriceFeed
	GetPriceBySymbol(ctx context.Context, symbol string) (*pricefeedTypes.QueryGetPriceResponse, error)
	GetAllSymbolPrices(ctx context.Context, pagination *query.PageRequest) (*pricefeedTypes.QueryAllPriceResponse, error)
	GetSymbolRequest(ctx context.Context, symbol string) (*pricefeedTypes.QueryGetSymbolRequestResponse, error)
	GetAllSymbolRequest(ctx context.Context, pagination *query.PageRequest) (*pricefeedTypes.QueryAllSymbolRequestResponse, error)
	GetPriceFeederInfo(ctx context.Context, priceFeederName string) (*pricefeedTypes.QueryGetPriceFeederInfoResponse, error)
	GetAllPriceFeederInfo(ctx context.Context, pagination *query.PageRequest) (*pricefeedTypes.QueryAllPriceFeederInfoResponse, error)
	GetGasPriceByChainId(ctx context.Context, chainId string) (*pricefeedTypes.QueryGetGasPriceResponse, error)
	GetAllGasPrice(ctx context.Context, pagination *query.PageRequest) (*pricefeedTypes.QueryAllGasPriceResponse, error)
	GetAllWhitelistedIBCChannels(ctx context.Context) (*pricefeedTypes.QueryWhitelistedIBCChannelsResponse, error)

	// Crosschain
	GetCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetCrosschainRequestResponse, error)
	GetAllCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllCrosschainRequestResponse, error)
	GetValidatedCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetValidCrosschainRequestResponse, error)
	GetAllValidCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllValidCrosschainRequestResponse, error)
	GetNativeTransferedCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetNativeTransferedCrosschainRequestResponse, error)
	GetAllNativeTransferedCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllNativeTransferedCrosschainRequestResponse, error)
	GetReadyToExecuteCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetReadyToExecuteCrosschainRequestResponse, error)
	GetAllReadyToExecuteCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllReadyToExecuteCrosschainRequestResponse, error)
	GetAllReadyToExecuteCrosschainRequestsByWorkflow(ctx context.Context, workflowType crosschainTypes.WorkflowType, pagination *query.PageRequest) (*crosschainTypes.QueryReadyToExecuteCrosschainRequestByWorkflowResponse, error)
	GetAllReadyToExecuteCrosschainRequestsByWorkflowAndRelayer(ctx context.Context, workflowType crosschainTypes.WorkflowType, relayerType crosschainTypes.RelayerType, pagination *query.PageRequest) (*crosschainTypes.QueryReadyToExecuteCrosschainRequestByWorkflowAndRelayerResponse, error)
	GetAllReadyToExecuteCrosschainAckRequestsByWorkflow(ctx context.Context, ackWorkflowType crosschainTypes.WorkflowType, pagination *query.PageRequest) (*crosschainTypes.QueryReadyToExecuteCrosschainAckRequestByWorkflowResponse, error)
	GetAllReadyToExecuteCrosschainAckRequestsByWorkflowAndRelayer(ctx context.Context, ackWorkflowType crosschainTypes.WorkflowType, ackRelayerType crosschainTypes.RelayerType, pagination *query.PageRequest) (*crosschainTypes.QueryReadyToExecuteCrosschainAckRequestByWorkflowAndRelayerResponse, error)
	GetBlockedCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetBlockedCrosschainRequestResponse, error)
	GetAllBlockedCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllBlockedCrosschainRequestResponse, error)
	GetExpiredCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetExpiredCrosschainRequestResponse, error)
	GetAllExpiredCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllExpiredCrosschainRequestResponse, error)
	GetExecutedCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetExecutedCrosschainRequestResponse, error)
	GetAllExecutedCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllExecutedCrosschainRequestResponse, error)
	GetFeesSettledCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetFeesSettledCrosschainRequestResponse, error)
	GetAllFeesSettledCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllFeesSettledCrosschainRequestResponse, error)
	GetCompletedCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetCompletedCrosschainRequestResponse, error)
	GetAllCompletedCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllCompletedCrosschainRequestResponse, error)

	// CrosschainAck
	GetCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetCrosschainAckRequestResponse, error)
	GetAllCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllCrosschainAckRequestResponse, error)
	GetValidCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetValidCrosschainAckRequestResponse, error)
	GetAllValidCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllValidCrosschainAckRequestResponse, error)
	GetBlockedCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetBlockedCrosschainAckRequestResponse, error)
	GetAllBlockedCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllBlockedCrosschainAckRequestResponse, error)
	GetReadyToExecuteCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetReadyToExecuteCrosschainAckRequestResponse, error)
	GetAllReadyToExecuteCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllReadyToExecuteCrosschainAckRequestResponse, error)
	GetExecutedCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetExecutedCrosschainAckRequestResponse, error)
	GetAllExecutedCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllExecutedCrosschainAckRequestResponse, error)
	GetFeesSettledCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetFeesSettledCrosschainAckRequestResponse, error)
	GetAllFeesSettledCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllFeesSettledCrosschainAckRequestResponse, error)
	GetCompletedCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetCompletedCrosschainAckRequestResponse, error)
	GetAllCompletedCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllCompletedCrosschainAckRequestResponse, error)
	GetExpiredCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetExpiredCrosschainAckRequestResponse, error)
	GetAllExpiredCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllExpiredCrosschainAckRequestResponse, error)

	GetAllCrosschainRequestConfirmations(ctx context.Context, pagination *query.PageRequest, sourceChainId string, requestIdentifier uint64, claimHash []byte) (*crosschainTypes.QueryAllCrosschainRequestConfirmResponse, error)
	GetAllCrosschainRequestAckConfirmations(ctx context.Context, pagination *query.PageRequest, ackSrcChainId string, ackRequestIdentifier uint64, claimHash []byte) (*crosschainTypes.QueryAllCrosschainAckRequestConfirmResponse, error)
	GetCrosschainRequestConfirmation(ctx context.Context, pagination *query.PageRequest, sourceChainId string, requestIdentifier uint64, claimHash []byte, orchestrator string) (*crosschainTypes.QueryGetCrosschainRequestConfirmResponse, error)
	GetCrosschainAckRequestConfirmation(ctx context.Context, pagination *query.PageRequest, ackSrcChainId string, ackRequestIdentifier uint64, claimHash []byte, orchestrator string) (*crosschainTypes.QueryGetCrosschainAckRequestConfirmResponse, error)
	GetCrosschainAckReceipt(ctx context.Context, ackReceiptSrcChainId string, ackReceiptIdentifier uint64) (*crosschainTypes.QueryGetCrosschainAckReceiptResponse, error)

	// MetaStore
	GetAllMetaInfo(ctx context.Context, pagination *query.PageRequest) (*metastoreTypes.QueryAllMetaInfoResponse, error)
	GetMetaInfo(ctx context.Context, chainId string, dappAddress []byte) (*metastoreTypes.QueryAllMetaInfoResponseByChainAndAddress, error)

	// Wasm
	StoreCode(file string, sender sdk.AccAddress) (int64, error)
	InstantiateContract(codeID uint64, label string, amountStr string, initMsg string, adminStr string, noAdmin bool, sender sdk.AccAddress) (string, error)
	ExecuteContract(amountStr string, sender sdk.AccAddress, contractAddr string, execMsg string) error
	SmartContractState(ctx context.Context, contractAddress string, queryData []byte) (*wasmTypes.QuerySmartContractStateResponse, error)
	RawContractState(ctx context.Context, contractAddress string, queryData []byte) (*wasmTypes.QueryRawContractStateResponse, error)

	GetGasFee() (string, error)
	Close()

	// Txlookup
	GetAllAdhocRequests(ctx context.Context, pagination *query.PageRequest) (*txlookupTypes.QueryListAdhocResponse, error)
}

type chainClient struct {
	ctx       client.Context
	opts      *common.ClientOptions
	logger    log.Logger
	conn      *grpc.ClientConn
	txFactory tx.Factory

	fromAddress sdk.AccAddress
	doneC       chan bool
	msgC        chan sdk.Msg
	syncMux     *sync.Mutex

	accNum    uint64
	accSeq    uint64
	gasWanted uint64
	gasFee    string

	sessionCookie  string
	sessionEnabled bool

	txClient           txtypes.ServiceClient
	authQueryClient    authtypes.QueryClient
	bankQueryClient    banktypes.QueryClient
	stakingQueryClient stakingtypes.QueryClient

	// Custom Query clients
	multichainQueryClient  multichainTypes.QueryClient
	attestationQueryClient attestationTypes.QueryClient
	pricefeedQueryClient   pricefeedTypes.QueryClient
	wasmQueryClient        wasmTypes.QueryClient
	metastoreQueryClient   metastoreTypes.QueryClient
	crosschainQueryClient  crosschainTypes.QueryClient
	voyagerQueryClient     voyagerTypes.QueryClient
	txlookupQueryClient    txlookupTypes.QueryClient

	closed  int64
	canSign bool
}

func InitialiseChainClient(network common.Network, keyringFrom string, passphrase string, privateKey string, keyringDir string, keyringBackend string) ChainClient {
	// network, err := common.LoadNetwork(network.Name, "k8s")
	// if err != nil {
	// 	fmt.Println("Error while loading network from default tmRPC Endpoint ", "rpc", network.TmEndpoint)
	// }
	log.Infoln("InitialiseChainClient|Network: ", network)

	networkChainId := network.ChainId
	tmEndpoint := network.TmEndpoint
	chainGrpcEndpoint := network.ChainGrpcEndpoint

	tmRPC, err := rpchttp.New(tmEndpoint, "/websocket")
	if err != nil {
		log.Errorln("Error while creating tmRPC client ", err)
	}

	if keyringDir == "" {
		keyringDir = os.Getenv("HOME") + "/.routerd"
	}

	senderAddress, cosmosKeyring, err := InitCosmosKeyring(
		keyringDir,
		"routerd",
		keyringBackend,
		keyringFrom,
		passphrase,
		privateKey, // keyring will be used if pk not provided
		false,
	)

	if err != nil {
		log.Errorln("Error while initializing keyring ", err)
	}

	// initialize grpc client
	clientCtx, err := NewClientContext(
		networkChainId,
		senderAddress.String(),
		cosmosKeyring,
	)

	if err != nil {
		log.Errorln("Error while initializing client context ", err)
		panic(err)
	}

	clientCtx = clientCtx.WithNodeURI(tmEndpoint).WithClient(tmRPC)

	routerchainClient, err := NewChainClient(
		clientCtx,
		chainGrpcEndpoint,
		// common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000route"),
	)

	if err != nil {
		log.Errorln("Error while initializing chain client ", err)
	}

	return routerchainClient
}

// NewCosmosClient creates a new gRPC client that communicates with gRPC server at protoAddr.
// protoAddr must be in form "tcp://127.0.0.1:8080" or "unix:///tmp/test.sock", protocol is required.
func NewChainClient(
	ctx client.Context,
	protoAddr string,
	options ...common.ClientOption,
) (ChainClient, error) {
	// process options
	opts := common.DefaultClientOptions()
	for _, opt := range options {
		if err := opt(opts); err != nil {
			err = errors.Wrap(err, "error in client option")
			return nil, err
		}
	}

	txFactory := NewTxFactory(ctx)
	if len(opts.GasPrices) > 0 {
		txFactory = txFactory.WithGasPrices(opts.GasPrices)
	}

	var conn *grpc.ClientConn
	var err error
	stickySessionEnabled := true
	if opts.TLSCert != nil {
		conn, err = grpc.Dial(protoAddr, grpc.WithTransportCredentials(opts.TLSCert), grpc.WithContextDialer(common.DialerFunc))
	} else {
		conn, err = grpc.Dial(protoAddr, grpc.WithInsecure(), grpc.WithContextDialer(common.DialerFunc))
		stickySessionEnabled = false
	}
	if err != nil {
		err := errors.Wrapf(err, "failed to connect to the gRPC: %s", protoAddr)
		return nil, err
	}

	// build client
	cc := &chainClient{
		ctx:  ctx,
		opts: opts,

		logger: log.WithFields(log.Fields{
			"module": "sdk-go",
			"svc":    "chainClient",
		}),

		conn:      conn,
		txFactory: txFactory,
		canSign:   ctx.Keyring != nil,
		syncMux:   new(sync.Mutex),
		msgC:      make(chan sdk.Msg, msgCommitBatchSizeLimit),
		doneC:     make(chan bool, 1),

		sessionEnabled: stickySessionEnabled,

		txClient:           txtypes.NewServiceClient(conn),
		authQueryClient:    authtypes.NewQueryClient(conn),
		bankQueryClient:    banktypes.NewQueryClient(conn),
		stakingQueryClient: stakingtypes.NewQueryClient(conn),

		multichainQueryClient:  multichainTypes.NewQueryClient(conn),
		attestationQueryClient: attestationTypes.NewQueryClient(conn),
		pricefeedQueryClient:   pricefeedTypes.NewQueryClient(conn),
		metastoreQueryClient:   metastoreTypes.NewQueryClient(conn),
		crosschainQueryClient:  crosschainTypes.NewQueryClient(conn),
		voyagerQueryClient:     voyagerTypes.NewQueryClient(conn),
		txlookupQueryClient:    txlookupTypes.NewQueryClient(conn),
		wasmQueryClient:        wasmTypes.NewQueryClient(conn),
	}

	if cc.canSign {
		var err error

		cc.accNum, cc.accSeq, err = cc.txFactory.AccountRetriever().GetAccountNumberSequence(ctx, ctx.GetFromAddress())
		if err != nil {
			err = errors.Wrap(err, "failed to get initial account num and seq")
			return nil, err
		}

		go cc.runBatchBroadcast()
		go cc.syncTimeoutHeight()
	}

	// create file if not exist
	os.OpenFile(defaultChainCookieName, os.O_RDONLY|os.O_CREATE, 0666)

	// attempt to load from disk
	data, err := os.ReadFile(defaultChainCookieName)
	if err != nil {
		cc.logger.Errorln(err)
	} else {
		cc.sessionCookie = string(data)
		cc.logger.Infoln("chain session cookie loaded from disk")
	}

	return cc, nil
}

func (c *chainClient) syncNonce() {
	num, seq, err := c.txFactory.AccountRetriever().GetAccountNumberSequence(c.ctx, c.ctx.GetFromAddress())
	if err != nil {
		c.logger.WithError(err).Errorln("failed to get account seq")
		return
	} else if num != c.accNum {
		c.logger.WithFields(log.Fields{
			"expected": c.accNum,
			"actual":   num,
		}).Panic("account number changed during nonce sync")
	}

	c.accSeq = seq
}

func (c *chainClient) syncTimeoutHeight() {
	for {
		ctx := context.Background()
		block, err := c.ctx.Client.Block(ctx, nil)
		if err != nil {
			c.logger.WithError(err).Errorln("failed to get current block")
			return
		}
		c.txFactory.WithTimeoutHeight(uint64(block.Block.Height) + defaultTimeoutHeight)
		time.Sleep(defaultTimeoutHeightSyncInterval)
	}
}

// prepareFactory ensures the account defined by ctx.GetFromAddress() exists and
// if the account number and/or the account sequence number are zero (not set),
// they will be queried for and set on the provided Factory. A new Factory with
// the updated fields will be returned.
func (c *chainClient) prepareFactory(clientCtx client.Context, txf tx.Factory) (tx.Factory, error) {
	from := clientCtx.GetFromAddress()

	if err := txf.AccountRetriever().EnsureExists(clientCtx, from); err != nil {
		return txf, err
	}

	initNum, initSeq := txf.AccountNumber(), txf.Sequence()
	if initNum == 0 || initSeq == 0 {
		num, seq, err := txf.AccountRetriever().GetAccountNumberSequence(clientCtx, from)
		if err != nil {
			return txf, err
		}

		if initNum == 0 {
			txf = txf.WithAccountNumber(num)
		}

		if initSeq == 0 {
			txf = txf.WithSequence(seq)
		}
	}

	return txf, nil
}

func (c *chainClient) getAccSeq() uint64 {
	defer func() {
		c.accSeq += 1
	}()
	return c.accSeq
}

func (c *chainClient) setCookie(metadata metadata.MD) {
	if !c.sessionEnabled {
		return
	}
	md := metadata.Get("set-cookie")
	if len(md) > 0 {
		// write to client instance
		c.sessionCookie = md[0]
		// write to disk
		err := os.WriteFile(defaultChainCookieName, []byte(md[0]), 0644)
		if err != nil {
			c.logger.Errorln(err)
			return
		}
		c.logger.Infoln("chain session cookie saved to disk")
	}
}

func (c *chainClient) fetchCookie(ctx context.Context) context.Context {
	var header metadata.MD
	c.txClient.GetTx(context.Background(), &txtypes.GetTxRequest{}, grpc.Header(&header))
	c.setCookie(header)
	time.Sleep(defaultBlockTime)
	return metadata.NewOutgoingContext(ctx, metadata.Pairs("cookie", c.sessionCookie))
}

func (c *chainClient) getCookie(ctx context.Context) context.Context {
	md := metadata.Pairs("cookie", c.sessionCookie)
	if !c.sessionEnabled {
		return metadata.NewOutgoingContext(ctx, md)
	}

	// borrow http request to parse cookie
	header := http.Header{}
	header.Add("Cookie", c.sessionCookie)
	request := http.Request{Header: header}
	cookies := request.Cookies()

	if len(cookies) > 0 {
		// format cookie date into RFC1123 standard
		expiresAt := strings.Replace(cookies[1].Value, "-", " ", -1)
		yyyy := fmt.Sprintf("20%s", expiresAt[12:14])
		expiresAt = expiresAt[:12] + yyyy + expiresAt[14:]

		// parse expire field into unix timestamp
		expiresTimestamp, err := time.Parse(time.RFC1123, expiresAt)
		if err != nil {
			panic(err)
		}

		// renew session if timestamp diff < offset
		timestampDiff := expiresTimestamp.Unix() - time.Now().Unix()
		if timestampDiff < defaultSessionRenewalOffset {
			return c.fetchCookie(ctx)
		}
	} else {
		return c.fetchCookie(ctx)
	}

	return metadata.NewOutgoingContext(ctx, md)
}

func (c *chainClient) QueryClient() *grpc.ClientConn {
	return c.conn
}

func (c *chainClient) ClientContext() client.Context {
	return c.ctx
}

func (c *chainClient) CanSignTransactions() bool {
	return c.canSign
}

func (c *chainClient) FromAddress() sdk.AccAddress {
	if !c.canSign {
		return sdk.AccAddress{}
	}

	return c.ctx.FromAddress
}

func (c *chainClient) Close() {
	if !c.canSign {
		return
	}
	if atomic.CompareAndSwapInt64(&c.closed, 0, 1) {
		close(c.msgC)
	}
	<-c.doneC
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *chainClient) GetBankBalances(ctx context.Context, address string) (*banktypes.QueryAllBalancesResponse, error) {
	req := &banktypes.QueryAllBalancesRequest{
		Address: address,
	}
	return c.bankQueryClient.AllBalances(ctx, req)
}

func (c *chainClient) GetAccount(ctx context.Context, address string) (*authtypes.QueryAccountResponse, error) {
	req := &authtypes.QueryAccountRequest{
		Address: address,
	}
	return c.authQueryClient.Account(ctx, req)
}

func (c *chainClient) GetBankBalance(ctx context.Context, address string, denom string) (*banktypes.QueryBalanceResponse, error) {
	req := &banktypes.QueryBalanceRequest{
		Address: address,
		Denom:   denom,
	}
	return c.bankQueryClient.Balance(ctx, req)
}

// Staking
func (c *chainClient) GetValidator(ctx context.Context, validatorAddr string) (*stakingtypes.QueryValidatorResponse, error) {
	req := &stakingtypes.QueryValidatorRequest{
		ValidatorAddr: validatorAddr,
	}
	return c.stakingQueryClient.Validator(ctx, req)
}

func (c *chainClient) GetAllValidators(ctx context.Context, status string, pagination *query.PageRequest) (*stakingtypes.QueryValidatorsResponse, error) {
	req := &stakingtypes.QueryValidatorsRequest{
		Status:     status,
		Pagination: pagination,
	}
	return c.stakingQueryClient.Validators(ctx, req)
}

// ///////////////////////////////
// //    Wasm           //////////
// //////////////////////////////
func (c *chainClient) SmartContractState(ctx context.Context, contractAddress string, queryData []byte) (*wasmTypes.QuerySmartContractStateResponse, error) {
	return c.wasmQueryClient.SmartContractState(
		ctx,
		&wasmTypes.QuerySmartContractStateRequest{
			Address:   contractAddress,
			QueryData: queryData,
		},
	)
}

func (c *chainClient) RawContractState(ctx context.Context, contractAddress string, queryData []byte) (*wasmTypes.QueryRawContractStateResponse, error) {
	return c.wasmQueryClient.RawContractState(
		ctx,
		&wasmTypes.QueryRawContractStateRequest{
			Address:   contractAddress,
			QueryData: queryData,
		},
	)
}

func (c *chainClient) StoreCode(file string, sender sdk.AccAddress) (codeID int64, err error) {
	wasm, err := os.ReadFile(file)
	if err != nil {
		return
	}

	// gzip the wasm file
	if ioutils.IsWasm(wasm) {
		wasm, err = ioutils.GzipIt(wasm)

		if err != nil {
			return
		}
	} else if !ioutils.IsGzip(wasm) {
		err = fmt.Errorf("invalid input file. Use wasm binary or gzip")
		return
	}

	msg := wasmTypes.MsgStoreCode{
		Sender:       sender.String(),
		WASMByteCode: wasm,
	}

	if err = msg.ValidateBasic(); err != nil {
		return
	}

	txResponse, err := c.SyncBroadcastMsg(&msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, event := range txResponse.GetTxResponse().Events {
		for _, eventAttribute := range event.Attributes {
			var key, value = string(eventAttribute.Key), eventAttribute.Value
			// json.Unmarshal(eventAttribute.Value, &value)
			switch key {
			case "code_id":
				fmt.Println("Key: ", key, "Value", value, "StrValue", (string)(value))

				codeID, err = strconv.ParseInt((string)(value), 10, 64)
				fmt.Println("CodeID", codeID, "err", err)
			}
		}
	}

	time.Sleep(time.Second * 5)
	gasFee, err := c.GetGasFee()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("gas fee:", gasFee)
	return
}

func (c *chainClient) InstantiateContract(codeID uint64, label string, amountStr string, initMsg string, adminStr string, noAdmin bool, sender sdk.AccAddress) (contractAddress string, err error) {
	if label == "" {
		err = errors.New("label is required on all contracts")
		return
	}

	amount, err := sdk.ParseCoinsNormalized(amountStr)
	if err != nil {
		return
	}

	// ensure sensible admin is set (or explicitly immutable)
	if adminStr == "" && !noAdmin {
		err = fmt.Errorf("you must set an admin or explicitly pass --no-admin to make it immutible (wasmd issue #719)")
		return
	}
	if adminStr != "" && noAdmin {
		err = fmt.Errorf("you set an admin and passed --no-admin, those cannot both be true")
		return
	}

	// build and sign the transaction, then broadcast to Tendermint
	msg := wasmTypes.MsgInstantiateContract{
		Sender: sender.String(),
		CodeID: codeID,
		Label:  label,
		Funds:  amount,
		Msg:    []byte(initMsg),
		Admin:  adminStr,
	}

	if err = msg.ValidateBasic(); err != nil {
		return
	}

	txResponse, err := c.SyncBroadcastMsg(&msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, event := range txResponse.GetTxResponse().Events {
		for _, eventAttribute := range event.Attributes {
			var key, value = string(eventAttribute.Key), eventAttribute.Value
			// json.Unmarshal(eventAttribute.Value, &value)
			switch key {
			case "_contract_address":
				contractAddress = (string)(value)
			}
		}
	}

	time.Sleep(time.Second * 5)
	gasFee, err := c.GetGasFee()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("gas fee:", gasFee)
	return
}

func (c *chainClient) ExecuteContract(amountStr string, sender sdk.AccAddress, contractAddr string, execMsg string) error {
	amount, err := sdk.ParseCoinsNormalized(amountStr)
	if err != nil {
		return err
	}

	msg := wasmTypes.MsgExecuteContract{
		Sender:   sender.String(),
		Contract: contractAddr,
		Funds:    amount,
		Msg:      []byte(execMsg),
	}

	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	err = c.QueueBroadcastMsg(&msg)
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 5)
	gasFee, err := c.GetGasFee()
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("gas fee:", gasFee)
	return nil
}

// ///////////////////////////////
// //    Multichain           ////
// //////////////////////////////
func (c *chainClient) GetAllChainConfig(ctx context.Context, pagination *query.PageRequest) (*multichainTypes.QueryAllChainConfigResponse, error) {
	req := &multichainTypes.QueryAllChainConfigRequest{Pagination: pagination}
	return c.multichainQueryClient.ChainConfigAll(ctx, req)
}

func (c *chainClient) GetChainConfig(ctx context.Context, chainId string) (*multichainTypes.QueryGetChainConfigResponse, error) {
	req := &multichainTypes.QueryGetChainConfigRequest{
		ChainId: chainId,
	}
	return c.multichainQueryClient.ChainConfig(ctx, req)
}

func (c *chainClient) GetContractConfig(ctx context.Context, chainId string, contract string) (*multichainTypes.QueryGetContractConfigResponse, error) {
	req := &multichainTypes.QueryGetContractConfigRequest{
		ChainId:         chainId,
		ContractAddress: contract,
	}
	return c.multichainQueryClient.ContractConfig(ctx, req)
}

func (c *chainClient) GetAllContractConfig(ctx context.Context, pagination *query.PageRequest) (*multichainTypes.QueryAllContractConfigResponse, error) {
	req := &multichainTypes.QueryAllContractConfigRequest{Pagination: pagination}
	return c.multichainQueryClient.ContractConfigAll(ctx, req)
}

func (c *chainClient) GetAllContractConfigByChainId(ctx context.Context, chainId string) (*multichainTypes.QueryAllContractConfigByChainIdResponse, error) {
	req := &multichainTypes.QueryAllContractConfigByChainIdRequest{
		ChainId: chainId,
	}
	return c.multichainQueryClient.ContractConfigByChainId(ctx, req)
}

func (c *chainClient) GetNonceObservedStatus(ctx sdk.Context, chainId string, contractAddress string, nonce uint64) (*multichainTypes.QueryGetNonceObservedStatusResponse, error) {
	req := &multichainTypes.QueryGetNonceObservedStatusRequest{
		ChainId:         chainId,
		ContractAddress: contractAddress,
		EventNonce:      nonce,
	}
	return c.multichainQueryClient.NonceObservedStatus(ctx, req)
}

// ////////////////////	///////////
// //    PriceFeed           ////
// //////////////////////////////
func (c *chainClient) GetPriceBySymbol(ctx context.Context, symbol string) (*pricefeedTypes.QueryGetPriceResponse, error) {
	req := &pricefeedTypes.QueryGetPriceRequest{Symbol: symbol}
	return c.pricefeedQueryClient.Price(ctx, req)
}

func (c *chainClient) GetAllSymbolPrices(ctx context.Context, pagination *query.PageRequest) (*pricefeedTypes.QueryAllPriceResponse, error) {
	req := &pricefeedTypes.QueryAllPriceRequest{Pagination: pagination}
	return c.pricefeedQueryClient.PriceAll(ctx, req)
}

func (c *chainClient) GetSymbolRequest(ctx context.Context, symbol string) (*pricefeedTypes.QueryGetSymbolRequestResponse, error) {
	req := &pricefeedTypes.QueryGetSymbolRequestRequest{Symbol: symbol}
	return c.pricefeedQueryClient.SymbolRequest(ctx, req)
}

func (c *chainClient) GetAllSymbolRequest(ctx context.Context, pagination *query.PageRequest) (*pricefeedTypes.QueryAllSymbolRequestResponse, error) {
	req := &pricefeedTypes.QueryAllSymbolRequestRequest{Pagination: pagination}
	return c.pricefeedQueryClient.SymbolRequestAll(ctx, req)
}

func (c *chainClient) GetPriceFeederInfo(ctx context.Context, priceFeederName string) (*pricefeedTypes.QueryGetPriceFeederInfoResponse, error) {
	req := &pricefeedTypes.QueryGetPriceFeederInfoRequest{Name: priceFeederName}
	return c.pricefeedQueryClient.PriceFeederInfo(ctx, req)
}

func (c *chainClient) GetAllPriceFeederInfo(ctx context.Context, pagination *query.PageRequest) (*pricefeedTypes.QueryAllPriceFeederInfoResponse, error) {
	req := &pricefeedTypes.QueryAllPriceFeederInfoRequest{Pagination: pagination}
	return c.pricefeedQueryClient.PriceFeederInfoAll(ctx, req)
}

func (c *chainClient) GetGasPriceByChainId(ctx context.Context, chainId string) (*pricefeedTypes.QueryGetGasPriceResponse, error) {
	req := &pricefeedTypes.QueryGetGasPriceRequest{ChainId: chainId}
	return c.pricefeedQueryClient.GasPrice(ctx, req)
}

func (c *chainClient) GetAllGasPrice(ctx context.Context, pagination *query.PageRequest) (*pricefeedTypes.QueryAllGasPriceResponse, error) {
	req := &pricefeedTypes.QueryAllGasPriceRequest{Pagination: pagination}
	return c.pricefeedQueryClient.GasPriceAll(ctx, req)
}

func (c *chainClient) GetAllWhitelistedIBCChannels(ctx context.Context) (*pricefeedTypes.QueryWhitelistedIBCChannelsResponse, error) {
	req := &pricefeedTypes.QueryWhitelistedIBCChannelsRequest{}
	return c.pricefeedQueryClient.WhitelistedIBCChannels(ctx, req)
}

/////////////////////////////////
////    MetaStore           ////
////////////////////////////////

func (c *chainClient) GetAllMetaInfo(ctx context.Context, pagination *query.PageRequest) (*metastoreTypes.QueryAllMetaInfoResponse, error) {
	req := &metastoreTypes.QueryAllMetaInfoRequest{Pagination: pagination}
	return c.metastoreQueryClient.MetaInfoAll(ctx, req)
}

func (c *chainClient) GetMetaInfo(ctx context.Context, chainId string, dappAddress []byte) (*metastoreTypes.QueryAllMetaInfoResponseByChainAndAddress, error) {
	req := &metastoreTypes.QueryAllMetaInfoRequestByChainAndAddress{
		ChainIds: chainId,
		Address:  string(dappAddress),
	}
	return c.metastoreQueryClient.MetaInfoAllByChainAndAddress(ctx, req)
}

// func (c *chainClient) FeePayerApproval(chainType multichainTypes.ChainType, chainID string, dappAddress []byte) (err error) {
// 	msg := metastoreTypes.MsgApproveFeepayerRequest{
// 		ChainType:   chainType,
// 		ChainId:     chainID,
// 		DaapAddress: dappAddress,
// 		Feepayer:    c.FromAddress().String(),
// 	}

// 	if err = msg.ValidateBasic(); err != nil {
// 		return
// 	}

// 	txResponse, err := c.SyncBroadcastMsg(&msg)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println("txHash:", txResponse.TxResponse.TxHash)
// 	return
// }

// ///////////////////////////////
// //     Attestation           ////
// //////////////////////////////
func (c *chainClient) GetAllValsets(ctx context.Context, pagination *query.PageRequest) (*attestationTypes.QueryAllValsetResponse, error) {
	req := &attestationTypes.QueryAllValsetRequest{Pagination: pagination}
	return c.attestationQueryClient.ValsetAll(ctx, req)
}

func (c *chainClient) GetLatestValsetNonce(ctx context.Context) (*attestationTypes.QueryLatestValsetNonceResponse, error) {
	req := &attestationTypes.QueryLatestValsetNonceRequest{}
	return c.attestationQueryClient.LatestValsetNonce(ctx, req)
}

func (c *chainClient) GetLatestValset(ctx context.Context) (*attestationTypes.QueryLatestValsetResponse, error) {
	req := &attestationTypes.QueryLatestValsetRequest{}
	return c.attestationQueryClient.LatestValset(ctx, req)
}

func (c *chainClient) GetValsetByNonce(ctx context.Context, valsetNonce uint64) (*attestationTypes.QueryGetValsetResponse, error) {
	req := &attestationTypes.QueryGetValsetRequest{
		Nonce: valsetNonce,
	}
	return c.attestationQueryClient.Valset(ctx, req)
}
func (c *chainClient) GetLastEventByValidator(ctx context.Context, chainId string, contract string, validator sdk.ValAddress) (*attestationTypes.QueryLastEventNonceResponse, error) {
	req := &attestationTypes.QueryLastEventNonceRequest{
		ChainId:          chainId,
		Contract:         contract,
		ValidatorAddress: validator.String(),
	}
	return c.attestationQueryClient.LastEventNonce(ctx, req)
}

func (c *chainClient) GetValsetConfirm(ctx context.Context, valsetNonce uint64, destChainType multichainTypes.ChainType, orchestrator string) (*attestationTypes.QueryGetValsetConfirmationResponse, error) {
	req := &attestationTypes.QueryGetValsetConfirmationRequest{
		ValsetNonce:   valsetNonce,
		Orchestrator:  orchestrator,
		DestChainType: destChainType,
	}
	return c.attestationQueryClient.ValsetConfirmation(ctx, req)
}

func (c *chainClient) GetAllValsetConfirms(ctx context.Context, valsetNonce uint64, destChainType multichainTypes.ChainType, pagination *query.PageRequest) (*attestationTypes.QueryAllValsetConfirmationResponse, error) {
	req := &attestationTypes.QueryAllValsetConfirmationRequest{ValsetNonce: valsetNonce, DestChainType: destChainType, Pagination: pagination}
	return c.attestationQueryClient.ValsetConfirmationAll(ctx, req)
}

func (c *chainClient) GetAllOrchestrators(ctx context.Context) (*attestationTypes.QueryListOrchestratorsResponse, error) {
	req := &attestationTypes.QueryListOrchestratorsRequest{}
	return c.attestationQueryClient.ListOrchestrators(ctx, req)
}

func (c *chainClient) GetOrchestratorValidator(ctx context.Context, orchestratorAddr sdk.AccAddress) (*attestationTypes.QueryFetchOrchestratorValidatorResponse, error) {
	req := &attestationTypes.QueryFetchOrchestratorValidatorRequest{
		OrchestratorAddress: orchestratorAddr.String(),
	}
	return c.attestationQueryClient.FetchOrchestratorValidator(ctx, req)
}

func (c *chainClient) GetAllAttestations(ctx context.Context, pagination *query.PageRequest) (*attestationTypes.QueryAllAttestationResponse, error) {
	req := &attestationTypes.QueryAllAttestationRequest{Pagination: pagination}
	return c.attestationQueryClient.AttestationAll(ctx, req)
}

func (c *chainClient) GetAllObservedAttestations(ctx context.Context, pagination *query.PageRequest) (*attestationTypes.QueryAllObservedAttestationResponse, error) {
	req := &attestationTypes.QueryAllObservedAttestationRequest{Pagination: pagination}
	return c.attestationQueryClient.ObservedAttestationAll(ctx, req)
}

// ///////////////////////////////
// //     Crosschain           ////
// //////////////////////////////
func (c *chainClient) GetCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryGetCrosschainRequestRequest{SourceChainId: srcChainId, RequestIdentifier: requestIdentifier}
	return c.crosschainQueryClient.CrosschainRequest(ctx, req)
}

func (c *chainClient) GetAllCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryAllCrosschainRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.CrosschainRequestAll(ctx, req)
}

func (c *chainClient) GetValidatedCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetValidCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryGetValidCrosschainRequestRequest{SourceChainId: srcChainId, RequestIdentifier: requestIdentifier}
	return c.crosschainQueryClient.ValidCrosschainRequest(ctx, req)
}

func (c *chainClient) GetAllValidCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllValidCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryAllValidCrosschainRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.ValidCrosschainRequestAll(ctx, req)
}

func (c *chainClient) GetNativeTransferedCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetNativeTransferedCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryGetNativeTransferedCrosschainRequestRequest{SourceChainId: srcChainId, RequestIdentifier: requestIdentifier}
	return c.crosschainQueryClient.NativeTransferedCrosschainRequest(ctx, req)
}

func (c *chainClient) GetAllNativeTransferedCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllNativeTransferedCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryAllNativeTransferedCrosschainRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.NativeTransferedCrosschainRequestAll(ctx, req)
}

func (c *chainClient) GetReadyToExecuteCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetReadyToExecuteCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryGetReadyToExecuteCrosschainRequestRequest{SourceChainId: srcChainId, RequestIdentifier: requestIdentifier}
	return c.crosschainQueryClient.ReadyToExecuteCrosschainRequest(ctx, req)
}

func (c *chainClient) GetAllReadyToExecuteCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllReadyToExecuteCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryAllReadyToExecuteCrosschainRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.ReadyToExecuteCrosschainRequestAll(ctx, req)
}

func (c *chainClient) GetAllReadyToExecuteCrosschainRequestsByWorkflow(ctx context.Context, workflowType crosschainTypes.WorkflowType, pagination *query.PageRequest) (*crosschainTypes.QueryReadyToExecuteCrosschainRequestByWorkflowResponse, error) {
	req := &crosschainTypes.QueryReadyToExecuteCrosschainRequestByWorkflowRequest{WorkflowType: uint64(workflowType), Pagination: pagination}
	return c.crosschainQueryClient.ReadyToExecuteCrosschainRequestByWorkflow(ctx, req)
}

func (c *chainClient) GetAllReadyToExecuteCrosschainRequestsByWorkflowAndRelayer(ctx context.Context, workflowType crosschainTypes.WorkflowType, relayerType crosschainTypes.RelayerType, pagination *query.PageRequest) (*crosschainTypes.QueryReadyToExecuteCrosschainRequestByWorkflowAndRelayerResponse, error) {
	req := &crosschainTypes.QueryReadyToExecuteCrosschainRequestByWorkflowAndRelayerRequest{WorkflowType: uint64(workflowType), RelayerType: uint64(relayerType), Pagination: pagination}
	return c.crosschainQueryClient.ReadyToExecuteCrosschainRequestByWorkflowAndRelayer(ctx, req)
}

func (c *chainClient) GetAllReadyToExecuteCrosschainAckRequestsByWorkflow(ctx context.Context, ackWorkflowType crosschainTypes.WorkflowType, pagination *query.PageRequest) (*crosschainTypes.QueryReadyToExecuteCrosschainAckRequestByWorkflowResponse, error) {
	req := &crosschainTypes.QueryReadyToExecuteCrosschainAckRequestByWorkflowRequest{AckWorkflowType: uint64(ackWorkflowType), Pagination: pagination}
	return c.crosschainQueryClient.ReadyToExecuteCrosschainAckRequestByWorkflow(ctx, req)
}

func (c *chainClient) GetAllReadyToExecuteCrosschainAckRequestsByWorkflowAndRelayer(ctx context.Context, ackWorkflowType crosschainTypes.WorkflowType, ackRelayerType crosschainTypes.RelayerType, pagination *query.PageRequest) (*crosschainTypes.QueryReadyToExecuteCrosschainAckRequestByWorkflowAndRelayerResponse, error) {
	req := &crosschainTypes.QueryReadyToExecuteCrosschainAckRequestByWorkflowAndRelayerRequest{AckWorkflowType: uint64(ackWorkflowType), AckRelayerType: uint64(ackRelayerType), Pagination: pagination}
	return c.crosschainQueryClient.ReadyToExecuteCrosschainAckRequestByWorkflowAndRelayer(ctx, req)
}

func (c *chainClient) GetBlockedCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetBlockedCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryGetBlockedCrosschainRequestRequest{SourceChainId: srcChainId, RequestIdentifier: requestIdentifier}
	return c.crosschainQueryClient.BlockedCrosschainRequest(ctx, req)
}

func (c *chainClient) GetAllBlockedCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllBlockedCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryAllBlockedCrosschainRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.BlockedCrosschainRequestAll(ctx, req)
}

func (c *chainClient) GetExpiredCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetExpiredCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryGetExpiredCrosschainRequestRequest{SourceChainId: srcChainId, RequestIdentifier: requestIdentifier}
	return c.crosschainQueryClient.ExpiredCrosschainRequest(ctx, req)
}

func (c *chainClient) GetAllExpiredCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllExpiredCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryAllExpiredCrosschainRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.ExpiredCrosschainRequestAll(ctx, req)
}

func (c *chainClient) GetExecutedCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetExecutedCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryGetExecutedCrosschainRequestRequest{SourceChainId: srcChainId, RequestIdentifier: requestIdentifier}
	return c.crosschainQueryClient.ExecutedCrosschainRequest(ctx, req)
}

func (c *chainClient) GetAllExecutedCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllExecutedCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryAllExecutedCrosschainRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.ExecutedCrosschainRequestAll(ctx, req)
}

func (c *chainClient) GetFeesSettledCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetFeesSettledCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryGetFeesSettledCrosschainRequestRequest{SourceChainId: srcChainId, RequestIdentifier: requestIdentifier}
	return c.crosschainQueryClient.FeesSettledCrosschainRequest(ctx, req)
}

func (c *chainClient) GetAllFeesSettledCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllFeesSettledCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryAllFeesSettledCrosschainRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.FeesSettledCrosschainRequestAll(ctx, req)
}

func (c *chainClient) GetCompletedCrosschainRequest(ctx context.Context, srcChainId string, requestIdentifier uint64) (*crosschainTypes.QueryGetCompletedCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryGetCompletedCrosschainRequestRequest{SourceChainId: srcChainId, RequestIdentifier: requestIdentifier}
	return c.crosschainQueryClient.CompletedCrosschainRequest(ctx, req)
}

func (c *chainClient) GetAllCompletedCrosschainRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllCompletedCrosschainRequestResponse, error) {
	req := &crosschainTypes.QueryAllCompletedCrosschainRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.CompletedCrosschainRequestAll(ctx, req)
}

func (c *chainClient) GetAllCrosschainRequestConfirmations(ctx context.Context, pagination *query.PageRequest, sourceChainId string, requestIdentifier uint64, claimHash []byte) (*crosschainTypes.QueryAllCrosschainRequestConfirmResponse, error) {
	req := &crosschainTypes.QueryAllCrosschainRequestConfirmRequest{
		SourceChainId:     sourceChainId,
		RequestIdentifier: requestIdentifier,
		ClaimHash:         claimHash,
		Pagination:        pagination,
	}
	return c.crosschainQueryClient.CrosschainRequestConfirmAll(ctx, req)
}

func (c *chainClient) GetCrosschainRequestConfirmation(ctx context.Context, pagination *query.PageRequest, sourceChainId string, requestIdentifier uint64, claimHash []byte, orchestrator string) (*crosschainTypes.QueryGetCrosschainRequestConfirmResponse, error) {
	req := &crosschainTypes.QueryGetCrosschainRequestConfirmRequest{
		Orchestrator:      orchestrator,
		SourceChainId:     sourceChainId,
		RequestIdentifier: requestIdentifier,
		ClaimHash:         claimHash,
	}
	return c.crosschainQueryClient.CrosschainRequestConfirm(ctx, req)
}

// Crosschain Ack
func (c *chainClient) GetCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryGetCrosschainAckRequestRequest{AckSrcChainId: ackSrcChainId, AckRequestIdentifier: ackRequestIdentifier}
	return c.crosschainQueryClient.CrosschainAckRequest(ctx, req)
}

func (c *chainClient) GetAllCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryAllCrosschainAckRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.CrosschainAckRequestAll(ctx, req)
}

func (c *chainClient) GetValidCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetValidCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryGetValidCrosschainAckRequestRequest{AckSrcChainId: ackSrcChainId, AckRequestIdentifier: ackRequestIdentifier}
	return c.crosschainQueryClient.ValidCrosschainAckRequest(ctx, req)
}

func (c *chainClient) GetAllValidCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllValidCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryAllValidCrosschainAckRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.ValidCrosschainAckRequestAll(ctx, req)
}

func (c *chainClient) GetBlockedCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetBlockedCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryGetBlockedCrosschainAckRequestRequest{AckSrcChainId: ackSrcChainId, AckRequestIdentifier: ackRequestIdentifier}
	return c.crosschainQueryClient.BlockedCrosschainAckRequest(ctx, req)
}

func (c *chainClient) GetAllBlockedCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllBlockedCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryAllBlockedCrosschainAckRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.BlockedCrosschainAckRequestAll(ctx, req)
}

func (c *chainClient) GetReadyToExecuteCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetReadyToExecuteCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryGetReadyToExecuteCrosschainAckRequestRequest{AckSrcChainId: ackSrcChainId, AckRequestIdentifier: ackRequestIdentifier}
	return c.crosschainQueryClient.ReadyToExecuteCrosschainAckRequest(ctx, req)
}

func (c *chainClient) GetAllReadyToExecuteCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllReadyToExecuteCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryAllReadyToExecuteCrosschainAckRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.ReadyToExecuteCrosschainAckRequestAll(ctx, req)
}

func (c *chainClient) GetExecutedCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetExecutedCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryGetExecutedCrosschainAckRequestRequest{AckSrcChainId: ackSrcChainId, AckRequestIdentifier: ackRequestIdentifier}
	return c.crosschainQueryClient.ExecutedCrosschainAckRequest(ctx, req)
}

func (c *chainClient) GetAllExecutedCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllExecutedCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryAllExecutedCrosschainAckRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.ExecutedCrosschainAckRequestAll(ctx, req)
}

func (c *chainClient) GetFeesSettledCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetFeesSettledCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryGetFeesSettledCrosschainAckRequestRequest{AckSrcChainId: ackSrcChainId, AckRequestIdentifier: ackRequestIdentifier}
	return c.crosschainQueryClient.FeesSettledCrosschainAckRequest(ctx, req)
}

func (c *chainClient) GetAllFeesSettledCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllFeesSettledCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryAllFeesSettledCrosschainAckRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.FeesSettledCrosschainAckRequestAll(ctx, req)
}

func (c *chainClient) GetCompletedCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetCompletedCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryGetCompletedCrosschainAckRequestRequest{AckSrcChainId: ackSrcChainId, AckRequestIdentifier: ackRequestIdentifier}
	return c.crosschainQueryClient.CompletedCrosschainAckRequest(ctx, req)
}

func (c *chainClient) GetAllCompletedCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllCompletedCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryAllCompletedCrosschainAckRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.CompletedCrosschainAckRequestAll(ctx, req)
}

func (c *chainClient) GetExpiredCrosschainAckRequest(ctx context.Context, ackSrcChainId string, ackRequestIdentifier uint64) (*crosschainTypes.QueryGetExpiredCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryGetExpiredCrosschainAckRequestRequest{AckSrcChainId: ackSrcChainId, AckRequestIdentifier: ackRequestIdentifier}
	return c.crosschainQueryClient.ExpiredCrosschainAckRequest(ctx, req)
}

func (c *chainClient) GetAllExpiredCrosschainAckRequests(ctx context.Context, pagination *query.PageRequest) (*crosschainTypes.QueryAllExpiredCrosschainAckRequestResponse, error) {
	req := &crosschainTypes.QueryAllExpiredCrosschainAckRequestRequest{Pagination: pagination}
	return c.crosschainQueryClient.ExpiredCrosschainAckRequestAll(ctx, req)
}

func (c *chainClient) GetCrosschainAckRequestConfirmation(ctx context.Context, pagination *query.PageRequest, ackSrcChainId string, ackRequestIdentifier uint64, claimHash []byte, orchestrator string) (*crosschainTypes.QueryGetCrosschainAckRequestConfirmResponse, error) {
	req := &crosschainTypes.QueryGetCrosschainAckRequestConfirmRequest{
		Orchestrator:         orchestrator,
		AckSrcChainId:        ackSrcChainId,
		AckRequestIdentifier: ackRequestIdentifier,
		ClaimHash:            claimHash,
	}
	return c.crosschainQueryClient.CrosschainAckRequestConfirm(ctx, req)
}

func (c *chainClient) GetAllCrosschainRequestAckConfirmations(ctx context.Context, pagination *query.PageRequest, ackSrcChainId string, ackRequestIdentifier uint64, claimHash []byte) (*crosschainTypes.QueryAllCrosschainAckRequestConfirmResponse, error) {
	req := &crosschainTypes.QueryAllCrosschainAckRequestConfirmRequest{
		AckSrcChainId:        ackSrcChainId,
		AckRequestIdentifier: ackRequestIdentifier,
		ClaimHash:            claimHash,
		Pagination:           pagination,
	}
	return c.crosschainQueryClient.CrosschainAckRequestConfirmAll(ctx, req)
}

func (c *chainClient) GetCrosschainAckReceipt(ctx context.Context, ackReceiptSrcChainId string, ackReceiptIdentifier uint64) (*crosschainTypes.QueryGetCrosschainAckReceiptResponse, error) {
	req := &crosschainTypes.QueryGetCrosschainAckReceiptRequest{AckReceiptSrcChainId: ackReceiptSrcChainId, AckReceiptIdentifier: ackReceiptIdentifier}
	return c.crosschainQueryClient.CrosschainAckReceipt(ctx, req)
}

// Txlookup module
func (c *chainClient) GetAllAdhocRequests(ctx context.Context, pagination *query.PageRequest) (*txlookupTypes.QueryListAdhocResponse, error) {
	req := &txlookupTypes.QueryListAdhocRequest{Pagination: pagination}
	return c.txlookupQueryClient.ListAdhocRequests(ctx, req)
}

// SyncBroadcastMsg sends Tx to chain and waits until Tx is included in block.
func (c *chainClient) SyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	c.syncMux.Lock()
	defer c.syncMux.Unlock()

	c.txFactory = c.txFactory.WithSequence(c.accSeq)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	res, err := c.broadcastTx(c.ctx, c.txFactory, true, msgs...)

	if err != nil {
		if strings.Contains(err.Error(), "account sequence mismatch") {
			c.syncNonce()
			c.txFactory = c.txFactory.WithSequence(c.accSeq)
			c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
			log.Debugln("retrying broadcastTx with nonce", c.accSeq)
			res, err = c.broadcastTx(c.ctx, c.txFactory, true, msgs...)
		}
		if err != nil {
			resJSON, _ := json.MarshalIndent(res, "", "\t")
			c.logger.WithField("size", len(msgs)).WithError(err).Errorln("failed to commit msg batch:", string(resJSON))
			return nil, err
		}
	}

	c.accSeq++

	return res, nil
}

func (c *chainClient) SimulateMsg(clientCtx client.Context, msgs ...sdk.Msg) (*txtypes.SimulateResponse, error) {
	c.txFactory = c.txFactory.WithSequence(c.accSeq)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	txf, err := c.prepareFactory(clientCtx, c.txFactory)
	if err != nil {
		err = errors.Wrap(err, "failed to prepareFactory")
		return nil, err
	}

	simTxBytes, err := txf.BuildSimTx(msgs...)
	if err != nil {
		err = errors.Wrap(err, "failed to build sim tx bytes")
		return nil, err
	}

	ctx := context.Background()
	ctx = c.getCookie(ctx)
	var header metadata.MD
	simRes, err := c.txClient.Simulate(ctx, &txtypes.SimulateRequest{TxBytes: simTxBytes}, grpc.Header(&header))
	if err != nil {
		err = errors.Wrap(err, "failed to CalculateGas")
		return nil, err
	}

	return simRes, nil
}

// AsyncBroadcastMsg sends Tx to chain and doesn't wait until Tx is included in block. This method
// cannot be used for rapid Tx sending, it is expected that you wait for transaction status with
// external tools. If you want sdk to wait for it, use SyncBroadcastMsg.
func (c *chainClient) AsyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	c.syncMux.Lock()
	defer c.syncMux.Unlock()

	c.txFactory = c.txFactory.WithSequence(c.accSeq)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	res, err := c.broadcastTx(c.ctx, c.txFactory, false, msgs...)
	if err != nil {
		if strings.Contains(err.Error(), "account sequence mismatch") {
			c.syncNonce()
			c.txFactory = c.txFactory.WithSequence(c.accSeq)
			c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
			log.Debugln("retrying broadcastTx with nonce", c.accSeq)
			res, err = c.broadcastTx(c.ctx, c.txFactory, false, msgs...)
		}
		if err != nil {
			resJSON, _ := json.MarshalIndent(res, "", "\t")
			c.logger.WithField("size", len(msgs)).WithError(err).Errorln("failed to commit msg batch:", string(resJSON))
			return nil, err
		}
	}

	c.accSeq++

	return res, nil
}

func (c *chainClient) broadcastTx(
	clientCtx client.Context,
	txf tx.Factory,
	await bool,
	msgs ...sdk.Msg,
) (*txtypes.BroadcastTxResponse, error) {

	txf, err := c.prepareFactory(clientCtx, txf)
	if err != nil {
		err = errors.Wrap(err, "failed to prepareFactory")
		return nil, err
	}

	fmt.Println("prepare tx factory", "Simulate", clientCtx.Simulate)
	ctx := context.Background()
	if clientCtx.Simulate {
		simTxBytes, err := txf.BuildSimTx(msgs...)
		if err != nil {
			err = errors.Wrap(err, "failed to build sim tx bytes")
			return nil, err
		}
		ctx := c.getCookie(ctx)
		var header metadata.MD
		simRes, err := c.txClient.Simulate(ctx, &txtypes.SimulateRequest{TxBytes: simTxBytes}, grpc.Header(&header))
		if err != nil {
			err = errors.Wrap(err, "failed to CalculateGas")
			return nil, err
		}

		adjustedGas := uint64(txf.GasAdjustment() * float64(simRes.GasInfo.GasUsed))
		txf = txf.WithGas(adjustedGas)

		c.gasWanted = adjustedGas
	}

	fmt.Println("Build unsigned tx")
	txn, err := txf.BuildUnsignedTx(msgs...)

	if err != nil {
		err = errors.Wrap(err, "failed to BuildUnsignedTx")
		return nil, err
	}

	txn.SetFeeGranter(clientCtx.GetFeeGranterAddress())
	fmt.Println("Sign tx", "signer", clientCtx.GetFromName())
	err = tx.Sign(context.Background(), txf, clientCtx.GetFromName(), txn, true)
	if err != nil {
		err = errors.Wrap(err, "failed to Sign Tx")
		return nil, err
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(txn.GetTx())
	if err != nil {
		err = errors.Wrap(err, "failed TxEncoder to encode Tx")
		return nil, err
	}

	fmt.Println("Broadcast tx request")
	req := txtypes.BroadcastTxRequest{
		txBytes,
		txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
	}
	// use our own client to broadcast tx
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.txClient.BroadcastTx(ctx, &req, grpc.Header(&header))
	if !await || err != nil {
		return res, err
	}

	fmt.Println("Broadcasted tx")
	awaitCtx, cancelFn := context.WithTimeout(context.Background(), defaultBroadcastTimeout)
	defer cancelFn()

	txHash, _ := hex.DecodeString(res.TxResponse.TxHash)
	t := time.NewTimer(defaultBroadcastStatusPoll)

	fmt.Println("Broadcasted tx", "txHash", res.TxResponse.TxHash)

	for {
		select {
		case <-awaitCtx.Done():
			err := errors.Wrapf(ErrTimedOut, "%s", res.TxResponse.TxHash)
			fmt.Println("Error timedout", err)
			t.Stop()
			return nil, err
		case <-t.C:
			resultTx, err := clientCtx.Client.Tx(awaitCtx, txHash, false)
			if err != nil {
				if errRes := client.CheckTendermintError(err, txBytes); errRes != nil {
					fmt.Println("Tendermint error", err)
					return &txtypes.BroadcastTxResponse{TxResponse: errRes}, err
				}

				// log.WithError(err).Warningln("Tx Error for Hash:")

				t.Reset(defaultBroadcastStatusPoll)
				continue

			} else if resultTx.Height > 0 {
				resResultTx := sdk.NewResponseResultTx(resultTx, res.TxResponse.Tx, res.TxResponse.Timestamp)
				res = &txtypes.BroadcastTxResponse{TxResponse: resResultTx}
				t.Stop()
				return res, err
			}

			t.Reset(defaultBroadcastStatusPoll)
		}
	}
}

// QueueBroadcastMsg enqueues a list of messages. Messages will added to the queue
// and grouped into Txns in chunks. Use this method to mass broadcast Txns with efficiency.
func (c *chainClient) QueueBroadcastMsg(msgs ...sdk.Msg) error {
	if !c.canSign {
		return ErrReadOnly
	} else if atomic.LoadInt64(&c.closed) == 1 {
		return ErrQueueClosed
	}

	t := time.NewTimer(10 * time.Second)
	for _, msg := range msgs {
		select {
		case <-t.C:
			return ErrEnqueueTimeout
		case c.msgC <- msg:
		}
	}
	t.Stop()

	return nil
}

func (c *chainClient) runBatchBroadcast() {
	expirationTimer := time.NewTimer(msgCommitBatchTimeLimit)
	msgBatch := make([]sdk.Msg, 0, msgCommitBatchSizeLimit)

	submitBatch := func(toSubmit []sdk.Msg) {
		c.syncMux.Lock()
		defer c.syncMux.Unlock()
		c.txFactory = c.txFactory.WithSequence(c.accSeq)
		c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
		log.Debugln("broadcastTx with nonce", c.accSeq)
		log.Debugln("txf", "fees", c.txFactory.Fees(), "gas", c.txFactory.Gas(), "gasPrice", c.txFactory.GasPrices())
		res, err := c.broadcastTx(c.ctx, c.txFactory, true, toSubmit...)
		if err != nil {
			fmt.Println("Error broadcasting tx", err)
			if strings.Contains(err.Error(), "account sequence mismatch") {
				c.syncNonce()
				c.txFactory = c.txFactory.WithSequence(c.accSeq)
				c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
				log.Debugln("retrying broadcastTx with nonce", c.accSeq)
				res, err = c.broadcastTx(c.ctx, c.txFactory, true, toSubmit...)
			}
			if err != nil {
				resJSON, _ := json.MarshalIndent(res, "", "\t")
				c.logger.WithField("size", len(toSubmit)).WithError(err).Errorln("failed to commit msg batch:", string(resJSON))
				return
			}
		}

		if res.TxResponse.Code != 0 {
			err = errors.Errorf("error %d (%s): %s", res.TxResponse.Code, res.TxResponse.Codespace, res.TxResponse.RawLog)
			log.WithField("txHash", res.TxResponse.TxHash).WithError(err).Errorln("failed to commit msg batch")
		} else {
			log.WithField("txHash", res.TxResponse.TxHash).Debugln("msg batch committed successfully at height", res.TxResponse.Height)
		}

		c.accSeq++
		log.Debugln("nonce incremented to", c.accSeq)
		log.Debugln("gas wanted: ", c.gasWanted)
	}

	for {
		select {
		case msg, ok := <-c.msgC:
			if !ok {
				// exit required
				if len(msgBatch) > 0 {
					submitBatch(msgBatch)
				}

				close(c.doneC)
				return
			}

			msgBatch = append(msgBatch, msg)

			if len(msgBatch) >= msgCommitBatchSizeLimit {
				toSubmit := msgBatch
				msgBatch = msgBatch[:0]
				expirationTimer.Reset(msgCommitBatchTimeLimit)

				submitBatch(toSubmit)
			}
		case <-expirationTimer.C:
			if len(msgBatch) > 0 {
				toSubmit := msgBatch
				msgBatch = msgBatch[:0]
				expirationTimer.Reset(msgCommitBatchTimeLimit)
				submitBatch(toSubmit)
			} else {
				expirationTimer.Reset(msgCommitBatchTimeLimit)
			}
		}
	}
}

func (c *chainClient) GetGasFee() (string, error) {
	gasPrices := strings.Trim(c.opts.GasPrices, "route")

	gas, err := strconv.ParseFloat(gasPrices, 64)

	if err != nil {
		return "", err
	}

	gasFeeAdjusted := gas * float64(c.gasWanted) / math.Pow(10, 18)
	gasFeeFormatted := strconv.FormatFloat(gasFeeAdjusted, 'f', -1, 64)
	c.gasFee = gasFeeFormatted

	return c.gasFee, err
}
