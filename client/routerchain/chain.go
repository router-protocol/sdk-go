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
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/pkg/errors"
	"github.com/router-protocol/sdk-go/client/routerchain/common"
	log "github.com/xlab/suplog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	"github.com/CosmWasm/wasmd/x/wasm/ioutils"
	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"

	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	crosstalkTypes "github.com/router-protocol/sdk-go/routerchain/crosstalk/types"
	inboundTypes "github.com/router-protocol/sdk-go/routerchain/inbound/types"
	metastoreTypes "github.com/router-protocol/sdk-go/routerchain/metastore/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	oracleTypes "github.com/router-protocol/sdk-go/routerchain/oracle/types"
	outboundTypes "github.com/router-protocol/sdk-go/routerchain/outbound/types"
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

	GetBankBalances(ctx context.Context, address string) (*banktypes.QueryAllBalancesResponse, error)
	GetBankBalance(ctx context.Context, address string, denom string) (*banktypes.QueryBalanceResponse, error)
	GetAccount(ctx context.Context, address string) (*authtypes.QueryAccountResponse, error)

	// MultiChain
	GetAllChainConfig(ctx context.Context) (*multichainTypes.QueryAllChainConfigResponse, error)
	GetChainConfig(ctx context.Context, chainType uint64, chainId string) (*multichainTypes.QueryGetChainConfigResponse, error)

	// Attestation
	GetLatestValsetNonce(ctx context.Context) (*attestationTypes.QueryLatestValsetNonceResponse, error)
	GetAllValsets(ctx context.Context) (*attestationTypes.QueryAllValsetResponse, error)
	GetValsetByNonce(c context.Context, valsetNonce uint64) (*attestationTypes.QueryGetValsetResponse, error)
	GetLatestValset(ctx context.Context) (*attestationTypes.QueryLatestValsetResponse, error)
	GetLastEventByValidator(ctx context.Context, chainType multichainTypes.ChainType, chainId string, validator sdk.ValAddress) (*attestationTypes.QueryLastEventNonceResponse, error)
	GetAllOrchestrators(ctx context.Context) (*attestationTypes.QueryListOrchestratorsResponse, error)
	GetOrchestratorValidator(ctx context.Context, orchestratorAddr sdk.AccAddress) (*attestationTypes.QueryFetchOrchestratorValidatorResponse, error)

	// Inbound
	GetIncomingTx(ctx context.Context, chainType uint64, chainID string, eventNonce uint64) (*inboundTypes.QueryGetIncomingTxResponse, error)

	// Outbound
	GetAllOutgoingBatchTx(ctx context.Context, pagination *query.PageRequest) (*outboundTypes.QueryAllOutgoingBatchTxResponse, error)
	GetOutgoingBatchTx(ctx context.Context, destinationChainType uint64, destinationChainId string, sourceAddress string, batchNonce uint64) (*outboundTypes.QueryGetOutgoingBatchTxResponse, error)
	GetOutgoingBatchTxConfirm(ctx context.Context, destinationChainType uint64, destinationChainId string, sourceAddress string, batchNonce uint64, orchestrator string) (*outboundTypes.QueryGetOutgoingBatchConfirmResponse, error)
	GetAllOutgoingBatchTxConfirms(ctx context.Context, destinationChainType uint64, destinationChainId string, sourceAddress string, batchNonce uint64) (*outboundTypes.QueryAllOutgoingBatchConfirmResponse, error)
	GetLastOutgoingBatchNonce(ctx context.Context, destinationChainType multichainTypes.ChainType, destinationChainId string, sourceAddress string) (*outboundTypes.QueryLastOutboundBatchNonceResponse, error)

	// CrossTalk
	GetCrossTalkRequest(ctx context.Context, chainType uint64, chainID string, eventNonce uint64) (*crosstalkTypes.QueryGetCrossTalkRequestResponse, error)
	GetAllCrossTalkRequest(ctx context.Context, pagination *query.PageRequest) (*crosstalkTypes.QueryAllCrossTalkRequestResponse, error)
	GetAllCrosstalkRequestConfirmations(ctx context.Context, pagination *query.PageRequest, sourceChainType uint64, sourceChainId string, eventNonce uint64, claimHash []byte) (*crosstalkTypes.QueryAllCrosstalkRequestConfirmResponse, error)
	GetAllCrossTalkAckRequest(ctx context.Context, pagination *query.PageRequest) (*crosstalkTypes.QueryAllCrossTalkAckRequestResponse, error)
	GetCrossTalkAckRequest(ctx context.Context, chainType uint64, chainID string, eventNonce uint64) (*crosstalkTypes.QueryGetCrossTalkAckRequestResponse, error)
	GetAllCrosstalkAckRequestConfirmations(ctx context.Context, pagination *query.PageRequest, chainType uint64, chainId string, eventNonce uint64, claimHash []byte) (*crosstalkTypes.QueryAllCrosstalkAckRequestConfirmResponse, error)
	GetAllCrossTalkRequestByStatus(ctx context.Context, status crosstalkTypes.CrossTalkRequestStatus, pagination *query.PageRequest) (*crosstalkTypes.QueryCrosstalkRequestByStatusResponse, error)
	GetAllCrossTalkAckRequestByStatus(ctx context.Context, status crosstalkTypes.CrossTalkAckRequestStatus, pagination *query.PageRequest) (*crosstalkTypes.QueryAllCrosstalkAckRequestsByStatusResponse, error)

	// Wasm
	StoreCode(file string, sender sdk.AccAddress) (int64, error)
	InstantiateContract(codeID uint64, label string, amountStr string, initMsg string, adminStr string, noAdmin bool, sender sdk.AccAddress) (string, error)
	ExecuteContract(amountStr string, sender sdk.AccAddress, contractAddr string, execMsg string) error
	SmartContractState(ctx context.Context, contractAddress string, queryData []byte) (*wasmTypes.QuerySmartContractStateResponse, error)
	RawContractState(ctx context.Context, contractAddress string, queryData []byte) (*wasmTypes.QueryRawContractStateResponse, error)

	GetGasFee() (string, error)
	Close()
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

	txClient         txtypes.ServiceClient
	authQueryClient  authtypes.QueryClient
	bankQueryClient  banktypes.QueryClient
	authzQueryClient authztypes.QueryClient

	// Custom Query clients
	multichainQueryClient  multichainTypes.QueryClient
	attestationQueryClient attestationTypes.QueryClient
	inboundQueryClient     inboundTypes.QueryClient
	outboundQueryClient    outboundTypes.QueryClient
	crosstalkQueryClient   crosstalkTypes.QueryClient
	oracleQueryClient      oracleTypes.QueryClient
	wasmQueryClient        wasmTypes.QueryClient
	metastoreQueryClient   metastoreTypes.QueryClient

	closed  int64
	canSign bool
}

func InitialiseChainClient(networkName string, keyringFrom string, passphrase string, privateKey string) ChainClient {
	network := common.LoadNetwork(networkName, "k8s")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Network", network)
	senderAddress, cosmosKeyring, err := InitCosmosKeyring(
		os.Getenv("HOME")+"/.routerd",
		"routerd",
		"file",
		keyringFrom,
		passphrase,
		privateKey, // keyring will be used if pk not provided
		false,
	)

	if err != nil {
		panic(err)
	}

	// initialize grpc client
	clientCtx, err := NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)

	if err != nil {
		fmt.Println(err)
	}

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmRPC)

	routerchainClient, err := NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		// common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000route"),
	)

	if err != nil {
		fmt.Println(err)
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

		txClient:         txtypes.NewServiceClient(conn),
		authQueryClient:  authtypes.NewQueryClient(conn),
		bankQueryClient:  banktypes.NewQueryClient(conn),
		authzQueryClient: authztypes.NewQueryClient(conn),

		multichainQueryClient:  multichainTypes.NewQueryClient(conn),
		attestationQueryClient: attestationTypes.NewQueryClient(conn),
		inboundQueryClient:     inboundTypes.NewQueryClient(conn),
		outboundQueryClient:    outboundTypes.NewQueryClient(conn),
		crosstalkQueryClient:   crosstalkTypes.NewQueryClient(conn),
		oracleQueryClient:      oracleTypes.NewQueryClient(conn),
		metastoreQueryClient:   metastoreTypes.NewQueryClient(conn),
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

/////////////////////////////////
////    Wasm           //////////
////////////////////////////////
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

/////////////////////////////////
////    Multichain           ////
////////////////////////////////
func (c *chainClient) GetAllChainConfig(ctx context.Context) (*multichainTypes.QueryAllChainConfigResponse, error) {
	req := &multichainTypes.QueryAllChainConfigRequest{}
	return c.multichainQueryClient.ChainConfigAll(ctx, req)
}

func (c *chainClient) GetChainConfig(ctx context.Context, chainType uint64, chainId string) (*multichainTypes.QueryGetChainConfigResponse, error) {
	req := &multichainTypes.QueryGetChainConfigRequest{
		ChainType: chainType,
		ChainId:   chainId,
	}
	return c.multichainQueryClient.ChainConfig(ctx, req)
}

/////////////////////////////////
////     Attestation           ////
////////////////////////////////
func (c *chainClient) GetAllValsets(ctx context.Context) (*attestationTypes.QueryAllValsetResponse, error) {
	req := &attestationTypes.QueryAllValsetRequest{}
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
func (c *chainClient) GetLastEventByValidator(ctx context.Context, chainType multichainTypes.ChainType, chainId string, validator sdk.ValAddress) (*attestationTypes.QueryLastEventNonceResponse, error) {
	req := &attestationTypes.QueryLastEventNonceRequest{
		ChainType:        uint64(chainType),
		ChainId:          chainId,
		ValidatorAddress: validator.String(),
	}
	return c.attestationQueryClient.LastEventNonce(ctx, req)
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

/////////////////////////////////
////     Inbound           //////
////////////////////////////////
func (c *chainClient) GetIncomingTx(ctx context.Context, chainType uint64, chainID string, eventNonce uint64) (*inboundTypes.QueryGetIncomingTxResponse, error) {
	req := &inboundTypes.QueryGetIncomingTxRequest{
		ChainType:  chainType,
		ChainId:    chainID,
		EventNonce: eventNonce,
	}

	return c.inboundQueryClient.IncomingTx(ctx, req)
}

/////////////////////////////////
////     Outbound           ////
////////////////////////////////
func (c *chainClient) GetAllOutgoingBatchTx(ctx context.Context, pagination *query.PageRequest) (*outboundTypes.QueryAllOutgoingBatchTxResponse, error) {
	req := &outboundTypes.QueryAllOutgoingBatchTxRequest{Pagination: pagination}
	return c.outboundQueryClient.OutgoingBatchTxAll(ctx, req)
}

func (c *chainClient) GetOutgoingBatchTx(ctx context.Context, destinationChainType uint64, destinationChainId string, sourceAddress string, batchNonce uint64) (*outboundTypes.QueryGetOutgoingBatchTxResponse, error) {
	req := &outboundTypes.QueryGetOutgoingBatchTxRequest{
		DestinationChainType: destinationChainType,
		DestinationChainId:   destinationChainId,
		SourceAddress:        sourceAddress,
		Nonce:                batchNonce,
	}
	return c.outboundQueryClient.OutgoingBatchTx(ctx, req)
}

func (c *chainClient) GetAllOutgoingBatchTxConfirms(ctx context.Context, destinationChainType uint64, destinationChainId string, sourceAddress string, batchNonce uint64) (*outboundTypes.QueryAllOutgoingBatchConfirmResponse, error) {
	req := &outboundTypes.QueryAllOutgoingBatchConfirmRequest{
		DestinationChainType: destinationChainType,
		DestinationChainId:   destinationChainId,
		SourceAddress:        sourceAddress,
		Nonce:                batchNonce,
	}
	return c.outboundQueryClient.OutgoingBatchConfirmAll(ctx, req)
}

func (c *chainClient) GetOutgoingBatchTxConfirm(ctx context.Context, destinationChainType uint64, destinationChainId string, sourceAddress string, batchNonce uint64, orchestrator string) (*outboundTypes.QueryGetOutgoingBatchConfirmResponse, error) {
	req := &outboundTypes.QueryGetOutgoingBatchConfirmRequest{
		DestinationChainType: destinationChainType,
		DestinationChainId:   destinationChainId,
		SourceAddress:        sourceAddress,
		Nonce:                batchNonce,
		Orchestrator:         orchestrator,
	}
	return c.outboundQueryClient.OutgoingBatchConfirm(ctx, req)
}

func (c *chainClient) GetLastOutgoingBatchNonce(ctx context.Context, destinationChainType multichainTypes.ChainType, destinationChainId string, sourceAddress string) (*outboundTypes.QueryLastOutboundBatchNonceResponse, error) {
	req := &outboundTypes.QueryLastOutboundBatchNonceRequest{
		DestinationChainType: uint64(destinationChainType),
		DestinationChainId:   destinationChainId,
		SourceAddress:        sourceAddress,
	}
	return c.outboundQueryClient.LastOutboundBatchNonce(ctx, req)
}

/////////////////////////////////
////     Crosstalk           ////
////////////////////////////////
func (c *chainClient) GetAllCrossTalkRequest(ctx context.Context, pagination *query.PageRequest) (*crosstalkTypes.QueryAllCrossTalkRequestResponse, error) {
	req := &crosstalkTypes.QueryAllCrossTalkRequest{Pagination: pagination}
	return c.crosstalkQueryClient.CrossTalkRequestAll(ctx, req)
}

func (c *chainClient) GetAllCrossTalkRequestByStatus(ctx context.Context, status crosstalkTypes.CrossTalkRequestStatus, pagination *query.PageRequest) (*crosstalkTypes.QueryCrosstalkRequestByStatusResponse, error) {
	req := &crosstalkTypes.QueryAllCrosstalkRequestByStatusRequest{
		Status:     uint64(status),
		Pagination: pagination,
	}
	return c.crosstalkQueryClient.CrosstalkRequestByStatus(ctx, req)
}

func (c *chainClient) GetCrossTalkRequest(ctx context.Context, chainType uint64, chainID string, eventNonce uint64) (*crosstalkTypes.QueryGetCrossTalkRequestResponse, error) {
	req := &crosstalkTypes.QueryGetCrossTalkRequest{ChainType: chainType, ChainId: chainID, EventNonce: eventNonce}
	return c.crosstalkQueryClient.CrossTalkRequest(ctx, req)
}

func (c *chainClient) GetAllCrosstalkRequestConfirmations(ctx context.Context, pagination *query.PageRequest, sourceChainType uint64, sourceChainId string, eventNonce uint64, claimHash []byte) (*crosstalkTypes.QueryAllCrosstalkRequestConfirmResponse, error) {
	req := &crosstalkTypes.QueryAllCrosstalkRequestConfirmRequest{
		SourceChainType: sourceChainType,
		SourceChainId:   sourceChainId,
		EventNonce:      eventNonce,
		ClaimHash:       claimHash,
		Pagination:      pagination,
	}
	return c.crosstalkQueryClient.CrosstalkRequestConfirmAll(ctx, req)
}

func (c *chainClient) GetAllCrossTalkAckRequest(ctx context.Context, pagination *query.PageRequest) (*crosstalkTypes.QueryAllCrossTalkAckRequestResponse, error) {
	req := &crosstalkTypes.QueryAllCrossTalkAckRequest{Pagination: pagination}
	return c.crosstalkQueryClient.CrossTalkAckRequestAll(ctx, req)
}

func (c *chainClient) GetAllCrossTalkAckRequestByStatus(ctx context.Context, status crosstalkTypes.CrossTalkAckRequestStatus, pagination *query.PageRequest) (*crosstalkTypes.QueryAllCrosstalkAckRequestsByStatusResponse, error) {
	req := &crosstalkTypes.QueryAllCrosstalkAckRequestsByStatusRequest{
		Status:     uint64(status),
		Pagination: pagination,
	}
	return c.crosstalkQueryClient.AllCrosstalkAckRequestsByStatus(ctx, req)
}

func (c *chainClient) GetCrossTalkAckRequest(ctx context.Context, chainType uint64, chainID string, eventNonce uint64) (*crosstalkTypes.QueryGetCrossTalkAckRequestResponse, error) {
	req := &crosstalkTypes.QueryGetCrossTalkAckRequest{ChainType: chainType, ChainId: chainID, EventNonce: eventNonce}
	return c.crosstalkQueryClient.CrossTalkAckRequest(ctx, req)
}

func (c *chainClient) GetAllCrosstalkAckRequestConfirmations(ctx context.Context, pagination *query.PageRequest, chainType uint64, chainId string, eventNonce uint64, claimHash []byte) (*crosstalkTypes.QueryAllCrosstalkAckRequestConfirmResponse, error) {
	req := &crosstalkTypes.QueryAllCrosstalkAckRequestConfirmRequest{
		ChainType:  chainType,
		ChainId:    chainId,
		EventNonce: eventNonce,
		ClaimHash:  claimHash,
		Pagination: pagination,
	}
	return c.crosstalkQueryClient.CrosstalkAckRequestConfirmAll(ctx, req)
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

//AsyncBroadcastMsg sends Tx to chain and doesn't wait until Tx is included in block. This method
//cannot be used for rapid Tx sending, it is expected that you wait for transaction status with
//external tools. If you want sdk to wait for it, use SyncBroadcastMsg.
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
	err = tx.Sign(txf, clientCtx.GetFromName(), txn, true)
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

//QueueBroadcastMsg enqueues a list of messages. Messages will added to the queue
//and grouped into Txns in chunks. Use this method to mass broadcast Txns with efficiency.
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
