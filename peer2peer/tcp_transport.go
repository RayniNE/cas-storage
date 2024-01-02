package peer2peer

import (
	"fmt"
	"net"
	"sync"

	"github.com/RayniNE/cas-storage/interfaces"
	"github.com/RayniNE/cas-storage/models"
)

// TCPPeer represents a remote node that we have stablished a connection through TCP
type TCPPeer struct {
	connection net.Conn

	// If we dial and get accepted, outbound == true
	// If we accept a connection, outboud == false
	outbound bool
}

type TCPTransportConfig struct {
	ListenAddress string
	Decoder       Decoder
	HandshakeFunc HandshakeFunc
}

type TCPTransport struct {
	TCPTransportConfig
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]interfaces.Peer
}

func NewTCPTransport(config TCPTransportConfig) *TCPTransport {
	return &TCPTransport{
		TCPTransportConfig: config,
	}
}

func NewTCPPeer(connection net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		connection: connection,
		outbound:   outbound,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	listener, err := net.Listen("tcp", t.ListenAddress)
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
			fmt.Printf("TCPTransport, accept connection error: %s\n", err.Error())
		}

		go t.handleConnection(connection)
	}
}

func (t *TCPTransport) handleConnection(connection net.Conn) {
	peer := NewTCPPeer(connection, true)
	fmt.Printf("TCPTransport, new incoming connection: %+v\n", peer)

	err := t.HandshakeFunc(peer)
	if err != nil {
		connection.Close()
		fmt.Printf("TCPTransport, closing connection due to handshake error: %s\n", err.Error())
		return
	}

	message := &models.Message{}

	for {
		err = t.Decoder.Decode(connection, message)
		if err != nil {
			fmt.Printf("TCPTransport, decoding connection error: %s\n", err.Error())
			continue
		}

		message.From = connection.RemoteAddr()

		fmt.Printf("Message: %+v\n", message)
		fmt.Printf("Payload: %s\n", string(message.Payload))
	}
}
