package types

import (
	"context"

	gogogrpc "github.com/cosmos/gogoproto/grpc" // nolint: staticcheck
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	// nolint: staticcheck
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// baseAppSimulateFn is the signature of the Baseapp#Simulate function.
type baseAppSimulateFn func(txBytes []byte) (sdk.GasInfo, *sdk.Result, error)

// txServer is the server for the protobuf Tx service.
type routerTxServer struct {
	clientCtx         client.Context
	simulate          baseAppSimulateFn
	interfaceRegistry codectypes.InterfaceRegistry
}

// NewRouterTxServer creates a new Router Tx service server.
func NewRouterTxServer(clientCtx client.Context, simulate baseAppSimulateFn, interfaceRegistry codectypes.InterfaceRegistry) RouterTxRpcServer {
	return routerTxServer{
		clientCtx:         clientCtx,
		simulate:          simulate,
		interfaceRegistry: interfaceRegistry,
	}
}

var _ RouterTxRpcServer = routerTxServer{}

// TxsByEvents implements the ServiceServer.TxsByEvents RPC method.
func (s routerTxServer) GetTx(ctx context.Context, req *GetTxRequest) (*GetTxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	return &GetTxResponse{}, nil
}

func (s routerTxServer) PrepareTx(ctx context.Context, req *PrepareTxRequest) (*PrepareTxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	return &PrepareTxResponse{
		Data: "{venky:sdf}",
	}, nil
}

func (s routerTxServer) BroadcastTx(ctx context.Context, req *BroadcastTxRequest) (*BroadcastTxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	return &BroadcastTxResponse{}, nil
}

// RegisterRouterTxService registers the router tx service on the gRPC router.
func RegisterRouterTxService(
	qrt gogogrpc.Server,
	clientCtx client.Context,
	simulateFn baseAppSimulateFn,
	interfaceRegistry codectypes.InterfaceRegistry,
) {
	RegisterRouterTxRpcServer(
		qrt,
		NewRouterTxServer(clientCtx, simulateFn, interfaceRegistry),
	)
}

// RegisterGRPCGatewayRoutes mounts the tx service's GRPC-gateway routes on the
// given Mux.
func RegisterGRPCGatewayRoutes(clientConn gogogrpc.ClientConn, mux *runtime.ServeMux) {
	_ = RegisterRouterTxRpcHandlerClient(context.Background(), mux, NewRouterTxRpcClient(clientConn))
}
