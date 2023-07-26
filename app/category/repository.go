package category

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Category, error) {
	var categories []Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return categories, err
	}
	return categories, nil
}
