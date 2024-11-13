package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
	"voter_registration/pkg/voting"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type VoterHandler struct {
	client *ethclient.Client
	auth   *bind.TransactOpts
}

func NewVoterHandler(client *ethclient.Client, auth *bind.TransactOpts) *VoterHandler {
	return &VoterHandler{
		client: client,
		auth:   auth,
	}
}

func (vh *VoterHandler) RegisterVoterHandler(w http.ResponseWriter, r *http.Request) {
	// Check HTTP method
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	log.Println(r.Body);

	// Decode request body
	var voterData struct {
		VoterInfo string `json:"voterInfo"`
	}
	if err := json.NewDecoder(r.Body).Decode(&voterData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// generate voter hash
	voteHash := sha256.Sum256([]byte(voterData.VoterInfo))


	// get the deployed contract instance 
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := voting.NewVoting(contractAddress, vh.client)
	if err != nil {
		http.Error(w, "Failed to load contract instance", http.StatusInternalServerError)
		return
	}

	// call the registerVoter function
	result, err := instance.RegisterVoterHash(vh.auth, voteHash)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to register voter on the blockchain", http.StatusInternalServerError)
		return
	}

	log.Printf("Voter registration initiated with tx hash: %s", result.Hash().Hex())

	receipt, err := waitForReceipt(vh.client, result.Hash())
	if err != nil {
		log.Fatalf("Failed to get transaction receipt: %v", err)
	}

	//check the logs 
    log.Println("Voter Reg receipt:", receipt.Logs[0])

	//the logs need to be unpacked <- working on this

	// send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"transactionHash": result.Hash().Hex()})
}


//function polls to get event emitted by the transaction
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