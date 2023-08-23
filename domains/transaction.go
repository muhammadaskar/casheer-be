package domains

import "time"

type Transacation struct {
	ID              int
	MemberCode      string
	TransactionCode string
	ProductID       int
	Quantity        int
	Amount          int
	UserID          int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
