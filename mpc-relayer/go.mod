module github.com/qredo/fusionchain/mpc-relayer

go 1.21

replace github.com/qredo/fusionchain/go-client => ../go-client

// required by go-client
replace github.com/qredo/fusionchain => ../blockchain

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0
	github.com/dgraph-io/badger/v4 v4.2.0
	github.com/ethereum/go-ethereum v1.13.2
	github.com/gorilla/mux v1.8.0
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20190702054246-869f871628b6 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.5-0.20220116011046-fa5810519dcb // indirect
	github.com/google/flatbuffers v1.12.1 // indirect
	github.com/holiman/uint256 v1.2.3 // indirect
	github.com/klauspost/compress v1.15.15 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.opencensus.io v0.22.5 // indirect
	golang.org/x/crypto v0.12.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
