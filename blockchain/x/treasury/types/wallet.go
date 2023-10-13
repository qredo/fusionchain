package types

import (
	"fmt"
	"math/big"
)

type Wallet interface {
	// Address returns a human readable version of the address.
	Address() string
}

func NewWallet(k *Key, w WalletType) (Wallet, error) {
	switch w {
	case WalletType_WALLET_TYPE_QRDO:
		return NewFusionWallet(k)
	case WalletType_WALLET_TYPE_ETH:
		return NewEthereumWallet(k)
	case WalletType_WALLET_TYPE_ETH_SEPOLIA:
		return NewEthereumWallet(k)
	}
	return nil, fmt.Errorf("error in NewWallet: unknown wallet type")
}

// Transfer represents a generic transfer of tokens on a layer 1 blockchain.
// Ideally, this will be the object passed to Blackbird for applying policy.
type Transfer struct {
	// To uniquely identifies the recipient of the transfer.
	To []byte

	// Amount is the amount being transferred.
	Amount *big.Int

	// CoinIdentifier uniquely identifies the coin being transferred.
	CoinIdentifier []byte

	// DataForSigning is the data that will be signed by the key.
	DataForSigning []byte
}

// TxParser can be implemented by wallets that are able to parse unsigned
// transactions into the common Layer1Tx format.
//
// By doing that, wallets can expose more functionalities (i.e. Blackbird
// policies).
type TxParser interface {
	ParseTx(b []byte) (Transfer, error)
}
