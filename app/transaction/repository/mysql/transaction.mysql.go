package mysql

import (
	"github.com/muhammadaskar/casheer-be/domains"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domains.CustomTransaction, error)
	FindAllMember() ([]domains.CustomTransactionMember, error)
	GetAmountOneMonthAgo(currentTime string, oneMonthAgo string) (domains.CustomTransactionAmount, error)
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

func (r *repository) FindAllMember() ([]domains.CustomTransactionMember, error) {
	var transactions []domains.CustomTransactionMember

	query := `SELECT transactions.id, transactions.member_code, members.name AS member_name, transaction_code, transactions, amount, users.name as name, transactions.created_at 
	FROM transactions
	INNER JOIN users ON transactions.user_id = users.id
	INNER JOIN members ON transactions.member_code = members.member_code;`
	err := r.db.Raw(query).Scan(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) GetAmountOneMonthAgo(currentTime string, oneMonthAgo string) (domains.CustomTransactionAmount, error) {
	var transaction domains.CustomTransactionAmount

	query := `SELECT SUM(amount) 
	FROM transactions 
	WHERE created_at >= ?
	AND created_at <=  ?`
	err := r.db.Raw(query, oneMonthAgo, currentTime).Scan(&transaction.Amount).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
