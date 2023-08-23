package mysql

import (
	"github.com/muhammadaskar/casheer-be/domains"
	"gorm.io/gorm"
)

type Repository interface {
	Create(transaction domains.Transacation) (domains.Transacation, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(transacation domains.Transacation) (domains.Transacation, error) {
	err := r.db.Create(&transacation).Error
	if err != nil {
		return transacation, err
	}
	return transacation, nil
}
