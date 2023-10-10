package mpc

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/sirupsen/logrus"

	"github.com/qredo/fusionchain/mpc-relayer/pkg/common"
)

var _ Client = (*localMPC)(nil)

type localMPC struct {
	_logger     *logrus.Entry
	initVersion int
}

func newLocalClient(logger *logrus.Entry, initVersion int) *localMPC {
	return &localMPC{
		_logger:     logger,
		initVersion: initVersion,
	}
}

func (m *localMPC) logger(assetID []byte, traceID string) *logrus.Entry {
	return m._logger.WithFields(logrus.Fields{
		"dd":      TraceID{Trace: traceID},
		"keyID":   fmt.Sprintf("%x", assetID),
		"traceID": traceID,
	})
}

func (m *localMPC) PublicKey(keyID []byte, keyType CryptoSystem) ([]byte, string, error) {
	req := &KeysRequest{
		KeyID: hex.EncodeToString(keyID),
	}

	// Trace ID
	b, err := common.RandomBytes(16)
	if err != nil {
		return nil, "", err
	}
	traceID := fmt.Sprintf("%16x", b)

	m.logger(keyID, traceID).Info("mpcPubKey")

	//Fake a public Key from the MPC
	response, err := localMPCKeys(req, m.initVersion, keyType)
	if err != nil {
		return nil, traceID, err
	}

	// Check matching KeyID
	if req.KeyID != response.KeyID {
		return nil, "", fmt.Errorf("mpc keyID mismatch expected %v, got %v", req.KeyID, response.KeyID)
	}
	// Use decompressed PK if ETH
	pubKeyBytes, err := hex.DecodeString(response.Pk)
	if err != nil {
		return nil, "", ErrNoPubKey
	}
	return pubKeyBytes, traceID, nil
}

func (m *localMPC) PubkeySignature(pubKey, seedID []byte, keyType CryptoSystem) ([]byte, string, error) {

	h := sha256.Sum256(pubKey)
	dataToSign := h[:]

	req := &SigRequest{
		KeyID:   hex.EncodeToString(seedID),
		Message: hex.EncodeToString(dataToSign),
		IsKey:   isKey,
	}

	// Trace ID
	b, err := common.RandomBytes(16)
	if err != nil {
		return nil, "", err
	}
	traceID := fmt.Sprintf("%16x", b)

	m.logger(seedID, traceID).Info("mpcPubKeySign")
	//do the post to the MPC server
	response, err := localMPCSign(req, m.initVersion, keyType)
	if err != nil {
		return nil, traceID, err
	}

	// Check matching KeyID
	if req.KeyID != response.KeyID {
		return nil, "", fmt.Errorf("mpc keyID mismatch expected %v, got %v", req.KeyID, response.KeyID)
	}
	// verify Signature against message
	rsSerialised, pubKey, valid, err := validateResponse(response, keyType)
	if err != nil || !valid {
		return nil, "", ErrInvalidSignature
	}
	if !bytes.Equal(pubKey, pubKey) {
		return nil, "", ErrNoPubKey
	}

	return rsSerialised, traceID, err
}

func (m *localMPC) Signature(sigRequestData *SigRequestData, keyType CryptoSystem) (*SigResponse, string, error) {
	keyID := hex.EncodeToString(sigRequestData.KeyID)
	req := &SigRequest{
		KeyID:   keyID,
		Message: hex.EncodeToString(sigRequestData.SigHash),
		IsKey:   isNotKey,
	}

	// Trace ID
	b, err := common.RandomBytes(16)
	if err != nil {
		return nil, "", err
	}
	traceID := fmt.Sprintf("%16x", b)

	m.logger(sigRequestData.KeyID, traceID).Info("mpcSign")

	//do the post to the MPC server
	response, err := localMPCSign(req, m.initVersion, keyType)
	if err != nil {
		return nil, traceID, err
	}

	// Check matching KeyID
	if req.KeyID != response.KeyID {
		return nil, "", fmt.Errorf("mpc keyID mismatch expected %v, got %v", req.KeyID, response.KeyID)
	}
	// verify Signature against message
	_, _, valid, err := validateResponse(response, keyType)
	if err != nil {
		return nil, "", ErrInvalidSignature
	}
	if !valid {
		return nil, "", ErrInvalidSignature
	}
	return response, traceID, err
}

func (m *localMPC) CheckMPC() (bool, string) {
	b, err := common.RandomBytes(16)
	if err != nil {
		return false, ""
	}
	return true, fmt.Sprintf("%16x", b)
}

// localMPCKeys - emulate the MPCKeys request
func localMPCKeys(req *KeysRequest, salt int, keyType CryptoSystem) (resp *KeysResponse, err error) {
	keyID, err := hex.DecodeString(req.KeyID)
	if err != nil {
		return nil, err
	}
	// create seed directly from keyID
	seed := sha256.Sum256(append(keyID, byte(salt)))

	pubKeyBytes, err := generateKey(seed[:], keyType)
	if err != nil {
		return nil, fmt.Errorf("could not generate key, err=%v", err)
	}

	resp = &KeysResponse{
		Service: "mpcclientparent",
		Message: "OK",
		Version: "0.0.1",
		KeyID:   req.KeyID,
		Pk:      hex.EncodeToString(pubKeyBytes),
	}
	if keyType == EdDSA {
		resp.EdPk = resp.Pk
		resp.Pk = ""
	}
	return resp, nil
}

// localMPCSign - emulate the MPC sign request
func localMPCSign(req *SigRequest, salt int, keyType CryptoSystem) (resp *SigResponse, err error) {
	//AssetID - used for validation only
	keyID, err := hex.DecodeString(req.KeyID)
	if err != nil {
		return nil, err
	}
	// Handle engine key differently
	seed := sha256.Sum256(append(keyID, byte(salt)))

	m, err := hex.DecodeString(req.Message)
	if err != nil {
		return nil, err
	}

	sigBytes, pubKeyBytes, err := generateSignature(seed[:], m, keyType)
	if err != nil {
		return nil, err
	}
	if sigBytes == nil {
		return nil, fmt.Errorf("signature was nil: %v", sigBytes)
	}
	if len(sigBytes) < 64 {
		return nil, fmt.Errorf("invalid signature length: %v", len(sigBytes))
	}
	sigR := new(big.Int).SetBytes(sigBytes[0:32])
	sigS := new(big.Int).SetBytes(sigBytes[32:64])
	resp = &SigResponse{
		Service: "mpc cltctl (mock)",
		Message: "ok",
		Version: "1.0.0",
		KeyID:   req.KeyID,
	}
	switch keyType {
	case EcDSA:
		resp.EdMessage = req.Message
		resp.EdR = toHexInt(sigR)
		resp.EdS = toHexInt(sigS)
		resp.EdPk = hex.EncodeToString(pubKeyBytes)
	case EdDSA:
		resp.EcMessage = req.Message
		resp.EcR = toHexInt(sigR)
		resp.EcS = toHexInt(sigS)
		resp.Pk = hex.EncodeToString(pubKeyBytes)
	default:
		return nil, fmt.Errorf("key type '%v' not supported", keyType)
	}

	return resp, err
}

func toHexInt(n *big.Int) string {
	b := math.PaddedBigBytes(n, 32)
	return fmt.Sprintf("%x", b) // or %X or upper case
}

func padZerosRight(msg string, n int) string {
	for len(msg) < n {
		msg += "0"
	}
	return msg
}
