package usecase

import (
	"errors"
	"regexp"

	notificationRepo "github.com/muhammadaskar/casheer-be/app/notification/repository/mysql"
	"github.com/muhammadaskar/casheer-be/app/user"
	"github.com/muhammadaskar/casheer-be/app/user/repository/mysql"
	"github.com/muhammadaskar/casheer-be/domains"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Register(input user.RegisterInput) (domains.User, error)
	Login(input user.LoginInput) (domains.User, error)
	IsEmailAvailable(email string) (bool, error)
	IsUsernameAvailable(username string) (bool, error)
	GetUserById(ID int) (domains.User, error)
}

type usecase struct {
	userRepository         mysql.Repository
	notificationRepository notificationRepo.Repository
}

func NewUseCase(userRepository mysql.Repository, notificationRepository notificationRepo.Repository) *usecase {
	return &usecase{userRepository, notificationRepository}
}

func (u *usecase) Register(input user.RegisterInput) (domains.User, error) {
	user := domains.User{}

	checkUsername := checkUsername(input.Username)
	if !checkUsername {
		return user, errors.New("Username is not valid.")
	}

	isUsernameAvailable, err := u.IsUsernameAvailable(input.Username)
	if err != nil {
		return user, err
	}

	if isUsernameAvailable {
		return user, errors.New("Username is not available.")
	}

	isEmailAvailable, err := u.IsEmailAvailable(input.Email)
	if err != nil {
		return user, err
	}

	if isEmailAvailable {
		return user, errors.New("Email is not available.")
	}

	user.Name = input.Name
	user.Username = input.Username
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	user.Role = 1
	user.IsActive = 1

	newUser, err := u.userRepository.Save(user)

	if err != nil {
		return newUser, err
	}

	notification := domains.Notification{}
	notification.Name = newUser.Name + " baru saja melakukan registrasi"
	notification.UserId = newUser.ID
	notification.Type = 1
	notification.IsRead = 1

	_, err = u.notificationRepository.CreateNotification(notification)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (u *usecase) Login(input user.LoginInput) (domains.User, error) {
	username := input.Username
	password := input.Password

	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	if user.IsActive != 0 {
		return user, errors.New("Your account not active")
	}

	return user, nil
}

func (s *usecase) IsEmailAvailable(email string) (bool, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID != 0 {
		return true, nil
	}

	return false, nil
}

func (s *usecase) IsUsernameAvailable(username string) (bool, error) {
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return false, err
	}

	if user.ID != 0 {
		return true, nil
	}

	return false, nil
}

func (s *usecase) GetUserById(ID int) (domains.User, error) {
	user, err := s.userRepository.FindById(ID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func checkUsername(username string) bool {
	// Regular expression pattern to match a username without spaces
	pattern := "^[^\\s]+$"

	match, _ := regexp.MatchString(pattern, username)
	return match
}
