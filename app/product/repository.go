package product

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Product, error)
	Create(product Product) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Preload("User").Preload("Category").Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) Create(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
