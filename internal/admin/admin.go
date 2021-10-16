package admin

/**
Processo Administração: realiza abertura e fechamento de contas (para agências), autentica que contas já existem (tanto para agências e caixas automáticos) e também executa as operações de manipulação destas contas (saques e depósitos). Deve garantir semântica de execução exactely once para operações que sejam não-idempotentes;*/

type Account struct {
	Holder string
	Agency string
	AccountNumber int
	Money float64
}

type Bank struct {
	memory [100]Account
}

var pos int = 0;


/*
About all methods:
	Return an error. If something return besides nil the client will receive just the error, without the reply pointer
*/

// 
/**
Open:
	Is to open some account and to do this method just need receive the holder, start with money equal zero and the account number and agency depend of other thing.

	For know is with a generic agency.
*/
func (bank *Bank) Open(holder string, reply *string) error {
	A := Account{
		Holder: holder,
		Agency: "0001",
		AccountNumber: pos,
		Money: 0,
	}
	bank.memory[pos] = A
	*reply = "Holder: " + A.Holder + "\nAgency: " + A.Agency + "\nAccountNumber " +  string(pos) // add money here in the end

	return nil
}

/**
Close:
	Is to close account and to do this the program receive the account and verify if the account is right
*/
func (bank *Bank) Close(A *Account, reply *string) error {
	if A.Holder == bank.memory[A.AccountNumber].Holder{
		bank.memory[A.AccountNumber].Holder = ""
		bank.memory[A.AccountNumber].Agency = ""
		bank.memory[A.AccountNumber].Money = 0
		bank.memory[A.AccountNumber].AccountNumber = -1
		*reply = "Account is close"
	}else{
		*reply = "This account don't exist"
	}
	
	return nil
}

/**
	Authenticate:
	Authenticates that accounts already exist, so will search in the memory.
*/

/*
// deposit some value in account A
func (bank *Bank) Deposit(A *Account, deposit float64, reply float64) error {
	A.Money = A.Money + deposit
	//*reply = A.Money
	return nil

}
*/

