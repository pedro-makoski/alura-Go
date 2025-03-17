package main

import (
	"fmt"
	"banco/contas"
)

func main() {
	contaDoPedro := contas.ContaCorrente{
		Titular:"Pedro Makoski",
		NumeroAgencia: 589,
		NumeroConta: 123456,
		Saldo: 1700.90,
	}

	contaDoGustavo := contas.ContaCorrente{
		Titular: "Gustavo",
		Saldo: 100,
	}

	status := contaDoGustavo.Transferir(-100, &contaDoPedro)
	fmt.Println(status)
	fmt.Println(contaDoGustavo)
	fmt.Println(contaDoPedro)
}