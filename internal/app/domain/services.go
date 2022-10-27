package domain

import (
	dbAdapters "github.com/EricMesquita/tcc-golang/internal/app/adapters/database"
	"github.com/EricMesquita/tcc-golang/internal/app/adapters/database/postgres"
	"github.com/EricMesquita/tcc-golang/internal/app/domain/aluno"
	"github.com/EricMesquita/tcc-golang/internal/app/domain/book"
	"github.com/EricMesquita/tcc-golang/internal/app/domain/health"
	"github.com/EricMesquita/tcc-golang/internal/app/domain/materia"
	"github.com/EricMesquita/tcc-golang/internal/infra/database"
)

type Services struct {
	Materia *materia.Service
	Aluno   *aluno.Service
	Book    *book.Service
	Health  *health.Service
}

func NewServices(dbs *database.Databases) *Services {
	return &Services{
		Materia: materia.NewService(postgres.NewMateriaRepository(dbs)),
		Aluno:   aluno.NewService(postgres.NewAlunoRepository(dbs)),
		Book:    book.NewService(postgres.NewBookRepository(dbs)),
		Health:  health.NewService(dbAdapters.NewHealthRepository(dbs)),
	}
}
