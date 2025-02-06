package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		comando := getComando()

		switch comando {
			case 1:	
				iniciarMonitoramento()
			case 2:
				fmt.Println("Exibindo logs...")
			case 0:
				fmt.Println("Saindo do programa")
				os.Exit(0)
			default:
				fmt.Println("Não conheço este comando")
				os.Exit(-1)
		}
	}
	
}

func exibeIntroducao(){
	nome := "Pedro Makoski"
	versao := 1.1

	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - sair do programa")
}

func getComando() int {
	var comandoLido int 
	fmt.Print("Insira o comando: ")
	fmt.Scan(&comandoLido)

	fmt.Println("O comando escolhido foi:", comandoLido)

	return comandoLido	
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://httpbin.org/status/200"
	resposta, _ := http.Get(site)
	
	if resposta.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", resposta.StatusCode)
	}
}