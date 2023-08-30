package usecase

import (
	"errors"
	"math"
	"time"

	"github.com/muhammadaskar/casheer-be/app/product"
	"github.com/muhammadaskar/casheer-be/app/product/repository/mysql"
	"github.com/muhammadaskar/casheer-be/domains"
)

type ProductUseCase interface {
	FindAll(query product.GetProductsQueryInput) ([]domains.CustomResult, bool, error)
	GetAll() ([]domains.CustomProduct, error)
	CountAll() (int64, error)
	FindById(input product.GetProductDetailInput) (domains.CustomResult, error)
	Create(input product.CreateInput) (domains.Product, error)
	Update(inputID product.GetProductDetailInput, inputData product.CreateInput) (domains.Product, error)
	Delete(input product.GetProductDetailInput) (domains.Product, error)
	UpdateQuantity(inputID product.GetProductDetailInput, inputData product.UpdateQuantity) (domains.Product, error)
}

type usecase struct {
	repository mysql.Repository
}

func NewUseCase(repository mysql.Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) FindAll(query product.GetProductsQueryInput) ([]domains.CustomResult, bool, error) {
	if query.Query != "" {
		product, err := u.repository.FindAll(query.Query, query.Page, query.Limit, true)
		if err != nil {
			return product, true, err
		}

		return product, true, nil
	} else {
		product, err := u.repository.FindAll(query.Query, query.Page, query.Limit, false)
		if err != nil {
			return product, true, err
		}
		totalCount, err := u.CountAll()
		if err != nil {
			return product, true, err
		}

		perPage := query.Limit
		offset := (query.Page - 1) * perPage
		totalPages := int(math.Ceil(float64(totalCount) / float64(perPage)))

		// Hitung nomor halaman saat ini berdasarkan offset dan produk per halaman
		currentPage := (offset / perPage) + 1

		// Periksa apakah Anda berada di halaman terakhir
		isLastPage := currentPage == totalPages

		return product, isLastPage, nil
	}
}

func (u *usecase) GetAll() ([]domains.CustomProduct, error) {
	products, err := u.repository.GetAll()
	if err != nil {
		return products, err
	}

	return products, nil

}

func (u *usecase) CountAll() (int64, error) {
	count, err := u.repository.Count()
	if err != nil {
		return count, err
	}
	return count, nil
}

func (u *usecase) FindById(input product.GetProductDetailInput) (domains.CustomResult, error) {
	product, err := u.repository.FindById(input.ID)
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
	product.IsDeleted = 1
	// product.ExpiredAt = input.ExpiredAt

	newProduct, err := u.repository.Create(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (u *usecase) Update(inputID product.GetProductDetailInput, inputData product.CreateInput) (domains.Product, error) {
	product, err := u.repository.FindByProductID(inputID.ID)
	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		return product, errors.New("No product on that ID")
	}

	product.Name = inputData.Name
	product.Price = inputData.Price
	product.Quantity = inputData.Quantity
	product.Description = inputData.Description
	product.UserID = inputData.User.ID
	product.CategoryID = inputData.CategoryID

	now := time.Now()

	// Get the date by setting the time part to 00:00:00
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	product.Image = "https://img-global.cpcdn.com/recipes/93a46a53e22256b8/680x482cq70/songkolo-bagadang-ketan-serundeng-foto-resep-utama.jpg"
	product.EntryAt = today
	// product.ExpiredAt = input.ExpiredAt

	updateProduct, err := u.repository.Update(product)
	if err != nil {
		return updateProduct, err
	}

	return updateProduct, nil
}

func (u *usecase) UpdateQuantity(inputID product.GetProductDetailInput, inputData product.UpdateQuantity) (domains.Product, error) {
	product, err := u.repository.FindByProductID(inputID.ID)
	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		return product, errors.New("No product on that ID")
	}

	product.Quantity = inputData.Quantity

	now := time.Now()

	// Get the date by setting the time part to 00:00:00
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	product.Image = "https://img-global.cpcdn.com/recipes/93a46a53e22256b8/680x482cq70/songkolo-bagadang-ketan-serundeng-foto-resep-utama.jpg"
	product.EntryAt = today
	// product.ExpiredAt = input.ExpiredAt

	updateProduct, err := u.repository.Update(product)
	if err != nil {
		return updateProduct, err
	}

	return updateProduct, nil
}

func (u *usecase) Delete(input product.GetProductDetailInput) (domains.Product, error) {
	product, err := u.repository.FindByProductID(input.ID)
	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		return product, errors.New("No product on that ID")
	}

	if product.IsDeleted == 0 {
		return product, errors.New("The product has been previously removed")
	}

	product.IsDeleted = 0

	deleteProduct, err := u.repository.Update(product)
	if err != nil {
		return deleteProduct, err
	}

	return deleteProduct, nil
}
