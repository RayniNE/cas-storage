package peer2peer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTCPTransport(t *testing.T) {
	config := TCPTransportConfig{
		ListenAddress: ":4000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}

	transport := NewTCPTransport(config)

	assert.Equal(t, config.ListenAddress, transport.ListenAddress)

	assert.Nil(t, transport.ListenAndAccept())
}
