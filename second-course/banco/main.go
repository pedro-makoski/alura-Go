package main

import (
	"fmt"
	"banco/contas"
	"banco/clientes"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {
	clienteBruno := clientes.Titular{"Bruno Miranda","111.111.111-11", "Ator"}

	contaDoBruno := contas.ContaCorrente{
		Titular:clienteBruno,
		NumeroAgencia:123,
		NumeroConta: 12456,
	}
	contaDoBruno.Depositar(100);

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)
	PagarBoleto(&contaDoBruno, 20)
	fmt.Println(contaDoDenis.ObterSaldo())
	fmt.Println(contaDoBruno.ObterSaldo())
}