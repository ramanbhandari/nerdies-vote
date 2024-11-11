package main

import (
	"context"
	"crypto/sha256"
	"log"
	"math/big"
	"time"
	"voter_registration/pkg/auth"
	"voter_registration/pkg/voting"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    // Connect to Ganache
    client, err := ethclient.Dial("http://127.0.0.1:7545")
    if err != nil {
        log.Fatalf("Failed to connect to the local Ganache blockchain: %v", err)
    }
    log.Println("Connected to Ganache")

	//create an auth object to sign transactions 
	auth, err := auth.NewAuth()
    if err != nil {
        log.Fatalf("Failed to create auth object: %v", err)
    }

	log.Println(auth.From)

	// Contract address 
	contractAddress := common.HexToAddress("0xB248Fd56EdC2896F14BD4a7a539750870880a157")

	// Call the interaction function
	checkCandidateTest(client, contractAddress)
	registerVoter(auth, client, contractAddress)


    defer client.Close()
	
}

func checkCandidateTest(client *ethclient.Client, contractAddress common.Address) {
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

    log.Println("Function result:", result.VoteCount)
}

func registerVoter(auth *bind.TransactOpts,client *ethclient.Client, contractAddress common.Address) {
	// getting the instance of the contract 
    instance, err := voting.NewVoting(contractAddress, client)
    if err != nil {
        log.Fatalf("Failed to load contract instance: %v", err)
    }

	voteHash := sha256.Sum256([]byte("new voter 5"))
	
    // call register vote function on the contract
    result, err := instance.RegisterVoterHash(auth, voteHash)
    if err != nil {
        log.Fatalf("Failed to call contract function: %v", err)
    }

    log.Println("Voter Reg result:", result.Hash().Hex())

	// wait for the transaction receipt in a blocking manner
	receipt, err := waitForReceipt(client, result.Hash())
	if err != nil {
		log.Fatalf("Failed to get transaction receipt: %v", err)
	}

    log.Println("Voter Reg receipt:", receipt.Logs[0])


}

func waitForReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	// Loop until the transaction receipt is found
	for {
		// Try to get the receipt
		receipt, err := client.TransactionReceipt(context.Background() , txHash)
		if err != nil && err.Error() != "not found" {
			// If the error is not "not found", return the error
			return nil, err
		}

		// If receipt is not nil, transaction is mined
		if receipt != nil {
			return receipt, nil
		}

		// If not mined yet, wait before trying again
		time.Sleep(2 * time.Second) // Wait for 2 seconds before polling again
	}
}