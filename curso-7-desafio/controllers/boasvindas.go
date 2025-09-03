package controllers

import "github.com/gin-gonic/gin"

const ROUTES_BOAS_VINDAS = "/greetings"

func BoasVindas(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Boas vindas ao Go",
	})
}
