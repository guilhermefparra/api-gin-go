package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/guilhermefparra/api-gin-go/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("alunos", controller.ExibeTodosAlunos)
	
	r.Run()
}