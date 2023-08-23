package transaction

import "github.com/muhammadaskar/casheer-be/domains"

type CreateInput struct {
	MemberCode string `json:"member_code"`
	ProductID  int    `json:"product_id" binding:"required"`
	Quantity   int    `json:"quantity" binding:"required"`
	// Amount     int    `json:"amount" binding:"required"`
	User domains.User
}
