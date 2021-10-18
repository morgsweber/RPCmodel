package main

import (
	"fmt" // implements formatted I/O with functions analogous to C's printf and scanf.
	"log"
	"net/rpc" // remote procedure call
	"ppd/t2/internal/admin"
)



func Work1(){
	//Variavel para receber os resultados
	var reply string
	//Estrutura para enviar os numeros
	acc := admin.Account{
		Holder:        "CesarDeRose",
		Agency:        "0001",
		AccountNumber: 0,
		Money:         0,
	}

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

	holder := "CesardeRose"

	err = c.Call("Bank.Open", holder, &reply)
	if err != nil {
		log.Fatal("Bank error: ", err)
	}
	fmt.Printf("Bank:\n%s\n\n", reply)

	holder = "Isabel"

	err = c.Call("Bank.Open", holder, &reply)
	if err != nil {
		log.Fatal("Bank error: ", err)
	}
	fmt.Printf("Bank:\n%s\n\n", reply)

	go Work1()
	

}


