package main 

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if(len(os.Args) < 3) {
		fmt.Println("Uso: conversor <valores> <unidades>")
		os.Exit(1)
	}

	unidadeOrigem := os.Args[len(os.Args)-1]
	valoresOrigem := os.Args[1:len(os.Args)-1]
	var unidadeDestino string

	switch(unidadeOrigem) {
	case "celsius":
		unidadeDestino = "fahrenheit"
	case "quilometros":
		unidadeDestino = "milhas"
	case "fahrenheit":
		unidadeDestino = "celsius"
	case "milhas":
		unidadeDestino = "quilometros"
	default:
		fmt.Printf("%s não é uma unidade conhecida\n", unidadeOrigem)
		os.Exit(1)
	}

	convertAllValues(valoresOrigem, unidadeDestino, unidadeOrigem)
}

func convertAllValues(valoresOrigem []string, unidadeDestino string, unidadeOrigem string) {
	for i, v := range valoresOrigem {
		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("O valor %s da posicao %d não é um número valido\n", v, i)
			os.Exit(1)
		}

		var valorDestino float64

		switch unidadeOrigem {
		case "celsius":
			valorDestino = valorOrigem*1.8 + 32
		case "quilometros":
			valorDestino = valorOrigem/1.60934
		case "fahrenheit":
			valorDestino = (valorOrigem-32)/1.8
		case "milhas":
			valorDestino = valorOrigem-32*1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
	}
}