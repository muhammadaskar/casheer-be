package category

type Service interface {
	FindAll() ([]Category, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Category, error) {
	category, err := s.repository.FindAll()
	if err != nil {
		return category, err
	}
	return category, nil
}
