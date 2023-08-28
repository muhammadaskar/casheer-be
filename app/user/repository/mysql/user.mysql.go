package mysql

import (
	"github.com/muhammadaskar/casheer-be/domains"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user domains.User) (domains.User, error)
	FindByEmail(email string) (domains.User, error)
	FindByUsername(username string) (domains.User, error)
	FindById(ID int) (domains.User, error)
	GetTotalCasheer() (domains.CustomTotalCasheer, error)
	Update(user domains.User) (domains.User, error)
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

	err := r.db.Where("username = ?", username).Find(&user).Error
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

func (r *repository) GetTotalCasheer() (domains.CustomTotalCasheer, error) {
	var casheer domains.CustomTotalCasheer

	query := `SELECT COUNT(*) AS count FROM users
	WHERE role=1
	AND is_active=0;`

	err := r.db.Raw(query).Scan(&casheer.TotalCasheer).Error
	if err != nil {
		return casheer, err
	}
	return casheer, nil
}

func (r *repository) Update(user domains.User) (domains.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
