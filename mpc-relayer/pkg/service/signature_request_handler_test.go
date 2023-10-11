package service

import (
	"testing"

	"github.com/qredo/fusionchain/mpc-relayer/pkg/database"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/logger"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/mpc"
	"github.com/qredo/fusionchain/x/treasury/types"
)

func Test_NewSigController(t *testing.T) {
	// TODO
}

func Test_ExecuteSigQuery(t *testing.T) {
	tests := []struct {
		name      string
		item      signatureRequestQueueItem
		expectErr bool
	}{
		{
			"simple",
			signatureRequestQueueItem{
				maxTries: 5,
				request:  &types.SignRequest{Id: 1},
			},
			false,
		},
	}
	log, err := logger.NewLogger("error", "plain", false, "test")
	if err != nil {
		t.Fatal(err)
	}
	memoryDB, err := database.NewBadger("", true)
	if err != nil {
		t.Fatal(err)
	}
	cl := mpc.NewClient(mpc.Config{Mock: true}, log)
	k := newFusionSignatureController(log, memoryDB, make(chan *signatureRequestQueueItem), cl, mockTxClient{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := k.executeRequest(&tt.item)
			if (err != nil) != tt.expectErr {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
	close(k.queue)
	close(k.stop)
	if err := memoryDB.Close(); err != nil {
		t.Fatal(err)
	}
}
