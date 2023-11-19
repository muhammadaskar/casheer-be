package usecase

import (
	"errors"
	"fmt"

	"github.com/muhammadaskar/casheer-be/app/user_image/repository/mysql"
	"github.com/muhammadaskar/casheer-be/domains"
	customgenerate "github.com/muhammadaskar/casheer-be/utils/custom_generate"
	customstorage "github.com/muhammadaskar/casheer-be/utils/custom_storage"
)

type UserImageUseCase interface {
	FindByUserId(id int) (domains.UserImage, error)
	CreateOrUpdate(userImageInput domains.UserImageInput) (domains.UserImage, error)
}

type usecase struct {
	userImageRepository mysql.Repository
}

func NewUseCase(userImageRepository mysql.Repository) *usecase {
	return &usecase{userImageRepository}
}

func (u *usecase) FindByUserId(id int) (domains.UserImage, error) {
	userImage, err := u.userImageRepository.FindByUserId(id)
	if err != nil {
		fmt.Println(userImage)
		return userImage, err
	}
	if userImage.ID == 0 {
		return userImage, errors.New("DATA NOT FOUND")
	}

	if userImage.Image == "" {
		return userImage, errors.New("IMAGE NOT FOUND")
	}

	image, err := customstorage.GetFileImage("asset/image/user/", userImage.Image)
	if err != nil {
		return userImage, err
	}
	userImage.Image = image

	return userImage, nil
}

func (u *usecase) CreateOrUpdate(userImageInput domains.UserImageInput) (domains.UserImage, error) {
	userImage, err := u.userImageRepository.FindByUserId(userImageInput.UserID)
	if err != nil {
		return userImage, err
	}
	path := "asset/image/user"

	fmt.Println(userImage)

	if userImage.ID == 0 {
		if userImageInput.Image == "" {
			return userImage, errors.New("IMAGE IS REQUIRED")
		}

		fileName := customgenerate.GenerateCode(12)

		image, err := customstorage.Upload(path, fileName, userImageInput.Image)
		if err != nil {
			return userImage, err
		}

		dataUserImage := domains.UserImage{}
		dataUserImage.UserID = userImageInput.UserID
		dataUserImage.Image = image

		newUserImage, err := u.userImageRepository.Create(dataUserImage)
		if err != nil {
			return newUserImage, err
		}
		return newUserImage, nil
	}
	if userImageInput.Image == "" {
		return userImage, errors.New("IMAGE IS REQUIRED")
	}

	err = customstorage.Delete(path+"/", userImage.Image)
	if err != nil {
		return userImage, err
	}

	fileName := customgenerate.GenerateCode(12)

	image, err := customstorage.Upload(path, fileName, userImageInput.Image)
	if err != nil {
		return userImage, err
	}

	userImage.UserID = userImageInput.UserID
	userImage.Image = image

	newUserImage, err := u.userImageRepository.Update(userImage)
	if err != nil {
		return newUserImage, err
	}
	return newUserImage, nil
}
