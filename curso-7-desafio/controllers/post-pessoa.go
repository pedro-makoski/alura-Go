package controllers

import (
	"fmt"
	"pessoas/models"

	"github.com/gin-gonic/gin"
)

const ROUTE_POST_PESSOA = "/pessoas"

func PostPessoa(c *gin.Context) {
	var newPessoa models.Pessoa
	if err := c.ShouldBindJSON(&newPessoa); err != nil {
		c.JSON(400, gin.H{"erro": "Pessoa invalida"})
		return
	}

	fmt.Println(newPessoa)
	models.Pessoas = append(models.Pessoas, newPessoa)
	c.JSON(200, gin.H{
		"pessoa-criada": newPessoa,
		"foi-imprimido": true,
	})
}
