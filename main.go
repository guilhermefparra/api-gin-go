package main

import (
	"github.com/guilhermefparra/api-gin-go/models"
	"github.com/guilhermefparra/api-gin-go/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Guilherme", CPF: "00000000", RG: "32131231"},
		{Nome: "Guilherme2", CPF: "20000000", RG: "1111111"},
	}
	routes.HandleRequest()
}