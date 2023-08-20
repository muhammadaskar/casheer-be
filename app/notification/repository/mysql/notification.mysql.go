package mysql

import (
	"github.com/muhammadaskar/casheer-be/domains"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domains.Notification, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domains.Notification, error) {
	var notification []domains.Notification
	err := r.db.Find(&notification).Error
	if err != nil {
		return notification, err
	}

	return notification, nil
}