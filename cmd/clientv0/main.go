package main

import (
	"fmt" // implements formatted I/O with functions analogous to C's printf and scanf.
	"log"
	"net/rpc" // remote procedure call
	"ppd/t2/internal/admin"
)

func main() {
	/*
	   Inicializa o cliente na porta 4040 do localhost
	   utilizando o protocolo tcp. Se o servidor estiver
	   em outra maquina deve ser utilizado IP:porta no
	   segundo parametro.
	*/
	c, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Error dialing: ", err)
	}

	//Variavel para receber os resultados
	var reply string

	holder := "Cesar de Rose"
	//Estrutura para enviar os numeros
	acc := admin.Account{
		Holder:        "Tester test",
		Agency:        "0001",
		AccountNumber: 1,
		Money:         0,
	}

	/*
	   Call chama um metodo atrves da conexao estabelecida
	   anteriormente. Os parametros devem ser:
	   -Metodo a ser chamado no servidor no formato 'Tipo.Nome'.
	   Este parametro deve ser uma string
	   -Primeiro argumento do metodo
	   -Segundo argumento do metodo(ponteiro para receber a resposta)
	*/
	// Test all methods of admin
	err = c.Call("Bank.Open", holder, &reply)
	if err != nil {
		log.Fatal("Bank error: ", err)
	}
	fmt.Printf("Bank:\n%s\n\n", reply)

	// Try close an accont that don't exist
	err = c.Call("Bank.Close", acc, &reply)
	if err != nil {
		log.Fatal("Bank error: ", err)
	}
	fmt.Printf("Bank:\n%s\n\n", reply)

	// Try authenticate an accont that don't exist
	err = c.Call("Bank.Authenticate", acc, &reply)
	if err != nil {
		log.Fatal("Bank error: ", err)
	}
	fmt.Printf("Bank:\n%s\n\n", reply)

	acc = admin.Account{
		Holder:        "Cesar de Rose",
		Agency:        "0001",
		AccountNumber: 1,
		Money:         50, // value
	}

	err = c.Call("Bank.Sack", acc, &reply)
	if err != nil {
		log.Fatal("Bank error: ", err)
	}
	fmt.Printf("Bank:\n%s\n\n", reply)

	acc = admin.Account{
		Holder:        "Cesar de Rose",
		Agency:        "0001",
		AccountNumber: 1,
		Money:         100,
	}

	err = c.Call("Bank.Deposit", acc, &reply)
	if err != nil {
		log.Fatal("Bank error: ", err)
	}
	fmt.Printf("Bank:\n%s\n\n", reply)

	acc = admin.Account{
		Holder:        "Cesar de Rose",
		Agency:        "0001",
		AccountNumber: 1,
		Money:         50,
	}

	err = c.Call("Bank.Sack", acc, &reply)
	if err != nil {
		log.Fatal("Bank error: ", err)
	}
	fmt.Printf("Bank:\n%s\n\n", reply)

}
