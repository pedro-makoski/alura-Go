package main

import (
	"github.com/pedro-makoski/alura-Go/curso-6-gin/database"
	"github.com/pedro-makoski/alura-Go/curso-6-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
