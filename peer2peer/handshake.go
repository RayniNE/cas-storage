package peer2peer

import "github.com/RayniNE/cas-storage/interfaces"

type HandshakeFunc func(interfaces.Peer) error

func NOPHandshakeFunc(interfaces.Peer) error {
	return nil
}
