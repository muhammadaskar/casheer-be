package category

import "errors"

type Service interface {
	FindAll() ([]Category, error)
	FindById(input GetCategoryInputID) (Category, error)
	Create(input CreateCategoryInput) (Category, error)
	UpdateCategory(inputId GetCategoryInputID, input CreateCategoryInput) (Category, error)
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

func (s *service) FindById(input GetCategoryInputID) (Category, error) {
	category, err := s.repository.FindById(input.ID)
	if err != nil {
		return category, err
	}

	if category.ID == 0 {
		return category, errors.New("category not found")
	}

	return category, nil
}

func (s *service) Create(input CreateCategoryInput) (Category, error) {
	category := Category{}
	category.Name = input.Name

	newCategory, err := s.repository.Save(category)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *service) UpdateCategory(inputID GetCategoryInputID, input CreateCategoryInput) (Category, error) {
	category, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return category, err
	}

	category.Name = input.Name

	updateCategory, err := s.repository.Update(category)
	if err != nil {
		return updateCategory, err
	}

	return updateCategory, nil
}
