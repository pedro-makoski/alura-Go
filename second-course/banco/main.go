package main

import "fmt"

type ContaCorrente struct {
	titular 		string 
	numeroAgencia 	int 
	numeroConta 	int
	saldo 			float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque >= 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}

	if valorDoSaque < 0 {
		return "O valor não pode ser menor que 0"
	}
	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDeposito float64) string {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso"
	}

	return "O valor do depósito é menor que 0"
}

func main() {
	contaDoPedro := ContaCorrente{
		titular:"Pedro Makoski",
		numeroAgencia: 589,
		numeroConta: 123456,
		saldo: 1700.90,
	}

	fmt.Println(contaDoPedro.saldo)

	contaDoPedro.Depositar(-100)

	fmt.Println(contaDoPedro.saldo)
}