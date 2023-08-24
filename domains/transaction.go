package domains

import "time"

type Transacation struct {
	ID              int
	MemberCode      string
	TransactionCode string
	Transacations   string
	Amount          int
	UserID          int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type TransacationProductQuantity struct {
	ProductID int
	Quantity  int
}
