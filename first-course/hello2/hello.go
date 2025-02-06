package main

import (
	"fmt"
)

func main() {
	nome := "Pedro Makoski"
	versao := 1.1

	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - sair do programa")

	var comando int 
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi:", comando)
}
