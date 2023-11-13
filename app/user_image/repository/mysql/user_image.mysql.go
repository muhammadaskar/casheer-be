package mysql

import (
	"github.com/muhammadaskar/casheer-be/domains"
	"gorm.io/gorm"
)

type Repository interface {
	FindByUserId(user_id int) (domains.UserImage, error)
	Create(userImage domains.UserImage) (domains.UserImage, error)
	Update(userImage domains.UserImage) (domains.UserImage, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByUserId(user_id int) (domains.UserImage, error) {
	var userImage domains.UserImage
	err := r.db.Where("user_id = ?", user_id).Find(&userImage).Error
	if err != nil {
		return userImage, err
	}
	return userImage, nil
}

func (r *repository) Create(userImage domains.UserImage) (domains.UserImage, error) {
	err := r.db.Create(&userImage).Error
	if err != nil {
		return userImage, err
	}
	return userImage, nil
}

func (r *repository) Update(userImage domains.UserImage) (domains.UserImage, error) {
	err := r.db.Save(&userImage).Error
	if err != nil {
		return userImage, err
	}
	return userImage, nil
}
