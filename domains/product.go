package domains

import (
	"time"
	// "github.com/muhammadaskar/casheer-be/app/category"
	// "github.com/muhammadaskar/casheer-be/app/user"
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
	IsDeleted   int
	EntryAt     time.Time
	ExpiredAt   *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// User        user.User
	// Category    category.Category
}

type CustomResult struct {
	ID         uint      `json:"id"`
	CategoryID int       `json:"category_id"`
	Name       string    `json:"name"`
	Category   string    `json:"category"`
	Price      float64   `json:"price"`
	Quantity   int       `json:"quantity"`
	CreatedBy  string    `json:"created_by"`
	EntryAt    time.Time `json:"entry_at"`
	CreatedAt  time.Time `json:"created_at"`
}

type CustomProduct struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
