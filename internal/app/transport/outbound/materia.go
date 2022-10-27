package outbound

type (
	CreateMateriaResponse struct {
		Id                 int64  `json:"id"`
		Nome               string `json:"nome"`
		Capacidade         int    `json:"capacidade"`
		MateriaDependencia *int   `json:"materia_dependencia"`
	}
)
