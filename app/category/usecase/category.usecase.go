package usecase

import (
	"errors"

	"github.com/muhammadaskar/casheer-be/app/category"
	"github.com/muhammadaskar/casheer-be/app/category/repository/mysql"
	"github.com/muhammadaskar/casheer-be/domains"
)

type CategoryUseCase interface {
	FindAll() ([]domains.Category, error)
	FindById(input category.GetCategoryInputID) (domains.Category, error)
	Create(input category.CreateCategoryInput) (domains.Category, error)
	UpdateCategory(inputId category.GetCategoryInputID, input category.CreateCategoryInput) (domains.Category, error)
}

type usecase struct {
	repository mysql.Repository
}

func NewUseCase(repository mysql.Repository) *usecase {
	return &usecase{repository}
}

func (s *usecase) FindAll() ([]domains.Category, error) {
	category, err := s.repository.FindAll()
	if err != nil {
		return category, err
	}
	return category, nil
}

func (s *usecase) FindById(input category.GetCategoryInputID) (domains.Category, error) {
	category, err := s.repository.FindById(input.ID)
	if err != nil {
		return category, err
	}

	if category.ID == 0 {
		return category, errors.New("category not found")
	}

	return category, nil
}

func (s *usecase) Create(input category.CreateCategoryInput) (domains.Category, error) {
	category := domains.Category{}
	category.Name = input.Name

	newCategory, err := s.repository.Save(category)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *usecase) UpdateCategory(inputID category.GetCategoryInputID, input category.CreateCategoryInput) (domains.Category, error) {
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
