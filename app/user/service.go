package user

import (
	"errors"
	"regexp"

	"github.com/muhammadaskar/casheer-be/app/notification"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(email string) (bool, error)
	IsUsernameAvailable(username string) (bool, error)
	GetUserById(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(input RegisterInput) (User, error) {
	user := User{}

	checkUsername := checkUsername(input.Username)
	if !checkUsername {
		return user, errors.New("Username is not valid.")
	}

	isUsernameAvailable, err := s.IsUsernameAvailable(input.Username)
	if err != nil {
		return user, err
	}

	if isUsernameAvailable {
		return user, errors.New("Username is not available.")
	}

	isEmailAvailable, err := s.IsEmailAvailable(input.Email)
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

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	notification := notification.Notification{}
	notification.Name = newUser.Name + " baru saja melakukan registrasi"
	notification.UserId = newUser.ID
	notification.Type = 1

	_, err = s.repository.CreateNotification(notification)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	username := input.Username
	password := input.Password

	user, err := s.repository.FindByUsername(username)
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

func (s *service) IsEmailAvailable(email string) (bool, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID != 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) IsUsernameAvailable(username string) (bool, error) {
	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return false, err
	}

	if user.ID != 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) GetUserById(ID int) (User, error) {
	user, err := s.repository.FindById(ID)
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
