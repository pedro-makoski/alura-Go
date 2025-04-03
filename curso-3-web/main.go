package main

import (
	"net/http"
	"html/template"
)

type Produto struct {
	Nome string
	Descricao string 
	Preco float64
	Quantidade int	
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)	
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{
			Nome: "Camiseta",
			Descricao: "Azul, Bem bonita",
			Preco: 19,
			Quantidade: 15,
		},
		{
			Nome: "Tênis",
			Descricao: "Confortável",
			Preco: 90,
			Quantidade: 3,
		},
		{"Fone", "Muito bom", 60, 2},
		{"Produto novo", "Muito leal", 1999, 1},
	}

	templates.ExecuteTemplate(w, "index", produtos)
}