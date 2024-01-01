package peer2peer

import (
	"net"
	"sync"

	"github.com/RayniNE/cas-storage/interfaces"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]interfaces.Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	listener, err := net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	t.listener = listener

	return nil
}
