package oracle

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	routerclient "github.com/router-protocol/sdk-go/client/routerchain"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

type OracleProvider struct {
	clientContext     client.Context
	routerChainClient routerclient.ChainClient
}

func NewOracleProvider(routerChainClient routerclient.ChainClient) *OracleProvider {
	return &OracleProvider{
		routerChainClient: routerChainClient,
	}
// }

// func (oracleProvider *OracleProvider) SendGasPrices() {
// 	// Initialize
// 	gasPrice := pricefeedTypes.GasPrice{
// 		ChainType: multichainTypes.CHAIN_TYPE_EVM,
// 		ChainId:   "80001",
// 		GasPrice:  uint64(4000000),
// 		Decimals:  uint64(18),
// 	}
// 	gasPrices := []oracleTypes.GasPriceState{gasPrice}

// 	// Create and submit Tx
// 	msg := oracleTypes.NewMsgGasPrices(oracleProvider.routerChainClient.FromAddress().String(), gasPrices)
// 	err := oracleProvider.routerChainClient.QueueBroadcastMsg(msg)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	time.Sleep(time.Second * 5)
// }
