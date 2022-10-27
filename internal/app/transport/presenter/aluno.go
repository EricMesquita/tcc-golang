package presenter

import (
	"github.com/EricMesquita/tcc-golang/internal/app/domain/aluno"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/outbound"
)

func CreateAluno(aluno *aluno.Aluno) *outbound.CreateAlunoResponse {
	return &outbound.CreateAlunoResponse{
		Id:        aluno.Id,
		Nome:      aluno.Nome,
		Documento: aluno.Documento,
	}
}
