package main

import (
	"api/rest/models"
	"api/rest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "História 1"},
		{Id: 2, Nome: "Nome 2", Historia: "História 2"},
		{Id: 3, Nome: "Nome 3", Historia: "História 3"},
		{Id: 4, Nome: "Nome 4", Historia: "História 4"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
