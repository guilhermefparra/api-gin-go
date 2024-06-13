package routes

import (
	docs "github.com/api-gin-go/docs"
	"github.com/gin-gonic/gin"
	controller "github.com/guilhermefparra/api-gin-go/controllers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/alunos", controller.ExibeTodosAlunos)
	r.GET("/alunos/:id", controller.BuscaAlunoPorId)
	r.GET("/alunos/cpf/:cpf", controller.BuscaAlunoPorCPF)

	r.GET("/index", controller.ExibePaginaIndex)

	r.POST("/alunos", controller.CriaNovoAluno)

	r.DELETE("/alunos/:id", controller.DeletaAluno)

	r.PATCH("/alunos/:id", controller.EditaAluno)

	r.NoRoute(controller.RotaNaoEncontrada)

	r.Run()
}
