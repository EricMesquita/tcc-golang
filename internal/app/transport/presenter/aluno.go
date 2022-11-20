package presenter

import (
	"github.com/EricMesquita/tcc-golang/internal/app/domain/aluno"
	"github.com/EricMesquita/tcc-golang/internal/app/domain/matricula"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/outbound"
)

func CreateAluno(aluno *aluno.Aluno) *outbound.CreateAlunoResponse {
	return &outbound.CreateAlunoResponse{
		Id:        aluno.Id,
		Nome:      aluno.Nome,
		Documento: aluno.Documento,
	}
}

func AlunoAndMatricula(aluno *aluno.Aluno, matriculas *[]*matricula.Matricula) *outbound.AlunoAndMatriculasResponse {
	a := &outbound.AlunoAndMatriculasResponse{
		Id:        aluno.Id,
		Nome:      aluno.Nome,
		Documento: aluno.Documento,
	}
	if matriculas != nil {
		var rs []*outbound.CreateMatriculaResponse

		for _, response := range *matriculas {
			m := outbound.CreateMatriculaResponse{
				Id:                response.Id,
				AlunoId:           response.AlunoId,
				MateriaId:         response.MateriaID,
				DescricaoSemestre: response.DescricaoSemestre,
				Finalizado:        response.Finalizado,
			}
			rs = append(rs, &m)
		}
		a.Matriculas = rs
	}

	return a
}
