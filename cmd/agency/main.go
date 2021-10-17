package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"ppd/t2/internal/admin"
	"strconv"
	"strings"
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
	var holder string
	var agency string
	var accountNumber int
	var money float64

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
		input := strings.Split(text, " ")

		//Verify the operation
		if input[0] == "Open" {
			holder = input[1]
			agency = "0001"
			accountNumber = 1111 //TO DO: get position
			money = 0.0
		} else {
			a, e1 := strconv.ParseInt(input[1], 10, 64)
			if e1 != nil {
				log.Fatal(e1)
			}
			accountNumber = int(a)
			if input[0] == "Sack" {
				b, e2 := strconv.ParseFloat(input[2], 64)
				if e2 != nil {
					log.Fatal(e2)
				}
				agency = input[2]
				money = b
			}
		}

		//Cria a struct para enviar para o servidor
		args := admin.Account{
			Holder:        holder,
			Agency:        agency,
			AccountNumber: accountNumber,
			Money:         money,
		}

		err = c.Call("Bank."+input[0], args, &reply)

		if err != nil {
			log.Fatal("Bank error: ", err)
		}
		fmt.Printf("Result = %f\n", reply)
	}

}