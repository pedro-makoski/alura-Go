package handler

import (
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPizzas(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pizzas": data.GetPizzas(),
	})
}

func GetPizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Digite um número que faça sentido, não o texto: " + idParam,
		})
		return
	}

	for _, pizza := range data.GetPizzas() {
		if pizza.ID == id {
			c.JSON(http.StatusOK, pizza)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"erro": "Pizza não encontrada",
	})
}

func PostPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	newPizza.ID = len(data.GetPizzas()) + 1

	err := data.AddPiza(newPizza)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"item-criado": newPizza,
		"foi-criado":  true,
	})
}

func DeletePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	for idx, pizza := range data.GetPizzas() {
		if pizza.ID == id {
			data.RemovPizza(idx)
			c.JSON(http.StatusAccepted, gin.H{"message": "pizza deletada"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "nenhuma pizza achada"})
}

func UpdatePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	var updatedPizza models.Pizza
	if err := c.ShouldBindJSON(&updatedPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	for idx, pizza := range data.GetPizzas() {
		if pizza.ID == id {
			err := data.UpdatePizza(idx, updatedPizza, id)

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusAccepted, data.GetPizzas()[idx])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "nenhuma pizza achada"})
}
