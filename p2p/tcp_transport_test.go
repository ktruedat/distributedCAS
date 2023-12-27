package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTCPTransport(t *testing.T) {
	opts := TCPTransportOpts{
		ListenAddr:    ":4000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}
	tcpTr := NewTCPTransport(opts)
	assert.Equal(t, tcpTr.ListenAddr, ":4000")

	assert.Nil(t, tcpTr.ListenAndAccept())
}
