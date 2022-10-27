package materia

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
