package category

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Category, error)
	FindById(ID int) (Category, error)
	Save(category Category) (Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Category, error) {
	var categories []Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func (r *repository) FindById(ID int) (Category, error) {
	var category Category
	err := r.db.Where("id = ?", ID).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) Save(category Category) (Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
