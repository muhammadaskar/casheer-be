package mysql

import (
	"github.com/muhammadaskar/casheer-be/domains"
	"gorm.io/gorm"
)

type Repository interface {
	FindOne() (domains.Store, error)
	Create(store domains.Store) (domains.Store, error)
	Update(store domains.Store) (domains.Store, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindOne() (domains.Store, error) {
	var store domains.Store

	err := r.db.Where("id = 1").Find(&store).Error
	if err != nil {
		return store, err
	}
	return store, nil
}

func (r *repository) Create(store domains.Store) (domains.Store, error) {
	err := r.db.Create(&store).Error
	if err != nil {
		return store, err
	}
	return store, nil
}

func (r *repository) Update(store domains.Store) (domains.Store, error) {
	err := r.db.Save(&store).Error
	if err != nil {
		return store, err
	}
	return store, nil
}
