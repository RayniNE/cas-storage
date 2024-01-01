package peer2peer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTCPTransport(t *testing.T) {
	listenAddress := ":4000"

	transport := NewTCPTransport(listenAddress)

	assert.Equal(t, listenAddress, transport.listenAddress)
}
