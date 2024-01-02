package models

import "net"

// Message holds data that's being sent over the network
// by two connected nodes
type Message struct {
	From    net.Addr
	Payload []byte
}
