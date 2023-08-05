package product

import (
	"time"

	"github.com/muhammadaskar/casheer-be/app/category"
	"github.com/muhammadaskar/casheer-be/app/user"
)

type Product struct {
	ID          int
	UserID      int
	CategoryID  int
	Name        string
	Price       int
	Quantity    int
	Image       string
	Description string
	EntryAt     time.Time
	ExpiredAt   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        user.User
	Category    category.Category
}
