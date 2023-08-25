package mysql

import (
	"github.com/muhammadaskar/casheer-be/domains"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domains.CustomTransaction, error)
	Create(transaction domains.Transaction) (domains.Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(transaction domains.Transaction) (domains.Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) FindAll() ([]domains.CustomTransaction, error) {
	var transactions []domains.CustomTransaction

	query := `SELECT transactions.id, member_code, transaction_code, transactions, amount, users.name as name, transactions.created_at 
	FROM transactions
	LEFT JOIN users ON transactions.user_id = users.id`
	err := r.db.Raw(query).Scan(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
