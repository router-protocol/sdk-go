all:

copy-chain-types:
	cp ../sdk-go/routerchain/multichain/types/*.go routerchain/multichain/types
	rm -rf routerchain/multichain/types/*test.go  rm -rf routerchain/multichain/types/*gw.go
	
	cp ../sdk-go/routerchain/attestation/types/*.go routerchain/attestation/types
	rm -rf routerchain/attestation/types/*test.go  rm -rf routerchain/attestation/types/*gw.go
	
	cp ../sdk-go/routerchain/inbound/types/*.go routerchain/inbound/types
	rm -rf routerchain/inbound/types/*test.go  rm -rf routerchain/inbound/types/*gw.go

	cp ../sdk-go/routerchain/crosstalk/types/*.go routerchain/crosstalk/types
	rm -rf routerchain/crosstalk/types/*test.go  rm -rf routerchain/crosstalk/types/*gw.go

	cp ../sdk-go/routerchain/outbound/types/*.go routerchain/outbound/types
	rm -rf routerchain/outbound/types/*test.go  rm -rf routerchain/outbound/types/*gw.go

	cp ../sdk-go/routerchain/oracle/types/*.go routerchain/oracle/types
	rm -rf routerchain/oracle/types/*test.go  rm -rf routerchain/oracle/types/*gw.go

	cp -r ../sdk-go/routerchain/types/*.go routerchain/types
	rm -rf routerchain/types/*test.go

	cp ../sdk-go/routerchain/util/*.go routerchain/util
	cp -r ../router-chain/proto/* routerchain/proto/

	echo "👉 Replace sdk-go/routerchain/util with sdk-go/routerchain/util"
	echo "👉 Replace sdk-go/routerchain with sdk-go/routerchain"
	echo "👉 Replace sdk-go/routerchain/types with sdk-go/routerchain/types"