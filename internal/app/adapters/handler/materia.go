package handler

import (
	"github.com/EricMesquita/tcc-golang/internal/app/domain"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/inbound"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/mapper"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/presenter"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MateriaHandler struct {
	services *domain.Services
}

func NewMateriaHandler(services *domain.Services) *MateriaHandler {
	return &MateriaHandler{
		services: services,
	}
}

func (h *MateriaHandler) Configure(server *echo.Echo) {
	server.POST("/v1/materia", h.CreateMateria)
}

func (h *MateriaHandler) CreateMateria(c echo.Context) error {
	request := new(inbound.CreateMateriaRequest)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	materia := mapper.MateriaFromCreateMateriaRequest(request)
	err := h.services.Materia.Create(c.Request().Context(), materia)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, presenter.CreateMateria(materia))
}
