package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	palavras := os.Args[1:]
	estatisticas := colherEstatisticas(palavras)
	imprimir(estatisticas)
}

func colherEstatisticas(palavras []string) map[string]int {
	estatisticas := make(map[string]int)

	for _, palavra := range palavras {
		inicial := palavra[0]
		inicialUpper := strings.ToUpper(string(inicial))
		_, encontrado := estatisticas[inicialUpper]
		if encontrado {
			estatisticas[inicialUpper]++
			continue 
		}

		estatisticas[inicialUpper] = 1
	}

	return estatisticas
}

func imprimir(estatisticas map[string]int) {
	fmt.Println("Contagem de palavras iniciadas em cada letra:")

	for inicial, frequencia := range estatisticas {
		fmt.Printf("%s tem %d de frequência\n", inicial, frequencia)
	}
}