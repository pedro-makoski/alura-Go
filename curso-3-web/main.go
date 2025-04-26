package main

import (
	"fmt"
	"net/http"
	"alura-Go/sites/routes"
)


func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
	fmt.Println("Começou")
}