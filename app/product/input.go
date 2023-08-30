package product

import "github.com/muhammadaskar/casheer-be/domains"

type GetProductDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetProductsQueryInput struct {
	Query string `form:"query"`
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
}

type CreateInput struct {
	// UserID     int    `json:"user_id" bind:"required"`
	CategoryID int    `json:"category_id" bind:"required"`
	Name       string `json:"name" bind:"required"`
	Price      int    `json:"price" bind:"required"`
	Quantity   int    `json:"quantity" bind:"required"`
	// Image int `json:"image" bind:"required"`
	Description string `json:"description" bind:"required"`
	User        domains.User
}

type UpdateQuantity struct {
	Quantity int `json:"quantity" bind:"required"`
}
