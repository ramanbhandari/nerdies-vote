package main

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    // Connect to Ganache
    client, err := ethclient.Dial("http://127.0.0.1:7545")
    if err != nil {
        log.Fatalf("Failed to connect to the local Ganache blockchain: %v", err)
    }
    log.Println("Connected to Ganache")
    defer client.Close()
	
}