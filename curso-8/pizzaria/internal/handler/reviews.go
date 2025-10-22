package handler

import (
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostReview(c *gin.Context) {
	pizzaId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var newReview models.Review
	if err := c.ShouldBindJSON(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	for i, pizza := range data.GetPizzas() {
		if pizza.ID == pizzaId {
			pizza.Reviews = append(pizza.Reviews, newReview)

			err := data.UpdatePizza(i, pizza, pizzaId)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, pizza)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})
}
