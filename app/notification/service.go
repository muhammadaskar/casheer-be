package notification

type Service interface {
	FindAll() ([]Notification, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Notification, error) {
	notification, err := s.repository.FindAll()
	if err != nil {
		return notification, err
	}

	return notification, nil
}
