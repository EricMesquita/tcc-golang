package aluno

import (
	"context"
	"github.com/EricMesquita/tcc-golang/internal/app/domain/matricula"
)

type Repository interface {
	Create(ctx context.Context, book *Aluno) error
	ReadById(ctx context.Context, id int64) (*Aluno, error)
	ReadAndMatriculaById(ctx context.Context, id int64) (*Aluno, *[]*matricula.Matricula, error)
}
