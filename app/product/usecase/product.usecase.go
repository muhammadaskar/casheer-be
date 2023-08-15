package usecase

import (
	"time"

	"github.com/muhammadaskar/casheer-be/app/product"
	"github.com/muhammadaskar/casheer-be/app/product/repository/mysql"
	"github.com/muhammadaskar/casheer-be/domains"
)

type ProductUseCase interface {
	FindAll(page int) ([]domains.CustomResult, error)
	Create(input product.CreateInput) (domains.Product, error)
}

type usecase struct {
	repository mysql.Repository
}

func NewUseCase(repository mysql.Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) FindAll(page int) ([]domains.CustomResult, error) {
	product, err := u.repository.FindAll(page)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (u *usecase) Create(input product.CreateInput) (domains.Product, error) {
	product := domains.Product{}
	product.Name = input.Name
	product.Price = input.Price
	product.Quantity = input.Quantity
	product.Description = input.Description
	product.UserID = input.User.ID
	product.CategoryID = input.CategoryID

	now := time.Now()

	// Get the date by setting the time part to 00:00:00
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	product.Image = "https://img-global.cpcdn.com/recipes/93a46a53e22256b8/680x482cq70/songkolo-bagadang-ketan-serundeng-foto-resep-utama.jpg"
	product.EntryAt = today
	// product.ExpiredAt = input.ExpiredAt

	newProduct, err := u.repository.Create(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}
