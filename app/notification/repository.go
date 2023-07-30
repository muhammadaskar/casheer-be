package notification

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Notification, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Notification, error) {
	var notification []Notification
	err := r.db.Find(&notification).Error
	if err != nil {
		return notification, err
	}

	return notification, nil
}
