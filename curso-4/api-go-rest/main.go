package main

import (
	"api/rest/database"
	"api/rest/routes"
	"fmt"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
