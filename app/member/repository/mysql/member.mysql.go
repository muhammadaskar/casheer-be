package mysql

import (
	"github.com/muhammadaskar/casheer-be/domains"
	"gorm.io/gorm"
)

type Repository interface {
	FindByMemberCode(code string) (domains.Member, error)
	FindByPhoneNumber(number string) (domains.Member, error)
	Create(member domains.Member) (domains.Member, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByMemberCode(code string) (domains.Member, error) {
	var member domains.Member
	err := r.db.Where("member_code = ?", code).Find(&member).Error
	if err != nil {
		return member, err
	}
	return member, nil
}

func (r *repository) FindByPhoneNumber(number string) (domains.Member, error) {
	var member domains.Member
	err := r.db.Where("phone = ?", number).Find(&member).Error
	if err != nil {
		return member, err
	}
	return member, nil
}

func (r *repository) Create(member domains.Member) (domains.Member, error) {
	err := r.db.Create(&member).Error
	if err != nil {
		return member, err
	}
	return member, nil
}
