package handlers

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"vote_tallying/pkg/voting"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CountHandler struct {
	client *ethclient.Client
	auth   *bind.TransactOpts
}

func NewVoterHandler(client *ethclient.Client, auth *bind.TransactOpts) *CountHandler {
	return &CountHandler{
		client: client,
		auth:   auth,
	}
}

func (vh *CountHandler) CountHandler(w http.ResponseWriter, r *http.Request) {
	// Check HTTP method
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}


	// get the deployed contract instance 
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := voting.NewVoting(contractAddress, vh.client)
	if err != nil {
		http.Error(w, "Failed to load contract instance", http.StatusInternalServerError)
		return
	}

	// Define the options for the call 
	callOpts := &bind.CallOpts{
		From:  vh.auth.From,  // specify sender address
		BlockNumber: nil, // Use the latest block for the query
	}

	// call the getCount function
	result, err := instance.GetCount(callOpts)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get count from the blockchain", http.StatusInternalServerError)
		return
	}

	log.Println(result);



	// send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]*big.Int{"counts": result})
}
