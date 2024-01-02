package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RayniNE/cas-storage/peer2peer"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	LISTEN_ADDRESS := os.Getenv("LISTEN_ADDRESS")

	config := peer2peer.TCPTransportConfig{
		ListenAddress: fmt.Sprintf(":%s", LISTEN_ADDRESS),
		HandshakeFunc: peer2peer.NOPHandshakeFunc,
		Decoder:       peer2peer.DefaultDecoder{},
	}

	tcpTransport := peer2peer.NewTCPTransport(config)

	err := tcpTransport.ListenAndAccept()
	if err != nil {
		panic(err)
	}

	log.Printf("Application started. Running on PORT: %s\n", LISTEN_ADDRESS)

	select {}
}
