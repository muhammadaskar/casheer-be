package domains

import "time"

type Transaction struct {
	ID              int
	MemberCode      string
	TransactionCode string
	Transactions    string
	Amount          int
	UserID          int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type TransactionProductQuantity struct {
	ProductID int
	Quantity  int
}
