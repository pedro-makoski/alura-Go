package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Pizza struct {
	ID    int
	Nome  string
	Preco float64
}

func main() {
	router := gin.Default()

	pizzas := []Pizza{
		{
			ID:    1,
			Nome:  "Marguerita",
			Preco: 30,
		},
		{
			ID:    2,
			Nome:  "Toscana",
			Preco: 30,
		},
		{
			ID:    3,
			Nome:  "Atum Com Quiejo",
			Preco: 29.1,
		},
	}

	router.GET("/pizzas", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"pizzas": pizzas,
		})
	})

	fmt.Println(pizzas)
	router.Run(":8000")
}
