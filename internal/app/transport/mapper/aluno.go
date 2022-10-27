package mapper

import (
	"github.com/EricMesquita/tcc-golang/internal/app/domain/aluno"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/inbound"
)

func AlunoFromCreateAlunoRequest(request *inbound.CreateAlunoRequest) *aluno.Aluno {
	return &aluno.Aluno{
		Nome:      request.Nome,
		Documento: request.Documento,
	}
}
