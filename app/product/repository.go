package product

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]CustomResult, error)
	Create(product Product) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]CustomResult, error) {
	var products []CustomResult
	query := `SELECT products.id, products.name, products.image, categories.name as category, products.price, users.name as created_by, products.entry_at, products.created_at
	FROM products
	LEFT JOIN users ON products.user_id = users.id
	LEFT JOIN categories ON products.category_id = categories.id;`
	err := r.db.Raw(query).Scan(&products).Error

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
