package main

import (
	"fmt"
	"banco/contas"
	"banco/clientes"
)

func main() {
	clienteBruno := clientes.Titular{"Bruno Miranda","111.111.111-11", "Ator"}

	contaDoBruno := contas.ContaCorrente{
		Titular:clienteBruno,
		NumeroAgencia:123,
		NumeroConta: 12456,
	}
	contaDoBruno.Depositar(100);

	fmt.Println(contaDoBruno.ObterSaldo())
}