Smart Contract 
======================

The smart contract Voting is developed using the Truffle framework and tested/deployed locally using Ganache. Follow these steps to set up, test, and migrate the contract.

Prerequisites
-------------

-   **Node.js**: Install Node.js if not already installed. (<https://nodejs.org/>)
-   **Truffle**: Install Truffle globally via npm:

    `npm install -g truffle`

-   **Ganache**: For the local blockchain, use the Ganache desktop app or CLI. (https://trufflesuite.com/ganache/)

Getting Started
---------------

1.  **Clone the repository**:

    `git clone <repository-url>
    cd <repository-name>`

2.  **Install dependencies**:

    `npm install`

3.  **Start Ganache**:

    -   Open Ganache and ensure it is running on `http://127.0.0.1:7545` (the default RPC server and port).
    -   Note private keys provided by Ganache for account management.
  
4.  **Compile the smart contracts**:

    `truffle compile`

5.  **Run migrations**:

    `truffle migrate --network development`

    This will compile and deploy the contracts to the local Ganache blockchain.

    Note the contract address once the migration is successful

6.  **Run tests**:

    `truffle test`

    This will execute all tests in the `test/` directory against the local blockchain.

Project Structure
-----------------

-   **contracts/**: Directory containing Solidity smart contract **Voting**.
-   **migrations/**: Migration scripts for deploying the contract.
-   **test/**: JavaScript file for contract testing.
-   **truffle-config.js**: Configuration file for Truffle, including network setup.

Additional Commands
-------------------

-   **Reset Migrations**:

    `truffle migrate --reset --network development`

    This will redeploy the contract.

-   Use `truffle console --network development` to interact with the contracts directly in the Truffle console.
