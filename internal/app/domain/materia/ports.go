package materia

import "context"

type Repository interface {
	Create(ctx context.Context, materia *Materia) error
}
