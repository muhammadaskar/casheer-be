package user

import (
	"github.com/muhammadaskar/casheer-be/app/notification"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByUsername(username string) (User, error)
	FindById(ID int) (User, error)
	CreateNotification(notification notification.Notification) (notification.Notification, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByUsername(username string) (User, error) {
	var user User

	err := r.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindById(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) CreateNotification(notification notification.Notification) (notification.Notification, error) {
	err := r.db.Create(&notification).Error
	if err != nil {
		return notification, err
	}

	return notification, nil
}
