package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"strconv"
)

const MONITORAMENTOS = 10
const DELAY_IN_MINUTES = 5

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		comando := getComando()

		switch comando {
			case 1:	
				iniciarMonitoramento()
			case 2:
				imprimeLogs()
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
	sites := lerSitesDoArquivo()
	fmt.Println("Monitorando...")

	for i:=0; i < MONITORAMENTOS; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
			fmt.Println("")
		}

		time.Sleep(DELAY_IN_MINUTES * time.Minute)
	}
}

func testSite(site string) {
	resposta, err := http.Get(site)
	
	if err != nil {
		fmt.Println("Ocorreu o erro", err)
		return 
	}

	if resposta.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", resposta.StatusCode)
	}

	registrarLog(site, resposta.StatusCode == 200)
} 

func lerSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	defer arquivo.Close()

	if err != nil {
		fmt.Println("Ocorreu o erro", err)
		return sites 
	}
	
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites,linha)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Ocorreu o erro", err)
			return sites 
		}
	}

	return sites
}

func registrarLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer arquivo.Close()
	if err != nil {
		fmt.Println("Ocorreu o erro", err)
		return
	}

	_, err = arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online:" + strconv.FormatBool(status) + "\n")
	if err != nil {
		fmt.Println("Ocorreu o erro", err)
		return
	}
}

func imprimeLogs() {
	fmt.Println("Exibindo logs...")

	arquivo, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu o erro", err)
		return
	}

	fmt.Println(string(arquivo))
	
}