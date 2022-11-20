package matricula

import "context"

type Repository interface {
	Create(ctx context.Context, matricula *Matricula) error
}
