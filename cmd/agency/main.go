package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"strings"

	"ppd/t2/internal/admin"
)

/**
Processo Agência: Solicita abertura, autenticação e fechamento de contas e também pode solicitar depósito, retirada e consulta de saldo em conta existente. A abertura de conta, o depósito e a retirada devem ser operações garantidamente não-idempotentes (semântica de execução exactely once);
*/

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

	//Buffer para ler do terminal
	reader := bufio.NewReader(os.Stdin)

	for {
		//Leitura de uma linha do terminal
		text, e := reader.ReadString('\n')
		if e != nil {
			log.Fatal(e)
		}
		text = strings.Replace(text, "\r\n", "", -1) //Windows
		//text = strings.Replace(text, "\n", "", -1) //Unix

		//Separa a linha pelos espacos em branco
		input := strings.Split(text, " ") // Mult 4 2 > [Mult, 4, 2]
		// Open <> <>
		// Deposit CesardeRose 0001 1 100

		//Converte string para float
		a, e1 := strconv.ParseFloat(input[1], 64)
		if e1 != nil {
			log.Fatal(e1)
		}
		b, e2 := strconv.ParseFloat(input[2], 64)
		if e2 != nil {
			log.Fatal(e2)
		}

		//Cria a struct para enviar para o servidor
		args := admin.Account{
			Holder:        a, // string
			Agency:        b, // string
			AccountNumber: c, // int
			Money:         d, // float
		}

		/*
		   Call chama um metodo atrves da conexao estabelecida
		   anteriormente. Os parametros devem ser:
		   -Metodo a ser chamado no servidor no formato 'Tipo.Nome'.
		   Este parametro deve ser uma string
		   -Primeiro argumento do metodo
		   -Segundo argumento do metodo(ponteiro para receber a resposta)
		*/
		err = c.Call("Bank."+input[0], args, &reply)
		if err != nil {
			log.Fatal("Bank error: ", err)
		}
		fmt.Printf("Result = %f\n", reply)
	}

}
