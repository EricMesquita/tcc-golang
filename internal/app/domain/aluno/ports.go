package aluno

import "context"

type Repository interface {
	Create(ctx context.Context, book *Aluno) error
	ReadById(ctx context.Context, id int64) (*Aluno, error)
}
