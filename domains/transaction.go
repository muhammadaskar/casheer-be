package domains

import "time"

type Transaction struct {
	ID              int
	MemberCode      string
	TransactionCode string
	Transactions    string
	TotalQuantity   int
	Amount          int
	UserID          int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type TransactionProductQuantity struct {
	ProductID int
	Quantity  int
}

type CustomTransaction struct {
	ID              int       `json:"id"`
	MemberCode      string    `json:"member_code"`
	TransactionCode string    `json:"transaction_code"`
	Transactions    string    `json:"product_and_quantity"`
	TotalQuantity   int       `json:"total_quantity"`
	Amount          int       `json:"amount"`
	Name            string    `json:"casheer_name"`
	CreatedAt       time.Time `json:"created_at"`
}

type CustomTransactionMember struct {
	ID              int       `json:"id"`
	MemberCode      string    `json:"member_code"`
	TransactionCode string    `json:"transaction_code"`
	MemberName      string    `json:"member_name"`
	Transactions    string    `json:"product_and_quantity"`
	TotalQuantity   int       `json:"total_quantity"`
	Amount          int       `json:"amount"`
	Name            string    `json:"casheer_name"`
	CreatedAt       time.Time `json:"created_at"`
}

type CustomTransactionAmount struct {
	Amount int `json:"amount"`
}

type CustomTransactionTotalQuantity struct {
	TotalQuantity int `json:"total_quantity"`
}

type GetInputIdTransaction struct {
	ID int `uri:"id" binding:"required"`
}

type GetCountTransactionThisYear struct {
	Month string `json:"month"`
	Count int    `json:"count"`
}

type GetAmountTransactionThisYear struct {
	Month  string `json:"month"`
	Amount int    `json:"amount"`
}
