package service

import (
	"testing"

	"github.com/qredo/fusionchain/mpc-relayer/pkg/database"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/logger"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/mpc"
	"github.com/qredo/fusionchain/x/treasury/types"
)

func Test_SigControllerStart(t *testing.T) {
	k := testSetupSignatureController(t)
	if err := k.Start(); err != nil {
		t.Fatal(err)
	}
	if err := k.Stop(); err != nil {
		t.Fatal(err)
	}
	close(k.queue)
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
	k := testSetupSignatureController(t)

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
}

func testSetupSignatureController(t *testing.T) *signatureController {
	log, err := logger.NewLogger("error", "plain", false, "test")
	if err != nil {
		t.Fatal(err)
	}
	memoryDB, err := database.NewBadger("", true)
	if err != nil {
		t.Fatal(err)
	}
	cl := mpc.NewClient(mpc.Config{Mock: true}, log)
	return newFusionSignatureController(log, memoryDB, make(chan *signatureRequestQueueItem), cl, mockTxClient{})
}
