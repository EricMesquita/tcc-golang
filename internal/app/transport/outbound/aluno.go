package outbound

type (
	CreateAlunoResponse struct {
		Id        int64  `json:"id"`
		Nome      string `json:"nome"`
		Documento string `json:"documento"`
	}

	AlunoAndMatriculasResponse struct {
		Id         int64                      `json:"id"`
		Nome       string                     `json:"nome"`
		Documento  string                     `json:"documento"`
		Matriculas []*CreateMatriculaResponse `json:"matriculas"`
	}
)
