package usecase

import (
	"errors"
	"math"
	"time"

	notifRepo "github.com/muhammadaskar/casheer-be/app/notification/repository/mysql"
	"github.com/muhammadaskar/casheer-be/app/product"
	productRepo "github.com/muhammadaskar/casheer-be/app/product/repository/mysql"
	"github.com/muhammadaskar/casheer-be/domains"
)

type ProductUseCase interface {
	FindAll(query product.GetProductsQueryInput) ([]domains.CustomResult, bool, error)
	FindAllIsDeleted(query product.GetProductsQueryInput) ([]domains.CustomResult, bool, error)
	GetAll() ([]domains.CustomProduct, error)
	CountAll(is_deleted int) (int64, error)
	FindById(input product.GetProductDetailInput) (domains.CustomResult, error)
	Create(input product.CreateInput) (domains.Product, error)
	Update(inputID product.GetProductDetailInput, inputData product.CreateInput) (domains.Product, error)
	Delete(input product.GetProductDetailInput) (domains.Product, error)
	UpdateQuantity(inputID product.GetProductDetailInput, inputData product.UpdateQuantity) (domains.Product, error)
}

type usecase struct {
	productRepository      productRepo.Repository
	notificationRepository notifRepo.Repository
}

func NewUseCase(productRepository productRepo.Repository, notificationRepository notifRepo.Repository) *usecase {
	return &usecase{productRepository, notificationRepository}
}

func (u *usecase) FindAll(query product.GetProductsQueryInput) ([]domains.CustomResult, bool, error) {
	if query.Query != "" {
		products, err := u.productRepository.FindAll(query.Query, query.Page, query.Limit, true)
		if err != nil {
			return products, true, err
		}

		_, err = u.setNotification(products)
		if err != nil {
			return products, true, err
		}

		return products, true, nil
	} else {
		products, err := u.productRepository.FindAll(query.Query, query.Page, query.Limit, false)
		if err != nil {
			return products, true, err
		}

		_, err = u.setNotification(products)
		if err != nil {
			return products, true, err
		}

		totalCount, err := u.CountAll(1)
		if err != nil {
			return products, true, err
		}

		perPage := query.Limit
		offset := (query.Page - 1) * perPage
		totalPages := int(math.Ceil(float64(totalCount) / float64(perPage)))

		// Hitung nomor halaman saat ini berdasarkan offset dan produk per halaman
		currentPage := (offset / perPage) + 1

		// Periksa apakah Anda berada di halaman terakhir
		isLastPage := currentPage == totalPages

		return products, isLastPage, nil
	}
}

func (u *usecase) FindAllIsDeleted(query product.GetProductsQueryInput) ([]domains.CustomResult, bool, error) {
	if query.Query != "" {
		products, err := u.productRepository.FindAllIsDeleted(query.Query, query.Page, query.Limit, true)
		if err != nil {
			return products, true, err
		}

		_, err = u.setNotification(products)
		if err != nil {
			return products, true, err
		}

		return products, true, nil
	} else {
		products, err := u.productRepository.FindAllIsDeleted(query.Query, query.Page, query.Limit, false)
		if err != nil {
			return products, true, err
		}

		_, err = u.setNotification(products)
		if err != nil {
			return products, true, err
		}

		totalCount, err := u.CountAll(0)
		if err != nil {
			return products, true, err
		}

		perPage := query.Limit
		offset := (query.Page - 1) * perPage
		totalPages := int(math.Ceil(float64(totalCount) / float64(perPage)))

		// Hitung nomor halaman saat ini berdasarkan offset dan produk per halaman
		currentPage := (offset / perPage) + 1

		// Periksa apakah Anda berada di halaman terakhir
		isLastPage := currentPage == totalPages

		return products, isLastPage, nil
	}
}

func (u *usecase) GetAll() ([]domains.CustomProduct, error) {
	products, err := u.productRepository.GetAll()
	if err != nil {
		return products, err
	}

	for _, product := range products {
		if product.Quantity == 0 || product.Quantity < 5 {
			notif, err := u.notificationRepository.FindByProductID(int(product.ID))
			if err != nil {
				return products, err
			}
			if notif.ID == 0 {

				var notificationName string
				var notificationType int

				if product.Quantity == 0 {
					notificationName = product.Name + " stock habis"
					notificationType = 2
				} else if product.Quantity < 5 {
					notificationName = product.Name + " stock menipis"
					notificationType = 3
				}

				notification := domains.Notification{}
				notification.Name = notificationName
				notification.ProductId = int(product.ID)
				notification.Type = notificationType
				notification.IsRead = 1
				_, err = u.notificationRepository.CreateNotification(notification)
				if err != nil {
					return products, err
				}
			}
		}
	}

	return products, nil
}

func (u *usecase) CountAll(is_deleted int) (int64, error) {
	count, err := u.productRepository.Count(is_deleted)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (u *usecase) FindById(input product.GetProductDetailInput) (domains.CustomResult, error) {
	product, err := u.productRepository.FindById(input.ID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (u *usecase) Create(input product.CreateInput) (domains.Product, error) {
	product := domains.Product{}

	productCode, err := u.productRepository.FindByProductCode(input.Code)
	if err != nil {
		return product, err
	}

	if productCode.ID != 0 {
		return product, errors.New("Product code is available")
	}

	product.Code = input.Code
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

	newProduct, err := u.productRepository.Create(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (u *usecase) Update(inputID product.GetProductDetailInput, inputData product.CreateInput) (domains.Product, error) {
	product, err := u.productRepository.FindByProductID(inputID.ID)
	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		return product, errors.New("No product on that ID")
	}

	if inputData.Code != product.Code {
		productCode, err := u.productRepository.FindByProductCode(inputData.Code)
		if err != nil {
			return product, err
		}

		if productCode.ID != 0 {
			return product, errors.New("Product code is available")
		}
	}

	product.Code = inputData.Code
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

	updateProduct, err := u.productRepository.Update(product)
	if err != nil {
		return updateProduct, err
	}

	return updateProduct, nil
}

func (u *usecase) UpdateQuantity(inputID product.GetProductDetailInput, inputData product.UpdateQuantity) (domains.Product, error) {
	product, err := u.productRepository.FindByProductID(inputID.ID)
	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		return product, errors.New("No product on that ID")
	}
	updateProduct, err := u.productRepository.UpdateQty(int(product.ID), inputData.Quantity)
	if err != nil {
		return updateProduct, err
	}

	notification, err := u.notificationRepository.FindByProductID(product.ID)
	if err != nil {
		return updateProduct, err
	}

	if inputData.Quantity > 0 || inputData.Quantity >= 5 {
		if notification.Type == 2 || notification.Type == 3 {
			_, err := u.notificationRepository.DeleteNotification(notification)
			if err != nil {
				return updateProduct, err
			}
		}
	}

	return updateProduct, nil
}

func (u *usecase) Delete(input product.GetProductDetailInput) (domains.Product, error) {
	product, err := u.productRepository.FindByProductID(input.ID)
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

	deleteProduct, err := u.productRepository.Update(product)
	if err != nil {
		return deleteProduct, err
	}

	return deleteProduct, nil
}

func (u *usecase) setNotification(products []domains.CustomResult) ([]domains.CustomResult, error) {
	for _, product := range products {
		if product.Quantity == 0 || product.Quantity < 5 {
			notif, err := u.notificationRepository.FindByProductID(int(product.ID))
			if err != nil {
				return products, err
			}
			if notif.ID == 0 {

				var notificationName string
				var notificationType int

				if product.Quantity == 0 {
					notificationName = product.Name + " stock habis"
					notificationType = 2
				} else if product.Quantity < 5 {
					notificationName = product.Name + " stock menipis"
					notificationType = 3
				}

				notification := domains.Notification{}
				notification.Name = notificationName
				notification.ProductId = int(product.ID)
				notification.Type = notificationType
				notification.IsRead = 1
				_, err = u.notificationRepository.CreateNotification(notification)
				if err != nil {
					return products, err
				}
			}
		}
	}
	return products, nil
}
