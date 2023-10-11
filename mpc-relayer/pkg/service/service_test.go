package service

/*

import (
	"bytes"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestService_Start(t *testing.T) {
	tests := []struct {
		name        string
		walletIDs   []string
		failAttempt map[string]int64
		success     int64
	}{
		{
			name:      "success",
			walletIDs: []string{"1", "2"},
			failAttempt: map[string]int64{
				"1": 0,
				"2": 0,
			},
			success: 2,
		},
		{
			name:      "with retry",
			walletIDs: []string{"1", "2"},
			failAttempt: map[string]int64{
				"1": 0,
				"2": 1,
			},
			success: 2,
		},
		{
			name:      "with discard",
			walletIDs: []string{"1", "2"},
			failAttempt: map[string]int64{
				"1": 0,
				"2": defaultMaxTries + 1,
			},
			success: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := newFakeProcessor(tt.failAttempt)
			s := &Service{
				processor: processor,
				catchUpFunc: func(start int64) ([]*qredochain.QueueItem, error) {
					return nil, nil
				},
				retriever: fakeRetriever{
					walletIDs: tt.walletIDs,
					out:       make(chan *qredochain.QueueItem),
				},
				db:         database.NewIndex(store.NewMemory()),
				waiting:    make(chan *retryableQueueItem),
				stop:       make(chan struct{}),
				retrySleep: 1 * time.Millisecond,
				maxTries:   defaultMaxTries,
			}
			if err := s.Start(); err != nil {
				t.Fatal(err)
			}
			time.Sleep(time.Millisecond) // sleep to ensure retriever send item
			if got, want := processor.success.Load(), tt.success; got != want {
				t.Fatalf("got %v want %v", got, want)
			}
			s.Stop()
		})
	}
}

func TestService_Stop(t *testing.T) {
	tests := []struct {
		name        string
		walletIDs   []string
		failAttempt map[string]int64
	}{
		{
			name:      "stop with waiting tx",
			walletIDs: []string{"1"},
			failAttempt: map[string]int64{
				"1": time.Minute.Milliseconds(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := newFakeProcessor(tt.failAttempt)
			s := &Service{
				processor: processor,
				catchUpFunc: func(start int64) ([]*qredochain.QueueItem, error) {
					return nil, nil
				},
				retriever: fakeRetriever{
					walletIDs: tt.walletIDs,
					out:       make(chan *qredochain.QueueItem),
				},
				waiting:    make(chan *retryableQueueItem),
				db:         database.NewIndex(store.NewMemory()),
				stop:       make(chan struct{}),
				retrySleep: time.Millisecond,
				maxTries:   time.Minute.Milliseconds(),
				indexStart: 0,
			}
			if err := s.Start(); err != nil {
				t.Fatal(err)
			}
			time.Sleep(time.Millisecond) // sleep to ensure retriever send item
			s.Stop()
			time.Sleep(10 * time.Millisecond) // sleep to ensure all retry process exit
		})
	}
}

type fakeProcessor struct {
	success   atomic.Int64
	returnErr map[string]*atomic.Int64
}

func newFakeProcessor(returnErr map[string]int64) *fakeProcessor {
	returnErrAtomic := make(map[string]*atomic.Int64)
	for k, v := range returnErr {
		value := &atomic.Int64{}
		value.Store(v)
		returnErrAtomic[k] = value
	}
	return &fakeProcessor{
		returnErr: returnErrAtomic,
	}
}

func (f *fakeProcessor) SignOnlyProcess(walletID []byte, mutableIndex int64, description string) (signResult, error) {
	if f.returnErr[string(walletID)].Load() != 0 {
		f.returnErr[string(walletID)].Add(-1)
		return signResult{}, fmt.Errorf("test error")
	}
	f.success.Add(1)
	return signResult{}, nil
}

type fakeRetriever struct {
	walletIDs []string
	out       chan *qredochain.QueueItem
}

func (f fakeRetriever) Start(int64) (chan *qredochain.QueueItem, error) {
	go func() {
		for _, id := range f.walletIDs {
			f.out <- &qredochain.QueueItem{
				TxAsset: &assets.Wallet{
					SignedAsset: assets.SignedAsset{
						CurrentAsset: &protobuffer.PBSignedAsset{
							Asset: &protobuffer.PBAsset{
								ID: []byte(id),
							},
						},
					},
				},
				Index: 0,
			}
		}
	}()
	return f.out, nil
}

func (f fakeRetriever) Stop() {
}

func Test_TXDecoder(t *testing.T) {
	test := []struct {
		name     string
		txBytes  []byte
		methodID []byte
	}{
		{
			"MMI Example 1",
			hexutil.MustDecode("0xb9013302f9012f82a4b1268084080befc08317401e947d5ba536ab244aaa1ea42ab88428847f25e3e67680b90104c7ca587000000000000000000000000000000000000000000000000000000000000033b400000000000000000000000000000000000000000000000000a9afdc1df1dbb8000000000000000000000000000000000000000000000002c282da34e96ef8da0000000000000000000000000000000000000000000000000000000002a687120000000000000000000000000000000000000000000000000000000064fae732000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001c0808080"),
			hexutil.MustDecode("0xc7ca5870"),
		},
		{
			"MMI Example 2",
			hexutil.MustDecode("0xb87502f87282e70808849502f900849502f90e83030853940b3a25ae91de4825b52d51ca54dfc8867367c72a80b844441a3e7000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007907373e23ad7380c0808080"),
			hexutil.MustDecode("0x441a3e70"),
		},
		{
			"MMI Example 3",
			hexutil.MustDecode("0xb87502f87282e70809849502f900849502f90e830110e6940b15a5e3ca0d4b492c3b476d0f807535f9b7207980b844095ea7b3000000000000000000000000bf05db69176e47bf89a6b19f7492d50751d20452ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808080"),
			hexutil.MustDecode("0x095ea7b3"),
		},
	}

	for _, tt := range test {
		tx := &types.Transaction{}
		t.Run(tt.name, func(t *testing.T) {
			if err := decodeETHRLPTx(tx, tt.txBytes); err != nil {
				t.Fatal(err)
			}
			if g, w := tx.Data()[0:4], tt.methodID; !bytes.Equal(g, w) {
				t.Fatalf("incorrect methodID expected %x, got %x", w, g)
			}
		})
	}

}

*/
