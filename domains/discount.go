package domains

import (
	"time"
)

type Discount struct {
	ID        int
	Discount  int
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
}
