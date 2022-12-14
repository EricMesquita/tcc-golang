package handler

import (
	"github.com/EricMesquita/tcc-golang/internal/app/domain"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/inbound"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/mapper"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/presenter"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MatriculaHandler struct {
	services *domain.Services
}

func NewMatriculaHandler(services *domain.Services) *MatriculaHandler {
	return &MatriculaHandler{
		services: services,
	}
}

func (h *MatriculaHandler) Configure(server *echo.Echo) {
	server.POST("/v1/matricula", h.CreateMatricula)
}

func (h *MatriculaHandler) CreateMatricula(c echo.Context) error {
	request := new(inbound.CreateMatriculaRequest)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	matricula := mapper.MatriculaFromCreateMatriculaRequest(request)
	err := h.services.Matricula.Create(c.Request().Context(), matricula)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, presenter.CreateMatricula(matricula))
}
