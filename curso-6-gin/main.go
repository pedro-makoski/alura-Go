package main

import (
	"github.com/pedro-makoski/alura-Go/curso-5-gin/database"
	"github.com/pedro-makoski/alura-Go/curso-5-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
