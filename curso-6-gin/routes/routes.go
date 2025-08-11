package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pedro-makoski/alura-Go/curso-6-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	r.POST("/alunos", controllers.CriaNovoAluno)

	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	r.PATCH("/alunos/:id", controllers.EditaAluno)

	r.Run(":5000")
}
