package mysql

import (
	"github.com/muhammadaskar/casheer-be/domains"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(search string, page int, limit int, noPagination bool) ([]domains.CustomResult, error)
	FindAllIsDeleted(search string, page int, limit int, noPagination bool) ([]domains.CustomResult, error)
	GetAll() ([]domains.CustomProduct, error)
	Count(is_deleted int) (int64, error)
	FindById(id int) (domains.CustomResult, error)
	FindByProductID(id int) (domains.Product, error)
	FindByProductCode(id string) (domains.Product, error)
	Create(product domains.Product) (domains.Product, error)
	Update(product domains.Product) (domains.Product, error)
	UpdateQty(id int, qty int) (domains.Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(search string, page int, limit int, noPagination bool) ([]domains.CustomResult, error) {
	var products []domains.CustomResult

	if noPagination {
		queryString := "%" + search + "%"
		query := `SELECT products.id, products.code as code, categories.id as category_id, products.name, categories.name as category, products.price, products.quantity, products.is_deleted, users.name as created_by, products.entry_at, products.description, products.created_at
			FROM products
			LEFT JOIN users ON products.user_id = users.id
			LEFT JOIN categories ON products.category_id = categories.id
			WHERE products.name LIKE ?
			AND products.is_deleted = 1
			ORDER BY products.is_deleted DESC;`

		err := r.db.Raw(query, queryString).Scan(&products).Error

		if err != nil {
			return products, err
		}
	} else {
		perPage := limit
		offset := (page - 1) * perPage

		query := `SELECT products.id, products.code as code, products.name, categories.id as category_id, categories.name as category, products.price, products.quantity, products.is_deleted, users.name as created_by, products.entry_at, products.description, products.created_at
				FROM products
				LEFT JOIN users ON products.user_id = users.id
				LEFT JOIN categories ON products.category_id = categories.id
				WHERE products.is_deleted = 1
				LIMIT ? OFFSET ?;`

		err := r.db.Raw(query, perPage, offset).Scan(&products).Error

		if err != nil {
			return products, err
		}

		return products, nil
	}
	return products, nil
}

func (r *repository) FindAllIsDeleted(search string, page int, limit int, noPagination bool) ([]domains.CustomResult, error) {
	var products []domains.CustomResult

	if noPagination {
		queryString := "%" + search + "%"
		query := `SELECT products.id, products.code as code, categories.id as category_id, products.name, categories.name as category, products.price, products.quantity, products.is_deleted, users.name as created_by, products.entry_at, products.description, products.created_at
			FROM products
			LEFT JOIN users ON products.user_id = users.id
			LEFT JOIN categories ON products.category_id = categories.id
			WHERE products.name LIKE ?
			AND products.is_deleted = 0
			ORDER BY products.is_deleted DESC;`

		err := r.db.Raw(query, queryString).Scan(&products).Error

		if err != nil {
			return products, err
		}
	} else {
		perPage := limit
		offset := (page - 1) * perPage

		query := `SELECT products.id, products.code as code, products.name, categories.id as category_id, categories.name as category, products.price, products.quantity, products.is_deleted, users.name as created_by, products.entry_at, products.description, products.created_at
				FROM products
				LEFT JOIN users ON products.user_id = users.id
				LEFT JOIN categories ON products.category_id = categories.id
				WHERE products.is_deleted = 0
				LIMIT ? OFFSET ?;`

		err := r.db.Raw(query, perPage, offset).Scan(&products).Error

		if err != nil {
			return products, err
		}

		return products, nil
	}
	return products, nil
}

func (r *repository) GetAll() ([]domains.CustomProduct, error) {
	var products []domains.CustomProduct

	query := `SELECT products.id, products.code as code, products.name, products.price, products.quantity, products.description FROM products WHERE is_deleted = 1`

	err := r.db.Raw(query).Scan(&products).Error

	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) Count(is_deleted int) (int64, error) {
	var products []domains.Product
	var count int64
	var err error

	if is_deleted == 2 {
		err = r.db.Model(&products).Count(&count).Error
	} else {
		err = r.db.Where("is_deleted = ?", is_deleted).Model(&products).Count(&count).Error
	}

	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *repository) FindById(id int) (domains.CustomResult, error) {
	var product domains.CustomResult

	query := `SELECT products.id, products.code as code, products.name, products.image, categories.name as category, products.price, products.quantity, products.is_deleted, users.name as created_by, products.entry_at, products.description, products.created_at
	FROM products
	LEFT JOIN users ON products.user_id = users.id
	LEFT JOIN categories ON products.category_id = categories.id
	WHERE products.id = ?;`
	err := r.db.Raw(query, id).Scan(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindByProductID(id int) (domains.Product, error) {
	var product domains.Product

	err := r.db.Where("id = ?", id).Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindByProductCode(code string) (domains.Product, error) {
	var product domains.Product

	err := r.db.Where("code = ?", code).Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) Create(product domains.Product) (domains.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Update(product domains.Product) (domains.Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) UpdateQty(id int, qty int) (domains.Product, error) {
	var product domains.Product
	err := r.db.Model(&product).Where("id = ?", id).Update("quantity", qty).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
