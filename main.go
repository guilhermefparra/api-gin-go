package main

import (
	"github.com/guilhermefparra/api-gin-go/database"
	"github.com/guilhermefparra/api-gin-go/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequest()
}
