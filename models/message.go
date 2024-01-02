package models

import "net"

type Message struct {
	From    net.Addr
	Payload []byte
}
