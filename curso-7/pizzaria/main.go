package main

import (
	"fmt"
	"pizzaria/models"

	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{
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

func main() {
	router := gin.Default()

	router.GET("/pizzas", GetPizzas)

	fmt.Println(pizzas)
	router.Run(":8000")
}

func GetPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}
