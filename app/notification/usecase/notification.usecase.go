package usecase

import (
	"errors"

	"github.com/muhammadaskar/casheer-be/app/notification"
	"github.com/muhammadaskar/casheer-be/app/notification/repository/mysql"
	"github.com/muhammadaskar/casheer-be/domains"
)

type NotificationUseCase interface {
	FindAll() ([]domains.Notification, error)
	Update(inputID notification.GetInputID, inputData notification.InputUpdateNotification) (domains.Notification, error)
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

func (u *usecase) Update(inputID notification.GetInputID, inputData notification.InputUpdateNotification) (domains.Notification, error) {
	notification, err := u.repository.FindByID(inputID.ID)
	if err != nil {
		return notification, err
	}

	if notification.ID == 0 {
		return notification, errors.New("Notification is not available")
	}

	notification.IsRead = inputData.IsRead

	updateNotification, err := u.repository.UpdateNotification(notification)
	if err != nil {
		return updateNotification, err
	}
	return updateNotification, nil
}
