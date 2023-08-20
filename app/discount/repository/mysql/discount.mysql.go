package mysql

import (
	"github.com/muhammadaskar/casheer-be/domains"
	"gorm.io/gorm"
)

type Repository interface {
	FindById() (domains.Discount, error)
	Create(discount domains.Discount) (domains.Discount, error)
	Update(discount domains.Discount) (domains.Discount, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindById() (domains.Discount, error) {
	var discount domains.Discount
	err := r.db.Where("id = ?", 1).Preload("User").Find(&discount).Error
	if err != nil {
		return discount, err
	}
	return discount, nil
}

func (r *repository) Create(discount domains.Discount) (domains.Discount, error) {
	err := r.db.Create(&discount).Error
	if err != nil {
		return discount, err
	}

	return discount, nil
}

func (r *repository) Update(discount domains.Discount) (domains.Discount, error) {
	err := r.db.Save(&discount).Error
	if err != nil {
		return discount, err
	}

	return discount, nil
}
