package postgres

import (
	"context"
	"github.com/EricMesquita/tcc-golang/internal/app/domain/matricula"
	"github.com/EricMesquita/tcc-golang/internal/infra/database"
)

type MatriculaRepository struct {
	dbs *database.Databases
}

func NewMatriculaRepository(dbs *database.Databases) *MatriculaRepository {
	return &MatriculaRepository{
		dbs: dbs,
	}
}

func (r *MatriculaRepository) Create(ctx context.Context, matricula *matricula.Matricula) error {
	r.dbs.Write.Connection().QueryRowContext(ctx, createMatriculaSql, matricula.MateriaID, matricula.AlunoId, matricula.DescricaoSemestre, matricula.Finalizado).Scan(&matricula.Id)
	return nil
}

const (
	createMatriculaSql = "insert into matricula (id, materia_id, aluno_id, descricao_semestre, finalizado) values ((SELECT nextval('seq_matricula')), $1, $2, $3, $4 ) RETURNING id"
)
