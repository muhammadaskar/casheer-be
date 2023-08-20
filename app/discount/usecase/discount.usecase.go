package usecase

import (
	"errors"

	"github.com/muhammadaskar/casheer-be/app/discount"
	"github.com/muhammadaskar/casheer-be/app/discount/repository/mysql"
	"github.com/muhammadaskar/casheer-be/domains"
)

type DiscountUseCase interface {
	FindByID() (domains.Discount, error)
	Create(input discount.CreateInput) (domains.Discount, error)
	Update(input discount.CreateInput) (domains.Discount, error)
}

type usecase struct {
	repository mysql.Repository
}

func NewUseCase(repository mysql.Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) FindByID() (domains.Discount, error) {
	discount, err := u.repository.FindById()
	if err != nil {
		return discount, err
	}

	if discount.ID == 0 {
		return discount, errors.New("No discount")
	}

	return discount, nil
}

func (u *usecase) Create(input discount.CreateInput) (domains.Discount, error) {
	discount, err := u.repository.FindById()
	if err != nil {
		return discount, err
	}

	if discount.ID != 0 {
		return discount, errors.New("Dicount is available")
	}

	discount = domains.Discount{}
	discount.Discount = input.Discount
	discount.UserID = input.User.ID

	newDiscount, err := u.repository.Create(discount)
	if err != nil {
		return newDiscount, err
	}
	return newDiscount, nil
}

func (u *usecase) Update(input discount.CreateInput) (domains.Discount, error) {
	discount, err := u.repository.FindById()
	if err != nil {
		return discount, err
	}

	if discount.ID == 0 {
		return discount, errors.New("Dicount is not available")
	}

	discount.Discount = input.Discount
	discount.UserID = input.User.ID

	updateDiscount, err := u.repository.Update(discount)
	if err != nil {
		return updateDiscount, err
	}
	return updateDiscount, nil
}
