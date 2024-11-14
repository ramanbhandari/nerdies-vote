package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
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

	// decode request body
	var voterData struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Age       int    `json:"age"`
		Address   string `json:"address"`
	}

	if err := json.NewDecoder(r.Body).Decode(&voterData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Concatenate voter information into a single string
	data := fmt.Sprintf("%s:%s:%d:%s", voterData.FirstName, voterData.LastName, voterData.Age, voterData.Address)
	// generate voter hash
	voteHash := sha256.Sum256([]byte(data))

	log.Println("VoterID created", voteHash)


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

	eventEmitted := false;
	// Loop over each log in the transaction receipt to find emitted event 
	for _, tLog := range receipt.Logs {

		//this will parse only the log(event) that is VoterRegistered else err, like checking the event name
		registrationEvent, err := instance.ParseVoterRegistered(*tLog)
		if err == nil {
			log.Println("Event Emitted: VoterRegistered")
			// log.Println("Event Emitted: VoterRegistered", registrationEvent.VoterHash, voteHash)

			//the event is emitted check if the voter hashes match 
			eventEmitted = registrationEvent.VoterHash == voteHash;
		}
	}

	log.Println("Event emitted?", eventEmitted)

	// send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"eventSuccess": eventEmitted})
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