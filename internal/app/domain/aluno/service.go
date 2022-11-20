package aluno

import (
	"context"
	"github.com/EricMesquita/tcc-golang/internal/app/domain/matricula"
)

type Service struct {
	repository Repository
}

func (s *Service) Create(ctx context.Context, aluno *Aluno) error {
	return s.repository.Create(ctx, aluno)
}

func (s *Service) ReadById(ctx context.Context, id int64) (*Aluno, error) {
	return s.repository.ReadById(ctx, id)
}

func (s *Service) ReadAndMatriculaById(ctx context.Context, id int64) (*Aluno, *[]*matricula.Matricula, error) {
	return s.repository.ReadAndMatriculaById(ctx, id)
}
