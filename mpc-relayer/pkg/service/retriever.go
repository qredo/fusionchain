package service

/*
import (
	"github.com/tendermint/tendermint/rpc/client"

	"github.com/qredo/assets/libs/assets"
	"github.com/qredo/assets/libs/protobuffer"
	"github.com/qredo/assets/libs/watcher/qredochain"
)

type Retriever interface {
	Start(start int64) (chan *QueueItem, error)
	Stop()
}

type QueueItem struct {
	Tx    []byte
	Index int64
}

type FusionRetriever struct {
	sub       *qredochain.Subscription
	connector client.EventsClient
	out       chan *qredochain.QueueItem
}

var filterSignature = func(item *qredochain.QueueItem) (bool, error) {
	switch item.TxAsset.GetType() {
	case protobuffer.PBAssetType_Wallet:
		wallet := item.TxAsset.(*assets.Wallet)
		if wallet.CurrentAsset.Asset.TransferRuleType != protobuffer.PBTransferRuleType_SignatureRequest {
			return true, nil
		}
		walletPayload, err := wallet.Payload()
		if err != nil {
			return true, err
		}
		if walletPayload.WalletType != protobuffer.PBWalletType_ExternalWallet {
			return true, nil
		}
	case protobuffer.PBAssetType_Control:
		control := item.TxAsset.(*assets.Control)
		ctlPayload, err := control.Payload()
		if err != nil {
			return true, err
		}
		if ctlPayload.Type != protobuffer.PBControlType_Rescue {
			return true, nil
		}
	default:
		return true, nil
	}
	return false, nil
}

func NewRetriever(connector client.EventsClient) *QredochainRetriever {
	return &QredochainRetriever{
		connector: connector,
		out:       make(chan *qredochain.QueueItem),
	}
}

func (q *QredochainRetriever) Start(start int64) (chan *qredochain.QueueItem, error) {
	sub, err := qredochain.NewSubscription(q.connector, "signer", q.out, "watcher", qredochain.FilterIndex(start), filterSignature)
	if err != nil {
		return nil, err
	}
	q.sub = sub
	return q.out, nil
}

func (q *QredochainRetriever) Stop() {
	q.sub.Close()
}
