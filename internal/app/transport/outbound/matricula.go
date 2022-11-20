package outbound

type (
	CreateMatriculaResponse struct {
		Id                int64  `json:"id"`
		AlunoId           int64  `json:"aluno_id"`
		MateriaId         int64  `json:"materia_id"`
		DescricaoSemestre string `json:"descricao_semestre"`
		Finalizado        bool   `json:"finalizado"`
	}
)
