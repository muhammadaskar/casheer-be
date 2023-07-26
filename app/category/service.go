package category

import "errors"

type Service interface {
	FindAll() ([]Category, error)
	FindById(ID GetCategoryInputID) (Category, error)
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

func (s *service) FindById(ID GetCategoryInputID) (Category, error) {
	category, err := s.repository.FindById(ID)
	if err != nil {
		return category, err
	}

	if category.ID == 0 {
		return category, errors.New("category not found")
	}

	return category, nil
}
