package mapper

import (
	"github.com/EricMesquita/tcc-golang/internal/app/domain/matricula"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/inbound"
)

func MatriculaFromCreateMatriculaRequest(request *inbound.CreateMatriculaRequest) *matricula.Matricula {
	return &matricula.Matricula{
		AlunoId:           request.AlunoId,
		MateriaID:         request.MateriaId,
		DescricaoSemestre: request.DescricaoSemestre,
		Finalizado:        request.Finalizado,
	}
}
