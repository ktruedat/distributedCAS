package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	tcpTr := NewTCPTransport(listenAddr)
	assert.Equal(t, tcpTr.listenAddress, listenAddr)

	assert.Nil(t, tcpTr.ListenAndAccept())
}
