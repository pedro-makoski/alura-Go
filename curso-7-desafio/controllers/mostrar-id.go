package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const ID_NAME = "id"
const MOSTRAR_ID_ROUTE = "/mostrar-id/:" + ID_NAME

func MostrarId(c *gin.Context) {
	idParam := c.Param(ID_NAME)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "neste programa os ids são números"})
		return
	}

	c.JSON(200, gin.H{"message": gerarIdString(id)})
}

func gerarIdString(id int) string {
	return "ID recebido: " + strconv.Itoa(id)
}
