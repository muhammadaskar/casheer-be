package usecase

import (
	"errors"

	"github.com/muhammadaskar/casheer-be/app/store"
	"github.com/muhammadaskar/casheer-be/app/store/repository/mysql"
	"github.com/muhammadaskar/casheer-be/domains"
	customstorage "github.com/muhammadaskar/casheer-be/utils/custom_storage"
)

type StoreUseCase interface {
	FindOne() (domains.Store, error)
	Create(input store.CreateInput) (domains.Store, error)
	Update(input store.CreateInput) (domains.Store, error)
}

type usecase struct {
	repository mysql.Repository
}

func NewUseCase(repository mysql.Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) FindOne() (domains.Store, error) {
	store, err := u.repository.FindOne()
	if err != nil {
		return store, err
	}

	if store.ID == 0 {
		return store, errors.New("Data not found")
	}

	image, err := customstorage.GetFileImage("assets/image/store/", store.Image)
	if err != nil {
		return store, err
	}

	store.Image = image

	return store, nil
}

func (u *usecase) Create(input store.CreateInput) (domains.Store, error) {
	store, err := u.repository.FindOne()
	if err != nil {
		return store, err
	}

	if store.ID != 0 {
		return store, errors.New("Data not found")
	}

	dataStore := domains.Store{}
	dataStore.Name = input.Name

	if input.Image == "" {
		return dataStore, errors.New("Image is required")
	}

	image, err := customstorage.Upload("assets/image/store", "logo", input.Image)

	dataStore.Image = image

	newStore, err := u.repository.Create(dataStore)
	if err != nil {
		return newStore, err
	}

	return dataStore, nil
}

func (u *usecase) Update(input store.CreateInput) (domains.Store, error) {
	store, err := u.repository.FindOne()
	if err != nil {
		return store, err
	}

	if store.ID != 1 {
		return store, errors.New("Data not found")
	}

	store.Name = input.Name
	if input.Image != "" {
		path := "assets/image/store/"

		err = customstorage.Delete(path, store.Image)
		if err != nil {
			return store, err
		}

		image, err := customstorage.Upload(path, "logo", input.Image)
		if err != nil {
			return store, err
		}

		store.Image = image
	}

	updateStore, err := u.repository.Update(store)
	if err != nil {
		return updateStore, err
	}

	return updateStore, nil
}
