package domains

import "time"

type Member struct {
	ID         int
	MemberCode string
	Name       string
	IsActive   int
	Phone      string
	CreatedBy  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
