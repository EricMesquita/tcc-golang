package inbound

type (
	CreateAlunoRequest struct {
		Nome      string `json:"nome"`
		Documento string `json:"documento"`
	}
)
