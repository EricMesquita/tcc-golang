package handler

import (
	"github.com/EricMesquita/tcc-golang/internal/app/domain"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/inbound"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/mapper"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/presenter"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type AlunoHandler struct {
	services *domain.Services
}

func NewAlunoHandler(services *domain.Services) *AlunoHandler {
	return &AlunoHandler{
		services: services,
	}
}

func (h *AlunoHandler) Configure(server *echo.Echo) {
	server.POST("/v1/aluno", h.CreateAluno)
	server.GET("/v1/aluno/:id", h.ReadAlunoById)
	server.GET("/v1/aluno/:id/matricula", h.ReadAlunoAndMatriculaById)
}

func (h *AlunoHandler) CreateAluno(c echo.Context) error {
	request := new(inbound.CreateAlunoRequest)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	aluno := mapper.AlunoFromCreateAlunoRequest(request)
	err := h.services.Aluno.Create(c.Request().Context(), aluno)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	aluno2, err := h.services.Aluno.ReadById(c.Request().Context(), aluno.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, presenter.CreateAluno(aluno2))
}

func (h *AlunoHandler) ReadAlunoById(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	aluno, err := h.services.Aluno.ReadById(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if aluno == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, presenter.CreateAluno(aluno))
}

func (h *AlunoHandler) ReadAlunoAndMatriculaById(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	aluno, matriculas, err := h.services.Aluno.ReadAndMatriculaById(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if aluno == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, presenter.AlunoAndMatricula(aluno, matriculas))
}
