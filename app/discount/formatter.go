package discount

import (
	"time"

	"github.com/muhammadaskar/casheer-be/domains"
)

type DiscountFormatter struct {
	ID        int       `json:"id"`
	Discount  int       `json:"discount"`
	User      string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatDiscount(discount domains.Discount) DiscountFormatter {
	formatter := DiscountFormatter{}
	formatter.ID = discount.ID
	formatter.Discount = discount.Discount
	formatter.User = discount.User.Name
	formatter.CreatedAt = discount.CreatedAt
	return formatter
}
