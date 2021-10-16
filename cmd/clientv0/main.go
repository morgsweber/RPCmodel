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
		Holder: "Tester test",
		Agency: "0001",
		AccountNumber: 1,
		Money: 0,
	}

	/*
	   Call chama um metodo atrves da conexao estabelecida
	   anteriormente. Os parametros devem ser:
	   -Metodo a ser chamado no servidor no formato 'Tipo.Nome'.
	   Este parametro deve ser uma string
	   -Primeiro argumento do metodo
	   -Segundo argumento do metodo(ponteiro para receber a resposta)
	*/
	err = c.Call("Bank.Open", holder, &reply)
	if err != nil {
		log.Fatal("Bank error: ", err)
	}
	fmt.Printf("Bank:\n%s\n", reply)

	err = c.Call("Bank.Close", acc, &reply)
	if err != nil {
		log.Fatal("Bank error: ", err)
	}
	fmt.Printf("Bank:\n%s\n", reply)

}
