package main

import (
	"github.com/pedro-makoski/alura-Go/curso-5-gin/models"
	"github.com/pedro-makoski/alura-Go/curso-5-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{"Pedro Makoski", "00000000000", "4700000000"},
		{"Iksokam Ordep", "1234578912", "4300000000"},
	}
	routes.HandleRequests()
}
