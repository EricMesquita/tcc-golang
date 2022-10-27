package mapper

import (
	"github.com/EricMesquita/tcc-golang/internal/app/domain/materia"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/inbound"
)

func MateriaFromCreateMateriaRequest(request *inbound.CreateMateriaRequest) *materia.Materia {
	return &materia.Materia{
		Nome:               request.Nome,
		Capacidade:         request.Capacidade,
		MateriaDependencia: request.MateriaDependencia,
	}
}
