package postgres

import (
	"context"
	"github.com/EricMesquita/tcc-golang/internal/app/domain/materia"
	"github.com/EricMesquita/tcc-golang/internal/infra/database"
)

type MateriaRepository struct {
	dbs *database.Databases
}

func NewMateriaRepository(dbs *database.Databases) *MateriaRepository {
	return &MateriaRepository{
		dbs: dbs,
	}
}

func (r *MateriaRepository) Create(ctx context.Context, materia *materia.Materia) error {
	r.dbs.Write.Connection().QueryRowContext(ctx, createMateriaSql, materia.Nome, materia.Capacidade).Scan(&materia.Id)
	if materia.MateriaDependencia != nil {
		r.dbs.Write.Connection().QueryRowContext(ctx, createMateriaDepSql, materia.Id, materia.MateriaDependencia).Scan(&materia.Id)
	}
	return nil
}

const (
	createMateriaSql    = "insert into materia (id, nome, capacidade) values ((SELECT nextval('seq_materia')), $1, $2) RETURNING id"
	createMateriaDepSql = "insert into materia_dependencia (id, materia_id, materia_dependencia_id) values ((SELECT nextval('seq_materia_dependencia')), $1, $2) RETURNING id"
)
