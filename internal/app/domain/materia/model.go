package materia

type Materia struct {
	Id                 int64
	Nome               string
	Capacidade         int
	MateriaDependencia *int
}
