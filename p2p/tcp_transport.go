package p2p

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established
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

// Close implements the Peer interface
func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
	OnPeer        func(Peer) error
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpcch    chan RPC
	mu       sync.RWMutex
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{TCPTransportOpts: opts, rpcch: make(chan RPC)}
}

// Consume implements the transport interface which
// will return read-only channel for reading incoming messages received from
// another peer
func (transport *TCPTransport) Consume() <-chan RPC {
	return transport.rpcch
}

func (transport *TCPTransport) ListenAndAccept() error {
	var err error
	transport.listener, err = net.Listen("tcp", transport.ListenAddr)
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
	var err error
	defer func() {
		fmt.Printf("dropping peer connection: %v", err)
		conn.Close()
	}()
	peer := NewTCPPeer(conn, true)

	if err = transport.HandshakeFunc(peer); err != nil {
		return
	}

	if transport.OnPeer != nil {
		if err = transport.OnPeer(peer); err != nil {
			return
		}
	}

	// Read loop
	rpc := RPC{}
	for {
		err := transport.Decoder.Decode(conn, &rpc)
		if errors.Is(err, net.ErrClosed) {
			return
		}
		if err != nil {
			fmt.Printf("TCP read error: %s\n", err)
			continue
		}
		rpc.From = conn.RemoteAddr()
		transport.rpcch <- rpc
		fmt.Printf("RPC: %+v\n", rpc)
	}

}
