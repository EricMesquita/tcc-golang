package presenter

import (
	"github.com/EricMesquita/tcc-golang/internal/app/domain/materia"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/outbound"
)

func CreateMateria(materia *materia.Materia) *outbound.CreateMateriaResponse {
	return &outbound.CreateMateriaResponse{
		Id:                 materia.Id,
		Nome:               materia.Nome,
		Capacidade:         materia.Capacidade,
		MateriaDependencia: materia.MateriaDependencia,
	}
}
