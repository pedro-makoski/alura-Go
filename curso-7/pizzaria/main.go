package main

import (
	"fmt"
	"pizzaria/models"
	"strconv"

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
	router.GET("/pizzas/:id", GetPizzaById)

	router.POST("/pizzas", PostPizzas)

	fmt.Println(pizzas)
	router.Run(":8000")
}

func GetPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func GetPizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Digite um número que faça sentido, não o texto: " + idParam,
		})
		return
	}

	for _, pizza := range pizzas {
		if pizza.ID == id {
			c.JSON(200, pizza)
			return
		}
	}

	c.JSON(404, gin.H{
		"erro": "Pizza não encontrada",
	})
}

func PostPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	pizzas = append(pizzas, newPizza)
	c.JSON(202, gin.H{
		"item-criado": newPizza,
		"foi-criado":  true,
	})
}
