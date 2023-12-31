package mysql

import (
	"github.com/muhammadaskar/casheer-be/domains"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domains.CustomTransaction, error)
	FindAllMember() ([]domains.CustomTransactionMember, error)
	FindById(id int) (domains.CustomTransaction, error)
	FindMemberById(id int) (domains.CustomTransactionMember, error)
	GetAmountOneMonthAgo(currentTime string, oneMonthAgo string) (domains.CustomTransactionAmount, error)
	GetItemOneOutMonthAgo(currentTime string, oneMonthAgo string) (domains.CustomTransactionTotalQuantity, error)
	GetCountTransactionThisYear(start string, end string) ([]domains.GetCountTransactionThisYear, error)
	GetAmountTransactionThisYear(start string, end string) ([]domains.GetAmountTransactionThisYear, error)
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

	query := `SELECT transactions.id, member_code, transaction_code, transactions, total_quantity, amount, users.name as name, transactions.created_at 
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

	query := `SELECT transactions.id, transactions.member_code, members.name AS member_name, transaction_code, transactions, total_quantity, amount, users.name as name, transactions.created_at 
	FROM transactions
	INNER JOIN users ON transactions.user_id = users.id
	INNER JOIN members ON transactions.member_code = members.member_code;`
	err := r.db.Raw(query).Scan(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) FindById(id int) (domains.CustomTransaction, error) {
	var transaction domains.CustomTransaction

	query := `SELECT transactions.id, member_code, transaction_code, transactions, total_quantity, amount, users.name as name, transactions.created_at 
	FROM transactions
	LEFT JOIN users ON transactions.user_id = users.id
	WHERE transactions.id = ?`
	err := r.db.Raw(query, id).Scan(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) FindMemberById(id int) (domains.CustomTransactionMember, error) {
	var transaction domains.CustomTransactionMember

	query := `SELECT transactions.id, transactions.member_code, members.name AS member_name, transaction_code, transactions, total_quantity, amount, users.name as name, transactions.created_at 
	FROM transactions
	INNER JOIN users ON transactions.user_id = users.id
	INNER JOIN members ON transactions.member_code = members.member_code
	WHERE transactions.id = ?;`
	err := r.db.Raw(query, id).Scan(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
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

func (r *repository) GetItemOneOutMonthAgo(currentTime string, oneMonthAgo string) (domains.CustomTransactionTotalQuantity, error) {
	var transaction domains.CustomTransactionTotalQuantity

	query := `SELECT SUM(total_quantity) 
	FROM transactions 
	WHERE created_at >= ?
	AND created_at <=  ?`
	err := r.db.Raw(query, oneMonthAgo, currentTime).Scan(&transaction.TotalQuantity).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) GetCountTransactionThisYear(start string, end string) ([]domains.GetCountTransactionThisYear, error) {
	var transactions []domains.GetCountTransactionThisYear

	query := `SELECT months.month AS Month,
				IFNULL(SUM(p.count), 0) AS count
			FROM (
				SELECT 1 AS month
				UNION SELECT 2 UNION SELECT 3 UNION SELECT 4
				UNION SELECT 5 UNION SELECT 6 UNION SELECT 7
				UNION SELECT 8 UNION SELECT 9 UNION SELECT 10
				UNION SELECT 11 UNION SELECT 12
			) AS months
			LEFT JOIN (
				SELECT
					DATE_FORMAT(created_at, '%m') AS month,
					COUNT(*) AS count
				FROM
					transactions
				WHERE
					created_at >= ?
					AND created_at <= ?
				GROUP BY
					month
			) AS p ON months.month = p.month
			GROUP BY
				months.month;`

	err := r.db.Raw(query, start, end).Scan(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) GetAmountTransactionThisYear(start string, end string) ([]domains.GetAmountTransactionThisYear, error) {
	var transactions []domains.GetAmountTransactionThisYear

	query := `SELECT months.month AS Month,
				IFNULL(SUM(p.amount), 0) AS amount
			FROM (
				SELECT 1 AS month
				UNION SELECT 2 UNION SELECT 3 UNION SELECT 4
				UNION SELECT 5 UNION SELECT 6 UNION SELECT 7
				UNION SELECT 8 UNION SELECT 9 UNION SELECT 10
				UNION SELECT 11 UNION SELECT 12
			) AS months
			LEFT JOIN (
				SELECT
					DATE_FORMAT(created_at, '%m') AS month,
					SUM(amount) AS amount
				FROM
					transactions
				WHERE
					created_at >= ?
					AND created_at <= ?
				GROUP BY
					month
			) AS p ON months.month = p.month
			GROUP BY
				months.month;`

	err := r.db.Raw(query, start, end).Scan(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
