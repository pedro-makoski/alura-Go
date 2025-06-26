package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pedro-makoski/alura-Go/curso-5-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.Run(":5000")
}
