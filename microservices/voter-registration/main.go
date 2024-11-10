package main

import (
	"log"
	"math/big"
	"voter_registration/voting"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    // Connect to Ganache
    client, err := ethclient.Dial("http://127.0.0.1:7545")
    if err != nil {
        log.Fatalf("Failed to connect to the local Ganache blockchain: %v", err)
    }
    log.Println("Connected to Ganache")


	// Contract address 
	contractAddress := common.HexToAddress("0xB248Fd56EdC2896F14BD4a7a539750870880a157")

	// Call the interaction function
	interactWithContract(client, contractAddress)


    defer client.Close()
	
}

func interactWithContract(client *ethclient.Client, contractAddress common.Address) {
	// getting the instance of the contract 
    instance, err := voting.NewVoting(contractAddress, client)
    if err != nil {
        log.Fatalf("Failed to load contract instance: %v", err)
    }

    // call contract methods directly
    result, err := instance.Candidates(nil,big.NewInt(1)) 
    if err != nil {
        log.Fatalf("Failed to call contract function: %v", err)
    }

    log.Println("Function result:", result)
}