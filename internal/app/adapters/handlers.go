package adapters

import (
	"github.com/EricMesquita/tcc-golang/internal/app/adapters/handler"
	"github.com/EricMesquita/tcc-golang/internal/app/domain"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	aluno   *handler.AlunoHandler
	book    *handler.BookHandler
	health  *handler.HealthHandler
	materia *handler.MateriaHandler
}

func NewHandlers(services *domain.Services) *Handlers {
	return &Handlers{
		aluno:   handler.NewAlunoHandler(services),
		book:    handler.NewBookHandler(services),
		health:  handler.NewHealthHandler(services),
		materia: handler.NewMateriaHandler(services),
	}
}

func (h *Handlers) Configure(server *echo.Echo) {
	h.aluno.Configure(server)
	h.book.Configure(server)
	h.health.Configure(server)
	h.materia.Configure(server)
}
