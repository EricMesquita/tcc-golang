package inbound

type (
	CreateMateriaRequest struct {
		Nome               string `json:"nome"`
		Capacidade         int    `json:"capacidade"`
		MateriaDependencia *int   `json:"materia_dependencia"`
	}
)
