package matricula

import (
	"context"
)

type Service struct {
	repository Repository
}

func (s *Service) Create(ctx context.Context, matricula *Matricula) error {
	return s.repository.Create(ctx, matricula)
}
