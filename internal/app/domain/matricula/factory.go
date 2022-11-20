package matricula

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
