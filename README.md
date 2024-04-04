# RPC Model #
This repository was created to develop a university work on Fundamentals of Parallel and Distributed Processing.
2021/2

##### About the project
The objective of this work is to develop a distributed system based on the Client/Server model using RPC that simulates the creation of customer accounts and bank transactions in these accounts.
The following functionalities were implemented:
- Administration Process: opens and closes accounts (for branches), authenticates which accounts already exist (both for branches and ATMs) and also performs manipulation operations on these accounts (withdrawals and deposits). It must guarantee exactely once execution semantics for operations that are non-idempotent;
- Agency Process: Requests opening, authentication and closing of accounts and can also request deposit, withdrawal and balance consultation on an existing account. Account opening, deposit and withdrawal must be guaranteed non-idempotent operations (exactely once execution semantics);
- Automatic Teller Machine: Requests deposit, withdrawal and balance check on existing account. The first two must be guaranteed non-idempotent operations (exactely once execution semantics) even if an error occurs when confirming the operation (simulate with fault injection).
  
##### Requirements
* [Go Language](https://go.dev/)

##### Running the project

In a terminal, initilize the server running:

`go run .\cmd\server_v0\main.go`

After, open a new terminal an run the following command to start the client:

`go run .\cmd\agency\main.go`


##### Observations

- The holder's name need to be on word, for exemple: MorganaWeber, Morgana, Daniela...)
- The agency will recognize commands like: 
  - Open Account: `Open`
  - Authenticate Account: `Authenticate`
  - Close Account: `Close`
  - Sack: `Sack`
  - Deposit: `Deposit`
- The ATM will recognize commands like:
  - Authenticate Account: `Authenticate`
  - Consult: `Consult`
  - Sack: `Sack`
  - Deposit: `Deposit`
