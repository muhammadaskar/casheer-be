package product

import (
	"os/user"
	"time"
	// "github.com/muhammadaskar/casheer-be/app/user"
)

type Product struct {
	ID          int
	CategoryId  int
	Name        string
	Price       int
	Quantity    int
	Image       string
	Description string
	EntryAt     time.Time
	ExpiredAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        user.User
}
