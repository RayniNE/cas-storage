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

	tcpTransport := peer2peer.NewTCPTransport(fmt.Sprintf(":%s", LISTEN_ADDRESS))

	err := tcpTransport.ListenAndAccept()
	if err != nil {
		panic(err)
	}

	log.Printf("Application started. Running on PORT: %s", LISTEN_ADDRESS)

	select {}
}
