all:

copy-chain-types:
	cp ../router-chain/x/multichain/types/*.go routerchain/multichain/types
	rm -rf routerchain/multichain/types/*test.go  rm -rf routerchain/multichain/types/*gw.go
	cp ../router-chain/x/attestation/types/*.go routerchain/attestation/types
	rm -rf routerchain/attestation/types/*test.go  rm -rf routerchain/attestation/types/*gw.go
	cp ../router-chain/x/inbound/types/*.go routerchain/inbound/types
	rm -rf routerchain/inbound/types/*test.go rm -rf  rm -rf routerchain/inbound/types/*gw.go
	cp ../router-chain/x/inbound/types/*.go routerchain/inbound/types
	cp ../router-chain/util/*.go routerchain/util

	echo "👉 Replace router-chain/util with sdk-go/routerchain/util"
	echo "👉 Replace sdk-go/routerchain with sdk-go/routerchain
