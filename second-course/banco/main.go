package main

import "fmt"

type ContaCorrente struct {
	titular 		string 
	numeroAgencia 	int 
	numeroConta 	int
	saldo 			float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
}

func main() {
	contaDoPedro := ContaCorrente{
		titular:"Pedro Makoski",
		numeroAgencia: 589,
		numeroConta: 123456,
		saldo: 1700.90,
	}

	fmt.Println(contaDoPedro.saldo)

	fmt.Println(contaDoPedro.Sacar(-100))

	fmt.Println(contaDoPedro.saldo)
}