package admin

import (
	"fmt"
	"strconv"
)

/**
Processo Administração: realiza abertura e fechamento de contas (para agências), autentica que contas já existem (tanto para agências e caixas automáticos) e também executa as operações de manipulação destas contas (saques e depósitos). Deve garantir semântica de execução exactely once para operações que sejam não-idempotentes;*/

type Account struct {
	Holder        string
	Agency        string
	AccountNumber int
	Money         float64
}

type Bank struct {
	memory [100]Account
}

var pos int = 0

/*
About all methods:
	Return an error. If something return besides nil the client will receive just the error, without the reply pointer
*/

/**
Open:
	Is to open some account and to do this method just need receive the holder, start with money equal zero and the account number and agency depend of other thing.

	For know is with a generic agency.
*/
func (bank *Bank) Open(A Account, reply *string) error {
	bank.memory[pos] = A
	bank.memory[pos].Money = 0
	*reply = "Holder: " + A.Holder + "\nAgency: " + A.Agency + "\nAccountNumber " + strconv.Itoa(pos)
	pos++

	return nil
}

/**
Close:
	Is to close account and to do this the program receive the account and verify if the account is right
*/
func (bank *Bank) Close(A *Account, reply *string) error {
	if A.Holder == bank.memory[A.AccountNumber].Holder {
		bank.memory[A.AccountNumber].Holder = ""
		bank.memory[A.AccountNumber].Agency = ""
		bank.memory[A.AccountNumber].Money = 0
		bank.memory[A.AccountNumber].AccountNumber = -1
		*reply = "Account is close"
	} else {
		*reply = "This account don't exist"
	}

	return nil
}

/**
Authenticate:
Authenticates that accounts already exist, so will search in the memory.
*/
func (bank *Bank) Authenticate(A Account, reply *string) error {
	*reply = "This is a invalid account"
	for i := 0; i < len(bank.memory); i++ {
		if A.Holder == bank.memory[A.AccountNumber].Holder {
			if A.Agency == bank.memory[A.AccountNumber].Agency {
				if A.AccountNumber == bank.memory[A.AccountNumber].AccountNumber {
					*reply = "This is a valid account"
				}
			}
		}
	}
	return nil
}

/**
Sack:
	Sack some value in account A if have this value, receiving a account where Money is the value to sack.
*/
func (bank *Bank) Sack(A Account, reply *string) error {
	if A.Money < bank.memory[A.AccountNumber].Money {
		bank.memory[A.AccountNumber].Money = bank.memory[A.AccountNumber].Money - A.Money

		*reply = "Success, now you have: " + fmt.Sprint(bank.memory[A.AccountNumber].Money)
	} else {
		*reply = "Insufficient funds"
	}
	return nil
}

/**
Deposit:
	Deposit some value in account A, receiving a account where Money is the value to deposit.
*/
func (bank *Bank) Deposit(A Account, reply *string) error {
	bank.memory[A.AccountNumber].Money = bank.memory[A.AccountNumber].Money + A.Money
	*reply = "Sucesso, now you have " + fmt.Sprint(bank.memory[A.AccountNumber].Money)
	return nil
}

/**
Consult:
	Consult some value in account A, receiving a account where Money is the value to consult.
*/
func (bank *Bank) Consult(A Account, reply *string) error {
	*reply = fmt.Sprint(bank.memory[A.AccountNumber].Money)
	return nil
}
