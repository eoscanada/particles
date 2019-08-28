package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	eos "github.com/eoscanada/eos-go"
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/cipher"
	"github.com/perlin-network/noise/handshake"
	"github.com/perlin-network/noise/skademlia"
	"github.com/toolkits/cron"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

const (
	C1 = 1
	C2 = 1
)

var flagPort = flag.String("port", "3000", "Port to listen on")

func main() {
	flag.Parse()

	keyBag := eos.NewKeyBag()
	_ = keyBag.ImportFromFile("keys")
	fmt.Println("Loaded keys:")
	for _, key := range keyBag.Keys {
		fmt.Println("-", key.PublicKey().String())
	}
	eosAPI := eos.New("https://mainnet.eoscanada.com")
	eosAPI.SetSigner(keyBag)

	listener, err := net.Listen("tcp", ":"+*flagPort)
	if err != nil {
		panic(err)
	}

	keys, err := skademlia.NewKeys(C1, C2)
	if err != nil {
		panic(err)
	}

	addr := "127.0.0.1:" + *flagPort
	fmt.Printf("Listening for peers on: %s\n", addr)

	client := skademlia.NewClient(addr, keys, skademlia.WithC1(C1), skademlia.WithC2(C2))
	client.SetCredentials(noise.NewCredentials(addr, handshake.NewECDH(), cipher.NewAEAD(), client.Protocol()))
	client.OnPeerJoin(func(conn *grpc.ClientConn, id *skademlia.ID) {
		fmt.Println("New peer joined:", id)
	})
	client.OnPeerLeave(func(conn *grpc.ClientConn, id *skademlia.ID) {
		fmt.Println("Peer left:", id)
	})

	accel := NewAccelerator(client)

	go func() {
		server := client.Listen()
		RegisterAcceleratorServer(server, accel)

		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	for _, addr := range flag.Args() {
		if _, err := client.Dial(addr); err != nil {
			panic(err)
		}
	}

	client.Bootstrap()

	// Start generating transactions
	// Pile them in a stack that need to be sent
	// Start watching the target chain to see which refblock to use,
	// in order to craft a deterministic trx ID.
	// Sign it when ready.
	c := cron.New()
	c.AddFunc("0,5,10,15,20,25,30,35,40,45,50,55 * * * * *", func() {
		fmt.Println("New cron")
		for _, conn := range client.AllPeers() {
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
	})
	c.Start()

	<-make(chan bool)
}
