package peer2peer

import (
	"fmt"
	"net"
	"sync"

	"github.com/RayniNE/cas-storage/interfaces"
)

// TCPPeer represents a remote node that we have stablished a connection through TCP
type TCPPeer struct {
	connection net.Conn

	// If we dial and get accepted, outbound == true
	// If we accept a connection, outboud == false
	outbound bool
}

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

func NewTCPPeer(connection net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		connection: connection,
		outbound:   outbound,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	listener, err := net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	t.listener = listener

	go t.acceptLoop()

	return nil
}

func (t *TCPTransport) acceptLoop() {
	for {
		connection, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCPTransport Accept Connection error: %s\n", err.Error())
		}

		go t.handleConnection(connection)
	}
}

func (t *TCPTransport) handleConnection(connection net.Conn) {
	peer := NewTCPPeer(connection, true)

	fmt.Printf("New incoming connection: %+v\n", peer)
}
