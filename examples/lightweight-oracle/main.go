package main

import (
	chainclient "github.com/router-protocol/sdk-go/client/routerchain"
	"github.com/router-protocol/sdk-go/examples/lightweight-oracle/oracle"
)

const (
	// Tx Agrs
	NETWORK_NAME             = "local"
	ORACLE_PROVIDER_KEY_NAME = "siva"
	PASSPHRASE               = "12345678"
)

func main() {

	// INITIALIZE ORACLE PROVIDER
	routerchainClient := chainclient.InitialiseChainClient(NETWORK_NAME, ORACLE_PROVIDER_KEY_NAME, PASSPHRASE, "")
	oracleProvider := oracle.NewOracleProvider(routerchainClient)

	// TRIGGER ACTION
	oracleProvider.SendGasPrices()

}
