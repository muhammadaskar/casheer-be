package mysql

import (
	"github.com/muhammadaskar/casheer-be/domains"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domains.Category, error)
	FindById(ID int) (domains.Category, error)
	Save(category domains.Category) (domains.Category, error)
	Update(category domains.Category) (domains.Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domains.Category, error) {
	var categories []domains.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func (r *repository) FindById(ID int) (domains.Category, error) {
	var category domains.Category
	err := r.db.Where("id = ?", ID).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) Save(category domains.Category) (domains.Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) Update(category domains.Category) (domains.Category, error) {
	err := r.db.Save(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
