/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
*/

package service

/*
import (
	"fmt"
	"math/big"

	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/sync/errgroup"

	"github.com/qredo/assets/libs/assets"
	"github.com/qredo/assets/libs/layerone"
	"github.com/qredo/assets/libs/layerone/currencyutils"
	"github.com/qredo/assets/libs/layerone/ethereum/contract"
	"github.com/qredo/assets/libs/layerone/ethereum/helpers"
	"github.com/qredo/assets/libs/protobuffer"
	"github.com/qredo/assets/libs/watcher/mpc"
	"github.com/qredo/assets/libs/watcher/qredochain"
)

type Processor interface {
	GenerateKey() (keyResult, error)
	GenerateSignature() (signResult, error)
}

type QredoProcessor struct {
	qredoWatcher qredochain.KeySearcher
	mpcClient    mpc.Client
	writer       qredochain.Writer
}

func NewProcessor(qredoWatcher qredochain.KeySearcher, mpcClient mpc.Client, writer qredochain.Writer) *QredoProcessor {
	return &QredoProcessor{
		qredoWatcher: qredoWatcher,
		mpcClient:    mpcClient,
		writer:       writer,
	}
}

func (q *QredoProcessor) getSettlementTX(assetID []byte) (*protobuffer.PBSettlementTransaction, error) {
	data, err := q.qredoWatcher.KeySearch(assets.DB_SETTLEMENTUTXO_PREFIX_STR, string(assetID))
	if err != nil {
		return nil, fmt.Errorf("search failed key '%s': %w", string(assetID), err)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("cannot find settlement tx for wallet ID %x", assetID)
	}
	settlementTX := protobuffer.PBSettlementTransaction{}
	if err := assets.Unmarshal(data, &settlementTX); err != nil {
		return nil, fmt.Errorf("error unmarshaling: %v", err)
	}
	return &settlementTX, nil
}

// buildSignatureList - Routine for obtaining signatures from the MPC interface
func (q *QredoProcessor) buildSignatureList(assetID []byte, settlementTX *protobuffer.PBSettlementTransaction) ([]*mpc.Response, error) {
	var mpcSignatureList []*mpc.Response
	for index, txInput := range settlementTX.Input {
		mpcSigResult, _, err := q.mpcClient.ExternalInputSignature(index, txInput, assetID, settlementTX.Currency)
		if err != nil {
			return nil, err
		}
		mpcSignatureList = append(mpcSignatureList, mpcSigResult)
	}
	return mpcSignatureList, nil
}

func processSigResultBytes(settlementTX *protobuffer.PBSettlementTransaction, mpcResponse []*mpc.Response) ([]externalTxMsg, error) {
	if settlementTX.Input == nil {
		return nil, fmt.Errorf("no settlement data provided")
	}
	switch {
	case layerone.IsEVM(settlementTX.Currency):
		return extractEVMResultBytes(settlementTX, mpcResponse)
	case layerone.IsEdDSABased(settlementTX.Currency):
		return extractResultBytes(settlementTX, mpcResponse, mpc.ExtractEdDSASignedMessage)
	case layerone.IsStacks(settlementTX.Currency):
		return extractResultBytes(settlementTX, mpcResponse, mpc.ExtractSigWithRecoveryID)
	case layerone.IsCosmos(settlementTX.Currency):
		return extractResultBytes(settlementTX, mpcResponse, mpc.ExtractECDSASignedMessageWithoutRecID)
	default:
		return nil, fmt.Errorf("invalid connected wallet currency: %v", settlementTX.Currency.String())
	}
}

func extractEVMResultBytes(settlementTX *protobuffer.PBSettlementTransaction, mpcResponse []*mpc.Response) ([]externalTxMsg, error) {
	var signedBytes []externalTxMsg
	for i, input := range settlementTX.Input {
		signedMessage, err := mpc.ExtractECDSASignedMessage(mpcResponse[i])
		if err != nil {
			return nil, err
		}
		switch input.HashType {
		case assets.HashTypeMessage:
			// if the preimage for the hash being signed is an arbitrary message then we only need to return the signature
			signedBytes = append(signedBytes, externalTxMsg{
				rawSigPayload: signedMessage.Signature,
				sender:        fmt.Sprintf("%x", signedMessage.PublicKey),
			})

		case assets.HashTypeTransaction, assets.HashTypeTransactionFee:
			// if the preimage for the hash then we must rebuild the RLP transaction
			tx := ethTypes.Transaction{}
			b := settlementTX.MultiUTXO[i]

			if err := decodeETHRLPTx(&tx, b); err != nil {
				return nil, err
			}

			chainID := new(big.Int).SetBytes(input.Chain)
			if chainID.Int64() == 0 {
				// chainID not supplied by the PBSettlementTransaction, but instead inferred from the wallet currency
				chainID = helpers.ChainIDFor(settlementTX.Currency)
			}

			signer := ethTypes.LatestSignerForChainID(chainID)
			signedTx, err := tx.WithSignature(signer, signedMessage.Signature)
			if err != nil {
				return nil, err
			}

			txBytes, err := rlp.EncodeToBytes(signedTx)
			if err != nil {
				return nil, err
			}
			fromAddr, err := signer.Sender(signedTx)
			if err != nil {
				return nil, err
			}
			signedBytes = append(signedBytes, externalTxMsg{
				chainID:       chainID.Int64(),
				rawSigPayload: txBytes,
				sender:        fromAddr.Hex(),
				txid:          signedTx.Hash().Hex(),
			})

		case assets.HashTypeInvoke:
			// if the preimage for the hash being signed is an arbitrary message then we only need to return the signature
			tx := ethTypes.Transaction{}
			b := settlementTX.MultiUTXO[i]
			if err := decodeETHRLPTx(&tx, b); err != nil {
				return nil, err
			}
			c := settlementTX.Currency
			chainID := new(big.Int).SetBytes(input.Chain)
			if chainID.Int64() == 0 {
				// chainID not supplied by the PBSettlementTransaction, but instead inferred from the wallet currency
				chainID = helpers.ChainIDFor(c)
			}
			to := tx.To()
			if to == nil {
				return nil, fmt.Errorf("tx recipient is nil")
			}
			d, err := contract.InvokePayloadWithSig(signedMessage.Signature, []byte{0x0}, *to, tx.Value())
			if err != nil {
				return nil, err
			}
			address, err := currencyutils.GetAddressBuilder(c).BuildAddress(signedMessage.PublicKey, false)
			if err != nil {
				return nil, err
			}
			signedBytes = append(signedBytes, externalTxMsg{
				chainID:       chainID.Int64(),
				rawSigPayload: d,
				sender:        address,
				txid:          tx.Hash().Hex(), // does not uniquely identify the transaction
			})
		}
	}
	return signedBytes, nil
}

func extractResultBytes(settlementTX *protobuffer.PBSettlementTransaction, mpcResponse []*mpc.Response, mpcExtract func(mpcSigResult *mpc.Response) (*mpc.SignedMessage, error)) ([]externalTxMsg, error) {
	var signatures []externalTxMsg
	for i, input := range settlementTX.Input {
		signedMessage, err := mpcExtract(mpcResponse[i])
		if err != nil {
			return nil, err
		}

		switch input.HashType {
		case assets.HashTypeTransaction, assets.HashTypeMessage:
			signatures = append(signatures, externalTxMsg{
				rawSigPayload: signedMessage.Signature,
				sender:        fmt.Sprintf("%x", signedMessage.PublicKey),
			})
		default:
			return nil, fmt.Errorf("incorrect EdDSA sig request format")
		}
	}
	return signatures, nil
}

type signResult struct {
	currency     protobuffer.PBCryptoCurrency
	transactions []externalTxMsg
}

type externalTxMsg struct {
	chainID       int64
	txid          string
	sender        string
	rawSigPayload []byte // either completed/serialized tx or raw signature
}

// SignOnlyProcess processes an External signature request and return either a signed transaction or raw signature to the Qredochain
func (q *QredoProcessor) SignOnlyProcess(walletID []byte, mutableIndex int64, description string) (signResult, error) {
	settlementTX, err := q.getSettlementTX(walletID)
	if err != nil {
		return signResult{}, err
	}

	// Check external wallet signing request is for the supported network
	if !layerone.SupportsExternal(settlementTX.Currency) {
		return signResult{}, fmt.Errorf("external wallet: invalid currency %v", settlementTX.Currency.String())
	}

	mpcSignatureList, err := q.buildSignatureList(walletID, settlementTX)
	if err != nil {
		return signResult{}, fmt.Errorf("mpc sign failure: %v", err)
	}

	transactions, err := processSigResultBytes(settlementTX, mpcSignatureList)
	if err != nil {
		return signResult{}, fmt.Errorf("process signature error: %v", err)
	}

	group := errgroup.Group{}
	for index, t := range transactions {
		index, tx := index, t.rawSigPayload
		group.Go(func() error {
			if err := q.writer.MPCSign(qredochain.MPCSign{
				WalletID:          walletID,
				WalletIndex:       mutableIndex,
				WalletDescription: description,
				Transaction:       tx,
				Failed:            false,
				Currency:          settlementTX.Currency,
				Index:             int64(index),
				MPCType:           protobuffer.PBMPCType_Signature_Only,
			}); err != nil {
				return fmt.Errorf("currency %s signature request, send mpcsign fail, error: %v", settlementTX.Currency, err)
			}
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		return signResult{}, err
	}
	return signResult{
		currency:     settlementTX.Currency,
		transactions: transactions,
	}, nil
}

// decodeETHRLPTx handles a backwards compatibility issue allowing both EIP2718 typed transaction envelope and pre-EIP2718
// RLP encoded transactions to be successfully decoded
func decodeETHRLPTx(tx *ethTypes.Transaction, b []byte) error {
	if err := tx.UnmarshalBinary(b); err != nil {
		if err2 := rlp.DecodeBytes(b, &tx); err2 != nil {
			return fmt.Errorf("failed to unmarshal binary: err=%v, err2=%v", err, err2)
		}
	}
	return nil
}
