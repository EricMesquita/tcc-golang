package postgres

import (
	"context"
	"database/sql"
	"github.com/EricMesquita/tcc-golang/internal/app/domain/aluno"
	"github.com/EricMesquita/tcc-golang/internal/app/domain/matricula"
	"github.com/EricMesquita/tcc-golang/internal/infra/database"
)

type AlunoRepository struct {
	dbs *database.Databases
}

func NewAlunoRepository(dbs *database.Databases) *AlunoRepository {
	return &AlunoRepository{
		dbs: dbs,
	}
}

func (r *AlunoRepository) Create(ctx context.Context, aluno *aluno.Aluno) error {
	return r.dbs.Write.Connection().QueryRowContext(ctx, createAlunoSql, aluno.Nome, aluno.Documento).Scan(&aluno.Id)
}

func (r *AlunoRepository) ReadById(ctx context.Context, id int64) (*aluno.Aluno, error) {
	var model aluno.Aluno
	err := r.dbs.Read.Connection().QueryRowContext(ctx, readAlunoByIdSql, id).
		Scan(&model.Nome, &model.Documento)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	model.Id = id
	return &model, nil
}

func (r *AlunoRepository) ReadAndMatriculaById(ctx context.Context, id int64) (*aluno.Aluno, *[]*matricula.Matricula, error) {
	var model aluno.Aluno
	err := r.dbs.Read.Connection().QueryRowContext(ctx, readAlunoByIdSql, id).
		Scan(&model.Nome, &model.Documento)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, nil
		}

		return nil, nil, err
	}

	model.Id = id

	rows, err := r.dbs.Read.Connection().Query(readMatriculaByIdSql, id)

	if err != nil {
		return nil, nil, err
	}
	var matriculas []*matricula.Matricula

	defer rows.Close()

	for rows.Next() {
		var m matricula.Matricula
		if err := rows.Scan(&m.Id, &m.AlunoId, &m.MateriaID, &m.DescricaoSemestre, &m.Finalizado); err != nil {
			return nil, nil, err
		}
		matriculas = append(matriculas, &m)
	}

	return &model, &matriculas, nil
}

const (
	createAlunoSql       = "insert into aluno (id, nome, documento) values ((SELECT nextval('seq_aluno')), $1, $2) RETURNING id"
	readAlunoByIdSql     = "select nome, documento from aluno where id = $1"
	readMatriculaByIdSql = "select id, aluno_id, materia_id, descricao_semestre, finalizado from matricula where aluno_id = $1"
)
