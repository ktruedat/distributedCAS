package main

import (
	"github.com/ktruedat/distributedCAS/p2p"
	"log"
)

func main() {
	transport := p2p.NewTCPTransport(":3000")
	if err := transport.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
