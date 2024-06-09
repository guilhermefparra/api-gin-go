package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermefparra/api-gin-go/models"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}