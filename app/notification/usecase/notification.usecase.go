package usecase

import (
	"github.com/muhammadaskar/casheer-be/app/notification/repository/mysql"
	"github.com/muhammadaskar/casheer-be/domains"
)

type NotificationUseCase interface {
	FindAll() ([]domains.Notification, error)
}

type usecase struct {
	repository mysql.Repository
}

func NewUseCase(repository mysql.Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) FindAll() ([]domains.Notification, error) {
	notification, err := u.repository.FindAll()
	if err != nil {
		return notification, err
	}

	return notification, nil
}
