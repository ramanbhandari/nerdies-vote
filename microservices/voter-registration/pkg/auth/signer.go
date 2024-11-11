package auth

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

// NewAuth creates and returns an auth object for transaction signing
func NewAuth() (*bind.TransactOpts, error) {
    // Load the .env file if not already loaded
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found; ensure PRIVATE_KEY and CHAIN_ID are set in environment")
    }

    // Retrieve private key and chain ID from environment
	
    privateKeyStr := os.Getenv("PRIVATE_KEY")
    if privateKeyStr == "" {
        return nil, fmt.Errorf("PRIVATE_KEY not set in environment")
    }

    chainIDStr := os.Getenv("CHAIN_ID")
    if chainIDStr == "" {
        return nil, fmt.Errorf("CHAIN_ID not set in environment")
    }

    // Parse the chain ID to *big.Int
    chainID := new(big.Int)
    _, ok := chainID.SetString(chainIDStr, 10)
    if !ok {
        return nil, fmt.Errorf("Invalid CHAIN_ID: %s", chainIDStr)
    }

    // Convert the private key from hex
    privateKey, err := crypto.HexToECDSA(privateKeyStr)
    if err != nil {
        return nil, fmt.Errorf("invalid private key: %v", err)
    }

    // Create the auth object with the private key and chain ID
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    if err != nil {
        return nil, fmt.Errorf("failed to create auth object: %v", err)
    }

    return auth, nil
}
