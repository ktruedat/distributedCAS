package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP estblished
// connection.
type TCPPeer struct {
	// conn is teh underlying connection of the peer
	conn net.Conn
	// if we dial and retrieve a connection => outbound == true
	// if we accept and retrieve a connection => outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{conn: conn, outbound: outbound}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandshakeFunc
	decoder       Decoder
	mu            sync.RWMutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		shakeHands:    NOPHandshakeFunc,
		listenAddress: listenAddr,
	}
}

func (transport *TCPTransport) ListenAndAccept() error {
	var err error
	transport.listener, err = net.Listen("tcp", transport.listenAddress)
	if err != nil {
		return err
	}
	go transport.startAcceptLoop()
	return nil
}

func (transport *TCPTransport) startAcceptLoop() {
	for {
		conn, err := transport.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %v", err)
		}
		fmt.Printf("new incoming connection: %+v\n", conn)
		go transport.handleConn(conn)
	}
}

type Temp struct{}

func (transport *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	if err := transport.shakeHands(peer); err != nil {
		conn.Close()
		return
	}

	// Read loop
	msg := &Temp{}
	for {
		if err := transport.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}
	}

}