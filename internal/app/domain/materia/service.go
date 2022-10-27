package materia

import (
	"context"
)

type Service struct {
	repository Repository
}

func (s *Service) Create(ctx context.Context, aluno *Materia) error {
	return s.repository.Create(ctx, aluno)
}
