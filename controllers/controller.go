package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermefparra/api-gin-go/database"
	"github.com/guilhermefparra/api-gin-go/models"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// ExibeTodosAlunos godoc
// @Summary      Exibe todos os alunos cadastrados.
// @Description  Rota para exibir todos os alunos cadastrados.
// @Tags         alunos
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /alunos [get]
func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

// BuscaAlunoPorId godoc
// @Summary      Busca aluno por id.
// @Description  Rota que retorna o aluno com aquele id.
// @Tags         alunos
// @Accept       json
// @Produce      json
// @Param			id	path		int	true	"ID do aluno"
// @Success      200  {object}   models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /alunos/{id} [get]
func BuscaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	result := database.DB.First(&aluno, id)

	if result.Error != nil || aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

// CriaNovoAluno godoc
// @Summary      Cria um novo aluno.
// @Description  Rota para criar um novo aluno.
// @Tags         alunos
// @Accept       json
// @Produce      json
// @Param			aluno	body		models.Aluno	true	"Add account"
// @Success      200  {object}   models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /alunos [post]
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&aluno)

	c.JSON(http.StatusOK, aluno)
}

// DeletaAluno godoc
//
//	@Summary		Deleta um aluno
//	@Description	Rota para deletar um aluno
//	@Tags			alunos
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"ID do aluno"	Format(int64)
//	@Success		204	{object}	models.Aluno
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/alunos/{id} [delete]
func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, aluno)
}

// EditaAluno godoc
//
//	@Summary		Edita os dados de um aluno
//	@Description	Rota que edita os dados de um aluno pelo seu ID
//	@Tags			alunos
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"ID do aluno"
//	@Param			alunos	body		models.Aluno	true	"Edita aluno"
//	@Success		200		{object}	models.Aluno
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/alunos/{id} [patch]
func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorCPF godoc
// @Summary      Busca aluno pelo CPF.
// @Description  Rota que retorna o aluno com aquele CPF.
// @Tags         alunos
// @Accept       json
// @Produce      json
// @Param			cpf	path		int	true	"CPF do aluno"
// @Success      200  {object}   models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /alunos/{cpf} [get]
func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")

	result := database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if result.Error != nil || aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno

	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
