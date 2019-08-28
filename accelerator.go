package main

import (
	"fmt"
	"sync"
	"time"

	eos "github.com/eoscanada/eos-go"
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/skademlia"
	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
)

type Accelerator struct {
	sync.Mutex

	client     *skademlia.Client
	pendingTrx map[string]*PendingTrx // trxPrefix to pending Trx
}

func NewAccelerator(client *skademlia.Client) *Accelerator {
	return &Accelerator{
		client:     client,
		pendingTrx: make(map[string]*PendingTrx),
	}
}

type PendingTrx struct {
	sync.Mutex

	id             string
	accountWeights map[string]int // accounts to weight
	threshold      int
	expiration     time.Time
	seenBy         map[string]bool
	seenAccounts   map[string]bool // hmm.. just trusting them can load our memory

	signatures []string
	trx        *eos.SignedTransaction
	packedTrx  *eos.PackedTransaction
}

func (t *PendingTrx) thresholdMet() bool {
	var threshold int
	for acct := range t.seenAccounts {
		threshold += t.accountWeights[acct]
	}
	if threshold >= t.threshold {
		return true
	}
}

func (a *Accelerator) PushInAccelerator(trx *eos.SignedTransaction, signature string, account string, accountWeights map[string]int, threshold int) error {
	packedTrx, err := trx.Pack(eos.CompressionNone)
	if err != nil {
		return err
	}

	idHash, err := packedTrx.ID()
	if err != nil {
		return err
	}

	id := idHash.String()
	trxPrefix := id[:12]
	pendingTrx := &PendingTrx{
		expiration:     trx.Expiration.Time,
		accountWeights: accountWeights,
		id:             id,
		trx:            trx,
		packedTrx:      packedTrx,
		signatures:     []string{signature},
		seenBy:         map[string]bool{},
		seenAccounts:   map[string]bool{account: true},
	}

	a.Lock()
	a.pendingTrx[trxPrefix] = pendingTrx
	a.Unlock()

	// disperse to nearest neighbor
	return nil
}

func (a *Accelerator) AddSignature(peerID string, trxPrefix string, account string, signature string) error {
	a.Lock()
	trx, ok := a.pendingTrx[trxPrefix]
	a.Unlock()
	if !ok {
		return fmt.Errorf("not found")
	}

	trx.Lock()
	defer trx.Unlock()

	if time.Since(trx.expiration) < 0 {
		// have another mechanism to take it out
		delete(a.pendingTrx, trxPrefix)
		return fmt.Errorf("expired")
	}

	if trx.seenBy[peerID] {
		return fmt.Errorf("already seen by peer")
	}

	if _, ok := trx.accountWeights[account]; !ok {
		return fmt.Errorf("unexpected account")
	}

	// Add who has seen it (the ID)
	trx.seenBy[peerID] = true
	trx.seenAccounts[account] = true

	if trx.thresholdMet() {
		// TODO: submit transaction to target network
	} else {

		go func() {
			time.Sleep(50 * time.Millisecond)

			a.Propagate(trxPrefix)
		}()
	}
	// TODO: propagate through ClosestPeers that have not yet seen it

	return nil
}

func (a *Accelerator) Propagate(trxPrefix string) {
	a.Lock()
	trx, ok := a.pendingTrx[trxPrefix]
	a.Unlock()
	if !ok {
		return
	}

	trx.Lock()
	defer trx.Unlock()

	// check all connections
	for _, id := range a.client.ClosestPeerIDs() {
		sid := id.String()
		if trx.seenBy[sid] {
			continue
		}

		conn, err := a.client.Dial(id.Address())
		if err != nil {
			fmt.Println("failed to connect for propagation:", err)
			continue
		}

		accel := NewAcceleratorClient(conn)

		fmt.Println("Dispersing the thing")
		resp, err := accel.Disperse(context.Background(), &TransactionSignature{
			TrxPrefix: "123123",
			Signature: "EOS123123123",
		})
		if err != nil {
			fmt.Println("Interrupted", err)
			continue
		}

		fmt.Println("Dispersion response:", resp.Ok)
	}
}

// Disperse serves the Disperse gRPC calls
func (a *Accelerator) Disperse(ctx context.Context, req *TransactionSignature) (*Receipt, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		panic("cannot get peer from context")
	}

	info := noise.InfoFromPeer(p)
	if info == nil {
		panic("cannot get info from peer")
	}

	id := info.Get(skademlia.KeyID)
	if id == nil {
		panic("cannot get id from peer")
	}

	fmt.Printf("%s> prefix=%s sig=%s\n", id, req.TrxPrefix, req.Signature)

	err := a.AddSignature(id.(*skademlia.ID).String(), req.TrxPrefix, req.Account, req.Signature)
	if err != nil {
		fmt.Println("Signature NOT contributed:", err)
		return &Receipt{Ok: false}, nil
	}
	fmt.Println("Signature contributed")
	return &Receipt{Ok: true}, nil
}
