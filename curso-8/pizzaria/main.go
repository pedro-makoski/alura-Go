package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{}

func main() {
	LoadPizzas()
	router := gin.Default()

	router.GET("/pizzas", GetPizzas)
	router.GET("/pizzas/:id", GetPizzaById)

	router.POST("/pizzas", PostPizzas)

	router.DELETE("/pizzas/:id", deletePizzaById)

	router.PUT("/pizzas/:id", updatePizzaById)
	fmt.Println(pizzas)
	router.Run(":8000")
}

func LoadPizzas() {
	file, err := os.Open("dados/pizzas.json")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de pizzas.json")
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Erro ao dessirealizar o JSON")
		return
	}
}

func SavePizza() {
	file, err := os.Create("dados/pizzas.json")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de pizzas.json")
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent(" ", "  ")
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Error encoding JSON: ", err.Error())
	}
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

	newPizza.ID = len(pizzas) + 1

	pizzas = append(pizzas, newPizza)
	SavePizza()
	c.JSON(202, gin.H{
		"item-criado": newPizza,
		"foi-criado":  true,
	})
}

func deletePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	for idx, pizza := range pizzas {
		if pizza.ID == id {
			pizzas = append(pizzas[:idx], pizzas[idx+1:]...)
			SavePizza()
			c.JSON(200, gin.H{"message": "pizza deletada"})
			return
		}
	}

	c.JSON(404, gin.H{"message": "nenhuma pizza achada"})
}

func updatePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	var updatedPizza models.Pizza
	if err := c.ShouldBindJSON(&updatedPizza); err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	for idx, pizza := range pizzas {
		if pizza.ID == id {
			pizzas[idx] = updatedPizza
			pizzas[idx].ID = id
			SavePizza()
			c.JSON(200, pizzas[idx])
			return
		}
	}
	c.JSON(404, gin.H{"message": "nenhuma pizza achada"})
}
