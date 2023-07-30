package notification

import (
	"time"
	// "github.com/muhammadaskar/casheer-be/app/product"
	// "github.com/muhammadaskar/casheer-be/app/user"
)

type Notification struct {
	ID        int
	Name      string
	Type      int
	UserId    int
	ProductId int
	CreatedAt time.Time
	UpdatedAt time.Time
	// User      user.User
	// Product   product.Product
}
