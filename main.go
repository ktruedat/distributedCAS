package main

import (
	"fmt"
	"github.com/ktruedat/distributedCAS/p2p"
	"log"
)

func onPeer(peer p2p.Peer) error {
	fmt.Println("doing onPeer func logic outside tcp transport")
	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        onPeer,
	}
	transport := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-transport.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()
	if err := transport.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
