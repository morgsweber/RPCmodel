package admin

/**
Processo Administração: realiza abertura e fechamento de contas (para agências), autentica que contas já existem (tanto para agências e caixas automáticos) e também executa as operações de manipulação destas contas (saques e depósitos). Deve garantir semântica de execução exactely once para operações que sejam não-idempotentes;*/

type Account struct {
	Holder string
	Agency int
	AccountNumber int
	Money float64
}

type Bank struct {
	memory [100]Account
}

var pos int = 0;


/*
	Return an error. If something return besides nil the client will receive just the error, without the reply pointer
*/

// 
/**
Open:
	To open some account the user just can decide the holder, start with money equal zero and the account number and agency depend of other thing
*/
func (bank *Bank) Open(Holder string, reply *float64) error {
	//Account A
//	a.memory[pos] = 
	//pos := pos + 1;
	return nil
}

/**
Close accounts:
	To close some account the program receive the account and verify if the account is right
*/
func (bank *Bank) Close(A *Account, reply *float64) error {
	if A.Holder == bank.memory[A.AccountNumber].Holder{
		bank.memory[A.AccountNumber].Holder = ""
		bank.memory[A.AccountNumber].Agency = -1
		bank.memory[A.AccountNumber].Money = 0
		bank.memory[A.AccountNumber].AccountNumber = -1
	}
	return nil
}


// deposit some value in account A
func (bank *Bank) Deposit(A *Account, deposit float64, reply float64) error {
	A.Money = A.Money + deposit
	//*reply = A.Money
	return nil
}