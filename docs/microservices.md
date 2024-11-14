Voter Registration Microservice
===============================

This service provides a backend API for handling administrative voter registration requests in a blockchain-based voting system. It connects to a local Ganache blockchain, allowing registration transactions to be signed and sent to a deployed smart contract.

Structure
---------

```/voter-registration
├── main.go
├── /pkg
│   ├── /auth
│   │   └── signer.go
│   └── /voting
│       └── voting.go
├── /handlers
│   └── voterHandler.go
└── .env 
```

### main.go

This is the entry point for the Voter Registration microservice.

-   **Server Setup**: Initializes an HTTP server that connects to a local Ganache blockchain instance.
-   **Endpoint**: Registers an HTTP route `/registerVoter` that triggers `RegisterVoterHandler` to process voter registration request. The server listens on port 8080.
-   **Request Body Format**: The request payload should be in JSON format as follows:

    ```
    {
      "firstName": "John",
      "lastName": "Doe",
      "age": 30,
      "address": "123 Main St"
    }
    ```

### Packages

#### auth/signer.go

-   Provides a helper function to generate a signed transaction object (`auth`) that includes the private key and chain ID, which are necessary for interacting with Ethereum smart contracts.
-   **Signing Transactions**: Ensures that all transactions are securely signed by the sender's private key, verifying ownership of the transaction's sender address.

#### voting/voting.go

-   **Generated with abigen**: Uses Go-Ethereum's `abigen` tool to generate a Go package based on the Ethereum contract's ABI and bytecode, providing type-safe functions to interact with the contract.
-   **Purpose**: This package simplifies interactions with the contract by abstracting the complexities of transaction building, signing, and contract calls.

### Handlers

#### voterHandler.go

This handler manages the core functionality of voter registration.

-   **Request Handling**: Validates the request type (POST), parses JSON data from the request body, and constructs a Voter ID by concatenating and hashing voter information.
-   **Blockchain Interaction**: Connects to the deployed smart contract, calling the `registerVoterHash` function with the Voter ID.
-   **Event Polling**: Polls the blockchain for a confirmation event. Once the event is emitted, it sends a success response back to the caller.

### .env File

The `.env` file holds critical environment variables needed for configuration:

-   **GANACHE_URL**: URL for Ganache (default: `http://127.0.0.1:7545`)
-   **CONTRACT_ADDRESS**: Address of the deployed smart contract
-   **PRIVATE_KEY**: Private key of the admin account (typically the first account in Ganache)
-   **CHAIN_ID**: Chain ID used by Ganache (default: `1337`)

Running the Service
-------------------

To start the Voter Registration service, run the following command:

`go run main.go`

Ensure that all environment variables are correctly configured in the `.env` file before starting the service.


Vote Tallying Microservice
===============================

This service provides a backend API for retrieving vote count for all the candidates. It connects to a local Ganache blockchain to retrieve counts from the deployed smart contract.

Structure
---------

```/vote-tallying
├── main.go
├── /pkg
│   ├── /auth
│   │   └── signer.go
│   └── /voting
│       └── voting.go
├── /handlers
│   └── countHandler.go
└── .env 
```

### main.go

This is the entry point for the Vote Tallying microservice.

-   **Server Setup**: Initializes an HTTP server that connects to a local Ganache blockchain instance.
-   **Endpoint**: Registers an HTTP route `/voteCount` that triggers `CountHandler` to process the get request for vote count. The server listens on port 8080.

### Packages

#### auth/signer.go

-   Provides a helper function to generate a signed transaction object (`auth`) that includes the private key and chain ID, which are necessary for interacting with Ethereum smart contracts.
-  **Purpose**: **signing Transactions**, ensures that all transactions are securely signed by the sender's private key, verifying ownership of the transaction's sender address.

#### voting/voting.go

-   **Generated with abigen**: Uses Go-Ethereum's `abigen` tool to generate a Go package based on the Ethereum contract's ABI and bytecode, providing type-safe functions to interact with the contract.
-   **Purpose**: This package simplifies interactions with the contract by abstracting the complexities of transaction building, signing, and contract calls.

### Handlers

#### countHandler.go

This handler manages the core functionality of retrieving vote counts.

-   **Request Handling**: Validates the request type (GET).
-   **Blockchain Interaction**: Connects to the deployed smart contract, calling the `getCount` function.
-   **Response**: Returns an array of vote counts for each candidate. Sends the array as a JSON response back to the caller.

### .env File

The `.env` file holds critical environment variables needed for configuration:

-   **GANACHE_URL**: URL for Ganache (default: `http://127.0.0.1:7545`)
-   **CONTRACT_ADDRESS**: Address of the deployed smart contract
-   **PRIVATE_KEY**: Private key of the admin account (typically the first account in Ganache)
-   **CHAIN_ID**: Chain ID used by Ganache (default: `1337`)

Running the Service
-------------------

To start the Voter Registration service, run the following command:

`go run main.go`

Ensure that all environment variables are correctly configured in the `.env` file before starting the service.
