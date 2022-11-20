package presenter

import (
	"github.com/EricMesquita/tcc-golang/internal/app/domain/matricula"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/outbound"
)

func CreateMatricula(matricula *matricula.Matricula) *outbound.CreateMatriculaResponse {
	return &outbound.CreateMatriculaResponse{
		Id:                matricula.Id,
		AlunoId:           matricula.AlunoId,
		MateriaId:         matricula.MateriaID,
		DescricaoSemestre: matricula.DescricaoSemestre,
		Finalizado:        matricula.Finalizado,
	}
}
