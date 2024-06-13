package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	controller "github.com/guilhermefparra/api-gin-go/controllers"
	"github.com/guilhermefparra/api-gin-go/database"
	"github.com/guilhermefparra/api-gin-go/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()

	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "11111111111", RG: "111111111"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()

	r.GET("/alunos", controller.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Status code informando não está correto!")
}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()

	r.GET("/alunos/:id", controller.BuscaAlunoPorId)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var alunoMock models.Aluno

	json.Unmarshal(response.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "11111111111", alunoMock.CPF, "Os CPFs devem ser iguais")
	assert.Equal(t, "111111111", alunoMock.RG, "Os RGs devem ser iguais")
	assert.Equal(t, http.StatusOK, response.Code, "Status code informando não está correto!")
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()

	r.GET("/alunos/cpf/:cpf", controller.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/11111111111", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var alunoMock models.Aluno

	json.Unmarshal(response.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "11111111111", alunoMock.CPF, "Os CPFs devem ser iguais")
	assert.Equal(t, "111111111", alunoMock.RG, "Os RGs devem ser iguais")
	assert.Equal(t, http.StatusOK, response.Code, "Status code informando não está correto!")
}

func TestEditaUmAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controller.EditaAluno)

	aluno := models.Aluno{Nome: "Nome do Aluno Teste Editado", CPF: "11111111111", RG: "111111111"}
	valorJson, _ := json.Marshal(aluno)
	path := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(valorJson))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var alunoMockadoEditado models.Aluno

	json.Unmarshal(response.Body.Bytes(), &alunoMockadoEditado)
	assert.Equal(t, "Nome do Aluno Teste Editado", alunoMockadoEditado.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "11111111111", alunoMockadoEditado.CPF, "Os CPFs devem ser iguais")
	assert.Equal(t, "111111111", alunoMockadoEditado.RG, "Os RGs devem ser iguais")
	assert.Equal(t, http.StatusOK, response.Code, "Status code informando não está correto!")
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controller.DeletaAluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Status code informando não está correto!")
}

func TestCriaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()

	r.POST("/alunos", controller.CriaNovoAluno)
	aluno := models.Aluno{Nome: "Nome do Aluno Teste Editado", CPF: "11111111111", RG: "111111111"}
	valorJson, _ := json.Marshal(aluno)

	req, _ := http.NewRequest("POST", "/alunos", bytes.NewBuffer(valorJson))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	var alunoMockadoCriado models.Aluno

	json.Unmarshal(response.Body.Bytes(), &alunoMockadoCriado)
	assert.Equal(t, "Nome do Aluno Teste Editado", alunoMockadoCriado.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "11111111111", alunoMockadoCriado.CPF, "Os CPFs devem ser iguais")
	assert.Equal(t, "111111111", alunoMockadoCriado.RG, "Os RGs devem ser iguais")
	assert.Equal(t, http.StatusOK, response.Code, "Status code informando não está correto!")
}
