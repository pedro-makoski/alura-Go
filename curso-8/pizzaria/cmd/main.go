package main

import (
	"pizzaria/internal/data"
	"pizzaria/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadPizzas()
	router := gin.Default()

	router.GET("/pizzas", handler.GetPizzas)
	router.GET("/pizzas/:id", handler.GetPizzaById)

	router.POST("/pizzas", handler.PostPizzas)

	router.DELETE("/pizzas/:id", handler.DeletePizzaById)

	router.PUT("/pizzas/:id", handler.UpdatePizzaById)
	router.POST("/pizzas/:id/reviews", handler.PostReview)
	router.Run(":8000")
}
