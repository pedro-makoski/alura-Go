package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pedro-makoski/alura-Go/curso-5-gin/models"
)

func ExibeTodosOsAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}
