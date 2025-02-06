package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Pedro Makoski"
	idade := 15
	versao := 1.1

	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variável nome é ", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável idade é ", reflect.TypeOf(idade))
	fmt.Println("O tipo da variável versao é ", reflect.TypeOf(versao))
}