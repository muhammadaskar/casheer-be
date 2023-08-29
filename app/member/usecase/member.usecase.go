package usecase

import (
	"errors"
	"math/rand"
	"time"

	"github.com/muhammadaskar/casheer-be/app/member"
	"github.com/muhammadaskar/casheer-be/app/member/repository/mysql"
	"github.com/muhammadaskar/casheer-be/domains"
)

type MemberUseCase interface {
	FindAll() ([]domains.Member, error)
	Create(input member.CreateInput) (domains.Member, error)
	Update(inputID member.GetMemberIDInput, inputData member.CreateInput) (domains.Member, error)
}

type usecase struct {
	repository mysql.Repository
}

func NewUseCase(repository mysql.Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) FindAll() ([]domains.Member, error) {
	members, err := u.repository.FindAll()
	if err != nil {
		return members, err
	}
	return members, nil
}

func (u *usecase) Create(input member.CreateInput) (domains.Member, error) {
	member := domains.Member{}

	memberCode := generateMemberCode()
	isMemberAvailable, err := u.isMemberCodeAvailable(memberCode)
	if err != nil {
		return member, err
	}

	if isMemberAvailable {
		return member, errors.New("Member code is not available")
	}

	member.MemberCode = memberCode

	member.Name = input.Name

	isPhoneAvailable, err := u.isMemberPhoneAvailable(input.Phone)
	if err != nil {
		return member, err
	}

	if isPhoneAvailable {
		return member, errors.New("Member phone is not available")
	}

	member.Phone = input.Phone
	member.CreatedBy = input.User.ID
	member.IsActive = 0

	newMember, err := u.repository.Create(member)
	if err != nil {
		return newMember, err
	}

	return newMember, nil
}

func (u *usecase) Update(inputID member.GetMemberIDInput, inputData member.CreateInput) (domains.Member, error) {
	member, err := u.repository.FindByID(inputID.ID)
	if err != nil {
		return member, err
	}

	member.Name = inputData.Name
	isPhoneAvailable, err := u.isMemberPhoneAvailable(inputData.Phone)
	if err != nil {
		return member, err
	}

	if isPhoneAvailable && inputData.Phone != member.Phone {
		return member, errors.New("Member phone is not available")
	}
	member.Phone = inputData.Phone

	updateMember, err := u.repository.Update(member)
	if err != nil {
		return updateMember, err
	}

	return updateMember, nil
}

func generateMemberCode() string {
	rand.Seed(time.Now().UnixNano())
	length := 12

	const availableDigits = "0123456789"
	numericPart := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(availableDigits))
		numericPart[i] = availableDigits[randomIndex]
	}

	return string(numericPart)
}

func (u *usecase) isMemberCodeAvailable(code string) (bool, error) {
	member, err := u.repository.FindByMemberCode(code)
	if err != nil {
		return false, err
	}

	if member.ID != 0 {
		return true, nil
	}

	return false, nil
}

func (u *usecase) isMemberPhoneAvailable(number string) (bool, error) {
	member, err := u.repository.FindByPhoneNumber(number)
	if err != nil {
		return false, err
	}

	if member.ID != 0 {
		return true, nil
	}

	return false, nil
}
