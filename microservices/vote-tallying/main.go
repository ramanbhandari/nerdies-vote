package main

import (
	"log"
	"net/http"
	"os"
	"vote_tallying/handlers"
	"vote_tallying/pkg/auth"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	//load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connect to Ganache
	client, err := ethclient.Dial(os.Getenv("GANACHE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the local Ganache blockchain: %v", err)
	}
	defer client.Close()
	log.Println("Connected to Ganache")

	// Set up auth
	auth, err := auth.NewAuth()
	if err != nil {
		log.Fatalf("Failed to create auth object: %v", err)
	}

	// set up the handler
	countHandler := handlers.NewCountHandler(client, auth)

	http.HandleFunc("/voteCount", countHandler.CountHandler)

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
