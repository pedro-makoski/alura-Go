package main

import "fmt"

type ContaCorrente struct {
	titular 		string 
	numeroAgencia 	int 
	numeroConta 	int
	saldo 			float64
}

func main() {
	contaDoPedro := ContaCorrente{
		titular:"Pedro Makoski",
		numeroAgencia: 589,
		numeroConta: 123456,
		saldo: 1700.90,
	}
	contaDaBia := ContaCorrente{"Beatriz Makoski", 111, 198, 200}
 	fmt.Println(contaDoPedro)
 	fmt.Println(contaDaBia)
}