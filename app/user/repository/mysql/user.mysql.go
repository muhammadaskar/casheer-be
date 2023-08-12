package mysql

import (
	"github.com/muhammadaskar/casheer-be/app/notification"
	"github.com/muhammadaskar/casheer-be/domains"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user domains.User) (domains.User, error)
	FindByEmail(email string) (domains.User, error)
	FindByUsername(username string) (domains.User, error)
	FindById(ID int) (domains.User, error)
	CreateNotification(notification notification.Notification) (notification.Notification, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user domains.User) (domains.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (domains.User, error) {
	var user domains.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByUsername(username string) (domains.User, error) {
	var user domains.User

	err := r.db.Select("id, name, username, email, password").Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindById(ID int) (domains.User, error) {
	var user domains.User

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
