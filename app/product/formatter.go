package product

import "time"

type ProductFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatProduct(product Product) ProductFormatter {
	formatter := ProductFormatter{}
	formatter.ID = product.ID
	formatter.Name = product.Name
	formatter.Price = product.Price
	formatter.Quantity = product.Quantity
	formatter.CreatedAt = product.CreatedAt
	return formatter
}
